package service

import (
	"TikTokLite_v2/util"
	"TikTokLite_v2/video/dal"
	"TikTokLite_v2/video/pb"
	"TikTokLite_v2/video/remote_call/call_fav_com"
	"TikTokLite_v2/video/remote_call/call_user_follow"
	"TikTokLite_v2/video/setting"
	bytes2 "bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
)

func PublishVideo(ctx context.Context, data []byte, filename string, userId int64, video dal.Video) error {
	//获取uuid，拼接视频名称，方便调试就先加上user_id和视频名称
	uuid := util.GetUUID()
	//先把视频保存本地，再制作封面，再一起上传到七牛云，完成后删除本地视频和封面
	//增加本地视频压缩（改变码率），再上传
	oldVideoName := fmt.Sprintf("old_%s_%d_%s", uuid, userId, filename)
	oldVideoPath := setting.Conf.VideoPathPrefix + oldVideoName
	videoName := fmt.Sprintf("%s_%d_%s", uuid, userId, filename)
	videoPath := setting.Conf.VideoPathPrefix + videoName
	//先保存本地然后压缩后再取出第一帧,(后可选一起上传至七牛云)
	if err := SaveUploadedFile(data, oldVideoPath); err != nil {
		log.Println("本地存储video失败", err)
		return err
	}
	//先截取第一帧做封面，再进行压缩
	coverName, err := getCoverName(videoName)
	if err != nil {
		log.Println("获取coverName失败：", err)
		return err
	}
	coverPath := setting.Conf.CoverPathPrefix + coverName
	cmd := exec.Command("ffmpeg", "-i", oldVideoPath, "-y", "-f", "mjpeg", "-ss", "0.1", "-t", "0.001", coverPath)
	if err := cmd.Run(); err != nil {
		log.Println("执行ffmpeg截取封面失败：", err)
		return err
	}
	var playUrl, coverUrl string
	//上传至七牛云
	if setting.Conf.PublishConfig.Mode {
		playUrl = setting.Conf.QiNiuCloudPlayUrlPrefix + videoName
		coverUrl = setting.Conf.QiNiuCloudCoverUrlPrefix + coverName
		video.PlayUrl = playUrl
		video.CoverUrl = coverUrl
		go func() {
			//压缩视频
			compressedVideo(oldVideoPath, videoPath)
			//上传
			err = uploadVideoToQiNiuCloud(ctx, videoName, coverName, videoPath, coverPath, video)
			if err != nil {
				log.Println("七牛云上传失败：", err)
			}
		}()
		return nil
	}
	playUrl = fmt.Sprintf("http://%s:%d/static/videos/?name=%s", setting.Conf.LocalIP, setting.Conf.PublishConfig.Port, videoName)
	coverUrl = fmt.Sprintf("http://%s:%d/static/covers/?name=%s", setting.Conf.LocalIP, setting.Conf.PublishConfig.Port, coverName)
	video.PlayUrl = playUrl
	video.CoverUrl = coverUrl
	//fmt.Println("!!!!!!!", coverUrl)
	go compressedVideo(oldVideoPath, videoPath) //异步压缩视频，否则很容易超时导致ctx取消。
	//这里不使用异步的原因是，CreateVideo还会进行RPC调用，如果这里提前返回，ctx就会被cancel，会出现级联cancel
	err = CreateVideo(ctx, &video)

	return err
}
func SaveUploadedFile(data []byte, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	dataReader := bytes.NewReader(data)
	n, err := io.Copy(f, dataReader)
	if err != nil {
		return err
	}
	if n < int64(len(data)) {
		return errors.New("io.Copy File incomplete")
	}
	return nil
}
func compressedVideo(oldVideoPath, videoPath string) {
	defer os.Remove(oldVideoPath)
	//压缩视频（减小码率）
	cmd := exec.Command("ffmpeg", "-i", oldVideoPath, "-b:v", "1.5M", videoPath)
	if err := cmd.Run(); err != nil {
		log.Println("执行ffmpeg压缩视频失败：", err)
		return
	}
}

func uploadVideoToQiNiuCloud(ctx context.Context, videoName, coverName, videoPath, coverPath string, video dal.Video) error {
	videoData, err := os.Open(videoPath)
	if err != nil {
		log.Println("创建cover失败：", err)
		return err
	}
	cover, err := os.Open(coverPath)
	if err != nil {
		log.Println("创建cover失败：", err)
		return err
	}
	defer os.Remove(videoPath)
	//因为先进后出，所以得先关闭链接之后再删除
	defer os.Remove(coverPath)
	defer cover.Close()
	defer videoData.Close()
	//最后上传至七牛云
	videoDataStat, err := videoData.Stat()
	if err != nil {
		log.Println("打开videoData.Stat失败：", err)
		return err
	}
	coverStat, err := cover.Stat()
	if err != nil {
		log.Println("打开cover.Stat失败：", err)
		return err
	}
	//上传的路径+文件名
	videoKey := fmt.Sprintf("videos/%s", videoName)
	coverKey := fmt.Sprintf("covers/%s", coverName)
	//上传凭证
	mac := qbox.NewMac(setting.Conf.AccessKey, setting.Conf.PublishConfig.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: setting.Conf.BucketName,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		// 空间对应的机房
		Zone: &storage.ZoneHuanan,
		// 是否使用https域名
		UseHTTPS: true,
		// 上传是否使用CDN上传加速
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	//额外参数
	putExtra := storage.PutExtra{
		//Params: map[string]string{
		//	"x:name": "github logo",
		//},
	}
	err = formUploader.Put(context.Background(), &ret, upToken, videoKey, videoData, videoDataStat.Size(), &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Println(ret.Key, ret.Hash) //打印此次上传的一些信息
	err = formUploader.Put(context.Background(), &ret, upToken, coverKey, cover, coverStat.Size(), &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	CreateVideo(ctx, &video)
	return nil
}

func uploadVideoToCloud(videoPath, videoName string) error {
	buf := bytes2.Buffer{}
	bodyWriter := multipart.NewWriter(&buf)
	fileWriter, _ := bodyWriter.CreateFormFile("video", videoPath)
	f, _ := os.Open(videoPath)
	defer f.Close()
	io.Copy(fileWriter, f)
	contenType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := fmt.Sprintf("http://0.0.0.0:0000/upload_video?video_name=%s", videoName)
	http.Post(url, contenType, &buf)
	return nil
}
func uploadCoverToCloud(coverPath, coverName string) error {
	buf := bytes2.Buffer{}
	bodyWriter := multipart.NewWriter(&buf)
	fileWriter, _ := bodyWriter.CreateFormFile("cover", coverPath)
	f, _ := os.Open(coverPath)
	defer f.Close()
	io.Copy(fileWriter, f)
	contenType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := fmt.Sprintf("http://0.0.0.0:0000/upload_cover?cover_name=%s", coverName)
	http.Post(url, contenType, &buf)
	return nil
}

func getCoverName(s string) (string, error) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			return fmt.Sprintf("%s.jpeg", s[:i]), nil
		}
	}
	return "", errors.New("文件名格式不合法")
}

func CreateVideo(ctx context.Context, v *dal.Video) error {
	err := v.Create(ctx)
	if err != nil {
		return err
	}
	//初始化redis中的点赞数和评论数
	call_fav_com.SetFavoriteNum(ctx, int64(v.ID), 0)
	call_fav_com.SetCommentNum(ctx, int64(v.ID), 0)
	now := v.CreatedAt.UnixMilli()
	//更新authorfeed
	_ = dal.InsertAuthorFeed(ctx, v.AuthorId, int64(v.ID), now)
	//更新userfeed
	var followers []int64
	followers, err = call_user_follow.GetFollowerListID(ctx, v.AuthorId)
	if err != nil {
		log.Println("fail to get follower list by", v.AuthorId)
		return err
	}
	_ = dal.PushNewVideoToActiveUsersFeed(ctx, followers, v.AuthorId, int64(v.ID), now)
	return nil
}

func GetVideoList(ctx context.Context, userId, toUserId int64) (*pb.PublishListResponse, error) {
	//数据库表格式的videos
	videos, err := dal.GetVideosByUserId(ctx, toUserId)
	if err != nil {
		log.Println("getVideosByUserId failed:", err)
		return &pb.PublishListResponse{}, err
	}
	videoList := make([]pb.Video, len(videos))
	//user复用一个就行
	userResp, err := call_user_follow.GetUser(ctx, userId, toUserId)
	if err != nil {
		return &pb.PublishListResponse{}, err
	}
	author := userResp.User
	//不能直接用video
	resp := &pb.PublishListResponse{}
	resp.VideoList = make([]*pb.Video, len(videos))

	for i := range videoList {
		video := pb.Video{}
		videoId := int64(videos[i].ID)
		video.Id = videoId
		video.Author = author
		video.Title = videos[i].Title
		video.PlayUrl = videos[i].PlayUrl
		video.CoverUrl = videos[i].CoverUrl
		video.FavoriteCount, video.CommentCount, video.IsFavorite =
			call_fav_com.GetFavoriteAndCommentInfo(ctx, userId, videoId)
		resp.VideoList[i] = &video
	}
	return resp, nil
}
func GetTotalWorkCount(ctx context.Context, userId int64) (*pb.GetTotalWorkCountResponse, error) {
	resp := &pb.GetTotalWorkCountResponse{}
	resp.Count = dal.GetTotalWorkCount(ctx, userId)
	return resp, nil
}
func GetVideoIDsByUser(ctx context.Context, userId int64) (*pb.GetVideoIDListOfUserResponse, error) {
	resp := &pb.GetVideoIDListOfUserResponse{}
	resp.VideoIdList = dal.GetVideoIDsByUser(ctx, userId)
	return resp, nil
}

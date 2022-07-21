package service

import (
	"TikTokLite_v2/favorite_comment/dal"
	"TikTokLite_v2/favorite_comment/dal/redb"
	"TikTokLite_v2/favorite_comment/pb"
	"TikTokLite_v2/favorite_comment/remote_call/call_user_follow"
	"TikTokLite_v2/favorite_comment/remote_call/call_video"
	pb3 "TikTokLite_v2/user_follow/user/pb"
	pb2 "TikTokLite_v2/video/pb"
	"context"
	"github.com/gistao/RedisGo-Async/redis"
)

func SetFavorite(ctx context.Context, videoID, userID int64) error {
	f := dal.Favorite{
		VideoID: videoID,
		UserID:  userID,
	}
	return f.UniqueInsert(ctx)
}

// CancelFavorite 取消点赞
func CancelFavorite(ctx context.Context, videoID, userID int64) error {
	f := dal.Favorite{
		VideoID: videoID,
		UserID:  userID,
	}
	return f.Delete(ctx)
}

// GetFavoriteList 获取喜欢列表
func GetFavoriteList(ctx context.Context, userID int64) (*pb.FavoriteListResponse, error) {
	rows, err := dal.GetFavoriteRes(ctx, userID)
	if err != nil {
		return &pb.FavoriteListResponse{}, err
	}
	resp := &pb.FavoriteListResponse{}
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	var rets []redis.AsyncRet
	for rows.Next() {
		var videoRes pb2.Video
		var authorName string
		videoRes.Author = &pb3.User{}
		err := rows.Scan(&videoRes.Id, &videoRes.Author.Id, &videoRes.PlayUrl,
			&videoRes.CoverUrl, &videoRes.Title, &authorName)
		if err != nil {
			return &pb.FavoriteListResponse{}, err
		}
		videoRes.IsFavorite = true
		//videoRes.FavoriteCount = dal.GetFavoriteNumRedis(ctx, videoRes.Id)
		//videoRes.CommentCount = dal.GetCommentNumRedis(ctx, videoRes.Id)
		ret, _ := conn.AsyncDo("ZSCORE", "favoriteSet", dal.ID2FavoriteKey(videoRes.Id))
		rets = append(rets, ret)
		ret, _ = conn.AsyncDo("ZSCORE", "commentSet", dal.ID2CommentKey(videoRes.Id))
		rets = append(rets, ret)
		videoRes.Author = &pb3.User{
			Id:   videoRes.Author.Id,
			Name: authorName,

			WorkCount: call_video.GetTotalWorkCount(ctx, videoRes.Author.Id),
		}
		videoRes.Author.TotalFavorited, videoRes.Author.FavoriteCount =
			GetTotalFavorited(ctx, videoRes.Author.Id), GetUserFavoriteVideoNum(ctx, videoRes.Author.Id)
		videoRes.Author.FollowCount, videoRes.Author.FollowerCount, videoRes.Author.IsFollow =
			call_user_follow.GetUserRelationInfo(ctx, userID, videoRes.Author.Id)
		resp.VideoList = append(resp.VideoList, &videoRes)
	}
	for i := 0; i > len(rets); i += 2 {
		out, _ := rets[i].Get()
		resp.VideoList[i].FavoriteCount = out.(int64)
		out, _ = rets[i+1].Get()
		resp.VideoList[i].CommentCount, _ = out.(int64)
	}
	return resp, err
}
func SetFavoriteNum(ctx context.Context, videoId, num int64) {
	dal.SetFavoriteNumRedis(ctx, videoId, num)
}
func GetVideoFavoritedNum(ctx context.Context, videoId int64) int64 {
	return dal.GetFavoriteNum(ctx, videoId)
}
func IsFavorite(ctx context.Context, userId, videdId int64) bool {
	return dal.IsFavorite(ctx, userId, videdId)
}
func GetUserFavoriteVideoNum(ctx context.Context, userId int64) int64 {
	return dal.GetUserFavoriteNum(ctx, userId)
}
func GetTotalFavorited(ctx context.Context, userId int64) (count int64) {
	videoIDList := call_video.GetVideoIDListOfUser(ctx, userId)
	return dal.GetTotalFavoritedRedis(ctx, videoIDList)
}

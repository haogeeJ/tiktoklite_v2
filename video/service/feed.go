package service

import (
	"TikTokLite_v2/common/redis_tracing"
	"TikTokLite_v2/video/dal"
	"TikTokLite_v2/video/dal/redb"
	"TikTokLite_v2/video/pb"
	"TikTokLite_v2/video/remote_call/call_fav_com"
	"TikTokLite_v2/video/remote_call/call_user_follow"
	"context"
	"fmt"
	"github.com/gistao/RedisGo-Async/redis"
	"log"
	"strconv"
	"time"
)

func BuildVideo(ctx context.Context, userID int64, _video dal.Video) pb.Video {
	var video pb.Video
	videoID := int64(_video.ID)

	video.Id = videoID
	resp, err := call_user_follow.GetUser(ctx, userID, _video.AuthorId)
	if err != nil {
		log.Fatal(err)
	}
	video.Author = resp.User
	video.Title = _video.Title
	video.PlayUrl = _video.PlayUrl
	video.CoverUrl = _video.CoverUrl
	video.FavoriteCount,
		video.CommentCount,
		video.IsFavorite = call_fav_com.GetFavoriteAndCommentInfo(ctx, userID, videoID)
	return video
}
func GetUserFeed(ctx context.Context, latestTime time.Time, userId int64) (*pb.FeedResponse, error) {
	//latestTime = latestTime.Add(5 * time.Minute)
	videoIDs, _ := dal.GetUserFeedRedis(ctx, latestTime, userId)
	//conn := model.RedisCache.Conn()
	//defer conn.Close()
	resp := &pb.FeedResponse{}

	var v dal.Video
	//log.Println("videoID:", videoIDs)
	for i := 0; i < len(videoIDs); i += 2 {
		var video pb.Video
		id := videoIDs[i]
		var err error
		v, err = dal.GetVideoByID(ctx, id)
		if err != nil {
			log.Println("get video by id failed:", err.Error())
			return nil, err
		}
		video = BuildVideo(ctx, userId, v)
		//log.Println(video)
		resp.VideoList = append(resp.VideoList, &video)
	}
	//log.Println(videos)
	resp.NextTime = time.Now().UnixMilli()
	if len(resp.VideoList) > 0 && len(resp.VideoList) == dal.FeedSize {
		resp.NextTime = videoIDs[len(videoIDs)-1]
	}
	//log.Println(latestTime.UnixMilli(), nextTime, videoIDs, "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"+
	//	"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	return resp, nil
}

func UserFeedInit(ctx context.Context, userID int64) {
	follows, err := call_user_follow.GetFollowListID(ctx, userID)
	if err != nil {
		log.Println("get followlist failed:", err.Error())
		return
	}
	conn := redb.RedisCache.AsynConn()
	defer conn.Close()
	userFeedKey := fmt.Sprintf("%s:%s", strconv.FormatInt(userID, 10), "userfeed")
	for _, id := range follows {
		authorkey := fmt.Sprintf("%s:%s", strconv.FormatInt(id, 10), "authorfeed")
		vals, _ := redis.Values(conn.Do("ZREVRANGEBYSCORE", authorkey, time.Now().UnixMilli(), 0, "withscores", "limit", 0, 10))
		for i := 0; i < len(vals); i += 2 {
			k, _ := redis.Int64(vals[i], nil)
			v, _ := redis.Int64(vals[i+1], nil)
			_, err = conn.AsyncDo("ZADD", userFeedKey, v, k)
			if err != nil {
				log.Println("userfeed set failed:", err.Error())
			}
		}
	}
	hots := dal.PullHotFeed(ctx, 20)
	for i := 0; i < len(hots); i++ {
		createTime := dal.GetVideoCreateTime(ctx, hots[i])
		_, _ = conn.AsyncDo("ZADD", userFeedKey, createTime, hots[i])
	}
}

func AuthorFeedPushToNewFollower(ctx context.Context, authorID, followerID int64) {
	//conn := model.RedisCache.Conn()
	conn := redb.RedisCache.AsynConn()
	defer conn.Close()
	authorFeedKey := fmt.Sprintf("%s:%s", strconv.FormatInt(authorID, 10), "authorfeed")
	videos, err := redis.Int64s(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZREVRANGEBYSCORE", authorFeedKey, "+inf", "-inf", "withscores", "limit", 0, 10))
	if err != nil {
		log.Println("get authorfeed error:", err.Error())
		return
	}
	userFeedKey := fmt.Sprintf("%s:%s", strconv.FormatInt(followerID, 10), "userfeed")
	for i := 0; i < len(videos); i += 2 {
		//conn.Send("ZADD", userFeedKey, videos[i+1], videos[i])
		_, err = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", userFeedKey, videos[i+1], videos[i])
	}
	//redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "")
}

func UpdateUnLoginFeed(ctx context.Context) {
	conn := redb.RedisCache.AsynConn()
	defer conn.Close()
	userFeedKey := "-1:userfeed"
	hots := dal.PullHotFeed(ctx, 20)
	for i := 0; i < len(hots); i++ {
		createTime := dal.GetVideoCreateTime(ctx, hots[i])
		_, _ = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", userFeedKey, createTime, hots[i])
	}
}

// BuildHotFeed 重建HotFeed
func BuildHotFeed(ctx context.Context) {
	resp, err := call_fav_com.GetHotFeed(ctx)
	if err != nil {
		log.Println("remote_call.GetHotFeed from favorite_comment error:", err)
		return
	}
	var h dal.HotCounter
	//h = make([]dal.HotCounter, len(resp.HotCounts))
	for _, hh := range resp.HotCounts {
		h.Vid, h.Favorite, h.Comment = hh.Vid, int(hh.FavoriteNum), int(hh.CommentNum)
		h.Time = dal.GetVideoCreateTime(ctx, h.Vid)
		dal.InsertHotFeed(ctx, h.Vid, h.ToScore())
	}
}
func CheckAliveUserAndPushHotFeed(ctx context.Context) {
	dal.CheckAliveUserAndPushHotFeed(ctx)
}

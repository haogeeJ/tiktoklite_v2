package dal

import (
	"TikTokLite_v2/common/redis_tracing"
	"TikTokLite_v2/video/dal/db"
	"TikTokLite_v2/video/dal/redb"
	"context"
	"fmt"
	"github.com/gistao/RedisGo-Async/redis"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

const (
	FeedSize  = 30
	aliveTime = time.Hour * 24
)

type Video struct {
	gorm.Model
	AuthorId int64  `gorm:"index:idx_author_id"`
	Title    string `gorm:"type:varchar(255)" ,json:"title"`
	PlayUrl  string `gorm:"type:varchar(255)" ,json:"play_url"`
	CoverUrl string `gorm:"type:varchar(255)" ,json:"cover_url"`
}

func (v *Video) Create(ctx context.Context) error {
	session := db.DB.WithContext(ctx)
	err := session.Create(&v).Error
	if err != nil {
		return err
	}

	return nil
}

func GetVideosByUserId(ctx context.Context, userId int64) ([]Video, error) {
	session := db.DB.WithContext(ctx)
	var videos []Video
	query := session.Where("author_id = ?", userId).Find(&videos)
	return videos, query.Error
}

func GetVideosByLatestTime(ctx context.Context, latestTime time.Time) ([]Video, error) {
	session := db.DB.WithContext(ctx)
	var videos []Video
	query := session.Order("created_at desc").Where("created_at > ?", latestTime).Limit(FeedSize).Find(&videos)
	return videos, query.Error
}

func GetTheLatestNVideos(ctx context.Context) ([]Video, error) {
	session := db.DB.WithContext(ctx)
	var videos []Video
	query := session.Order("created_at desc").Limit(FeedSize).Find(&videos)
	return videos, query.Error
}

func GetLatestVideo(ctx context.Context) (Video, error) {
	session := db.DB.WithContext(ctx)
	var video Video
	query := session.Last(&video)
	return video, query.Error
}

func GetVideoCreateTime(ctx context.Context, videoID int64) int64 {
	session := db.DB.WithContext(ctx)
	var t time.Time
	session.Model(&Video{}).Where("id = ?", videoID).Select("created_at").Scan(&t)
	return t.UnixMilli()
}

// Authorfeed增加新的视频
func InsertAuthorFeed(ctx context.Context, userID, videoID, now int64) (err error) {
	authorFeedKey := fmt.Sprintf("%s:%s", strconv.FormatInt(userID, 10), "authorfeed")
	conn := redb.RedisCache.Conn()
	defer conn.Close()
	_, err = redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZADD", authorFeedKey, now, videoID)
	return
}
func PushNewVideoToActiveUsersFeed(ctx context.Context, followers []int64, userID, videoID, now int64) (err error) {
	var loginTimeKey, userFeedKey, ids string
	var loginTime int64
	conn := redb.RedisCache.AsynConn()
	defer conn.Close()
	for i := 0; i < len(followers); i++ {
		ids = strconv.FormatInt(followers[i], 10)
		loginTimeKey = fmt.Sprintf("%s", ids)
		loginTime, err = redis.Int64(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "HGET", "aliveUser", loginTimeKey))
		if err != nil {
			if err == redis.ErrNil {
				continue
			}
			return err
		}
		//检查登录是否超时
		if time.UnixMilli(loginTime).Add(aliveTime).After(time.Now()) {
			userFeedKey = fmt.Sprintf("%s:%s", ids, "userfeed")
			_, err = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", userFeedKey, now, videoID)
			if err != nil {
				log.Println("userFeed Push failed:", err.Error())
			}
		} else {
			_, err = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", userFeedKey, now, videoID)
		}
	}
	userFeedKey = fmt.Sprintf("%s:%s", strconv.FormatInt(userID, 10), "userfeed")
	_, err = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", userFeedKey, now, videoID)
	return nil
}
func GetUserFeedRedis(ctx context.Context, latestTime time.Time, userId int64) ([]int64, error) {
	id := strconv.FormatInt(userId, 10)
	key := fmt.Sprintf("%s:%s", id, "userfeed")
	conn := redb.RedisCache.AsynConn()
	defer conn.Close()
	var err error
	var offset, timeStamp int64
	//存储对应用户feed流的偏移量
	userOffset := fmt.Sprintf("%s:%s", id, "offset")
	offset, err = redis.Int64(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "get", userOffset))
	if err == redis.ErrNil {
		offset = 0
	}
	//存储对应用户feed流的起始时间
	userFeedTimeStamp := fmt.Sprintf("%s:%s", id, "feedtimestamp")
	timeStamp, err = redis.Int64(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "GET", userFeedTimeStamp))
	if err == redis.ErrNil || offset == 0 {
		timeStamp = time.Now().UnixMilli()
		_, _ = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "SET", userFeedTimeStamp, timeStamp)
	}
	var vals []int64
	vals, err = redis.Int64s(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZREVRANGEBYSCORE", key, timeStamp, 0, "withscores", "limit", offset, FeedSize))
	offset += FeedSize
	//意味着feed流已查询到底
	if len(vals) < FeedSize*2 {
		offset = 0
		//如果没视频，则offset置0再拉取一遍
		if len(vals) == 0 {
			timeStamp = time.Now().UnixMilli()
			vals, err = redis.Int64s(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZREVRANGEBYSCORE", key, timeStamp, 0, "withscores", "limit", offset, FeedSize))
			offset += FeedSize
			if len(vals) < FeedSize*2 {
				offset = 0
			}
		}
	}
	_, _ = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "SET", userOffset, offset)
	return vals, nil
}

func GetVideoByID(ctx context.Context, videoID int64) (Video, error) {
	var video Video
	session := db.DB.WithContext(ctx)
	err := session.Where("id=?", videoID).First(&video).Error
	return video, err
}

func GetTotalWorkCount(ctx context.Context, userID int64) (count int64) {
	session := db.DB.WithContext(ctx)
	session.Model(&Video{}).Where("author_id = ?", userID).Count(&count)
	return
}

func GetVideoIDsByUser(ctx context.Context, userID int64) (ids []int64) {
	session := db.DB.WithContext(ctx)
	session.Model(&Video{}).Where("author_id = ?", userID).Select("id").Scan(&ids)
	return
}

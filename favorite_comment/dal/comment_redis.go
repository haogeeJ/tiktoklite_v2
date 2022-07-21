package dal

import (
	"TikTokLite_v2/common/redis_tracing"
	"TikTokLite_v2/favorite_comment/dal/redb"
	"TikTokLite_v2/video/pb"
	"context"
	"github.com/gistao/RedisGo-Async/redis"
	"log"
	"strconv"
)

//在redis中保存的hash结构，里面对应个视频的评论数
const CommentSet = "commentSet"

// GetCommentNumRedis 优先从redis里获取评论数
func GetCommentNumRedis(ctx context.Context, videoID int64) (count int64) {
	conn := redb.RedisCache.Conn()
	defer func() {
		_ = conn.Close()
	}()
	commentKey := ID2CommentKey(videoID)
	num, err := redis.Int64(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZSCORE", CommentSet, commentKey))
	//num, err := redis.Int64(conn.Do("ZSCORE", CommentSet, commentKey))
	if err != nil {
		count = GetCommentNum(ctx, videoID)
		SetCommentNumRedis(ctx, videoID, count)
		return
	}
	return num
}
func GetCommentsNumRedis(ctx context.Context, videos []*pb.Video) {
	conn := redb.RedisCache.Conn()
	defer func() {
		_ = conn.Close()
	}()
	for _, v := range videos {
		commentKey := ID2CommentKey(v.Id)
		conn.Send("ZSCORE", CommentSet, commentKey)
	}
	conn.Flush()
	for _, v := range videos {
		v.CommentCount, _ = redis.Int64(conn.Receive())
	}
}

// ID2CommentKey videoID转CommentKey
func ID2CommentKey(videoID int64) string {
	return "comment:" + strconv.FormatInt(videoID, 10)
}

func CommentKey2ID(key string) int64 {
	res, _ := strconv.Atoi(key[8:])
	return int64(res)
}

func SetCommentNumRedis(ctx context.Context, videoID int64, num int64) {
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	commentKey := ID2CommentKey(videoID)
	_, err := redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", CommentSet, num, commentKey)
	//_, err := conn.AsyncDo("ZADD", CommentSet, num, commentKey)
	if err != nil {
		log.Print("err in SetCommentNumRedis:", err)
		return
	}
}

func IncrCommentRedis(ctx context.Context, videoID int64) {
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	favoriteKey := ID2CommentKey(videoID)
	_, err := redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZINCRBY", CommentSet, 1, favoriteKey)
	//_, err := conn.AsyncDo("ZINCRBY", CommentSet, 1, favoriteKey)
	if err != nil {
		log.Print("err in IncrCommentRedis:", err)
		return
	}
}

func DecrCommentRedis(ctx context.Context, videoID int64) {
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	favoriteKey := ID2CommentKey(videoID)
	_, err := redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZINCRBY", CommentSet, -1, favoriteKey)
	//_, err := conn.AsyncDo("ZINCRBY", CommentSet, -1, favoriteKey)
	if err != nil {
		log.Print("err in DecrCommentRedis:", err)
		return
	}
}

func GetTopComment(ctx context.Context, n int) (top map[int64]int64) {
	conn := redb.RedisCache.Conn()
	defer func() {
		_ = conn.Close()
	}()
	values, err := redis.Values(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZREVRANGE", CommentSet, 0, n, "WITHSCORES"))
	//values, err := redis.Values(conn.Do("ZREVRANGE", CommentSet, 0, n, "WITHSCORES"))
	if err != nil {
		log.Println("err in GetTopComment:", err)
		return nil
	}
	top = make(map[int64]int64)
	for i := 0; i < len(values); i += 2 {
		key, _ := redis.String(values[i], nil)
		v, _ := redis.Int64(values[i+1], nil)
		if CommentKey2ID(key) == 0 || v == 0 {
			continue
		}
		top[CommentKey2ID(key)] = v
	}
	return top
}

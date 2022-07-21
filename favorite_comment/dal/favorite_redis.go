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

//在redis中保存的hash结构，成员是每个视频对应的点赞数
const FavoriteSortedSet = "favoriteSet"

//GetFavoriteNumRedis 如果redis中hash结构FavoriteSortedSet里查不到对应video，则查数据库，然后在redis中保存了视频被点赞次数。
func GetFavoriteNumRedis(ctx context.Context, videoID int64) (count int64) {
	conn := redb.RedisCache.Conn()
	defer func() {
		_ = conn.Close()
	}()
	favoriteKey := ID2FavoriteKey(videoID)
	num, err := redis.Int64(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZSCORE", FavoriteSortedSet, favoriteKey))
	if err != nil {
		count = GetFavoriteNum(ctx, videoID)
		SetFavoriteNumRedis(ctx, videoID, count)
		return
	}
	return num
}
func GetFavoritesNumRedis(ctx context.Context, videos []*pb.Video) {
	conn := redb.RedisCache.Conn()
	defer func() {
		_ = conn.Close()
	}()
	for _, v := range videos {
		favoriteKey := ID2FavoriteKey(v.Id)
		conn.Send("ZSCORE", FavoriteSortedSet, favoriteKey)

	}
	conn.Flush()
	for _, v := range videos {
		v.FavoriteCount, _ = redis.Int64(conn.Receive())
	}
}
func ID2FavoriteKey(videoID int64) string {
	return "favorite:" + strconv.FormatInt(videoID, 10)
}

func FavoriteKey2ID(key string) int64 {
	res, _ := strconv.Atoi(key[9:])
	return int64(res)
}

//SetFavoriteNumRedis 设置视频被点赞次数到FavoriteSortedSet中
func SetFavoriteNumRedis(ctx context.Context, videoID int64, num int64) {
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	favoriteKey := ID2FavoriteKey(videoID)
	_, err := redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", FavoriteSortedSet, num, favoriteKey)
	//_, err := conn.AsyncDo("ZADD", FavoriteSortedSet, num, favoriteKey)
	if err != nil {
		log.Print("err in SetFavoriteNumRedis:", err)
		return
	}
}

//IncrFavoriteRedis 视频每次被点赞则自增1
func IncrFavoriteRedis(ctx context.Context, videoID int64) {
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	favoriteKey := ID2FavoriteKey(videoID)
	_, err := redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZINCRBY", FavoriteSortedSet, 1, favoriteKey)
	//_, err := conn.AsyncDo("ZINCRBY", FavoriteSortedSet, 1, favoriteKey)
	if err != nil {
		log.Print("err in IncrFavoriteRedis:", err)
		return
	}
}

//DecrFavoriteRedis 取消点赞，自减1
func DecrFavoriteRedis(ctx context.Context, videoID int64) {
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	favoriteKey := ID2FavoriteKey(videoID)
	_, err := redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZINCRBY", FavoriteSortedSet, -1, favoriteKey)
	//_, err := conn.AsyncDo("ZINCRBY", FavoriteSortedSet, -1, favoriteKey)
	if err != nil {
		log.Print("err in DecrFavoriteRedis:", err)
		return
	}
}

func GetTopFavorite(ctx context.Context, n int) (top map[int64]int64) {
	conn := redb.RedisCache.Conn()
	defer func() {
		_ = conn.Close()
	}()
	values, err := redis.Values(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZREVRANGE", FavoriteSortedSet, 0, n, "WITHSCORES"))
	//values, err := redis.Values(conn.Do("ZREVRANGE", FavoriteSortedSet, 0, n, "WITHSCORES"))
	if err != nil {
		log.Println("err in GetTopFavorite:", err)
		return nil
	}
	top = make(map[int64]int64)
	for i := 0; i < len(values); i += 2 {
		key, _ := redis.String(values[i], nil)
		v, _ := redis.Int64(values[i+1], nil)
		if FavoriteKey2ID(key) == 0 {
			continue
		}
		top[FavoriteKey2ID(key)] = v
	}
	return top
}

//获取用户被点赞的总数
func GetTotalFavoritedRedis(ctx context.Context, videoIDList []int64) (count int64) {
	if len(videoIDList) == 0 {
		return 0
	}
	for _, val := range videoIDList {
		count += GetFavoriteNumRedis(ctx, val)
	}
	return count
}

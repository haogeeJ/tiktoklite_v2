package dal

import (
	"TikTokLite_v2/common/redis_tracing"
	"TikTokLite_v2/video/dal/redb"
	"context"
	"fmt"
	"github.com/gistao/RedisGo-Async/redis"
	"log"
	"time"
)

const HotFeedKey = "hot"

// HotCounter 暂时设定top20
type HotCounter struct {
	Vid      int64
	Favorite int
	Comment  int
	Time     int64
}

func InsertHotFeed(ctx context.Context, vid int64, score int) {
	conn := redb.RedisCache.AsynConn()
	defer func() {
		_ = conn.Close()
	}()
	_, err := redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", HotFeedKey, score, vid)
	//_, err := conn.AsyncDo("ZADD", HotFeedKey, score, vid)
	if err != nil {
		log.Println("err in InsertHotFeed:", err)
	}
}

func PullHotFeed(ctx context.Context, n int) []int64 {
	conn := redb.RedisCache.Conn()
	defer func() {
		_ = conn.Close()
	}()
	res, err := redis.Int64s(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "ZREVRANGEBYSCORE", HotFeedKey, "+inf", "-inf"))
	//res, err := redis.Int64s(conn.Do("ZREVRANGEBYSCORE", HotFeedKey, "+inf", "-inf"))
	if err != nil {
		fmt.Println("err in PullHotFeed:", err)
		return nil
	}
	if n > len(res) {
		n = len(res)
	}
	return res[:n]
}

func (h *HotCounter) ToScore() int {
	return h.Favorite + (h.Comment * 2)
}

func CheckAliveUserAndPushHotFeed(ctx context.Context) {
	conn := redb.RedisCache.AsynConn()
	defer conn.Close()
	vals, err := redis.Int64Map(redis_tracing.SyncDoAndTracing(ctx, conn.Do, "HGETALL", "aliveUser"))
	if err != nil {
		log.Println("push hotfeed error:", err.Error())
		return
	}
	hots := PullHotFeed(ctx, 20)
	var userFeedKey string
	for k, v := range vals {
		if time.UnixMilli(v).Add(aliveTime).After(time.Now()) {
			userFeedKey = fmt.Sprintf("%s:%s", k, "userfeed")
			for i := 0; i < len(hots); i++ {
				createTime := GetVideoCreateTime(ctx, hots[i])
				_, _ = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "ZADD", userFeedKey, createTime, hots[i])
			}
		} else {
			_, _ = redis_tracing.AsyncDoAndTracing(ctx, conn.AsyncDo, "HDEL", "aliveUser", k)
		}
	}
}

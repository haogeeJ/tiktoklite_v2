package video

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/controller/middleware"
	"TikTokLite_v2/controller/remote_call/call_video"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed 返回限制时间后发布的n个视频，如果限制时间后没有新的投稿，就按照最新的顺序返回n个视频
func Feed(c *gin.Context) {
	//resp := pb.FeedResponse{}
	//返回的信息还得按照用户是否登陆再去判断是否查询is_follow，is_favorite
	_, isLogin := c.GetQuery("token")
	var userId int64 = -1
	if isLogin {
		middleware.ValidDataTokenMiddleWare(c)
		ifUserId, _ := c.Get("user_id")
		userId = ifUserId.(int64)
	}
	latestTime := time.Now().UnixMilli()
	strLatestTime, exist := c.GetQuery("latest_time")
	if exist {
		latestTime, _ = strconv.ParseInt(strLatestTime, 10, 64)
	}
	resp, err := call_video.GetUserFeed(c.Request.Context(), latestTime, userId)
	if resp != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
	}
	c.JSON(http.StatusOK, resp)
}

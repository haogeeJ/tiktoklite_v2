package favorite_comment

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/controller/remote_call/call_fav_com"
	"TikTokLite_v2/favorite_comment/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	videoIDQuery, _ := c.GetQuery("video_id")
	actionTypeQuery, _ := c.GetQuery("action_type")
	userIDToken, _ := c.Get("user_id")
	userID := userIDToken.(int64)
	videoID, _ := strconv.Atoi(videoIDQuery)
	actionType, _ := strconv.Atoi(actionTypeQuery)
	var resp *pb.FavoriteActionResponse
	var err error
	if actionType == 1 {
		//点赞
		resp, err = call_fav_com.SetFavorite(c.Request.Context(), int64(videoID), userID)
		if resp != nil {
			resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "setFavorite failed",
			})
			return
		}

	} else if actionType == 2 {
		//取消点赞
		resp, err = call_fav_com.CancelFavorite(c.Request.Context(), int64(videoID), userID)
		if resp != nil {
			resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 2,
				"status_msg":  "cancelFavorite failed",
			})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 3,
			"status_msg":  "invalid action_type",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "ok",
	})
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	userIDToken, _ := c.GetQuery("user_id")
	userID, _ := strconv.ParseInt(userIDToken, 10, 64)
	resp, err := call_fav_com.GetFavoriteList(c.Request.Context(), userID)
	if resp != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "getFavoriteList failed",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

package follow

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/controller/remote_call/call_user_follow"
	pb2 "TikTokLite_v2/user_follow/follow/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(c *gin.Context) {

	var req pb2.RelationActionRequest
	req.Token, _ = c.GetQuery("token")
	toUserIDS, _ := c.GetQuery("to_user_id")
	actionTpS, _ := c.GetQuery("action_type")
	actionTp, _ := strconv.ParseInt(actionTpS, 10, 32)
	req.ActionType = int32(actionTp)
	req.ToUserId, _ = strconv.ParseInt(toUserIDS, 10, 64)
	userID, _ := c.Get("user_id")
	req.UserId = userID.(int64)
	resp, err := call_user_follow.RelationAction(c.Request.Context(), &req)
	if resp != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
	}
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "failed",
			})
			return
		}
		return
	}
	c.JSON(http.StatusOK, resp)

}

//BindFollowListRequest 从url中读取参数
func BindFollowListRequest(c *gin.Context) pb2.RelationFollowListRequest {
	var req pb2.RelationFollowListRequest
	req.Token, _ = c.GetQuery("token")
	userIDS, _ := c.GetQuery("user_id")
	req.UserId, _ = strconv.ParseInt(userIDS, 10, 64)
	return req
}

func FollowList(c *gin.Context) {
	req := BindFollowListRequest(c)
	resp, err := call_user_follow.GetFollowList(c.Request.Context(), &req)
	if resp != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
	}

	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "failed",
			})
			return
		}
		return
	}
	c.JSON(http.StatusOK, resp)
}

func FollowerList(c *gin.Context) {
	req := BindFollowListRequest(c)
	resp, err := call_user_follow.GetFollowerList(c.Request.Context(), &req)
	if resp != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "failed",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

package user

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/controller/common"
	"TikTokLite_v2/controller/remote_call/call_user_follow"
	"TikTokLite_v2/controller/remote_call/call_video"
	"TikTokLite_v2/user_follow/user/pb"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {

	var req pb.UserLoginOrRegisterRequest
	req.Name, _ = c.GetQuery("username")
	password, _ := c.GetQuery("password")
	has := md5.Sum([]byte(password))
	req.Password = fmt.Sprintf("%X", has)
	resp, err := call_user_follow.UserLogin(c.Request.Context(), &req)
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
	//初始化feed流
	call_video.InitUserFeed(c.Request.Context(), resp.UserId)
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})
}

//Register 用户注册
func Register(c *gin.Context) {

	var req pb.UserLoginOrRegisterRequest
	req.Name, _ = c.GetQuery("username")
	password, _ := c.GetQuery("password")
	if len(password) > 20 {
		statusCode, statusMsg := service.BuildResponse(errors.New("password is too long,limit <= 20"))
		common.ErrResponse(c, statusCode, statusMsg)
	}
	has := md5.Sum([]byte(password))
	req.Password = fmt.Sprintf("%X", has)
	resp, err := call_user_follow.UserRegister(c.Request.Context(), &req)
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
	//初始化feed流
	call_video.InitUserFeed(c.Request.Context(), resp.UserId)
	c.JSON(http.StatusOK, resp)
}

//UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	idStr, _ := c.GetQuery("user_id")
	userId, _ := strconv.ParseInt(idStr, 10, 64)
	//token, _ := c.GetQuery("token")
	resp, err := call_user_follow.GetUser(c.Request.Context(), userId, userId)
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

package video

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/controller/remote_call/call_video"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Publish(c *gin.Context) {
	userId, _ := c.Get("user_id")
	title := c.PostForm("title")
	//data, err := c.FormFile("data")

	form, err := c.MultipartForm()
	f, _ := form.File["data"][0].Open()
	data, _ := ioutil.ReadAll(f)
	_ = f.Close()
	filename := form.File["data"][0].Filename
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "cannot parse multipartForm",
		})
		return
	}
	//截取封面，上传视频和封面并返回外链

	_, err = call_video.PublishVideo(context.Background(), data, userId.(int64), filename, title)
	if err != nil {
		log.Println("publish failed，err：", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  fmt.Sprintf("userID:%d,title:%s, uploaded successfully", userId, title),
	})
}

func PublishList(c *gin.Context) {
	//token 里面解析出来的为发出请求的用户id
	userId, _ := c.Get("user_id")
	//query中的user_id才是需要被查询的用户id
	strToUserId, _ := c.GetQuery("user_id")
	toUserId, _ := strconv.ParseInt(strToUserId, 10, 64)

	resp, err := call_video.GetVideoList(c.Request.Context(), userId.(int64), toUserId)
	resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
	c.JSON(http.StatusOK, resp)
}

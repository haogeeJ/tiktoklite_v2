package favorite_comment

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/controller/remote_call/call_fav_com"
	"TikTokLite_v2/favorite_comment/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentAction 评论操作-评论/删除评论
func CommentAction(c *gin.Context) {
	videoIDQuery, _ := c.GetQuery("video_id")
	actionTypeQuery, _ := c.GetQuery("action_type")
	commentTextQuery, _ := c.GetQuery("comment_text")
	commentIDQuery, _ := c.GetQuery("comment_id")
	userIDToken, _ := c.Get("user_id")

	videoID, _ := strconv.Atoi(videoIDQuery)
	commentID, _ := strconv.Atoi(commentIDQuery)
	actionType, _ := strconv.Atoi(actionTypeQuery)
	userID := userIDToken.(int64)
	//resComment := service.Comment{}
	res := pb.CommentActionResponse{}
	if actionType == 1 {
		//评论过滤器，检测敏感词
		var err error
		comment, _ := call_fav_com.CommentFilter(c.Request.Context(), commentTextQuery)
		res.Comment, err = call_fav_com.CreateComment(c.Request.Context(), int64(videoID), userID, comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 5,
				"status_msg":  "err in CreateComment",
			})
			return
		}
		userName, _ := c.Get("user_name")
		res.Comment.User.Name = userName.(string)
	} else if actionType == 2 {
		err := call_fav_com.DeleteComment(c.Request.Context(), userID, int64(commentID), int64(videoID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 6,
				"status_msg":  "err in DeleteComment",
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
	res.StatusCode, res.StatusMsg = service.BuildResponse(nil)
	c.JSON(http.StatusOK, res)
}

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	videoIDQuery, _ := c.GetQuery("video_id")
	videoID, _ := strconv.Atoi(videoIDQuery)
	userIDToken, _ := c.Get("user_id")
	userID := userIDToken.(int64)
	commentListResp, err := call_fav_com.GetCommentList(c.Request.Context(), int64(videoID), userID)
	commentListResp.StatusCode, commentListResp.StatusMsg = service.BuildResponse(err)
	if err != nil {

		c.JSON(http.StatusInternalServerError, commentListResp)
		return
	}
	c.JSON(http.StatusOK, commentListResp)
}

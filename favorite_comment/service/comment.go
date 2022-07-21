package service

import (
	"TikTokLite_v2/favorite_comment/dal"
	"TikTokLite_v2/favorite_comment/pb"
	"TikTokLite_v2/favorite_comment/remote_call/call_user_follow"
	"TikTokLite_v2/favorite_comment/remote_call/call_video"
	pb2 "TikTokLite_v2/user_follow/user/pb"
	"TikTokLite_v2/util"
	"context"
	"gorm.io/gorm"
	"time"
)

// GetCommentByJoin 改用联查的版本
func GetCommentByJoin(ctx context.Context, videoID int64, userID int64) (*pb.CommentListResponse, error) {
	rows, err := dal.GetCommentRes(ctx, videoID)
	if err != nil {
		return nil, err
	}
	resp := &pb.CommentListResponse{}
	for rows.Next() {
		var comment pb.Comment
		var createDate time.Time
		comment.User = &pb2.User{}
		err := rows.Scan(&comment.Id, &comment.Content, &createDate, &comment.User.Id, &comment.User.Name)
		if err != nil {
			return nil, err
		}
		comment.CreateDate = util.Time2String(createDate)
		comment.User.FollowerCount, comment.User.FollowCount, comment.User.IsFollow =
			call_user_follow.GetUserRelationInfo(ctx, userID, comment.User.Id)

		comment.User.TotalFavorited, comment.User.FavoriteCount =
			GetTotalFavorited(ctx, comment.User.Id), GetUserFavoriteVideoNum(ctx, comment.User.Id)
		comment.User.WorkCount = call_video.GetTotalWorkCount(ctx, comment.User.Id)
		resp.CommentList = append(resp.CommentList, &comment)
	}
	return resp, nil
}

func CreateComment(ctx context.Context, videoID, userID int64, text string) (dal.Comment, error) {
	c := dal.Comment{
		VideoID: videoID,
		UserID:  userID,
		Content: text,
	}
	return c, c.Create(ctx)
}

func DeleteComment(ctx context.Context, userID, commentID, videoID int64) error {
	c := dal.Comment{
		Model: gorm.Model{
			ID: uint(commentID),
		},
		UserID:  userID,
		VideoID: videoID,
	}
	return c.DeleteByUser(ctx)
}
func SetCommentNum(ctx context.Context, videoId, num int64) {
	dal.SetCommentNumRedis(ctx, videoId, num)
}

// CommentFilter 评论过滤器，过滤敏感词
func CommentFilter(commentMsg string) (string, bool) {
	return util.Filtration(commentMsg)
}
func GetCommentNum(ctx context.Context, videoId int64) int64 {
	return dal.GetCommentNum(ctx, videoId)
}

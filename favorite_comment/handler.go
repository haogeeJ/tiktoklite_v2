package main

import (
	"TikTokLite_v2/favorite_comment/pb"
	"TikTokLite_v2/favorite_comment/remote_call/call_user_follow"
	"TikTokLite_v2/favorite_comment/remote_call/call_video"
	"TikTokLite_v2/favorite_comment/service"
	pb2 "TikTokLite_v2/user_follow/user/pb"
	"TikTokLite_v2/util"
	"context"
	"errors"
)

type CommentService struct {
}

func (s *CommentService) GetCommentList(ctx context.Context, req *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	resp, err := service.GetCommentByJoin(ctx, req.VideoId, req.UserId)
	//span := opentracing.SpanFromContext(ctx)
	//jaectx := span.Context().(jaeger.SpanContext)
	//fmt.Println(jaectx.TraceID())
	if err != nil {
		return &pb.CommentListResponse{}, err
	}
	return resp, err
}
func (s *CommentService) CreateComment(ctx context.Context, req *pb.CommentActionRequest) (*pb.CommentActionResponse, error) {
	resp := &pb.CommentActionResponse{}
	comment, err := service.CreateComment(ctx, req.VideoId, req.UserId, req.CommentText)
	if err != nil {
		return nil, err
	}
	respComment := pb.Comment{}
	respComment.User = &pb2.User{}
	respComment.Id = int64(comment.ID)
	respComment.User.Id = comment.UserID
	respComment.Content = comment.Content
	respComment.CreateDate = util.Time2String(comment.CreatedAt)
	respComment.User.FollowerCount, respComment.User.FollowCount, respComment.User.IsFollow =
		call_user_follow.GetUserRelationInfo(ctx, req.UserId, respComment.User.Id)
	//暂时不支持，后续补上调用
	respComment.User.TotalFavorited, respComment.User.FavoriteCount =
		service.GetTotalFavorited(ctx, respComment.User.Id), service.GetUserFavoriteVideoNum(ctx, respComment.User.Id)
	respComment.User.WorkCount = call_video.GetTotalWorkCount(ctx, respComment.User.Id)
	resp.Comment = &respComment
	return resp, nil
}
func (s *CommentService) DeleteComment(ctx context.Context, req *pb.CommentActionRequest) (*pb.CommentActionResponse, error) {
	err := service.DeleteComment(ctx, req.UserId, req.CommentId, req.VideoId)
	if err != nil {
		return nil, err
	}
	return &pb.CommentActionResponse{}, nil
}
func (s *CommentService) CommentFilter(ctx context.Context, req *pb.CommentFilterRequest) (*pb.CommentFilterResponse, error) {
	newMsg, ok := service.CommentFilter(req.CommentMsg)
	if !ok {
		return &pb.CommentFilterResponse{}, errors.New("CommentFilter error")
	}
	return &pb.CommentFilterResponse{CommentMsg: newMsg}, nil
}
func (s *CommentService) SetCommentNum(ctx context.Context, req *pb.SetCommentNumRequest) (*pb.SetCommentNumResponse, error) {
	service.SetCommentNum(ctx, req.VideoId, req.Num)
	return &pb.SetCommentNumResponse{}, nil
}

type FavoriteService struct {
}

func (s *FavoriteService) SetFavorite(ctx context.Context, req *pb.FavoriteActionRequest) (*pb.FavoriteActionResponse, error) {
	err := service.SetFavorite(ctx, req.VideoId, req.UserId)
	if err != nil {
		return &pb.FavoriteActionResponse{}, err
	}
	return &pb.FavoriteActionResponse{}, nil
}
func (s *FavoriteService) CancelFavorite(ctx context.Context, req *pb.FavoriteActionRequest) (*pb.FavoriteActionResponse, error) {
	err := service.CancelFavorite(ctx, req.VideoId, req.UserId)
	if err != nil {
		return &pb.FavoriteActionResponse{}, err
	}
	return &pb.FavoriteActionResponse{}, nil
}
func (s *FavoriteService) GetFavoriteList(ctx context.Context, req *pb.FavoriteListRequest) (*pb.FavoriteListResponse, error) {
	resp, err := service.GetFavoriteList(ctx, req.UserId)
	return resp, err
}
func (s *FavoriteService) SetFavoriteNum(ctx context.Context, req *pb.SetFavoriteNumRequest) (*pb.SetFavoriteNumResponse, error) {
	service.SetFavoriteNum(ctx, req.VideoId, req.Num)
	return &pb.SetFavoriteNumResponse{}, nil
}
func (s *FavoriteService) UserFavoriteInfo(ctx context.Context, req *pb.UserFavoriteInfoRequest) (*pb.UserFavoriteInfoResponse, error) {
	resp := &pb.UserFavoriteInfoResponse{}
	resp.FavoriteCount = service.GetUserFavoriteVideoNum(ctx, req.UserId)
	resp.TotalFavorited = service.GetTotalFavorited(ctx, req.UserId)
	return resp, nil
}

type HotFeedService struct {
}

func (s *HotFeedService) GetHotFeed(ctx context.Context, req *pb.HotFeedRequest) (*pb.HotFeedResponse, error) {
	resp, err := service.BuildHotFeed(ctx, req)
	return resp, err
}

type VideoInfoService struct {
}

func (s *VideoInfoService) GetVideoInfoAboutFavAndCom(ctx context.Context, req *pb.VideoInfoRequest) (*pb.VideoInfoResponse, error) {
	return service.GetVideoInfoAboutFavAndCom(ctx, req)
}

package main

import (
	"TikTokLite_v2/video/dal"
	"TikTokLite_v2/video/pb"
	"TikTokLite_v2/video/service"
	"context"
	"time"
)

type VideoService struct {
}

func (s *VideoService) PublishVideo(ctx context.Context, req *pb.PublishActionRequest) (*pb.PublishActionResponse, error) {
	err := service.PublishVideo(ctx, req.Data, req.Filename, req.UserId, dal.Video{AuthorId: req.UserId, Title: req.Title})
	return &pb.PublishActionResponse{}, err
}
func (s *VideoService) GetVideoList(ctx context.Context, req *pb.PublishListRequest) (*pb.PublishListResponse, error) {
	return service.GetVideoList(ctx, req.UserId, req.ToUserId)
}
func (s *VideoService) GetUserFeed(ctx context.Context, req *pb.FeedRequest) (*pb.FeedResponse, error) {
	return service.GetUserFeed(ctx, time.UnixMilli(req.LatestTime), req.UserId)
}
func (s *VideoService) GetVideoIDListOfUser(ctx context.Context, req *pb.GetVideoIDListOfUserRequest) (*pb.GetVideoIDListOfUserResponse, error) {
	return service.GetVideoIDsByUser(ctx, req.UserId)
}
func (s *VideoService) GetTotalWorkCount(ctx context.Context, req *pb.GetTotalWorkCountRequest) (*pb.GetTotalWorkCountResponse, error) {
	return service.GetTotalWorkCount(ctx, req.UserId)
}
func (s *VideoService) InitUserFeed(ctx context.Context, req *pb.InitUserFeedRequest) (*pb.InitUserFeedResponse, error) {
	service.UserFeedInit(ctx, req.UserId)
	return &pb.InitUserFeedResponse{}, nil
}
func (s *VideoService) AuthorFeedPushToNewFollower(ctx context.Context, req *pb.AuthorFeedPushToNewFollowerRequest) (*pb.AuthorFeedPushToNewFollowerResponse, error) {
	//不关心返回值，采用直接返回
	service.AuthorFeedPushToNewFollower(ctx, req.AuthorId, req.FollowerId)
	return &pb.AuthorFeedPushToNewFollowerResponse{}, nil
}

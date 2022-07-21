package follow

import (
	"TikTokLite_v2/user_follow/follow/pb"
	"TikTokLite_v2/user_follow/follow/service"
	"context"
	"google.golang.org/grpc"
)

type Service struct {
	sv *service.FollowService
}

func (s *Service) RelationAction(ctx context.Context, req *pb.RelationActionRequest) (*pb.RelationActionResponse, error) {
	return s.sv.RelationAction(ctx, req)
}
func (s *Service) FollowList(ctx context.Context, req *pb.RelationFollowListRequest) (*pb.RelationFollowListResponse, error) {
	return s.sv.FollowList(ctx, req.UserId)
}
func (s *Service) FollowerList(ctx context.Context, req *pb.RelationFollowListRequest) (*pb.RelationFollowListResponse, error) {
	return s.sv.FollowerList(ctx, req.UserId)
}
func (s *Service) UserRelationInfo(ctx context.Context, req *pb.UserRelationInfoRequest) (*pb.UserRelationInfoResponse, error) {
	return s.sv.UserRelationInfo(ctx, req)
}
func (s *Service) FollowListID(ctx context.Context, req *pb.FollowListIDRequest) (*pb.FollowListIDResponse, error) {
	return s.sv.FollowListID(ctx, req.UserId)
}
func (s *Service) FollowerListID(ctx context.Context, req *pb.FollowListIDRequest) (*pb.FollowListIDResponse, error) {
	return s.sv.FollowerListID(ctx, req.UserId)
}
func RegisterService(grpcServer *grpc.Server) {
	pb.RegisterFollowServiceServer(grpcServer, &Service{service.NewFollowService()})
}

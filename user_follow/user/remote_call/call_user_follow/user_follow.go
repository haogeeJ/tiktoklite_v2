package call_user_follow

import (
	gtrace "TikTokLite_v2/common/grpc_jaeger"
	pb3 "TikTokLite_v2/user_follow/follow/pb"
	"TikTokLite_v2/user_follow/setting"
	pb2 "TikTokLite_v2/user_follow/user/pb"
	"context"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var conn *grpc.ClientConn

func Init() {
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)}

	var err error
	tracer := opentracing.GlobalTracer()
	if tracer != nil {
		dialOpts = append(dialOpts, gtrace.DialOption(tracer))
	}
	conn, err = grpc.Dial(

		setting.Conf.Consul.ApiHealthCheck.Targets[1],
		dialOpts...,
	)
	if err != nil {
		log.Fatal("dial user_follow-server error:", err)
	}
}
func GetUser(ctx context.Context, userId, toUserId int64) (*pb2.UserInfoResponse, error) {
	grpcClient := pb2.NewUserServiceClient(conn)
	req := pb2.UserInfoRequest{
		UserId:   userId,
		ToUserId: toUserId,
	}
	resp, err := grpcClient.UserInfo(ctx, &req)
	if err != nil {
		//log.Fatal("remote_call UserInfo error:", err)
		return resp, err
	}

	return resp, nil
}
func GetFollowListID(ctx context.Context, uid int64) ([]int64, error) {

	var err error
	grpcClient := pb3.NewFollowServiceClient(conn)
	req := pb3.FollowListIDRequest{
		UserId: uid,
	}
	var resp *pb3.FollowListIDResponse
	resp, err = grpcClient.FollowListID(ctx, &req)
	return resp.FollowList, err
}
func GetFollowerListID(ctx context.Context, uid int64) ([]int64, error) {
	var err error
	grpcClient := pb3.NewFollowServiceClient(conn)
	req := pb3.FollowListIDRequest{
		UserId: uid,
	}
	var resp *pb3.FollowListIDResponse
	resp, err = grpcClient.FollowerListID(ctx, &req)
	return resp.FollowList, err
}
func GetUserRelationInfo(ctx context.Context, userId, authorId int64) (followCount, followerCount int64, isFollow bool) {

	grpcClient := pb3.NewFollowServiceClient(conn)

	req := &pb3.UserRelationInfoRequest{
		UserId:   userId,
		AuthorId: authorId,
	}
	resp, err := grpcClient.UserRelationInfo(ctx, req)
	if err != nil {
		return -1, -1, false
	}
	return resp.FollowCount, resp.FollowerCount, resp.IsFollow
}
func GetFollowerList(ctx context.Context, req *pb3.RelationFollowListRequest) (*pb3.RelationFollowListResponse, error) {
	grpcClient := pb3.NewFollowServiceClient(conn)
	return grpcClient.FollowerList(ctx, req)
}
func GetFollowList(ctx context.Context, req *pb3.RelationFollowListRequest) (*pb3.RelationFollowListResponse, error) {
	grpcClient := pb3.NewFollowServiceClient(conn)
	return grpcClient.FollowList(ctx, req)
}

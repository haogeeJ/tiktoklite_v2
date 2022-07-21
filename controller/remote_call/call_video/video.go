package call_video

import (
	gtrace "TikTokLite_v2/common/grpc_jaeger"
	"TikTokLite_v2/controller/setting"
	"TikTokLite_v2/video/pb"
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
	tracer := opentracing.GlobalTracer()
	var err error

	if tracer != nil {
		dialOpts = append(dialOpts, gtrace.DialOption(tracer))
	}
	conn, err = grpc.Dial(
		setting.Conf.Consul.ApiHealthCheck.Targets[0],
		dialOpts...,
	)
	if err != nil {
		log.Fatal("dial video-server error:", err)
	}
}
func GetUserFeed(ctx context.Context, latestTime, userId int64) (*pb.FeedResponse, error) {
	var resp *pb.FeedResponse
	var err error

	grpcClient := pb.NewVideoServiceClient(conn)
	req := pb.FeedRequest{
		UserId:     userId,
		LatestTime: latestTime,
	}
	resp, err = grpcClient.GetUserFeed(ctx, &req)
	return resp, err
}
func PublishVideo(ctx context.Context, data []byte, userId int64, filename, title string) (*pb.PublishActionResponse, error) {
	var resp *pb.PublishActionResponse
	var err error

	grpcClient := pb.NewVideoServiceClient(conn)
	req := pb.PublishActionRequest{
		Filename: filename,
		Data:     data,
		UserId:   userId,
		Title:    title,
	}

	resp, err = grpcClient.PublishVideo(ctx, &req)
	return resp, err
}
func GetVideoList(ctx context.Context, userId, toUserId int64) (*pb.PublishListResponse, error) {
	var resp *pb.PublishListResponse
	var err error

	grpcClient := pb.NewVideoServiceClient(conn)
	req := pb.PublishListRequest{
		UserId:   userId,
		ToUserId: toUserId,
	}
	resp, err = grpcClient.GetVideoList(ctx, &req)
	return resp, err
}
func GetTotalWorkCount(ctx context.Context, userId int64) (count int64) {
	var resp *pb.GetTotalWorkCountResponse

	grpcClient := pb.NewVideoServiceClient(conn)
	req := pb.GetTotalWorkCountRequest{
		UserId: userId,
	}
	resp, _ = grpcClient.GetTotalWorkCount(ctx, &req)
	return resp.Count
}
func GetVideoIDListOfUser(ctx context.Context, userId int64) (videoIdList []int64) {
	var resp *pb.GetVideoIDListOfUserResponse
	grpcClient := pb.NewVideoServiceClient(conn)
	req := pb.GetVideoIDListOfUserRequest{
		UserId: userId,
	}
	resp, _ = grpcClient.GetVideoIDListOfUser(ctx, &req)
	return resp.VideoIdList
}
func InitUserFeed(ctx context.Context, userId int64) {
	grpcClient := pb.NewVideoServiceClient(conn)
	req := pb.InitUserFeedRequest{
		UserId: userId,
	}
	grpcClient.InitUserFeed(ctx, &req)
}
func AuthorFeedPushToNewFollower(ctx context.Context, authorId, followerId int64) {
	grpcClient := pb.NewVideoServiceClient(conn)
	req := pb.AuthorFeedPushToNewFollowerRequest{
		AuthorId:   authorId,
		FollowerId: followerId,
	}
	grpcClient.AuthorFeedPushToNewFollower(ctx, &req)
}

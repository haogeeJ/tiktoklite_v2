package call_fav_com

import (
	gtrace "TikTokLite_v2/common/grpc_jaeger"
	"TikTokLite_v2/controller/setting"
	"TikTokLite_v2/favorite_comment/pb"
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

		setting.Conf.Consul.ApiHealthCheck.Targets[2],
		dialOpts...,
	)
	if err != nil {
		log.Fatal("dial favorite_comment-server error:", err)
	}
}
func GetHotFeed(ctx context.Context) (*pb.HotFeedResponse, error) {
	grpcClient := pb.NewHotFeedClient(conn)
	req := pb.HotFeedRequest{}
	req.Num = 20
	resp, err := grpcClient.GetHotFeed(ctx, &req)
	return resp, err
}
func SetFavoriteNum(ctx context.Context, videoId, num int64) {
	grpcClient := pb.NewFavoriteServiceClient(conn)
	req := pb.SetFavoriteNumRequest{
		VideoId: videoId,
		Num:     num,
	}
	grpcClient.SetFavoriteNum(ctx, &req)
}
func SetCommentNum(ctx context.Context, videoId, num int64) {
	grpcClient := pb.NewCommentServiceClient(conn)
	req := pb.SetCommentNumRequest{
		VideoId: videoId,
		Num:     num,
	}
	grpcClient.SetCommentNum(ctx, &req)
}
func GetFavoriteAndCommentInfo(ctx context.Context, userId, videoId int64) (favoriteNum, commentNum int64, isFavorite bool) {
	var err error
	grpcClient := pb.NewVideoInfoClient(conn)
	req := pb.VideoInfoRequest{
		UserId:  userId,
		VideoId: videoId,
	}
	var resp *pb.VideoInfoResponse
	resp, err = grpcClient.GetVideoInfoAboutFavAndCom(ctx, &req)
	if err != nil {
		return 0, 0, false
	}
	return resp.FavoriteNum, resp.CommentNum, resp.IsFavorite
}
func CommentFilter(ctx context.Context, commentText string) (filterText string, ok bool) {
	var err error
	grpcClient := pb.NewCommentServiceClient(conn)
	req := pb.CommentFilterRequest{
		CommentMsg: commentText,
	}
	var resp *pb.CommentFilterResponse
	resp, err = grpcClient.CommentFilter(ctx, &req)
	if err != nil {
		return "", false
	}
	return resp.CommentMsg, true
}
func CreateComment(ctx context.Context, videoId, userId int64, commentText string) (*pb.Comment, error) {
	var err error
	grpcClient := pb.NewCommentServiceClient(conn)
	req := pb.CommentActionRequest{
		VideoId:     videoId,
		UserId:      userId,
		CommentText: commentText,
	}
	var resp *pb.CommentActionResponse
	resp, err = grpcClient.CreateComment(ctx, &req)
	return resp.Comment, err
}
func DeleteComment(ctx context.Context, userId, commentId, videoId int64) error {
	var err error
	grpcClient := pb.NewCommentServiceClient(conn)
	req := pb.CommentActionRequest{
		VideoId:   videoId,
		UserId:    userId,
		CommentId: commentId,
	}
	_, err = grpcClient.DeleteComment(ctx, &req)
	return err

}
func GetCommentList(ctx context.Context, videoId, userId int64) (*pb.CommentListResponse, error) {
	var err error
	grpcClient := pb.NewCommentServiceClient(conn)
	req := pb.CommentListRequest{
		VideoId: videoId,
		UserId:  userId,
	}
	var resp *pb.CommentListResponse
	resp, err = grpcClient.GetCommentList(ctx, &req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func SetFavorite(ctx context.Context, videoId, userId int64) (*pb.FavoriteActionResponse, error) {
	resp := &pb.FavoriteActionResponse{}
	var err error
	grpcClient := pb.NewFavoriteServiceClient(conn)
	req := pb.FavoriteActionRequest{VideoId: videoId, UserId: userId}
	resp, err = grpcClient.SetFavorite(ctx, &req)
	return resp, err
}
func CancelFavorite(ctx context.Context, videoId, userId int64) (*pb.FavoriteActionResponse, error) {
	resp := &pb.FavoriteActionResponse{}
	var err error
	grpcClient := pb.NewFavoriteServiceClient(conn)
	req := pb.FavoriteActionRequest{VideoId: videoId, UserId: userId}
	_, err = grpcClient.CancelFavorite(ctx, &req)
	return resp, err
}
func GetFavoriteList(ctx context.Context, userId int64) (*pb.FavoriteListResponse, error) {
	var err error
	grpcClient := pb.NewFavoriteServiceClient(conn)
	req := pb.FavoriteListRequest{
		UserId: userId,
	}
	var resp *pb.FavoriteListResponse
	resp, err = grpcClient.GetFavoriteList(ctx, &req)
	return resp, err
}
func UserFavoriteInfo(ctx context.Context, userId int64) (totalFavorited, favoriteCount int64, err error) {
	grpcClient := pb.NewFavoriteServiceClient(conn)
	req := pb.UserFavoriteInfoRequest{
		UserId: userId,
	}
	var resp *pb.UserFavoriteInfoResponse
	resp, err = grpcClient.UserFavoriteInfo(ctx, &req)
	return resp.TotalFavorited, resp.FavoriteCount, err
}

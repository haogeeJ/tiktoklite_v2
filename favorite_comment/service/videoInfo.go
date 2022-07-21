package service

import (
	"TikTokLite_v2/favorite_comment/dal"
	"TikTokLite_v2/favorite_comment/pb"
	"context"
)

func GetVideoInfoAboutFavAndCom(ctx context.Context, req *pb.VideoInfoRequest) (*pb.VideoInfoResponse, error) {
	resp := &pb.VideoInfoResponse{}
	resp.FavoriteNum = dal.GetFavoriteNum(ctx, req.VideoId)
	resp.CommentNum = dal.GetCommentNum(ctx, req.VideoId)
	resp.IsFavorite = dal.IsFavorite(ctx, req.UserId, req.VideoId)
	return resp, nil
}

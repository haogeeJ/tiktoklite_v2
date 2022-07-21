package service

import (
	"TikTokLite_v2/favorite_comment/dal"
	"TikTokLite_v2/favorite_comment/pb"
	"context"
)

// BuildHotFeed 每ns触发一次
func BuildHotFeed(ctx context.Context, req *pb.HotFeedRequest) (*pb.HotFeedResponse, error) {
	var set map[int64]bool
	//去重
	set = make(map[int64]bool)
	topf := dal.GetTopFavorite(ctx, 20)
	for key, _ := range topf {
		set[key] = true
	}
	topc := dal.GetTopComment(ctx, 20)
	for key, _ := range topc {
		set[key] = true
	}
	var hotFeed pb.HotFeedResponse
	for i, _ := range set {
		favoriteNum, ok := topf[i]
		if !ok || favoriteNum == 0 {
			favoriteNum = dal.GetFavoriteNumRedis(ctx, i)
		}
		commentNum, ok := topc[i]
		if !ok || commentNum == 0 {
			commentNum = dal.GetCommentNumRedis(ctx, i)
		}
		h := pb.HotCount{
			Vid:         i,
			FavoriteNum: favoriteNum,
			CommentNum:  commentNum,
		}
		hotFeed.HotCounts = append(hotFeed.HotCounts, &h)
	}
	return &hotFeed, nil
}

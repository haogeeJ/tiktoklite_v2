package service

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/user_follow/follow/dal"
	"TikTokLite_v2/user_follow/follow/pb"
	"TikTokLite_v2/user_follow/user/remote_call/call_video"
	service2 "TikTokLite_v2/user_follow/user/service"
	"context"
)

type FollowService struct {
	FollowRepository dal.IFollowRepository
}

func NewFollowService() *FollowService {
	return &FollowService{
		FollowRepository: dal.NewFollowManagerRepository(),
	}
}
func (s *FollowService) UserRelationInfo(ctx context.Context, req *pb.UserRelationInfoRequest) (resp *pb.UserRelationInfoResponse, err error) {
	resp = &pb.UserRelationInfoResponse{}
	resp.FollowerCount = s.FollowRepository.RedisFollowerCount(req.UserId)
	resp.FollowCount = s.FollowRepository.RedisFollowCount(req.UserId)
	resp.IsFollow = s.FollowRepository.RedisIsFollow(req.UserId, req.AuthorId)
	return
}
func (s *FollowService) RelationAction(ctx context.Context, req *pb.RelationActionRequest) (resp *pb.RelationActionResponse, err error) {
	if req.ActionType == 1 {
		err = s.FollowRepository.RedisInsert(req.ToUserId, req.UserId)
		call_video.AuthorFeedPushToNewFollower(ctx, req.ToUserId, req.UserId)
		//time.Sleep(time.Millisecond * 10)
	} else {
		err = s.FollowRepository.RedisDelete(req.ToUserId, req.UserId)
	}
	resp = new(pb.RelationActionResponse)
	resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
	return
}

func (s *FollowService) FollowList(ctx context.Context, uid int64) (resp *pb.RelationFollowListResponse, err error) {
	var follows []int64
	follows, err = s.FollowRepository.RedisGetFollowList(uid)
	if err != nil {
		return nil, err
	}
	resp = new(pb.RelationFollowListResponse)
	resp.UserList = service2.BuildUserList(ctx, uid, follows, s.FollowRepository)
	return
}

func (s *FollowService) FollowerList(ctx context.Context, uid int64) (resp *pb.RelationFollowListResponse, err error) {
	var followers []int64
	followers, err = s.FollowRepository.RedisGetFollowerList(uid)
	if err != nil {
		return nil, err
	}
	resp = new(pb.RelationFollowListResponse)
	resp.UserList = service2.BuildUserList(ctx, uid, followers, s.FollowRepository)
	return
}
func (s *FollowService) FollowListID(ctx context.Context, uid int64) (resp *pb.FollowListIDResponse, err error) {
	resp = &pb.FollowListIDResponse{}
	resp.FollowList, err = s.FollowRepository.RedisGetFollowList(uid)
	return resp, err
}
func (s *FollowService) FollowerListID(ctx context.Context, uid int64) (resp *pb.FollowListIDResponse, err error) {
	resp = &pb.FollowListIDResponse{}
	resp.FollowList, err = s.FollowRepository.RedisGetFollowerList(uid)
	//fmt.Println("!!!!!!!????????", resp)
	return resp, err
}

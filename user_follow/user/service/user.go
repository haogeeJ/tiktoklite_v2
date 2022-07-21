package service

import (
	"TikTokLite_v2/common/service"
	dal2 "TikTokLite_v2/user_follow/follow/dal"
	"TikTokLite_v2/user_follow/user/dal"
	"TikTokLite_v2/user_follow/user/dal/redb"
	"TikTokLite_v2/user_follow/user/pb"
	"TikTokLite_v2/user_follow/user/remote_call/call_fav_com"
	"TikTokLite_v2/user_follow/user/remote_call/call_video"
	"TikTokLite_v2/util"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

//UserService 实现service层用户信息操作的接口方法
type UserService struct {
	//数据库操作接口
	UserRepository dal.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: dal.NewUserManagerRepository(),
	}
}

//创建新用户
func (s *UserService) createUser(ctx context.Context, u *dal.User) error {
	//insert一个实例
	return s.UserRepository.Insert(ctx, u)
}

//根据id获取用户信息
func (s *UserService) getUserById(ctx context.Context, u *dal.User, id uint) (statusCode int32, statusMsg string, err error) {
	err = s.UserRepository.GetById(ctx, u, id)
	statusCode, statusMsg = service.BuildResponse(err)
	return
}

//UpdateUser 更新用户基本信息
func (s *UserService) UpdateUser(ctx context.Context, u *dal.User) (statusCode int32, statusMsg string, err error) {
	err = s.UserRepository.Update(ctx, u)
	statusCode, statusMsg = service.BuildResponse(err)
	return
}

//检查用户名是否存在
func (s *UserService) userIsExists(ctx context.Context, username string) error {
	return s.UserRepository.IsExists(ctx, username)
}

//检查用户名密码
func (s *UserService) checkUser(ctx context.Context, req *pb.UserLoginOrRegisterRequest) (u *dal.User, err error) {
	u = &dal.User{}
	err = s.UserRepository.GetByName(ctx, u, req.Name)
	if err != nil {
		return
	}
	if u.Password != req.Password {
		err = errors.New("password error")
		return
	}
	return
}

//UserLogin 用户登录，先检查用户名密码的正确性，再获取对应token
func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginOrRegisterRequest) (resp *pb.UserLoginOrRegisterResponse, err error) {
	resp = &pb.UserLoginOrRegisterResponse{}
	var u *dal.User
	u, err = s.checkUser(ctx, req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
		return
	}
	resp.UserId = int64(u.ID)
	resp.Token, resp.StatusCode, resp.StatusMsg, err = util.GetToken(u)
	conn := redb.RedisCache.AsynConn()
	defer conn.Close()
	loginTimeKey := fmt.Sprintf("%s", strconv.FormatInt(resp.UserId, 10))
	_, err = conn.AsyncDo("HSET", "aliveUser", loginTimeKey, time.Now().UnixMilli())
	return
}

//UserRegister 用户注册，先判断username是否已存在，再插入user实例，获取对应token
func (s *UserService) UserRegister(ctx context.Context, req *pb.UserLoginOrRegisterRequest) (resp *pb.UserLoginOrRegisterResponse, err error) {
	resp = &pb.UserLoginOrRegisterResponse{}
	if err = s.userIsExists(ctx, req.Name); err != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
		return
	}
	var u dal.User
	u.Name = req.Name
	u.Password = req.Password
	if err = s.createUser(ctx, &u); err != nil {
		resp.StatusCode, resp.StatusMsg = service.BuildResponse(err)
		return
	}
	resp.UserId = int64(u.ID)
	//GetToken生成token
	resp.Token, resp.StatusCode, resp.StatusMsg, err = util.GetToken(&u)
	return
}

//UserInfo 获取用户基本信息
func (s *UserService) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (resp *pb.UserInfoResponse, err error) {
	resp = &pb.UserInfoResponse{}
	resp.StatusCode, resp.StatusMsg = service.BuildResponse(nil)
	user := BuildUser(ctx, req.UserId, req.ToUserId, dal2.NewFollowManagerRepository())
	resp.User = &user
	return
}

//BuildUserList 通用接口，传入user_id集合和follow仓库接口，返回[]User（Id，Name，FollowCount，FollowerCount，IsFollow）
func BuildUserList(ctx context.Context, userID int64, userIDList []int64, m dal2.IFollowRepository) []*pb.User {
	Users := make([]*pb.User, len(userIDList))
	for i := 0; i < len(userIDList); i++ {
		user := BuildUser(ctx, userID, userIDList[i], m)
		Users[i] = &user
	}
	return Users
}

/*BuildUser 返回User（Id，Name，FollowCount，FollowerCount，IsFollow）
userID是当前用户的id，toUserID是要查询的ID，m是follow仓库的接口*/
func BuildUser(ctx context.Context, userID, toUserID int64, m dal2.IFollowRepository) pb.User {
	var user pb.User
	user.Id = toUserID
	user.Name = m.GetName(toUserID)
	user.IsFollow = m.RedisIsFollow(userID, toUserID)
	user.FollowCount = m.RedisFollowCount(toUserID)
	user.FollowerCount = m.RedisFollowerCount(toUserID)
	user.TotalFavorited, user.FavoriteCount, _ = call_fav_com.UserFavoriteInfo(ctx, toUserID)
	user.WorkCount = call_video.GetTotalWorkCount(ctx, toUserID)
	return user
}

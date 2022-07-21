// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: follow.proto

package pb

import (
	"TikTokLite_v2/user_follow/user/pb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RelationActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ToUserId   int64  `protobuf:"varint,3,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
	ActionType int32  `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

func (x *RelationActionRequest) Reset() {
	*x = RelationActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionRequest) ProtoMessage() {}

func (x *RelationActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionRequest.ProtoReflect.Descriptor instead.
func (*RelationActionRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{0}
}

func (x *RelationActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RelationActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *RelationActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *RelationActionResponse) Reset() {
	*x = RelationActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResponse) ProtoMessage() {}

func (x *RelationActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResponse.ProtoReflect.Descriptor instead.
func (*RelationActionResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type RelationFollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RelationFollowListRequest) Reset() {
	*x = RelationFollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListRequest) ProtoMessage() {}

func (x *RelationFollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowListRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{2}
}

func (x *RelationFollowListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFollowListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationFollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserList   []*pb.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *RelationFollowListResponse) Reset() {
	*x = RelationFollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListResponse) ProtoMessage() {}

func (x *RelationFollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowListResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{3}
}

func (x *RelationFollowListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFollowListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFollowListResponse) GetUserList() []*pb.User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type UserRelationInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthorId int64 `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *UserRelationInfoRequest) Reset() {
	*x = UserRelationInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationInfoRequest) ProtoMessage() {}

func (x *UserRelationInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationInfoRequest.ProtoReflect.Descriptor instead.
func (*UserRelationInfoRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{4}
}

func (x *UserRelationInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRelationInfoRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type UserRelationInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerCount int64 `protobuf:"varint,1,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	FollowCount   int64 `protobuf:"varint,2,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	IsFollow      bool  `protobuf:"varint,3,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *UserRelationInfoResponse) Reset() {
	*x = UserRelationInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationInfoResponse) ProtoMessage() {}

func (x *UserRelationInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationInfoResponse.ProtoReflect.Descriptor instead.
func (*UserRelationInfoResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{5}
}

func (x *UserRelationInfoResponse) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *UserRelationInfoResponse) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *UserRelationInfoResponse) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

type FollowListIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FollowListIDRequest) Reset() {
	*x = FollowListIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowListIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowListIDRequest) ProtoMessage() {}

func (x *FollowListIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowListIDRequest.ProtoReflect.Descriptor instead.
func (*FollowListIDRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{6}
}

func (x *FollowListIDRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FollowListIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowList []int64 `protobuf:"varint,1,rep,packed,name=follow_list,json=followList,proto3" json:"follow_list,omitempty"`
}

func (x *FollowListIDResponse) Reset() {
	*x = FollowListIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowListIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowListIDResponse) ProtoMessage() {}

func (x *FollowListIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowListIDResponse.ProtoReflect.Descriptor instead.
func (*FollowListIDResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{7}
}

func (x *FollowListIDResponse) GetFollowList() []int64 {
	if x != nil {
		return x.FollowList
	}
	return nil
}

var File_follow_proto protoreflect.FileDescriptor

var file_follow_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x22, 0x4a, 0x0a, 0x19, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x85, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12,
	0x27, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x2e, 0x0a,
	0x13, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a,
	0x14, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x87, 0x04, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x57, 0x0a, 0x0c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x21, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1f, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follow_proto_rawDescOnce sync.Once
	file_follow_proto_rawDescData = file_follow_proto_rawDesc
)

func file_follow_proto_rawDescGZIP() []byte {
	file_follow_proto_rawDescOnce.Do(func() {
		file_follow_proto_rawDescData = protoimpl.X.CompressGZIP(file_follow_proto_rawDescData)
	})
	return file_follow_proto_rawDescData
}

var file_follow_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_follow_proto_goTypes = []interface{}{
	(*RelationActionRequest)(nil),      // 0: follow.RelationActionRequest
	(*RelationActionResponse)(nil),     // 1: follow.RelationActionResponse
	(*RelationFollowListRequest)(nil),  // 2: follow.RelationFollowListRequest
	(*RelationFollowListResponse)(nil), // 3: follow.RelationFollowListResponse
	(*UserRelationInfoRequest)(nil),    // 4: follow.UserRelationInfoRequest
	(*UserRelationInfoResponse)(nil),   // 5: follow.UserRelationInfoResponse
	(*FollowListIDRequest)(nil),        // 6: follow.FollowListIDRequest
	(*FollowListIDResponse)(nil),       // 7: follow.FollowListIDResponse
	(*pb.User)(nil),                    // 8: user.User
}
var file_follow_proto_depIdxs = []int32{
	8, // 0: follow.RelationFollowListResponse.user_list:type_name -> user.User
	0, // 1: follow.FollowService.RelationAction:input_type -> follow.RelationActionRequest
	2, // 2: follow.FollowService.FollowList:input_type -> follow.RelationFollowListRequest
	2, // 3: follow.FollowService.FollowerList:input_type -> follow.RelationFollowListRequest
	4, // 4: follow.FollowService.UserRelationInfo:input_type -> follow.UserRelationInfoRequest
	6, // 5: follow.FollowService.FollowListID:input_type -> follow.FollowListIDRequest
	6, // 6: follow.FollowService.FollowerListID:input_type -> follow.FollowListIDRequest
	1, // 7: follow.FollowService.RelationAction:output_type -> follow.RelationActionResponse
	3, // 8: follow.FollowService.FollowList:output_type -> follow.RelationFollowListResponse
	3, // 9: follow.FollowService.FollowerList:output_type -> follow.RelationFollowListResponse
	5, // 10: follow.FollowService.UserRelationInfo:output_type -> follow.UserRelationInfoResponse
	7, // 11: follow.FollowService.FollowListID:output_type -> follow.FollowListIDResponse
	7, // 12: follow.FollowService.FollowerListID:output_type -> follow.FollowListIDResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_follow_proto_init() }
func file_follow_proto_init() {
	if File_follow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowListIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowListIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_follow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follow_proto_goTypes,
		DependencyIndexes: file_follow_proto_depIdxs,
		MessageInfos:      file_follow_proto_msgTypes,
	}.Build()
	File_follow_proto = out.File
	file_follow_proto_rawDesc = nil
	file_follow_proto_goTypes = nil
	file_follow_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FollowServiceClient is the client API for FollowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FollowServiceClient interface {
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error)
	FollowList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error)
	FollowerList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error)
	UserRelationInfo(ctx context.Context, in *UserRelationInfoRequest, opts ...grpc.CallOption) (*UserRelationInfoResponse, error)
	FollowListID(ctx context.Context, in *FollowListIDRequest, opts ...grpc.CallOption) (*FollowListIDResponse, error)
	FollowerListID(ctx context.Context, in *FollowListIDRequest, opts ...grpc.CallOption) (*FollowListIDResponse, error)
}

type followServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowServiceClient(cc grpc.ClientConnInterface) FollowServiceClient {
	return &followServiceClient{cc}
}

func (c *followServiceClient) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error) {
	out := new(RelationActionResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/RelationAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error) {
	out := new(RelationFollowListResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/FollowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowerList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error) {
	out := new(RelationFollowListResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/FollowerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) UserRelationInfo(ctx context.Context, in *UserRelationInfoRequest, opts ...grpc.CallOption) (*UserRelationInfoResponse, error) {
	out := new(UserRelationInfoResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/UserRelationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowListID(ctx context.Context, in *FollowListIDRequest, opts ...grpc.CallOption) (*FollowListIDResponse, error) {
	out := new(FollowListIDResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/FollowListID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowerListID(ctx context.Context, in *FollowListIDRequest, opts ...grpc.CallOption) (*FollowListIDResponse, error) {
	out := new(FollowListIDResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/FollowerListID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServiceServer is the server API for FollowService service.
type FollowServiceServer interface {
	RelationAction(context.Context, *RelationActionRequest) (*RelationActionResponse, error)
	FollowList(context.Context, *RelationFollowListRequest) (*RelationFollowListResponse, error)
	FollowerList(context.Context, *RelationFollowListRequest) (*RelationFollowListResponse, error)
	UserRelationInfo(context.Context, *UserRelationInfoRequest) (*UserRelationInfoResponse, error)
	FollowListID(context.Context, *FollowListIDRequest) (*FollowListIDResponse, error)
	FollowerListID(context.Context, *FollowListIDRequest) (*FollowListIDResponse, error)
}

// UnimplementedFollowServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFollowServiceServer struct {
}

func (*UnimplementedFollowServiceServer) RelationAction(context.Context, *RelationActionRequest) (*RelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (*UnimplementedFollowServiceServer) FollowList(context.Context, *RelationFollowListRequest) (*RelationFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowList not implemented")
}
func (*UnimplementedFollowServiceServer) FollowerList(context.Context, *RelationFollowListRequest) (*RelationFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}
func (*UnimplementedFollowServiceServer) UserRelationInfo(context.Context, *UserRelationInfoRequest) (*UserRelationInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRelationInfo not implemented")
}
func (*UnimplementedFollowServiceServer) FollowListID(context.Context, *FollowListIDRequest) (*FollowListIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowListID not implemented")
}
func (*UnimplementedFollowServiceServer) FollowerListID(context.Context, *FollowListIDRequest) (*FollowListIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowerListID not implemented")
}

func RegisterFollowServiceServer(s *grpc.Server, srv FollowServiceServer) {
	s.RegisterService(&_FollowService_serviceDesc, srv)
}

func _FollowService_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/RelationAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).RelationAction(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/FollowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowList(ctx, req.(*RelationFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/FollowerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowerList(ctx, req.(*RelationFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_UserRelationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRelationInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).UserRelationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/UserRelationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).UserRelationInfo(ctx, req.(*UserRelationInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowListID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowListIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowListID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/FollowListID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowListID(ctx, req.(*FollowListIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowerListID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowListIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowerListID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/FollowerListID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowerListID(ctx, req.(*FollowListIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FollowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "follow.FollowService",
	HandlerType: (*FollowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RelationAction",
			Handler:    _FollowService_RelationAction_Handler,
		},
		{
			MethodName: "FollowList",
			Handler:    _FollowService_FollowList_Handler,
		},
		{
			MethodName: "FollowerList",
			Handler:    _FollowService_FollowerList_Handler,
		},
		{
			MethodName: "UserRelationInfo",
			Handler:    _FollowService_UserRelationInfo_Handler,
		},
		{
			MethodName: "FollowListID",
			Handler:    _FollowService_FollowListID_Handler,
		},
		{
			MethodName: "FollowerListID",
			Handler:    _FollowService_FollowerListID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow.proto",
}
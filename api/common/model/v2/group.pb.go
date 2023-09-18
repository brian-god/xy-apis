// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: common/model/v2/group.proto

package v2

import (
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

// 群列表项
type GroupItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID  string `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`   // 用户id
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"` // 昵称
	FaceURL  string `protobuf:"bytes,3,opt,name=faceURL,proto3" json:"faceURL,omitempty"`   // 头像
}

func (x *GroupItem) Reset() {
	*x = GroupItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupItem) ProtoMessage() {}

func (x *GroupItem) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupItem.ProtoReflect.Descriptor instead.
func (*GroupItem) Descriptor() ([]byte, []int) {
	return file_common_model_v2_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupItem) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupItem) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *GroupItem) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID                string `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`
	GroupName              string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Code                   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Notification           string `protobuf:"bytes,4,opt,name=notification,proto3" json:"notification,omitempty"`
	Introduction           string `protobuf:"bytes,5,opt,name=introduction,proto3" json:"introduction,omitempty"`
	FaceURL                string `protobuf:"bytes,6,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	OwnerUserID            string `protobuf:"bytes,7,opt,name=ownerUserID,proto3" json:"ownerUserID,omitempty"`
	CreateTime             uint32 `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	MemberCount            uint32 `protobuf:"varint,9,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	Ex                     string `protobuf:"bytes,10,opt,name=ex,proto3" json:"ex,omitempty"`
	Status                 int32  `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	CreatorUserID          string `protobuf:"bytes,12,opt,name=creatorUserID,proto3" json:"creatorUserID,omitempty"`
	GroupType              int32  `protobuf:"varint,13,opt,name=groupType,proto3" json:"groupType,omitempty"`
	NeedVerification       int32  `protobuf:"varint,14,opt,name=needVerification,proto3" json:"needVerification,omitempty"`
	LookMemberInfo         int32  `protobuf:"varint,15,opt,name=lookMemberInfo,proto3" json:"lookMemberInfo,omitempty"`
	ApplyMemberFriend      int32  `protobuf:"varint,16,opt,name=applyMemberFriend,proto3" json:"applyMemberFriend,omitempty"`
	NotificationUpdateTime uint32 `protobuf:"varint,17,opt,name=notificationUpdateTime,proto3" json:"notificationUpdateTime,omitempty"`
	NotificationUserID     string `protobuf:"bytes,18,opt,name=notificationUserID,proto3" json:"notificationUserID,omitempty"`
	IsReal                 bool   `protobuf:"varint,19,opt,name=isReal,proto3" json:"isReal,omitempty"`
	IsOpen                 uint32 `protobuf:"varint,20,opt,name=isOpen,proto3" json:"isOpen,omitempty"`
	AllowPrivateChat       uint32 `protobuf:"varint,21,opt,name=allowPrivateChat,proto3" json:"allowPrivateChat,omitempty"`
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupInfo.ProtoReflect.Descriptor instead.
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return file_common_model_v2_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupInfo) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupInfo) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GroupInfo) GetNotification() string {
	if x != nil {
		return x.Notification
	}
	return ""
}

func (x *GroupInfo) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *GroupInfo) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

func (x *GroupInfo) GetOwnerUserID() string {
	if x != nil {
		return x.OwnerUserID
	}
	return ""
}

func (x *GroupInfo) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GroupInfo) GetMemberCount() uint32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *GroupInfo) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *GroupInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GroupInfo) GetCreatorUserID() string {
	if x != nil {
		return x.CreatorUserID
	}
	return ""
}

func (x *GroupInfo) GetGroupType() int32 {
	if x != nil {
		return x.GroupType
	}
	return 0
}

func (x *GroupInfo) GetNeedVerification() int32 {
	if x != nil {
		return x.NeedVerification
	}
	return 0
}

func (x *GroupInfo) GetLookMemberInfo() int32 {
	if x != nil {
		return x.LookMemberInfo
	}
	return 0
}

func (x *GroupInfo) GetApplyMemberFriend() int32 {
	if x != nil {
		return x.ApplyMemberFriend
	}
	return 0
}

func (x *GroupInfo) GetNotificationUpdateTime() uint32 {
	if x != nil {
		return x.NotificationUpdateTime
	}
	return 0
}

func (x *GroupInfo) GetNotificationUserID() string {
	if x != nil {
		return x.NotificationUserID
	}
	return ""
}

func (x *GroupInfo) GetIsReal() bool {
	if x != nil {
		return x.IsReal
	}
	return false
}

func (x *GroupInfo) GetIsOpen() uint32 {
	if x != nil {
		return x.IsOpen
	}
	return 0
}

func (x *GroupInfo) GetAllowPrivateChat() uint32 {
	if x != nil {
		return x.AllowPrivateChat
	}
	return 0
}

type GroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo      *PublicUserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	GroupInfo     *GroupInfo      `protobuf:"bytes,2,opt,name=groupInfo,proto3" json:"groupInfo,omitempty"`
	HandleResult  int32           `protobuf:"varint,3,opt,name=handleResult,proto3" json:"handleResult,omitempty"`
	ReqMsg        string          `protobuf:"bytes,4,opt,name=reqMsg,proto3" json:"reqMsg,omitempty"`
	HandleMsg     string          `protobuf:"bytes,5,opt,name=handleMsg,proto3" json:"handleMsg,omitempty"`
	ReqTime       uint32          `protobuf:"varint,6,opt,name=reqTime,proto3" json:"reqTime,omitempty"`
	HandleUserID  string          `protobuf:"bytes,7,opt,name=handleUserID,proto3" json:"handleUserID,omitempty"`
	HandleTime    uint32          `protobuf:"varint,8,opt,name=handleTime,proto3" json:"handleTime,omitempty"`
	Ex            string          `protobuf:"bytes,9,opt,name=ex,proto3" json:"ex,omitempty"`
	JoinSource    int32           `protobuf:"varint,10,opt,name=joinSource,proto3" json:"joinSource,omitempty"`
	InviterUserID string          `protobuf:"bytes,11,opt,name=inviterUserID,proto3" json:"inviterUserID,omitempty"`
}

func (x *GroupRequest) Reset() {
	*x = GroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRequest) ProtoMessage() {}

func (x *GroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRequest.ProtoReflect.Descriptor instead.
func (*GroupRequest) Descriptor() ([]byte, []int) {
	return file_common_model_v2_group_proto_rawDescGZIP(), []int{2}
}

func (x *GroupRequest) GetUserInfo() *PublicUserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *GroupRequest) GetGroupInfo() *GroupInfo {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

func (x *GroupRequest) GetHandleResult() int32 {
	if x != nil {
		return x.HandleResult
	}
	return 0
}

func (x *GroupRequest) GetReqMsg() string {
	if x != nil {
		return x.ReqMsg
	}
	return ""
}

func (x *GroupRequest) GetHandleMsg() string {
	if x != nil {
		return x.HandleMsg
	}
	return ""
}

func (x *GroupRequest) GetReqTime() uint32 {
	if x != nil {
		return x.ReqTime
	}
	return 0
}

func (x *GroupRequest) GetHandleUserID() string {
	if x != nil {
		return x.HandleUserID
	}
	return ""
}

func (x *GroupRequest) GetHandleTime() uint32 {
	if x != nil {
		return x.HandleTime
	}
	return 0
}

func (x *GroupRequest) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *GroupRequest) GetJoinSource() int32 {
	if x != nil {
		return x.JoinSource
	}
	return 0
}

func (x *GroupRequest) GetInviterUserID() string {
	if x != nil {
		return x.InviterUserID
	}
	return ""
}

type UserIDResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UserIDResult) Reset() {
	*x = UserIDResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIDResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDResult) ProtoMessage() {}

func (x *UserIDResult) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDResult.ProtoReflect.Descriptor instead.
func (*UserIDResult) Descriptor() ([]byte, []int) {
	return file_common_model_v2_group_proto_rawDescGZIP(), []int{3}
}

func (x *UserIDResult) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserIDResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GroupMemberFullInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID        string `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`
	UserID         string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	RoleLevel      int32  `protobuf:"varint,3,opt,name=roleLevel,proto3" json:"roleLevel,omitempty"`
	JoinTime       int32  `protobuf:"varint,4,opt,name=joinTime,proto3" json:"joinTime,omitempty"`
	Nickname       string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FaceURL        string `protobuf:"bytes,6,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	AppMangerLevel int32  `protobuf:"varint,7,opt,name=appMangerLevel,proto3" json:"appMangerLevel,omitempty"` //if >0
	JoinSource     int32  `protobuf:"varint,8,opt,name=joinSource,proto3" json:"joinSource,omitempty"`
	OperatorUserID string `protobuf:"bytes,9,opt,name=operatorUserID,proto3" json:"operatorUserID,omitempty"`
	Ex             string `protobuf:"bytes,10,opt,name=ex,proto3" json:"ex,omitempty"`
	MuteEndTime    uint32 `protobuf:"varint,11,opt,name=muteEndTime,proto3" json:"muteEndTime,omitempty"`
	InviterUserID  string `protobuf:"bytes,12,opt,name=inviterUserID,proto3" json:"inviterUserID,omitempty"`
	BackgroundUrl  string `protobuf:"bytes,13,opt,name=backgroundUrl,proto3" json:"backgroundUrl,omitempty"`
}

func (x *GroupMemberFullInfo) Reset() {
	*x = GroupMemberFullInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_v2_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMemberFullInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMemberFullInfo) ProtoMessage() {}

func (x *GroupMemberFullInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_v2_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMemberFullInfo.ProtoReflect.Descriptor instead.
func (*GroupMemberFullInfo) Descriptor() ([]byte, []int) {
	return file_common_model_v2_group_proto_rawDescGZIP(), []int{4}
}

func (x *GroupMemberFullInfo) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupMemberFullInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GroupMemberFullInfo) GetRoleLevel() int32 {
	if x != nil {
		return x.RoleLevel
	}
	return 0
}

func (x *GroupMemberFullInfo) GetJoinTime() int32 {
	if x != nil {
		return x.JoinTime
	}
	return 0
}

func (x *GroupMemberFullInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *GroupMemberFullInfo) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

func (x *GroupMemberFullInfo) GetAppMangerLevel() int32 {
	if x != nil {
		return x.AppMangerLevel
	}
	return 0
}

func (x *GroupMemberFullInfo) GetJoinSource() int32 {
	if x != nil {
		return x.JoinSource
	}
	return 0
}

func (x *GroupMemberFullInfo) GetOperatorUserID() string {
	if x != nil {
		return x.OperatorUserID
	}
	return ""
}

func (x *GroupMemberFullInfo) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *GroupMemberFullInfo) GetMuteEndTime() uint32 {
	if x != nil {
		return x.MuteEndTime
	}
	return 0
}

func (x *GroupMemberFullInfo) GetInviterUserID() string {
	if x != nil {
		return x.InviterUserID
	}
	return ""
}

func (x *GroupMemberFullInfo) GetBackgroundUrl() string {
	if x != nil {
		return x.BackgroundUrl
	}
	return ""
}

var File_common_model_v2_group_proto protoreflect.FileDescriptor

var file_common_model_v2_group_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x32, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x32, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x22, 0xcf, 0x05,
	0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x65,
	0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x65, 0x65, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6c,
	0x6f, 0x6f, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a,
	0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x6c, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x73, 0x4f,
	0x70, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x22,
	0x9b, 0x03, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x22, 0x0a, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x71, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6a, 0x6f, 0x69,
	0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3e, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa5, 0x03,
	0x0a, 0x13, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x75, 0x6c,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x4d, 0x61,
	0x6e, 0x67, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x61, 0x70, 0x70, 0x4d, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x75, 0x74, 0x65, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x75,
	0x74, 0x65, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x24, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x72, 0x6c,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x55, 0x72, 0x6c, 0x42, 0x4c, 0x0a, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x69, 0x61, 0x6e,
	0x2d, 0x67, 0x6f, 0x64, 0x2f, 0x78, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x32,
	0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_model_v2_group_proto_rawDescOnce sync.Once
	file_common_model_v2_group_proto_rawDescData = file_common_model_v2_group_proto_rawDesc
)

func file_common_model_v2_group_proto_rawDescGZIP() []byte {
	file_common_model_v2_group_proto_rawDescOnce.Do(func() {
		file_common_model_v2_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_model_v2_group_proto_rawDescData)
	})
	return file_common_model_v2_group_proto_rawDescData
}

var file_common_model_v2_group_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_model_v2_group_proto_goTypes = []interface{}{
	(*GroupItem)(nil),           // 0: api.common.model.v2.GroupItem
	(*GroupInfo)(nil),           // 1: api.common.model.v2.GroupInfo
	(*GroupRequest)(nil),        // 2: api.common.model.v2.GroupRequest
	(*UserIDResult)(nil),        // 3: api.common.model.v2.UserIDResult
	(*GroupMemberFullInfo)(nil), // 4: api.common.model.v2.GroupMemberFullInfo
	(*PublicUserInfo)(nil),      // 5: api.common.model.v2.PublicUserInfo
}
var file_common_model_v2_group_proto_depIdxs = []int32{
	5, // 0: api.common.model.v2.GroupRequest.userInfo:type_name -> api.common.model.v2.PublicUserInfo
	1, // 1: api.common.model.v2.GroupRequest.groupInfo:type_name -> api.common.model.v2.GroupInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_model_v2_group_proto_init() }
func file_common_model_v2_group_proto_init() {
	if File_common_model_v2_group_proto != nil {
		return
	}
	file_common_model_v2_friend_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_model_v2_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupItem); i {
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
		file_common_model_v2_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupInfo); i {
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
		file_common_model_v2_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRequest); i {
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
		file_common_model_v2_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIDResult); i {
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
		file_common_model_v2_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMemberFullInfo); i {
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
			RawDescriptor: file_common_model_v2_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_model_v2_group_proto_goTypes,
		DependencyIndexes: file_common_model_v2_group_proto_depIdxs,
		MessageInfos:      file_common_model_v2_group_proto_msgTypes,
	}.Build()
	File_common_model_v2_group_proto = out.File
	file_common_model_v2_group_proto_rawDesc = nil
	file_common_model_v2_group_proto_goTypes = nil
	file_common_model_v2_group_proto_depIdxs = nil
}

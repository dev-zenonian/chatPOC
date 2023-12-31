// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: groupService.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetGroupWithIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetGroupWithIDRequest) Reset() {
	*x = GetGroupWithIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupWithIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupWithIDRequest) ProtoMessage() {}

func (x *GetGroupWithIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupWithIDRequest.ProtoReflect.Descriptor instead.
func (*GetGroupWithIDRequest) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupWithIDRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetGroupWithIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are assignable to Data:
	//
	//	*GetGroupWithIDResponse_Group
	//	*GetGroupWithIDResponse_Msg
	Data isGetGroupWithIDResponse_Data `protobuf_oneof:"data"`
}

func (x *GetGroupWithIDResponse) Reset() {
	*x = GetGroupWithIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupWithIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupWithIDResponse) ProtoMessage() {}

func (x *GetGroupWithIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupWithIDResponse.ProtoReflect.Descriptor instead.
func (*GetGroupWithIDResponse) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupWithIDResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (m *GetGroupWithIDResponse) GetData() isGetGroupWithIDResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *GetGroupWithIDResponse) GetGroup() *GroupInformation {
	if x, ok := x.GetData().(*GetGroupWithIDResponse_Group); ok {
		return x.Group
	}
	return nil
}

func (x *GetGroupWithIDResponse) GetMsg() string {
	if x, ok := x.GetData().(*GetGroupWithIDResponse_Msg); ok {
		return x.Msg
	}
	return ""
}

type isGetGroupWithIDResponse_Data interface {
	isGetGroupWithIDResponse_Data()
}

type GetGroupWithIDResponse_Group struct {
	Group *GroupInformation `protobuf:"bytes,2,opt,name=group,proto3,oneof"`
}

type GetGroupWithIDResponse_Msg struct {
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3,oneof"`
}

func (*GetGroupWithIDResponse_Group) isGetGroupWithIDResponse_Data() {}

func (*GetGroupWithIDResponse_Msg) isGetGroupWithIDResponse_Data() {}

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName string   `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	ClientIDs []string `protobuf:"bytes,2,rep,name=clientIDs,proto3" json:"clientIDs,omitempty"`
	AdminID   string   `protobuf:"bytes,3,opt,name=adminID,proto3" json:"adminID,omitempty"`
	IsPrivate bool     `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGroupRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *CreateGroupRequest) GetClientIDs() []string {
	if x != nil {
		return x.ClientIDs
	}
	return nil
}

func (x *CreateGroupRequest) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

func (x *CreateGroupRequest) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are assignable to Data:
	//
	//	*CreateGroupResponse_Group
	//	*CreateGroupResponse_Msg
	Data isCreateGroupResponse_Data `protobuf_oneof:"data"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGroupResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (m *CreateGroupResponse) GetData() isCreateGroupResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CreateGroupResponse) GetGroup() *GroupInformation {
	if x, ok := x.GetData().(*CreateGroupResponse_Group); ok {
		return x.Group
	}
	return nil
}

func (x *CreateGroupResponse) GetMsg() string {
	if x, ok := x.GetData().(*CreateGroupResponse_Msg); ok {
		return x.Msg
	}
	return ""
}

type isCreateGroupResponse_Data interface {
	isCreateGroupResponse_Data()
}

type CreateGroupResponse_Group struct {
	Group *GroupInformation `protobuf:"bytes,2,opt,name=group,proto3,oneof"`
}

type CreateGroupResponse_Msg struct {
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3,oneof"`
}

func (*CreateGroupResponse_Group) isCreateGroupResponse_Data() {}

func (*CreateGroupResponse_Msg) isCreateGroupResponse_Data() {}

type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID string `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`
	AdminID string `protobuf:"bytes,2,opt,name=adminID,proto3" json:"adminID,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGroupRequest) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *DeleteGroupRequest) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

type DeleteGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeleteGroupResponse) Reset() {
	*x = DeleteGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupResponse) ProtoMessage() {}

func (x *DeleteGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGroupResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *DeleteGroupResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_groupService_proto protoreflect.FileDescriptor

var file_groupService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x89, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x22, 0x78, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x2f, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x12, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x48, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb1, 0x02, 0x0a, 0x0c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x5e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a,
	0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x5b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groupService_proto_rawDescOnce sync.Once
	file_groupService_proto_rawDescData = file_groupService_proto_rawDesc
)

func file_groupService_proto_rawDescGZIP() []byte {
	file_groupService_proto_rawDescOnce.Do(func() {
		file_groupService_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupService_proto_rawDescData)
	})
	return file_groupService_proto_rawDescData
}

var file_groupService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_groupService_proto_goTypes = []interface{}{
	(*GetGroupWithIDRequest)(nil),  // 0: proto.GetGroupWithIDRequest
	(*GetGroupWithIDResponse)(nil), // 1: proto.GetGroupWithIDResponse
	(*CreateGroupRequest)(nil),     // 2: proto.CreateGroupRequest
	(*CreateGroupResponse)(nil),    // 3: proto.CreateGroupResponse
	(*DeleteGroupRequest)(nil),     // 4: proto.DeleteGroupRequest
	(*DeleteGroupResponse)(nil),    // 5: proto.DeleteGroupResponse
	(*GroupInformation)(nil),       // 6: proto.GroupInformation
}
var file_groupService_proto_depIdxs = []int32{
	6, // 0: proto.GetGroupWithIDResponse.group:type_name -> proto.GroupInformation
	6, // 1: proto.CreateGroupResponse.group:type_name -> proto.GroupInformation
	0, // 2: proto.GroupService.GetGroupWithID:input_type -> proto.GetGroupWithIDRequest
	2, // 3: proto.GroupService.CreateGroup:input_type -> proto.CreateGroupRequest
	4, // 4: proto.GroupService.DeleteGroup:input_type -> proto.DeleteGroupRequest
	1, // 5: proto.GroupService.GetGroupWithID:output_type -> proto.GetGroupWithIDResponse
	3, // 6: proto.GroupService.CreateGroup:output_type -> proto.CreateGroupResponse
	5, // 7: proto.GroupService.DeleteGroup:output_type -> proto.DeleteGroupResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_groupService_proto_init() }
func file_groupService_proto_init() {
	if File_groupService_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_groupService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupWithIDRequest); i {
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
		file_groupService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupWithIDResponse); i {
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
		file_groupService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_groupService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_groupService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_groupService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupResponse); i {
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
	file_groupService_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetGroupWithIDResponse_Group)(nil),
		(*GetGroupWithIDResponse_Msg)(nil),
	}
	file_groupService_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CreateGroupResponse_Group)(nil),
		(*CreateGroupResponse_Msg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_groupService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groupService_proto_goTypes,
		DependencyIndexes: file_groupService_proto_depIdxs,
		MessageInfos:      file_groupService_proto_msgTypes,
	}.Build()
	File_groupService_proto = out.File
	file_groupService_proto_rawDesc = nil
	file_groupService_proto_goTypes = nil
	file_groupService_proto_depIdxs = nil
}

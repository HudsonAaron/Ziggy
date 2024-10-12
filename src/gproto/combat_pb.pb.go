// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: combat_pb.proto

package gproto

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

// 请求入局（进房间）
type ApplyToCombatC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode   int32      `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`    // 战斗模式，大乱斗、组队PVE，组队对抗、闯关
	Myself *ActorData `protobuf:"bytes,2,opt,name=myself,proto3" json:"myself,omitempty"` // 我自己的开局选用角色上传给服务器
}

func (x *ApplyToCombatC) Reset() {
	*x = ApplyToCombatC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyToCombatC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyToCombatC) ProtoMessage() {}

func (x *ApplyToCombatC) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyToCombatC.ProtoReflect.Descriptor instead.
func (*ApplyToCombatC) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{0}
}

func (x *ApplyToCombatC) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *ApplyToCombatC) GetMyself() *ActorData {
	if x != nil {
		return x.Myself
	}
	return nil
}

// 回复入局（进房间）
type ApplyToCombatS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 返回码
	Mode   int32      `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"` // 战斗模式，大乱斗、组队PVE，组队对抗、闯关
	RoomId int32      `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Myself *ActorData `protobuf:"bytes,4,opt,name=myself,proto3" json:"myself,omitempty"` // 我自己
}

func (x *ApplyToCombatS) Reset() {
	*x = ApplyToCombatS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyToCombatS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyToCombatS) ProtoMessage() {}

func (x *ApplyToCombatS) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyToCombatS.ProtoReflect.Descriptor instead.
func (*ApplyToCombatS) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{1}
}

func (x *ApplyToCombatS) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ApplyToCombatS) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *ApplyToCombatS) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ApplyToCombatS) GetMyself() *ActorData {
	if x != nil {
		return x.Myself
	}
	return nil
}

// 加入战斗--告诉服务器我的出生位置
type JoinInCombatC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actors []*ActorData `protobuf:"bytes,1,rep,name=actors,proto3" json:"actors,omitempty"` //
}

func (x *JoinInCombatC) Reset() {
	*x = JoinInCombatC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinInCombatC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinInCombatC) ProtoMessage() {}

func (x *JoinInCombatC) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinInCombatC.ProtoReflect.Descriptor instead.
func (*JoinInCombatC) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{2}
}

func (x *JoinInCombatC) GetActors() []*ActorData {
	if x != nil {
		return x.Actors
	}
	return nil
}

// 加入战斗--推送要加入的角色
type JoinInCombatS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Actors []*ActorData `protobuf:"bytes,2,rep,name=actors,proto3" json:"actors,omitempty"` //
}

func (x *JoinInCombatS) Reset() {
	*x = JoinInCombatS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinInCombatS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinInCombatS) ProtoMessage() {}

func (x *JoinInCombatS) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinInCombatS.ProtoReflect.Descriptor instead.
func (*JoinInCombatS) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{3}
}

func (x *JoinInCombatS) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *JoinInCombatS) GetActors() []*ActorData {
	if x != nil {
		return x.Actors
	}
	return nil
}

// 离开战斗--告诉服务器我要离开
type LeaveCombatC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actors []*ActorData `protobuf:"bytes,1,rep,name=actors,proto3" json:"actors,omitempty"` //
}

func (x *LeaveCombatC) Reset() {
	*x = LeaveCombatC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveCombatC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveCombatC) ProtoMessage() {}

func (x *LeaveCombatC) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveCombatC.ProtoReflect.Descriptor instead.
func (*LeaveCombatC) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{4}
}

func (x *LeaveCombatC) GetActors() []*ActorData {
	if x != nil {
		return x.Actors
	}
	return nil
}

// 离开战斗--推送要离开的角色
type LeaveCombatS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Actors []*ActorData `protobuf:"bytes,2,rep,name=actors,proto3" json:"actors,omitempty"` //
}

func (x *LeaveCombatS) Reset() {
	*x = LeaveCombatS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveCombatS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveCombatS) ProtoMessage() {}

func (x *LeaveCombatS) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveCombatS.ProtoReflect.Descriptor instead.
func (*LeaveCombatS) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{5}
}

func (x *LeaveCombatS) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LeaveCombatS) GetActors() []*ActorData {
	if x != nil {
		return x.Actors
	}
	return nil
}

// 房间战斗---客户端上传帧
type RoomTickC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick   int32        `protobuf:"varint,1,opt,name=tick,proto3" json:"tick,omitempty"`    // 客户端当前帧
	Actors []*ActorData `protobuf:"bytes,2,rep,name=actors,proto3" json:"actors,omitempty"` // 变化的角色
}

func (x *RoomTickC) Reset() {
	*x = RoomTickC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomTickC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomTickC) ProtoMessage() {}

func (x *RoomTickC) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomTickC.ProtoReflect.Descriptor instead.
func (*RoomTickC) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{6}
}

func (x *RoomTickC) GetTick() int32 {
	if x != nil {
		return x.Tick
	}
	return 0
}

func (x *RoomTickC) GetActors() []*ActorData {
	if x != nil {
		return x.Actors
	}
	return nil
}

// 房间战斗---服务器下发帧
type RoomTickS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code            int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Tick            int32        `protobuf:"varint,2,opt,name=tick,proto3" json:"tick,omitempty"`                                                       // 当前帧
	Actors          []*ActorData `protobuf:"bytes,3,rep,name=actors,proto3" json:"actors,omitempty"`                                                    // 有变化的角色
	RemoveUniqueIds []int32      `protobuf:"varint,4,rep,packed,name=remove_unique_ids,json=removeUniqueIds,proto3" json:"remove_unique_ids,omitempty"` //离开的角色
}

func (x *RoomTickS) Reset() {
	*x = RoomTickS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomTickS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomTickS) ProtoMessage() {}

func (x *RoomTickS) ProtoReflect() protoreflect.Message {
	mi := &file_combat_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomTickS.ProtoReflect.Descriptor instead.
func (*RoomTickS) Descriptor() ([]byte, []int) {
	return file_combat_pb_proto_rawDescGZIP(), []int{7}
}

func (x *RoomTickS) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RoomTickS) GetTick() int32 {
	if x != nil {
		return x.Tick
	}
	return 0
}

func (x *RoomTickS) GetActors() []*ActorData {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *RoomTickS) GetRemoveUniqueIds() []int32 {
	if x != nil {
		return x.RemoveUniqueIds
	}
	return nil
}

var File_combat_pb_proto protoreflect.FileDescriptor

var file_combat_pb_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x11, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x5f, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x79, 0x73, 0x65, 0x6c, 0x66, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6d, 0x79, 0x73, 0x65, 0x6c, 0x66, 0x22,
	0x80, 0x01, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6d,
	0x62, 0x61, 0x74, 0x5f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x79, 0x73, 0x65, 0x6c, 0x66,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6d, 0x79, 0x73, 0x65,
	0x6c, 0x66, 0x22, 0x3e, 0x0a, 0x10, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x5f, 0x63, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x22, 0x52, 0x0a, 0x10, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x5f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x3c, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x5f,
	0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x5f, 0x63, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x22, 0x50, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x5f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x5f, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x5f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x2a, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x49, 0x64, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x67, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_combat_pb_proto_rawDescOnce sync.Once
	file_combat_pb_proto_rawDescData = file_combat_pb_proto_rawDesc
)

func file_combat_pb_proto_rawDescGZIP() []byte {
	file_combat_pb_proto_rawDescOnce.Do(func() {
		file_combat_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_combat_pb_proto_rawDescData)
	})
	return file_combat_pb_proto_rawDescData
}

var file_combat_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_combat_pb_proto_goTypes = []interface{}{
	(*ApplyToCombatC)(nil), // 0: gproto.apply_to_combat_c
	(*ApplyToCombatS)(nil), // 1: gproto.apply_to_combat_s
	(*JoinInCombatC)(nil),  // 2: gproto.join_in_combat_c
	(*JoinInCombatS)(nil),  // 3: gproto.join_in_combat_s
	(*LeaveCombatC)(nil),   // 4: gproto.leave_combat_c
	(*LeaveCombatS)(nil),   // 5: gproto.leave_combat_s
	(*RoomTickC)(nil),      // 6: gproto.room_tick_c
	(*RoomTickS)(nil),      // 7: gproto.room_tick_s
	(*ActorData)(nil),      // 8: gproto.actor_data
}
var file_combat_pb_proto_depIdxs = []int32{
	8, // 0: gproto.apply_to_combat_c.myself:type_name -> gproto.actor_data
	8, // 1: gproto.apply_to_combat_s.myself:type_name -> gproto.actor_data
	8, // 2: gproto.join_in_combat_c.actors:type_name -> gproto.actor_data
	8, // 3: gproto.join_in_combat_s.actors:type_name -> gproto.actor_data
	8, // 4: gproto.leave_combat_c.actors:type_name -> gproto.actor_data
	8, // 5: gproto.leave_combat_s.actors:type_name -> gproto.actor_data
	8, // 6: gproto.room_tick_c.actors:type_name -> gproto.actor_data
	8, // 7: gproto.room_tick_s.actors:type_name -> gproto.actor_data
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_combat_pb_proto_init() }
func file_combat_pb_proto_init() {
	if File_combat_pb_proto != nil {
		return
	}
	file_common_pb_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_combat_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyToCombatC); i {
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
		file_combat_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyToCombatS); i {
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
		file_combat_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinInCombatC); i {
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
		file_combat_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinInCombatS); i {
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
		file_combat_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveCombatC); i {
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
		file_combat_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveCombatS); i {
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
		file_combat_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomTickC); i {
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
		file_combat_pb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomTickS); i {
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
			RawDescriptor: file_combat_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_combat_pb_proto_goTypes,
		DependencyIndexes: file_combat_pb_proto_depIdxs,
		MessageInfos:      file_combat_pb_proto_msgTypes,
	}.Build()
	File_combat_pb_proto = out.File
	file_combat_pb_proto_rawDesc = nil
	file_combat_pb_proto_goTypes = nil
	file_combat_pb_proto_depIdxs = nil
}
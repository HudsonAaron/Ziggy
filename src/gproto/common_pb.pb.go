// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: common_pb.proto

// 通用模块

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

// 每条协议的格式
type Comdata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoId int32  `protobuf:"varint,1,opt,name=proto_id,json=protoId,proto3" json:"proto_id,omitempty"` // 协议号
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                       // 协议内容
}

func (x *Comdata) Reset() {
	*x = Comdata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comdata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comdata) ProtoMessage() {}

func (x *Comdata) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comdata.ProtoReflect.Descriptor instead.
func (*Comdata) Descriptor() ([]byte, []int) {
	return file_common_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Comdata) GetProtoId() int32 {
	if x != nil {
		return x.ProtoId
	}
	return 0
}

func (x *Comdata) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 发送协议包
type ProtoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comdata []*Comdata `protobuf:"bytes,1,rep,name=comdata,proto3" json:"comdata,omitempty"` // 协议列表
}

func (x *ProtoData) Reset() {
	*x = ProtoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoData) ProtoMessage() {}

func (x *ProtoData) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoData.ProtoReflect.Descriptor instead.
func (*ProtoData) Descriptor() ([]byte, []int) {
	return file_common_pb_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoData) GetComdata() []*Comdata {
	if x != nil {
		return x.Comdata
	}
	return nil
}

// 世界坐标,浮点数×0.001f
type WorldSpaceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position   []int32 `protobuf:"varint,1,rep,packed,name=position,proto3" json:"position,omitempty"`     // 世界坐标位置，x,y,z
	Quaternion []int32 `protobuf:"varint,2,rep,packed,name=quaternion,proto3" json:"quaternion,omitempty"` // 世界坐标三轴朝向，x,y,z
	Scale      []int32 `protobuf:"varint,3,rep,packed,name=scale,proto3" json:"scale,omitempty"`           // 缩放比例，x,y,z
}

func (x *WorldSpaceData) Reset() {
	*x = WorldSpaceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldSpaceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldSpaceData) ProtoMessage() {}

func (x *WorldSpaceData) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldSpaceData.ProtoReflect.Descriptor instead.
func (*WorldSpaceData) Descriptor() ([]byte, []int) {
	return file_common_pb_proto_rawDescGZIP(), []int{2}
}

func (x *WorldSpaceData) GetPosition() []int32 {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *WorldSpaceData) GetQuaternion() []int32 {
	if x != nil {
		return x.Quaternion
	}
	return nil
}

func (x *WorldSpaceData) GetScale() []int32 {
	if x != nil {
		return x.Scale
	}
	return nil
}

// 战斗单元 基本配置
type CombatantData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId int32 `protobuf:"varint,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"` // 角色配置表ID
	Level    int32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`                       // 角色等级
	Type     int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`                         // 类型( 1: 真人, 2: 宠物, 3: 怪物、4：队长、5：boss)
	CampId   int32 `protobuf:"varint,4,opt,name=campId,proto3" json:"campId,omitempty"`                     //1红方、2蓝方、3中立阵营
}

func (x *CombatantData) Reset() {
	*x = CombatantData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombatantData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombatantData) ProtoMessage() {}

func (x *CombatantData) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombatantData.ProtoReflect.Descriptor instead.
func (*CombatantData) Descriptor() ([]byte, []int) {
	return file_common_pb_proto_rawDescGZIP(), []int{3}
}

func (x *CombatantData) GetConfigId() int32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *CombatantData) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CombatantData) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CombatantData) GetCampId() int32 {
	if x != nil {
		return x.CampId
	}
	return 0
}

// 释放技能信息
type FireData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FireId    int32   `protobuf:"varint,1,opt,name=fire_id,json=fireId,proto3" json:"fire_id,omitempty"`          // 角色配置表ID
	FireLevel int32   `protobuf:"varint,2,opt,name=fire_level,json=fireLevel,proto3" json:"fire_level,omitempty"` // 角色等级
	FireState int32   `protobuf:"varint,3,opt,name=fire_state,json=fireState,proto3" json:"fire_state,omitempty"` //1是开火，2是持续，3是中途打断
	Targets   []int32 `protobuf:"varint,4,rep,packed,name=targets,proto3" json:"targets,omitempty"`               // 技能目标
}

func (x *FireData) Reset() {
	*x = FireData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FireData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FireData) ProtoMessage() {}

func (x *FireData) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FireData.ProtoReflect.Descriptor instead.
func (*FireData) Descriptor() ([]byte, []int) {
	return file_common_pb_proto_rawDescGZIP(), []int{4}
}

func (x *FireData) GetFireId() int32 {
	if x != nil {
		return x.FireId
	}
	return 0
}

func (x *FireData) GetFireLevel() int32 {
	if x != nil {
		return x.FireLevel
	}
	return 0
}

func (x *FireData) GetFireState() int32 {
	if x != nil {
		return x.FireState
	}
	return 0
}

func (x *FireData) GetTargets() []int32 {
	if x != nil {
		return x.Targets
	}
	return nil
}

// 摇杆信息
type JoystickData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JoyX     int32 `protobuf:"varint,1,opt,name=joy_x,json=joyX,proto3" json:"joy_x,omitempty"`             // 角色配置表ID
	JoyY     int32 `protobuf:"varint,2,opt,name=joy_y,json=joyY,proto3" json:"joy_y,omitempty"`             // 角色等级
	JoyState int32 `protobuf:"varint,3,opt,name=joy_state,json=joyState,proto3" json:"joy_state,omitempty"` // 按下0，按住1，弹起2
}

func (x *JoystickData) Reset() {
	*x = JoystickData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoystickData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoystickData) ProtoMessage() {}

func (x *JoystickData) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoystickData.ProtoReflect.Descriptor instead.
func (*JoystickData) Descriptor() ([]byte, []int) {
	return file_common_pb_proto_rawDescGZIP(), []int{5}
}

func (x *JoystickData) GetJoyX() int32 {
	if x != nil {
		return x.JoyX
	}
	return 0
}

func (x *JoystickData) GetJoyY() int32 {
	if x != nil {
		return x.JoyY
	}
	return 0
}

func (x *JoystickData) GetJoyState() int32 {
	if x != nil {
		return x.JoyState
	}
	return 0
}

// 战斗中一个角色的世界坐标信息
type ActorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId int32           `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"` // 玩家ID，不同玩家不同ID，AI类型角色值是0
	UniqueId int32           `protobuf:"varint,2,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"` // 在这场战斗中的角色（玩家、npc）唯一ID，
	RData    *CombatantData  `protobuf:"bytes,3,opt,name=r_data,json=rData,proto3" json:"r_data,omitempty"`           // 角色基本信息
	SData    *FireData       `protobuf:"bytes,4,opt,name=s_data,json=sData,proto3" json:"s_data,omitempty"`
	WsData   *WorldSpaceData `protobuf:"bytes,5,opt,name=ws_data,json=wsData,proto3" json:"ws_data,omitempty"`    //世界坐标
	JoyData  *JoystickData   `protobuf:"bytes,6,opt,name=joy_data,json=joyData,proto3" json:"joy_data,omitempty"` //按键信息
}

func (x *ActorData) Reset() {
	*x = ActorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorData) ProtoMessage() {}

func (x *ActorData) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorData.ProtoReflect.Descriptor instead.
func (*ActorData) Descriptor() ([]byte, []int) {
	return file_common_pb_proto_rawDescGZIP(), []int{6}
}

func (x *ActorData) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *ActorData) GetUniqueId() int32 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

func (x *ActorData) GetRData() *CombatantData {
	if x != nil {
		return x.RData
	}
	return nil
}

func (x *ActorData) GetSData() *FireData {
	if x != nil {
		return x.SData
	}
	return nil
}

func (x *ActorData) GetWsData() *WorldSpaceData {
	if x != nil {
		return x.WsData
	}
	return nil
}

func (x *ActorData) GetJoyData() *JoystickData {
	if x != nil {
		return x.JoyData
	}
	return nil
}

var File_common_pb_proto protoreflect.FileDescriptor

var file_common_pb_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x10,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x71, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0a, 0x71, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x22, 0x6f, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x61, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x61, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x61, 0x6d,
	0x70, 0x49, 0x64, 0x22, 0x7c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x66, 0x69, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x22, 0x56, 0x0a, 0x0d, 0x6a, 0x6f, 0x79, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x6a, 0x6f, 0x79, 0x5f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6a, 0x6f, 0x79, 0x58, 0x12, 0x13, 0x0a, 0x05, 0x6a, 0x6f, 0x79, 0x5f, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6a, 0x6f, 0x79, 0x59, 0x12, 0x1b, 0x0a, 0x09,
	0x6a, 0x6f, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6a, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x84, 0x02, 0x0a, 0x0a, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x07, 0x77,
	0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x77, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30,
	0x0a, 0x08, 0x6a, 0x6f, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x6f, 0x79, 0x73, 0x74, 0x69,
	0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6a, 0x6f, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_pb_proto_rawDescOnce sync.Once
	file_common_pb_proto_rawDescData = file_common_pb_proto_rawDesc
)

func file_common_pb_proto_rawDescGZIP() []byte {
	file_common_pb_proto_rawDescOnce.Do(func() {
		file_common_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pb_proto_rawDescData)
	})
	return file_common_pb_proto_rawDescData
}

var file_common_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_pb_proto_goTypes = []interface{}{
	(*Comdata)(nil),        // 0: gproto.comdata
	(*ProtoData)(nil),      // 1: gproto.proto_data
	(*WorldSpaceData)(nil), // 2: gproto.world_space_data
	(*CombatantData)(nil),  // 3: gproto.combatant_data
	(*FireData)(nil),       // 4: gproto.fire_data
	(*JoystickData)(nil),   // 5: gproto.joystick_data
	(*ActorData)(nil),      // 6: gproto.actor_data
}
var file_common_pb_proto_depIdxs = []int32{
	0, // 0: gproto.proto_data.comdata:type_name -> gproto.comdata
	3, // 1: gproto.actor_data.r_data:type_name -> gproto.combatant_data
	4, // 2: gproto.actor_data.s_data:type_name -> gproto.fire_data
	2, // 3: gproto.actor_data.ws_data:type_name -> gproto.world_space_data
	5, // 4: gproto.actor_data.joy_data:type_name -> gproto.joystick_data
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_pb_proto_init() }
func file_common_pb_proto_init() {
	if File_common_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comdata); i {
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
		file_common_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoData); i {
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
		file_common_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldSpaceData); i {
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
		file_common_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombatantData); i {
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
		file_common_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FireData); i {
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
		file_common_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoystickData); i {
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
		file_common_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorData); i {
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
			RawDescriptor: file_common_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pb_proto_goTypes,
		DependencyIndexes: file_common_pb_proto_depIdxs,
		MessageInfos:      file_common_pb_proto_msgTypes,
	}.Build()
	File_common_pb_proto = out.File
	file_common_pb_proto_rawDesc = nil
	file_common_pb_proto_goTypes = nil
	file_common_pb_proto_depIdxs = nil
}

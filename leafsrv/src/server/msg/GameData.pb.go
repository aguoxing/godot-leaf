// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.1
// source: GameData.proto

//option csharp_namespace = "LeafClient";

package msg

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

type PlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   string    `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	PlayerName string    `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
	Vel        *Velocity `protobuf:"bytes,3,opt,name=vel,proto3" json:"vel,omitempty"`
}

func (x *PlayerRequest) Reset() {
	*x = PlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRequest) ProtoMessage() {}

func (x *PlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GameData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRequest.ProtoReflect.Descriptor instead.
func (*PlayerRequest) Descriptor() ([]byte, []int) {
	return file_GameData_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerRequest) GetVel() *Velocity {
	if x != nil {
		return x.Vel
	}
	return nil
}

type PlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   string    `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	PlayerName string    `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
	Vel        *Velocity `protobuf:"bytes,3,opt,name=vel,proto3" json:"vel,omitempty"`
	Msg        string    `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PlayerResponse) Reset() {
	*x = PlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerResponse) ProtoMessage() {}

func (x *PlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GameData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerResponse.ProtoReflect.Descriptor instead.
func (*PlayerResponse) Descriptor() ([]byte, []int) {
	return file_GameData_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerResponse) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerResponse) GetVel() *Velocity {
	if x != nil {
		return x.Vel
	}
	return nil
}

func (x *PlayerResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Velocity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Velocity) Reset() {
	*x = Velocity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Velocity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Velocity) ProtoMessage() {}

func (x *Velocity) ProtoReflect() protoreflect.Message {
	mi := &file_GameData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Velocity.ProtoReflect.Descriptor instead.
func (*Velocity) Descriptor() ([]byte, []int) {
	return file_GameData_proto_rawDescGZIP(), []int{2}
}

func (x *Velocity) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Velocity) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	X        int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y        int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GameData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_GameData_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerInfo) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerInfo) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PlayerInfo) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type NewGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *NewGame) Reset() {
	*x = NewGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewGame) ProtoMessage() {}

func (x *NewGame) ProtoReflect() protoreflect.Message {
	mi := &file_GameData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewGame.ProtoReflect.Descriptor instead.
func (*NewGame) Descriptor() ([]byte, []int) {
	return file_GameData_proto_rawDescGZIP(), []int{4}
}

func (x *NewGame) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *NewGame) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type StartGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId     string        `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerList []*PlayerInfo `protobuf:"bytes,2,rep,name=playerList,proto3" json:"playerList,omitempty"`
}

func (x *StartGame) Reset() {
	*x = StartGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameData_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGame) ProtoMessage() {}

func (x *StartGame) ProtoReflect() protoreflect.Message {
	mi := &file_GameData_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGame.ProtoReflect.Descriptor instead.
func (*StartGame) Descriptor() ([]byte, []int) {
	return file_GameData_proto_rawDescGZIP(), []int{5}
}

func (x *StartGame) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *StartGame) GetPlayerList() []*PlayerInfo {
	if x != nil {
		return x.PlayerList
	}
	return nil
}

type ExitGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *ExitGame) Reset() {
	*x = ExitGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameData_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitGame) ProtoMessage() {}

func (x *ExitGame) ProtoReflect() protoreflect.Message {
	mi := &file_GameData_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitGame.ProtoReflect.Descriptor instead.
func (*ExitGame) Descriptor() ([]byte, []int) {
	return file_GameData_proto_rawDescGZIP(), []int{6}
}

func (x *ExitGame) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *ExitGame) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

var File_GameData_proto protoreflect.FileDescriptor

var file_GameData_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x6b, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x03, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x03, 0x76, 0x65,
	0x6c, 0x22, 0x7e, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x03, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x03, 0x76, 0x65, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x26, 0x0a, 0x08, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x44, 0x0a, 0x0a, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22,
	0x3d, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x08, 0x45, 0x78, 0x69, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GameData_proto_rawDescOnce sync.Once
	file_GameData_proto_rawDescData = file_GameData_proto_rawDesc
)

func file_GameData_proto_rawDescGZIP() []byte {
	file_GameData_proto_rawDescOnce.Do(func() {
		file_GameData_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameData_proto_rawDescData)
	})
	return file_GameData_proto_rawDescData
}

var file_GameData_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_GameData_proto_goTypes = []interface{}{
	(*PlayerRequest)(nil),  // 0: pb.PlayerRequest
	(*PlayerResponse)(nil), // 1: pb.PlayerResponse
	(*Velocity)(nil),       // 2: pb.Velocity
	(*PlayerInfo)(nil),     // 3: pb.PlayerInfo
	(*NewGame)(nil),        // 4: pb.NewGame
	(*StartGame)(nil),      // 5: pb.StartGame
	(*ExitGame)(nil),       // 6: pb.ExitGame
}
var file_GameData_proto_depIdxs = []int32{
	2, // 0: pb.PlayerRequest.vel:type_name -> pb.Velocity
	2, // 1: pb.PlayerResponse.vel:type_name -> pb.Velocity
	3, // 2: pb.StartGame.playerList:type_name -> pb.PlayerInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GameData_proto_init() }
func file_GameData_proto_init() {
	if File_GameData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GameData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRequest); i {
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
		file_GameData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerResponse); i {
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
		file_GameData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Velocity); i {
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
		file_GameData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
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
		file_GameData_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewGame); i {
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
		file_GameData_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGame); i {
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
		file_GameData_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitGame); i {
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
			RawDescriptor: file_GameData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GameData_proto_goTypes,
		DependencyIndexes: file_GameData_proto_depIdxs,
		MessageInfos:      file_GameData_proto_msgTypes,
	}.Build()
	File_GameData_proto = out.File
	file_GameData_proto_rawDesc = nil
	file_GameData_proto_goTypes = nil
	file_GameData_proto_depIdxs = nil
}
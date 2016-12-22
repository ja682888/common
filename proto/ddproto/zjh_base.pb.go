// Code generated by protoc-gen-go.
// source: zjh_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EProtoId int32

const (
	// //////////////////////////////////
	EProtoId_ZJH_PID_HEARTBEAT           EProtoId = 0
	EProtoId_ZJH_PID_QUICKCOnn           EProtoId = 1
	EProtoId_ZJH_PID_QUICKCONNA_ACK      EProtoId = 2
	EProtoId_ZJH_PID_GAME_LOGIN          EProtoId = 3
	EProtoId_ZJH_PID_GAME_LOGIN_ACK      EProtoId = 4
	EProtoId_ZJH_PID_GET_ROOM_LIST       EProtoId = 5
	EProtoId_ZJH_PID_ENTER_ROOM_LIST_ACK EProtoId = 6
	EProtoId_ZJH_PID_AUTO_ENTER_DESK     EProtoId = 7
	EProtoId_ZJH_PID_AUTO_ENTER_DESK_ACK EProtoId = 8
	EProtoId_ZJH_PID_SEND_GAMEINFO       EProtoId = 9
	EProtoId_ZJH_PID_BC_NEWENTER         EProtoId = 10
	EProtoId_ZJH_PID_BC_OPENING          EProtoId = 11
	EProtoId_ZJH_PID_REQ_READY           EProtoId = 12
	EProtoId_ZJH_PID_BC_READY            EProtoId = 13
	EProtoId_ZJH_PID_REQ_LEAVEDESK       EProtoId = 14
	EProtoId_ZJH_PID_BC_LEAVEDESK        EProtoId = 15
	EProtoId_ZJH_PID_BC_NEXT             EProtoId = 16
	EProtoId_ZJH_PID_REQ_KANPAI          EProtoId = 17
	EProtoId_ZJH_PID_BC_KANPAI           EProtoId = 18
	EProtoId_ZJH_PID_REQ_QIPAI           EProtoId = 19
	EProtoId_ZJH_PID_BC_QIPAI            EProtoId = 20
	EProtoId_ZJH_PID_REQ_GENZHU          EProtoId = 21
	EProtoId_ZJH_PID_BC_GENZHU           EProtoId = 22
	EProtoId_ZJH_PID_REQ_FAQI_XUEPIN     EProtoId = 23
	EProtoId_ZJH_PID_BC_FAQI_XUEPIN      EProtoId = 24
	EProtoId_ZJH_PID_BC_XUEPIN_END       EProtoId = 25
	EProtoId_ZJH_PID_REQ_JIAZHU          EProtoId = 26
	EProtoId_ZJH_PID_BC_JIAZHU           EProtoId = 27
	EProtoId_ZJH_PID_REQ_BIPAI           EProtoId = 28
	EProtoId_ZJH_PID_BC_BIPAI            EProtoId = 29
	EProtoId_ZJH_PID_BC_GAME_END         EProtoId = 30
)

var EProtoId_name = map[int32]string{
	0:  "ZJH_PID_HEARTBEAT",
	1:  "ZJH_PID_QUICKCOnn",
	2:  "ZJH_PID_QUICKCONNA_ACK",
	3:  "ZJH_PID_GAME_LOGIN",
	4:  "ZJH_PID_GAME_LOGIN_ACK",
	5:  "ZJH_PID_GET_ROOM_LIST",
	6:  "ZJH_PID_ENTER_ROOM_LIST_ACK",
	7:  "ZJH_PID_AUTO_ENTER_DESK",
	8:  "ZJH_PID_AUTO_ENTER_DESK_ACK",
	9:  "ZJH_PID_SEND_GAMEINFO",
	10: "ZJH_PID_BC_NEWENTER",
	11: "ZJH_PID_BC_OPENING",
	12: "ZJH_PID_REQ_READY",
	13: "ZJH_PID_BC_READY",
	14: "ZJH_PID_REQ_LEAVEDESK",
	15: "ZJH_PID_BC_LEAVEDESK",
	16: "ZJH_PID_BC_NEXT",
	17: "ZJH_PID_REQ_KANPAI",
	18: "ZJH_PID_BC_KANPAI",
	19: "ZJH_PID_REQ_QIPAI",
	20: "ZJH_PID_BC_QIPAI",
	21: "ZJH_PID_REQ_GENZHU",
	22: "ZJH_PID_BC_GENZHU",
	23: "ZJH_PID_REQ_FAQI_XUEPIN",
	24: "ZJH_PID_BC_FAQI_XUEPIN",
	25: "ZJH_PID_BC_XUEPIN_END",
	26: "ZJH_PID_REQ_JIAZHU",
	27: "ZJH_PID_BC_JIAZHU",
	28: "ZJH_PID_REQ_BIPAI",
	29: "ZJH_PID_BC_BIPAI",
	30: "ZJH_PID_BC_GAME_END",
}
var EProtoId_value = map[string]int32{
	"ZJH_PID_HEARTBEAT":           0,
	"ZJH_PID_QUICKCOnn":           1,
	"ZJH_PID_QUICKCONNA_ACK":      2,
	"ZJH_PID_GAME_LOGIN":          3,
	"ZJH_PID_GAME_LOGIN_ACK":      4,
	"ZJH_PID_GET_ROOM_LIST":       5,
	"ZJH_PID_ENTER_ROOM_LIST_ACK": 6,
	"ZJH_PID_AUTO_ENTER_DESK":     7,
	"ZJH_PID_AUTO_ENTER_DESK_ACK": 8,
	"ZJH_PID_SEND_GAMEINFO":       9,
	"ZJH_PID_BC_NEWENTER":         10,
	"ZJH_PID_BC_OPENING":          11,
	"ZJH_PID_REQ_READY":           12,
	"ZJH_PID_BC_READY":            13,
	"ZJH_PID_REQ_LEAVEDESK":       14,
	"ZJH_PID_BC_LEAVEDESK":        15,
	"ZJH_PID_BC_NEXT":             16,
	"ZJH_PID_REQ_KANPAI":          17,
	"ZJH_PID_BC_KANPAI":           18,
	"ZJH_PID_REQ_QIPAI":           19,
	"ZJH_PID_BC_QIPAI":            20,
	"ZJH_PID_REQ_GENZHU":          21,
	"ZJH_PID_BC_GENZHU":           22,
	"ZJH_PID_REQ_FAQI_XUEPIN":     23,
	"ZJH_PID_BC_FAQI_XUEPIN":      24,
	"ZJH_PID_BC_XUEPIN_END":       25,
	"ZJH_PID_REQ_JIAZHU":          26,
	"ZJH_PID_BC_JIAZHU":           27,
	"ZJH_PID_REQ_BIPAI":           28,
	"ZJH_PID_BC_BIPAI":            29,
	"ZJH_PID_BC_GAME_END":         30,
}

func (x EProtoId) Enum() *EProtoId {
	p := new(EProtoId)
	*p = x
	return p
}
func (x EProtoId) String() string {
	return proto.EnumName(EProtoId_name, int32(x))
}
func (x *EProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EProtoId_value, data, "EProtoId")
	if err != nil {
		return err
	}
	*x = EProtoId(value)
	return nil
}
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

// 玩家当前状态
type ZjhEnumPlayerGameStatus int32

const (
	ZjhEnumPlayerGameStatus_ZJH_TEMP ZjhEnumPlayerGameStatus = 0
)

var ZjhEnumPlayerGameStatus_name = map[int32]string{
	0: "ZJH_TEMP",
}
var ZjhEnumPlayerGameStatus_value = map[string]int32{
	"ZJH_TEMP": 0,
}

func (x ZjhEnumPlayerGameStatus) Enum() *ZjhEnumPlayerGameStatus {
	p := new(ZjhEnumPlayerGameStatus)
	*p = x
	return p
}
func (x ZjhEnumPlayerGameStatus) String() string {
	return proto.EnumName(ZjhEnumPlayerGameStatus_name, int32(x))
}
func (x *ZjhEnumPlayerGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumPlayerGameStatus_value, data, "ZjhEnumPlayerGameStatus")
	if err != nil {
		return err
	}
	*x = ZjhEnumPlayerGameStatus(value)
	return nil
}
func (ZjhEnumPlayerGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

// 桌面状态
type ZjhEnumDeskState int32

const (
	ZjhEnumDeskState_DESK_IS_GAMING ZjhEnumDeskState = 1
	ZjhEnumDeskState_DESK_IS_WAIT   ZjhEnumDeskState = 2
)

var ZjhEnumDeskState_name = map[int32]string{
	1: "DESK_IS_GAMING",
	2: "DESK_IS_WAIT",
}
var ZjhEnumDeskState_value = map[string]int32{
	"DESK_IS_GAMING": 1,
	"DESK_IS_WAIT":   2,
}

func (x ZjhEnumDeskState) Enum() *ZjhEnumDeskState {
	p := new(ZjhEnumDeskState)
	*p = x
	return p
}
func (x ZjhEnumDeskState) String() string {
	return proto.EnumName(ZjhEnumDeskState_name, int32(x))
}
func (x *ZjhEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumDeskState_value, data, "ZjhEnumDeskState")
	if err != nil {
		return err
	}
	*x = ZjhEnumDeskState(value)
	return nil
}
func (ZjhEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

// 房间玩家状态
type ZjhEnumUserState int32

const (
	ZjhEnumUserState_USER_IS_GAMING ZjhEnumUserState = 1
	ZjhEnumUserState_USER_IS_STAND  ZjhEnumUserState = 2
	ZjhEnumUserState_USER_IS_SITED  ZjhEnumUserState = 3
)

var ZjhEnumUserState_name = map[int32]string{
	1: "USER_IS_GAMING",
	2: "USER_IS_STAND",
	3: "USER_IS_SITED",
}
var ZjhEnumUserState_value = map[string]int32{
	"USER_IS_GAMING": 1,
	"USER_IS_STAND":  2,
	"USER_IS_SITED":  3,
}

func (x ZjhEnumUserState) Enum() *ZjhEnumUserState {
	p := new(ZjhEnumUserState)
	*p = x
	return p
}
func (x ZjhEnumUserState) String() string {
	return proto.EnumName(ZjhEnumUserState_name, int32(x))
}
func (x *ZjhEnumUserState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumUserState_value, data, "ZjhEnumUserState")
	if err != nil {
		return err
	}
	*x = ZjhEnumUserState(value)
	return nil
}
func (ZjhEnumUserState) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

// 房间类型
type ZjhEnumRoomType int32

const (
	ZjhEnumRoomType_ROOM_TYPE_FRIEND   ZjhEnumRoomType = 1
	ZjhEnumRoomType_ROOM_TYPE_NORMAL   ZjhEnumRoomType = 2
	ZjhEnumRoomType_ROOM_TYPE_REDBLACK ZjhEnumRoomType = 3
)

var ZjhEnumRoomType_name = map[int32]string{
	1: "ROOM_TYPE_FRIEND",
	2: "ROOM_TYPE_NORMAL",
	3: "ROOM_TYPE_REDBLACK",
}
var ZjhEnumRoomType_value = map[string]int32{
	"ROOM_TYPE_FRIEND":   1,
	"ROOM_TYPE_NORMAL":   2,
	"ROOM_TYPE_REDBLACK": 3,
}

func (x ZjhEnumRoomType) Enum() *ZjhEnumRoomType {
	p := new(ZjhEnumRoomType)
	*p = x
	return p
}
func (x ZjhEnumRoomType) String() string {
	return proto.EnumName(ZjhEnumRoomType_name, int32(x))
}
func (x *ZjhEnumRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumRoomType_value, data, "ZjhEnumRoomType")
	if err != nil {
		return err
	}
	*x = ZjhEnumRoomType(value)
	return nil
}
func (ZjhEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func init() {
	proto.RegisterEnum("ddproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("ddproto.ZjhEnumPlayerGameStatus", ZjhEnumPlayerGameStatus_name, ZjhEnumPlayerGameStatus_value)
	proto.RegisterEnum("ddproto.ZjhEnumDeskState", ZjhEnumDeskState_name, ZjhEnumDeskState_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserState", ZjhEnumUserState_name, ZjhEnumUserState_value)
	proto.RegisterEnum("ddproto.ZjhEnumRoomType", ZjhEnumRoomType_name, ZjhEnumRoomType_value)
}

var fileDescriptor16 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x53, 0x4b, 0x73, 0xda, 0x30,
	0x10, 0x0e, 0xd0, 0x26, 0x54, 0x25, 0x64, 0x11, 0xaf, 0x10, 0xfa, 0x38, 0x97, 0x43, 0x7f, 0x40,
	0x6f, 0x02, 0x2f, 0xa0, 0x00, 0xb2, 0xb1, 0xe5, 0x92, 0xe4, 0xa2, 0xa1, 0x83, 0x67, 0x3a, 0x6d,
	0x79, 0x0c, 0x8f, 0x43, 0xfa, 0x4f, 0xfa, 0x6f, 0x6b, 0xd9, 0x18, 0x84, 0x0e, 0x9e, 0xf1, 0x7c,
	0x8f, 0xd5, 0xb7, 0xd2, 0x2e, 0x29, 0xff, 0xfd, 0xf5, 0x53, 0xfd, 0x98, 0xef, 0xa2, 0xaf, 0x9b,
	0xed, 0x7a, 0xbf, 0xa6, 0x37, 0x8b, 0x45, 0xf2, 0xd3, 0xf9, 0x77, 0x4d, 0x8a, 0xe8, 0xe9, 0x5f,
	0xbe, 0xa0, 0x75, 0x52, 0x79, 0x79, 0x1c, 0x2a, 0x8f, 0x3b, 0x6a, 0x88, 0xcc, 0x97, 0x5d, 0x64,
	0x12, 0xae, 0x4c, 0x78, 0x1a, 0xf2, 0xde, 0xa8, 0xe7, 0xae, 0x56, 0x90, 0xa3, 0x0f, 0xa4, 0x61,
	0xc1, 0x42, 0x30, 0xc5, 0x7a, 0x23, 0xc8, 0xd3, 0x06, 0xa1, 0x19, 0x37, 0x60, 0x13, 0x54, 0x63,
	0x77, 0xc0, 0x05, 0x14, 0x4c, 0xcf, 0x19, 0x4f, 0x3c, 0x6f, 0x68, 0x8b, 0xd4, 0x4f, 0x1c, 0x4a,
	0xe5, 0xbb, 0xee, 0x44, 0x8d, 0x79, 0x20, 0xe1, 0x2d, 0xfd, 0x4c, 0xda, 0x19, 0x85, 0x42, 0xa2,
	0x7f, 0x26, 0x13, 0xef, 0x35, 0x6d, 0x93, 0x66, 0x26, 0x60, 0xa1, 0x74, 0x8f, 0x2a, 0x07, 0x83,
	0x11, 0xdc, 0x98, 0x6e, 0x8b, 0x4c, 0xdc, 0x45, 0xf3, 0xe4, 0x00, 0x45, 0x1a, 0x8d, 0x8b, 0xbe,
	0x0b, 0xef, 0x68, 0x93, 0x54, 0x33, 0xaa, 0xdb, 0x53, 0x02, 0x67, 0x89, 0x19, 0x88, 0xd9, 0x61,
	0x4c, 0xb8, 0x1e, 0x0a, 0x2e, 0x06, 0xf0, 0xde, 0xbc, 0x2c, 0x1f, 0xa7, 0xf1, 0xc7, 0x9c, 0x67,
	0x28, 0xd1, 0x1a, 0x01, 0x43, 0x9e, 0xa2, 0xb7, 0xe6, 0xc1, 0x5a, 0x3c, 0x46, 0xf6, 0x1d, 0x93,
	0xd0, 0x65, 0x7a, 0x4f, 0x6a, 0x86, 0xe1, 0xcc, 0xdc, 0xd1, 0x2a, 0xb9, 0xbb, 0x88, 0xf4, 0x24,
	0x01, 0xcc, 0x38, 0xba, 0xd2, 0x88, 0x09, 0x8f, 0x71, 0xa8, 0x98, 0x71, 0x62, 0xf1, 0x11, 0xa6,
	0x76, 0xca, 0x29, 0xd7, 0x70, 0xd5, 0x4a, 0x99, 0xa2, 0x35, 0xbb, 0xf6, 0x00, 0xc5, 0xcb, 0x30,
	0x84, 0xba, 0x55, 0xfb, 0x08, 0x37, 0xcc, 0xb7, 0xd0, 0xf2, 0x3e, 0x9b, 0x72, 0xf5, 0x14, 0xa2,
	0x17, 0x0f, 0x40, 0xd3, 0x1c, 0x80, 0xd8, 0x63, 0x72, 0xf7, 0xe6, 0x6d, 0xc4, 0x5c, 0x0a, 0xc7,
	0x8f, 0xe5, 0x40, 0xcb, 0x8e, 0xf0, 0xc8, 0x99, 0x3e, 0xeb, 0xc1, 0x8a, 0x70, 0x84, 0xdb, 0x76,
	0x7b, 0xdd, 0xa4, 0x91, 0x0f, 0x56, 0x7b, 0x29, 0xfa, 0xd1, 0x7a, 0xe2, 0x64, 0x2c, 0xf5, 0xa1,
	0x9f, 0x3a, 0x5f, 0x48, 0x4b, 0xaf, 0x4d, 0xb4, 0x3a, 0x2c, 0xd5, 0xe6, 0xcf, 0xfc, 0x35, 0xda,
	0x0e, 0xe6, 0xcb, 0x28, 0xd8, 0xcf, 0xf7, 0x87, 0x1d, 0x2d, 0x91, 0xa2, 0x76, 0x49, 0x9c, 0x78,
	0x70, 0xd5, 0xf9, 0x46, 0xe8, 0x49, 0xba, 0x88, 0x76, 0xbf, 0xb5, 0x28, 0xa2, 0x94, 0x94, 0x93,
	0x29, 0xe3, 0x81, 0x2e, 0xab, 0xe7, 0x23, 0x47, 0x81, 0x94, 0x32, 0x6c, 0xc6, 0xb8, 0x84, 0x7c,
	0x47, 0x18, 0xde, 0xc3, 0x2e, 0xda, 0x9e, 0xbc, 0x61, 0x10, 0x8f, 0xa9, 0xe9, 0xad, 0x90, 0xdb,
	0x0c, 0x0b, 0x24, 0x8b, 0x33, 0xe6, 0x2f, 0x20, 0x2e, 0xd1, 0x81, 0x42, 0x67, 0x46, 0x2a, 0xa7,
	0x7a, 0xdb, 0xf5, 0x7a, 0x29, 0x5f, 0x37, 0x91, 0x6e, 0x3d, 0xd9, 0x19, 0xf9, 0xec, 0xa1, 0xea,
	0xfb, 0x5c, 0x77, 0x98, 0xbb, 0x44, 0x85, 0xeb, 0x4f, 0xd8, 0x38, 0x5d, 0xde, 0x33, 0xea, 0xa3,
	0xd3, 0x1d, 0xeb, 0x35, 0x29, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x45, 0x2f, 0x3a, 0x45,
	0x04, 0x00, 0x00,
}

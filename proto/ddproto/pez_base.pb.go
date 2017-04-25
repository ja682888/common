// Code generated by protoc-gen-go.
// source: pez_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

type PezEnumProtoId int32

const (
	PezEnumProtoId_PEZ_PID_HEARTBEAT      PezEnumProtoId = 0
	PezEnumProtoId_PEZ_PID_CREATEROOM     PezEnumProtoId = 1
	PezEnumProtoId_PEZ_PID_CREATEROOM_ACK PezEnumProtoId = 2
	PezEnumProtoId_PEZ_PID_ENTER_ROOM     PezEnumProtoId = 3
	PezEnumProtoId_PEZ_PID_ENTER_ROOM_ACK PezEnumProtoId = 4
	PezEnumProtoId_PEZ_PID_SEND_GAMEINFO  PezEnumProtoId = 5
	PezEnumProtoId_PEZ_PID_READY          PezEnumProtoId = 6
	PezEnumProtoId_PEZ_PID_READY_ACK      PezEnumProtoId = 7
	PezEnumProtoId_PEZ_PID_OPENING        PezEnumProtoId = 8
	PezEnumProtoId_PEZ_PID_DEALCARDS      PezEnumProtoId = 9
	PezEnumProtoId_PEZ_PID_BET            PezEnumProtoId = 10
	PezEnumProtoId_PEZ_PID_BET_ACK        PezEnumProtoId = 11
)

var PezEnumProtoId_name = map[int32]string{
	0:  "PEZ_PID_HEARTBEAT",
	1:  "PEZ_PID_CREATEROOM",
	2:  "PEZ_PID_CREATEROOM_ACK",
	3:  "PEZ_PID_ENTER_ROOM",
	4:  "PEZ_PID_ENTER_ROOM_ACK",
	5:  "PEZ_PID_SEND_GAMEINFO",
	6:  "PEZ_PID_READY",
	7:  "PEZ_PID_READY_ACK",
	8:  "PEZ_PID_OPENING",
	9:  "PEZ_PID_DEALCARDS",
	10: "PEZ_PID_BET",
	11: "PEZ_PID_BET_ACK",
}
var PezEnumProtoId_value = map[string]int32{
	"PEZ_PID_HEARTBEAT":      0,
	"PEZ_PID_CREATEROOM":     1,
	"PEZ_PID_CREATEROOM_ACK": 2,
	"PEZ_PID_ENTER_ROOM":     3,
	"PEZ_PID_ENTER_ROOM_ACK": 4,
	"PEZ_PID_SEND_GAMEINFO":  5,
	"PEZ_PID_READY":          6,
	"PEZ_PID_READY_ACK":      7,
	"PEZ_PID_OPENING":        8,
	"PEZ_PID_DEALCARDS":      9,
	"PEZ_PID_BET":            10,
	"PEZ_PID_BET_ACK":        11,
}

func (x PezEnumProtoId) Enum() *PezEnumProtoId {
	p := new(PezEnumProtoId)
	*p = x
	return p
}
func (x PezEnumProtoId) String() string {
	return proto.EnumName(PezEnumProtoId_name, int32(x))
}
func (x *PezEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnumProtoId_value, data, "PezEnumProtoId")
	if err != nil {
		return err
	}
	*x = PezEnumProtoId(value)
	return nil
}
func (PezEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{0} }

type PezEnum_ErrorCode int32

const (
	PezEnum_ErrorCode_PEZ_EC_SUCCESS PezEnum_ErrorCode = 0
	// -101   -200	游戏异常
	PezEnum_ErrorCode_PEZ_EC_CREATE_DESK_DIAMOND_NOTENOUGH PezEnum_ErrorCode = -101
	PezEnum_ErrorCode_PEZ_EC_CREATE_DESK_USER_NOTFOUND     PezEnum_ErrorCode = -102
	PezEnum_ErrorCode_PEZ_EC_INTO_DESK_NOTFOUND            PezEnum_ErrorCode = -103
	PezEnum_ErrorCode_PEZ_EC_GAME_READY_REPEAT             PezEnum_ErrorCode = -110
)

var PezEnum_ErrorCode_name = map[int32]string{
	0:    "PEZ_EC_SUCCESS",
	-101: "PEZ_EC_CREATE_DESK_DIAMOND_NOTENOUGH",
	-102: "PEZ_EC_CREATE_DESK_USER_NOTFOUND",
	-103: "PEZ_EC_INTO_DESK_NOTFOUND",
	-110: "PEZ_EC_GAME_READY_REPEAT",
}
var PezEnum_ErrorCode_value = map[string]int32{
	"PEZ_EC_SUCCESS":                       0,
	"PEZ_EC_CREATE_DESK_DIAMOND_NOTENOUGH": -101,
	"PEZ_EC_CREATE_DESK_USER_NOTFOUND":     -102,
	"PEZ_EC_INTO_DESK_NOTFOUND":            -103,
	"PEZ_EC_GAME_READY_REPEAT":             -110,
}

func (x PezEnum_ErrorCode) Enum() *PezEnum_ErrorCode {
	p := new(PezEnum_ErrorCode)
	*p = x
	return p
}
func (x PezEnum_ErrorCode) String() string {
	return proto.EnumName(PezEnum_ErrorCode_name, int32(x))
}
func (x *PezEnum_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnum_ErrorCode_value, data, "PezEnum_ErrorCode")
	if err != nil {
		return err
	}
	*x = PezEnum_ErrorCode(value)
	return nil
}
func (PezEnum_ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{1} }

type PezEnum_Option int32

const (
	PezEnum_Option_LUN_BANK  PezEnum_Option = 1
	PezEnum_Option_DING_BANK PezEnum_Option = 2
	PezEnum_Option_TONG_BI   PezEnum_Option = 3
)

var PezEnum_Option_name = map[int32]string{
	1: "LUN_BANK",
	2: "DING_BANK",
	3: "TONG_BI",
}
var PezEnum_Option_value = map[string]int32{
	"LUN_BANK":  1,
	"DING_BANK": 2,
	"TONG_BI":   3,
}

func (x PezEnum_Option) Enum() *PezEnum_Option {
	p := new(PezEnum_Option)
	*p = x
	return p
}
func (x PezEnum_Option) String() string {
	return proto.EnumName(PezEnum_Option_name, int32(x))
}
func (x *PezEnum_Option) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnum_Option_value, data, "PezEnum_Option")
	if err != nil {
		return err
	}
	*x = PezEnum_Option(value)
	return nil
}
func (PezEnum_Option) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{2} }

// 房间类型信息：包含房间类型和对应的局数、封顶、玩法等信息
// 房间类型枚举
type Pez_RoomType int32

const (
	Pez_RoomType_roomType_pingerzhang Pez_RoomType = 0
)

var Pez_RoomType_name = map[int32]string{
	0: "roomType_pingerzhang",
}
var Pez_RoomType_value = map[string]int32{
	"roomType_pingerzhang": 0,
}

func (x Pez_RoomType) Enum() *Pez_RoomType {
	p := new(Pez_RoomType)
	*p = x
	return p
}
func (x Pez_RoomType) String() string {
	return proto.EnumName(Pez_RoomType_name, int32(x))
}
func (x *Pez_RoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Pez_RoomType_value, data, "Pez_RoomType")
	if err != nil {
		return err
	}
	*x = Pez_RoomType(value)
	return nil
}
func (Pez_RoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{3} }

// 麻将花色
type PezEnumMjColor int32

const (
	PezEnumMjColor_PEZ_TIAO PezEnumMjColor = 1
	PezEnumMjColor_PEZ_TONG PezEnumMjColor = 2
)

var PezEnumMjColor_name = map[int32]string{
	1: "PEZ_TIAO",
	2: "PEZ_TONG",
}
var PezEnumMjColor_value = map[string]int32{
	"PEZ_TIAO": 1,
	"PEZ_TONG": 2,
}

func (x PezEnumMjColor) Enum() *PezEnumMjColor {
	p := new(PezEnumMjColor)
	*p = x
	return p
}
func (x PezEnumMjColor) String() string {
	return proto.EnumName(PezEnumMjColor_name, int32(x))
}
func (x *PezEnumMjColor) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnumMjColor_value, data, "PezEnumMjColor")
	if err != nil {
		return err
	}
	*x = PezEnumMjColor(value)
	return nil
}
func (PezEnumMjColor) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{4} }

type PezEnum_PaiType int32

const (
	PezEnum_PaiType_PEZ_LAIZI PezEnum_PaiType = 0
	PezEnum_PaiType_PEZ_DUIZI PezEnum_PaiType = 1
	PezEnum_PaiType_PEZ_DUIJI PezEnum_PaiType = 2
)

var PezEnum_PaiType_name = map[int32]string{
	0: "PEZ_LAIZI",
	1: "PEZ_DUIZI",
	2: "PEZ_DUIJI",
}
var PezEnum_PaiType_value = map[string]int32{
	"PEZ_LAIZI": 0,
	"PEZ_DUIZI": 1,
	"PEZ_DUIJI": 2,
}

func (x PezEnum_PaiType) Enum() *PezEnum_PaiType {
	p := new(PezEnum_PaiType)
	*p = x
	return p
}
func (x PezEnum_PaiType) String() string {
	return proto.EnumName(PezEnum_PaiType_name, int32(x))
}
func (x *PezEnum_PaiType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnum_PaiType_value, data, "PezEnum_PaiType")
	if err != nil {
		return err
	}
	*x = PezEnum_PaiType(value)
	return nil
}
func (PezEnum_PaiType) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{5} }

type PezEnum_Base int32

const (
	PezEnum_Base_PEZ_LOW  PezEnum_Base = 0
	PezEnum_Base_PEZ_HIGH PezEnum_Base = 1
)

var PezEnum_Base_name = map[int32]string{
	0: "PEZ_LOW",
	1: "PEZ_HIGH",
}
var PezEnum_Base_value = map[string]int32{
	"PEZ_LOW":  0,
	"PEZ_HIGH": 1,
}

func (x PezEnum_Base) Enum() *PezEnum_Base {
	p := new(PezEnum_Base)
	*p = x
	return p
}
func (x PezEnum_Base) String() string {
	return proto.EnumName(PezEnum_Base_name, int32(x))
}
func (x *PezEnum_Base) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnum_Base_value, data, "PezEnum_Base")
	if err != nil {
		return err
	}
	*x = PezEnum_Base(value)
	return nil
}
func (PezEnum_Base) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{6} }

type PezEnum_UserGameStatus int32

const (
	PezEnum_UserGameStatus_PEZ_U_INIT    PezEnum_UserGameStatus = 0
	PezEnum_UserGameStatus_PEZ_U_READY   PezEnum_UserGameStatus = 1
	PezEnum_UserGameStatus_PEZ_U_BET     PezEnum_UserGameStatus = 2
	PezEnum_UserGameStatus_PEZ_U_PLAYING PezEnum_UserGameStatus = 3
	PezEnum_UserGameStatus_PEZ_U_FINISH  PezEnum_UserGameStatus = 4
)

var PezEnum_UserGameStatus_name = map[int32]string{
	0: "PEZ_U_INIT",
	1: "PEZ_U_READY",
	2: "PEZ_U_BET",
	3: "PEZ_U_PLAYING",
	4: "PEZ_U_FINISH",
}
var PezEnum_UserGameStatus_value = map[string]int32{
	"PEZ_U_INIT":    0,
	"PEZ_U_READY":   1,
	"PEZ_U_BET":     2,
	"PEZ_U_PLAYING": 3,
	"PEZ_U_FINISH":  4,
}

func (x PezEnum_UserGameStatus) Enum() *PezEnum_UserGameStatus {
	p := new(PezEnum_UserGameStatus)
	*p = x
	return p
}
func (x PezEnum_UserGameStatus) String() string {
	return proto.EnumName(PezEnum_UserGameStatus_name, int32(x))
}
func (x *PezEnum_UserGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnum_UserGameStatus_value, data, "PezEnum_UserGameStatus")
	if err != nil {
		return err
	}
	*x = PezEnum_UserGameStatus(value)
	return nil
}
func (PezEnum_UserGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{7} }

type PezEnum_DeskGameStatus int32

const (
	PezEnum_DeskGameStatus_PEZ_INIT    PezEnum_DeskGameStatus = 0
	PezEnum_DeskGameStatus_PEZ_FAPAI   PezEnum_DeskGameStatus = 1
	PezEnum_DeskGameStatus_PEZ_BET     PezEnum_DeskGameStatus = 2
	PezEnum_DeskGameStatus_PEZ_OPENPAI PezEnum_DeskGameStatus = 3
	PezEnum_DeskGameStatus_PEZ_PLAYING PezEnum_DeskGameStatus = 4
	PezEnum_DeskGameStatus_PEZ_FINISH  PezEnum_DeskGameStatus = 5
)

var PezEnum_DeskGameStatus_name = map[int32]string{
	0: "PEZ_INIT",
	1: "PEZ_FAPAI",
	2: "PEZ_BET",
	3: "PEZ_OPENPAI",
	4: "PEZ_PLAYING",
	5: "PEZ_FINISH",
}
var PezEnum_DeskGameStatus_value = map[string]int32{
	"PEZ_INIT":    0,
	"PEZ_FAPAI":   1,
	"PEZ_BET":     2,
	"PEZ_OPENPAI": 3,
	"PEZ_PLAYING": 4,
	"PEZ_FINISH":  5,
}

func (x PezEnum_DeskGameStatus) Enum() *PezEnum_DeskGameStatus {
	p := new(PezEnum_DeskGameStatus)
	*p = x
	return p
}
func (x PezEnum_DeskGameStatus) String() string {
	return proto.EnumName(PezEnum_DeskGameStatus_name, int32(x))
}
func (x *PezEnum_DeskGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnum_DeskGameStatus_value, data, "PezEnum_DeskGameStatus")
	if err != nil {
		return err
	}
	*x = PezEnum_DeskGameStatus(value)
	return nil
}
func (PezEnum_DeskGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor34, []int{8} }

// 麻将牌
type PezBase_PaiInfo struct {
	Type             *int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Value            *int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Id               *int32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PezBase_PaiInfo) Reset()                    { *m = PezBase_PaiInfo{} }
func (m *PezBase_PaiInfo) String() string            { return proto.CompactTextString(m) }
func (*PezBase_PaiInfo) ProtoMessage()               {}
func (*PezBase_PaiInfo) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{0} }

func (m *PezBase_PaiInfo) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *PezBase_PaiInfo) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *PezBase_PaiInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

// 拼二张的配置
type PezBase_PlayConf struct {
	PlayerCount      *int32 `protobuf:"varint,1,opt,name=playerCount" json:"playerCount,omitempty"`
	IgnoreBank       *bool  `protobuf:"varint,2,opt,name=ignoreBank" json:"ignoreBank,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PezBase_PlayConf) Reset()                    { *m = PezBase_PlayConf{} }
func (m *PezBase_PlayConf) String() string            { return proto.CompactTextString(m) }
func (*PezBase_PlayConf) ProtoMessage()               {}
func (*PezBase_PlayConf) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{1} }

func (m *PezBase_PlayConf) GetPlayerCount() int32 {
	if m != nil && m.PlayerCount != nil {
		return *m.PlayerCount
	}
	return 0
}

func (m *PezBase_PlayConf) GetIgnoreBank() bool {
	if m != nil && m.IgnoreBank != nil {
		return *m.IgnoreBank
	}
	return false
}

type PezBase_RoomTypeInfo struct {
	OwnerID          *uint32         `protobuf:"varint,1,opt,name=ownerID" json:"ownerID,omitempty"`
	BoardsCout       *int64          `protobuf:"varint,2,opt,name=boardsCout" json:"boardsCout,omitempty"`
	PlayOption       *PezEnum_Option `protobuf:"varint,3,opt,name=playOption,enum=ddproto.PezEnum_Option" json:"playOption,omitempty"`
	BaseValue        *PezEnum_Base   `protobuf:"varint,4,opt,name=baseValue,enum=ddproto.PezEnum_Base" json:"baseValue,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PezBase_RoomTypeInfo) Reset()                    { *m = PezBase_RoomTypeInfo{} }
func (m *PezBase_RoomTypeInfo) String() string            { return proto.CompactTextString(m) }
func (*PezBase_RoomTypeInfo) ProtoMessage()               {}
func (*PezBase_RoomTypeInfo) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{2} }

func (m *PezBase_RoomTypeInfo) GetOwnerID() uint32 {
	if m != nil && m.OwnerID != nil {
		return *m.OwnerID
	}
	return 0
}

func (m *PezBase_RoomTypeInfo) GetBoardsCout() int64 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *PezBase_RoomTypeInfo) GetPlayOption() PezEnum_Option {
	if m != nil && m.PlayOption != nil {
		return *m.PlayOption
	}
	return PezEnum_Option_LUN_BANK
}

func (m *PezBase_RoomTypeInfo) GetBaseValue() PezEnum_Base {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return PezEnum_Base_PEZ_LOW
}

type PezBase_PaiValue struct {
	Value            *int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PezBase_PaiValue) Reset()                    { *m = PezBase_PaiValue{} }
func (m *PezBase_PaiValue) String() string            { return proto.CompactTextString(m) }
func (*PezBase_PaiValue) ProtoMessage()               {}
func (*PezBase_PaiValue) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{3} }

func (m *PezBase_PaiValue) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

// 手里的牌
type PezBase_PlayerCard struct {
	HandCard         []*PezBase_PaiInfo `protobuf:"bytes,1,rep,name=handCard" json:"handCard,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	HandCardCount    *int32             `protobuf:"varint,3,opt,name=handCardCount" json:"handCardCount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PezBase_PlayerCard) Reset()                    { *m = PezBase_PlayerCard{} }
func (m *PezBase_PlayerCard) String() string            { return proto.CompactTextString(m) }
func (*PezBase_PlayerCard) ProtoMessage()               {}
func (*PezBase_PlayerCard) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{4} }

func (m *PezBase_PlayerCard) GetHandCard() []*PezBase_PaiInfo {
	if m != nil {
		return m.HandCard
	}
	return nil
}

func (m *PezBase_PlayerCard) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezBase_PlayerCard) GetHandCardCount() int32 {
	if m != nil && m.HandCardCount != nil {
		return *m.HandCardCount
	}
	return 0
}

type PezBase_PlayerInfo struct {
	IsBanker         *bool               `protobuf:"varint,1,opt,name=isBanker" json:"isBanker,omitempty"`
	PlayerCard       *PezBase_PlayerCard `protobuf:"bytes,2,opt,name=playerCard" json:"playerCard,omitempty"`
	HandPaiType      *PezEnum_PaiType    `protobuf:"varint,3,opt,name=HandPaiType,enum=ddproto.PezEnum_PaiType" json:"HandPaiType,omitempty"`
	Coin             *int64              `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	BetNum           *int32              `protobuf:"varint,5,opt,name=betNum" json:"betNum,omitempty"`
	NickName         *string             `protobuf:"bytes,6,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32              `protobuf:"varint,7,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32             `protobuf:"varint,8,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool               `protobuf:"varint,9,opt,name=isOwner" json:"isOwner,omitempty"`
	IsReady          *int32              `protobuf:"varint,10,opt,name=isReady" json:"isReady,omitempty"`
	IsBet            *int32              `protobuf:"varint,11,opt,name=isBet" json:"isBet,omitempty"`
	WxInfo           *WeixinInfo         `protobuf:"bytes,12,opt,name=wxInfo" json:"wxInfo,omitempty"`
	GameStatus       *int32              `protobuf:"varint,13,opt,name=GameStatus" json:"GameStatus,omitempty"`
	Ip               *string             `protobuf:"bytes,14,opt,name=ip" json:"ip,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *PezBase_PlayerInfo) Reset()                    { *m = PezBase_PlayerInfo{} }
func (m *PezBase_PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PezBase_PlayerInfo) ProtoMessage()               {}
func (*PezBase_PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{5} }

func (m *PezBase_PlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *PezBase_PlayerInfo) GetPlayerCard() *PezBase_PlayerCard {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *PezBase_PlayerInfo) GetHandPaiType() PezEnum_PaiType {
	if m != nil && m.HandPaiType != nil {
		return *m.HandPaiType
	}
	return PezEnum_PaiType_PEZ_LAIZI
}

func (m *PezBase_PlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PezBase_PlayerInfo) GetBetNum() int32 {
	if m != nil && m.BetNum != nil {
		return *m.BetNum
	}
	return 0
}

func (m *PezBase_PlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PezBase_PlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PezBase_PlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezBase_PlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PezBase_PlayerInfo) GetIsReady() int32 {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return 0
}

func (m *PezBase_PlayerInfo) GetIsBet() int32 {
	if m != nil && m.IsBet != nil {
		return *m.IsBet
	}
	return 0
}

func (m *PezBase_PlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PezBase_PlayerInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PezBase_PlayerInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

type PezBase_DeskGameInfo struct {
	GameStatus       *int32                `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	RoomTypeInfo     *PezBase_RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	PlayerNum        *int32                `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActionTime       *int32                `protobuf:"varint,4,opt,name=actionTime" json:"actionTime,omitempty"`
	DelayTime        *int32                `protobuf:"varint,5,opt,name=delayTime" json:"delayTime,omitempty"`
	NInitActionTime  *int32                `protobuf:"varint,6,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	NInitDelayTime   *int32                `protobuf:"varint,7,opt,name=nInitDelayTime" json:"nInitDelayTime,omitempty"`
	InitRoomCoin     *int64                `protobuf:"varint,8,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32                `protobuf:"varint,9,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32                `protobuf:"varint,10,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string               `protobuf:"bytes,11,opt,name=roomNumber" json:"roomNumber,omitempty"`
	RemainCards      *int32                `protobuf:"varint,12,opt,name=remainCards" json:"remainCards,omitempty"`
	Banker           *uint32               `protobuf:"varint,13,opt,name=banker" json:"banker,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PezBase_DeskGameInfo) Reset()                    { *m = PezBase_DeskGameInfo{} }
func (m *PezBase_DeskGameInfo) String() string            { return proto.CompactTextString(m) }
func (*PezBase_DeskGameInfo) ProtoMessage()               {}
func (*PezBase_DeskGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{6} }

func (m *PezBase_DeskGameInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetRoomTypeInfo() *PezBase_RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *PezBase_DeskGameInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetActionTime() int32 {
	if m != nil && m.ActionTime != nil {
		return *m.ActionTime
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetDelayTime() int32 {
	if m != nil && m.DelayTime != nil {
		return *m.DelayTime
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetNInitDelayTime() int32 {
	if m != nil && m.NInitDelayTime != nil {
		return *m.NInitDelayTime
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *PezBase_DeskGameInfo) GetRemainCards() int32 {
	if m != nil && m.RemainCards != nil {
		return *m.RemainCards
	}
	return 0
}

func (m *PezBase_DeskGameInfo) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func init() {
	proto.RegisterType((*PezBase_PaiInfo)(nil), "ddproto.pez_base_PaiInfo")
	proto.RegisterType((*PezBase_PlayConf)(nil), "ddproto.pez_base_PlayConf")
	proto.RegisterType((*PezBase_RoomTypeInfo)(nil), "ddproto.pez_base_RoomTypeInfo")
	proto.RegisterType((*PezBase_PaiValue)(nil), "ddproto.pez_base_PaiValue")
	proto.RegisterType((*PezBase_PlayerCard)(nil), "ddproto.pez_base_PlayerCard")
	proto.RegisterType((*PezBase_PlayerInfo)(nil), "ddproto.pez_base_PlayerInfo")
	proto.RegisterType((*PezBase_DeskGameInfo)(nil), "ddproto.pez_base_DeskGameInfo")
	proto.RegisterEnum("ddproto.PezEnumProtoId", PezEnumProtoId_name, PezEnumProtoId_value)
	proto.RegisterEnum("ddproto.PezEnum_ErrorCode", PezEnum_ErrorCode_name, PezEnum_ErrorCode_value)
	proto.RegisterEnum("ddproto.PezEnum_Option", PezEnum_Option_name, PezEnum_Option_value)
	proto.RegisterEnum("ddproto.Pez_RoomType", Pez_RoomType_name, Pez_RoomType_value)
	proto.RegisterEnum("ddproto.PezEnumMjColor", PezEnumMjColor_name, PezEnumMjColor_value)
	proto.RegisterEnum("ddproto.PezEnum_PaiType", PezEnum_PaiType_name, PezEnum_PaiType_value)
	proto.RegisterEnum("ddproto.PezEnum_Base", PezEnum_Base_name, PezEnum_Base_value)
	proto.RegisterEnum("ddproto.PezEnum_UserGameStatus", PezEnum_UserGameStatus_name, PezEnum_UserGameStatus_value)
	proto.RegisterEnum("ddproto.PezEnum_DeskGameStatus", PezEnum_DeskGameStatus_name, PezEnum_DeskGameStatus_value)
}

var fileDescriptor34 = []byte{
	// 1205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0x36, 0x25, 0xcb, 0x96, 0x46, 0x96, 0xbd, 0x59, 0x27, 0xfe, 0xd1, 0x41, 0x10, 0x18, 0x42,
	0x7e, 0x81, 0xab, 0xa2, 0x06, 0x1a, 0xb4, 0x40, 0x81, 0x14, 0x28, 0x28, 0x92, 0x96, 0xd9, 0x38,
	0xa4, 0xb0, 0x22, 0x1b, 0x24, 0x17, 0x82, 0x36, 0x37, 0x09, 0x1b, 0x89, 0x14, 0x28, 0xaa, 0x89,
	0xf3, 0x08, 0x3d, 0xf6, 0xd6, 0xf6, 0x05, 0xfa, 0x1a, 0x3d, 0xf4, 0x21, 0xfa, 0x20, 0x3d, 0xb7,
	0x98, 0x5d, 0x92, 0x5a, 0x59, 0xd1, 0x89, 0xf3, 0xed, 0x37, 0x7f, 0xf6, 0x9b, 0x99, 0x15, 0xec,
	0xcf, 0xf9, 0xc7, 0xf0, 0x2a, 0x5a, 0xf0, 0xb3, 0x79, 0x9e, 0x15, 0x19, 0xdd, 0x8d, 0x63, 0xf1,
	0x71, 0xff, 0xf0, 0x3a, 0x9b, 0xcd, 0xb2, 0x34, 0xbc, 0x9e, 0x26, 0x3c, 0x2d, 0xe4, 0x69, 0xff,
	0x12, 0x48, 0xc5, 0x0f, 0xc7, 0x51, 0xe2, 0xa4, 0xaf, 0x33, 0x4a, 0x61, 0xbb, 0xb8, 0x99, 0x73,
	0x5d, 0x3b, 0xd1, 0x4e, 0x5b, 0x4c, 0x7c, 0xd3, 0xbb, 0xd0, 0xfa, 0x29, 0x9a, 0x2e, 0xb9, 0xde,
	0x10, 0xa0, 0x34, 0xe8, 0x3e, 0x34, 0x92, 0x58, 0x6f, 0x0a, 0xa8, 0x91, 0xc4, 0xfd, 0x00, 0xee,
	0xac, 0xa2, 0x4d, 0xa3, 0x1b, 0x33, 0x4b, 0x5f, 0xd3, 0x13, 0xe8, 0xce, 0xa7, 0xd1, 0x0d, 0xcf,
	0xcd, 0x6c, 0x99, 0x16, 0x65, 0x54, 0x15, 0xa2, 0x0f, 0x01, 0x92, 0x37, 0x69, 0x96, 0xf3, 0x61,
	0x94, 0xbe, 0x13, 0x19, 0xda, 0x4c, 0x41, 0xfa, 0x7f, 0x6a, 0x70, 0xaf, 0x8e, 0xcb, 0xb2, 0x6c,
	0xe6, 0xdf, 0xcc, 0xb9, 0x28, 0x55, 0x87, 0xdd, 0xec, 0x7d, 0xca, 0x73, 0xc7, 0x12, 0x71, 0x7b,
	0xac, 0x32, 0x31, 0xe6, 0x55, 0x16, 0xe5, 0xf1, 0xc2, 0xcc, 0x96, 0x85, 0x88, 0xd9, 0x64, 0x0a,
	0x42, 0xbf, 0x01, 0xc0, 0x12, 0xbc, 0x79, 0x91, 0x64, 0xa9, 0xb8, 0xc2, 0xfe, 0x13, 0xfd, 0xac,
	0xd4, 0xea, 0x0c, 0xb3, 0xf1, 0x74, 0x39, 0x0b, 0xe5, 0x39, 0x53, 0xb8, 0xf4, 0x2b, 0xe8, 0x60,
	0x21, 0x3f, 0x08, 0x39, 0xb6, 0x85, 0xe3, 0xd1, 0xa6, 0xe3, 0x30, 0x5a, 0x70, 0xb6, 0x22, 0xf6,
	0x3f, 0x53, 0xa5, 0x89, 0x12, 0x01, 0xae, 0x54, 0xd5, 0x14, 0x55, 0xfb, 0x3f, 0x6b, 0x70, 0xb8,
	0x26, 0x23, 0xcf, 0xcd, 0x28, 0x8f, 0xe9, 0xd7, 0xd0, 0x7e, 0x1b, 0xa5, 0x31, 0x7e, 0xeb, 0xda,
	0x49, 0xf3, 0xb4, 0xfb, 0xe4, 0x78, 0x2d, 0xaf, 0xda, 0x44, 0x56, 0x53, 0xe9, 0x11, 0xec, 0x04,
	0x0b, 0x9e, 0x3b, 0xb1, 0x50, 0xa1, 0xc7, 0x4a, 0x8b, 0x3e, 0x82, 0x5e, 0xc5, 0x91, 0x9d, 0x91,
	0x7d, 0x5c, 0x07, 0xfb, 0x7f, 0x35, 0x37, 0x8a, 0x11, 0xca, 0xdf, 0x87, 0x76, 0xb2, 0xc0, 0xee,
	0xf0, 0x5c, 0x54, 0xdf, 0x66, 0xb5, 0x4d, 0xbf, 0x95, 0xda, 0xca, 0xb2, 0x45, 0xd6, 0xee, 0x93,
	0x07, 0x9f, 0x28, 0xb5, 0xe6, 0x30, 0x85, 0x4f, 0x9f, 0x42, 0xf7, 0x22, 0x4a, 0xe3, 0x71, 0x94,
	0x60, 0x9b, 0xcb, 0xd6, 0x1c, 0x6f, 0x2a, 0x5c, 0x12, 0x98, 0xca, 0xc6, 0xd9, 0xbd, 0xce, 0x92,
	0x54, 0xf4, 0xa5, 0xc9, 0xc4, 0x37, 0x0a, 0x70, 0xc5, 0x0b, 0x77, 0x39, 0xd3, 0x5b, 0xe2, 0x86,
	0xa5, 0x85, 0x57, 0x48, 0x93, 0xeb, 0x77, 0x6e, 0x34, 0xe3, 0xfa, 0xce, 0x89, 0x76, 0xda, 0x61,
	0xb5, 0x4d, 0x09, 0x34, 0x17, 0xfc, 0x83, 0xbe, 0x2b, 0x1c, 0xf0, 0x13, 0xa3, 0x2c, 0xa5, 0x8c,
	0x6d, 0x29, 0xa3, 0xb4, 0x70, 0x04, 0x93, 0x85, 0x87, 0x53, 0xa7, 0x77, 0x84, 0x0e, 0x95, 0x29,
	0x4f, 0x18, 0x8f, 0xe2, 0x1b, 0x1d, 0x44, 0x9c, 0xca, 0xc4, 0xbe, 0x27, 0x8b, 0x21, 0x2f, 0xf4,
	0xae, 0xec, 0xbb, 0x30, 0xe8, 0xe7, 0xb0, 0xf3, 0xfe, 0x03, 0x8a, 0xab, 0xef, 0x09, 0xc9, 0x0e,
	0xeb, 0x3b, 0xbf, 0xe0, 0xc9, 0x87, 0x24, 0x15, 0x7d, 0x2d, 0x29, 0x38, 0xdf, 0xa3, 0x68, 0xc6,
	0x27, 0x45, 0x54, 0x2c, 0x17, 0x7a, 0x4f, 0xc4, 0x51, 0x10, 0xb1, 0x9a, 0x73, 0x7d, 0x5f, 0x5c,
	0xab, 0x91, 0xcc, 0xfb, 0xff, 0x34, 0x95, 0x1d, 0xb2, 0xf8, 0xe2, 0x1d, 0x72, 0x3f, 0x11, 0x49,
	0xdb, 0x88, 0x34, 0x84, 0xbd, 0x5c, 0xd9, 0xb9, 0xb2, 0x9f, 0x0f, 0x37, 0xfb, 0xa9, 0x6e, 0x26,
	0x5b, 0xf3, 0xa1, 0x0f, 0xa0, 0x23, 0x3b, 0x8c, 0x5d, 0x90, 0x73, 0xb6, 0x02, 0xb0, 0x82, 0xe8,
	0x1a, 0x77, 0xcb, 0x4f, 0x66, 0x72, 0xa5, 0x5a, 0x4c, 0x41, 0xd0, 0x3b, 0xe6, 0xd3, 0xe8, 0x46,
	0x1c, 0xcb, 0x1e, 0xae, 0x00, 0x7a, 0x0a, 0x07, 0xa9, 0x93, 0x26, 0x85, 0xb1, 0x0a, 0xb1, 0x23,
	0x38, 0xb7, 0x61, 0xfa, 0x18, 0xf6, 0x05, 0x64, 0xd5, 0xc1, 0x64, 0x7f, 0x6f, 0xa1, 0xb4, 0x0f,
	0x7b, 0x49, 0x9a, 0x14, 0x78, 0x1f, 0x13, 0x87, 0xa9, 0x2d, 0x86, 0x69, 0x0d, 0xc3, 0xed, 0xb9,
	0x5e, 0xe6, 0xb9, 0x7c, 0xe5, 0x70, 0x7b, 0x3a, 0x72, 0x7b, 0xd6, 0x40, 0xcc, 0x58, 0x64, 0x45,
	0x34, 0x5d, 0xd1, 0xe4, 0x24, 0xdc, 0x42, 0x51, 0x01, 0xd4, 0xcb, 0x5d, 0xce, 0xae, 0x78, 0x2e,
	0xa6, 0xa2, 0xc3, 0x14, 0x04, 0xdf, 0xd0, 0x9c, 0xcf, 0xa2, 0x24, 0xc5, 0x0d, 0x59, 0x88, 0xf9,
	0x68, 0x31, 0x15, 0x12, 0x43, 0x2e, 0xb7, 0xb1, 0x27, 0xc7, 0x53, 0x5a, 0x83, 0x3f, 0x1a, 0xf2,
	0x85, 0x17, 0x2b, 0x23, 0x1a, 0xe6, 0xc4, 0xf4, 0x1e, 0xdc, 0x19, 0xdb, 0xaf, 0xc2, 0xb1, 0x63,
	0x85, 0x17, 0xb6, 0xc1, 0xfc, 0xa1, 0x6d, 0xf8, 0x64, 0x8b, 0x1e, 0x01, 0xad, 0x60, 0x93, 0xd9,
	0x86, 0x6f, 0x33, 0xcf, 0x7b, 0x4e, 0x34, 0x7a, 0x1f, 0x8e, 0x36, 0xf1, 0xd0, 0x30, 0x9f, 0x91,
	0x86, 0xea, 0x63, 0xbb, 0xbe, 0xcd, 0x42, 0xe1, 0xd3, 0x54, 0x7d, 0x56, 0xb8, 0xf0, 0xd9, 0xa6,
	0xc7, 0x70, 0xaf, 0x3a, 0x9b, 0xd8, 0xae, 0x15, 0x8e, 0x8c, 0xe7, 0xb6, 0xe3, 0x9e, 0x7b, 0xa4,
	0x45, 0xef, 0x40, 0xaf, 0x3a, 0x62, 0xb6, 0x61, 0xbd, 0x24, 0x3b, 0x6a, 0xb1, 0x02, 0x12, 0x41,
	0x76, 0xe9, 0x21, 0x1c, 0x54, 0xb0, 0x37, 0xb6, 0x5d, 0xc7, 0x1d, 0x91, 0xb6, 0xca, 0xb5, 0x6c,
	0xe3, 0xd2, 0x34, 0x98, 0x35, 0x21, 0x1d, 0x7a, 0x00, 0xdd, 0x0a, 0x1e, 0xda, 0x3e, 0x01, 0xd5,
	0x79, 0x68, 0xfb, 0x22, 0x62, 0x77, 0xf0, 0xb7, 0x06, 0xb4, 0x96, 0xca, 0xce, 0xf3, 0x2c, 0x37,
	0xb3, 0x18, 0x9f, 0x94, 0x7d, 0xe4, 0xda, 0x66, 0x38, 0x09, 0x4c, 0xd3, 0x9e, 0x4c, 0xc8, 0x16,
	0xfd, 0x12, 0x1e, 0x95, 0x98, 0x14, 0x24, 0xb4, 0xec, 0xc9, 0xb3, 0xd0, 0x72, 0x8c, 0xe7, 0x9e,
	0x6b, 0x85, 0xae, 0xe7, 0xdb, 0xae, 0x17, 0x8c, 0x2e, 0xc8, 0xef, 0xff, 0x96, 0x3f, 0x8d, 0x7e,
	0x01, 0x27, 0x9f, 0x70, 0x09, 0x26, 0x36, 0x43, 0xfe, 0xb9, 0x17, 0xb8, 0x16, 0xf9, 0x6d, 0x45,
	0x7f, 0x0c, 0xc7, 0x25, 0xdd, 0x71, 0x7d, 0x4f, 0x92, 0x6b, 0xde, 0xaf, 0x2b, 0xde, 0xff, 0x41,
	0x2f, 0x79, 0xa8, 0x62, 0xa9, 0x10, 0xb3, 0xc7, 0xd8, 0xd1, 0x5f, 0x6a, 0xda, 0xe0, 0x29, 0x1c,
	0xdc, 0xfa, 0x4f, 0xa3, 0x7b, 0xd0, 0xbe, 0x0c, 0xdc, 0x70, 0x68, 0xb8, 0xcf, 0x88, 0x46, 0x7b,
	0xd0, 0xb1, 0x1c, 0x77, 0x24, 0xcd, 0x06, 0xed, 0xc2, 0xae, 0xef, 0xa1, 0xe9, 0x90, 0xe6, 0xe0,
	0x14, 0xf6, 0xd0, 0xb9, 0xda, 0x6f, 0xaa, 0xc3, 0xdd, 0x6a, 0xbb, 0xc3, 0x79, 0x92, 0xbe, 0xe1,
	0xf9, 0xc7, 0xb7, 0x51, 0xfa, 0x86, 0x6c, 0x0d, 0xce, 0x94, 0x61, 0x9b, 0xfd, 0x68, 0x66, 0xd3,
	0x2c, 0xc7, 0x3c, 0x58, 0xa1, 0xef, 0x18, 0x1e, 0xd1, 0x6a, 0xcb, 0x73, 0x47, 0xa4, 0x31, 0xf8,
	0x4e, 0xe1, 0x57, 0x4f, 0x78, 0x0f, 0x3a, 0xc8, 0xb8, 0x34, 0x9c, 0x57, 0x0e, 0xd9, 0xaa, 0x4c,
	0x2b, 0x40, 0x53, 0x53, 0xcc, 0xef, 0x1d, 0xd2, 0x18, 0x0c, 0xa0, 0xb7, 0xf6, 0x97, 0x8b, 0x85,
	0x0b, 0x6f, 0xef, 0x05, 0xd9, 0xaa, 0x92, 0x5d, 0x38, 0xa3, 0x0b, 0xa2, 0x0d, 0x52, 0xf8, 0x5f,
	0xcd, 0xc5, 0xff, 0xc0, 0xb5, 0xd7, 0x12, 0x90, 0x18, 0x84, 0x8e, 0xeb, 0xe0, 0x26, 0x94, 0x03,
	0x13, 0x94, 0x43, 0x58, 0xa7, 0x0d, 0xc4, 0xfc, 0x34, 0xaa, 0x31, 0x0d, 0xc2, 0xf1, 0xa5, 0xf1,
	0x12, 0x47, 0xaf, 0x49, 0x09, 0xec, 0x49, 0xe8, 0xdc, 0x71, 0x9d, 0xc9, 0x05, 0xd9, 0x1e, 0x2c,
	0x95, 0x7c, 0xd5, 0x8b, 0x5b, 0xe6, 0x2b, 0x0b, 0x2b, 0xb3, 0x95, 0xc1, 0xcf, 0x8d, 0xb1, 0x81,
	0x57, 0x2c, 0xaf, 0x20, 0x33, 0x95, 0x95, 0xe0, 0x88, 0xe3, 0x69, 0xb3, 0x9e, 0xe5, 0x32, 0xf1,
	0x76, 0x55, 0x7b, 0x99, 0xb6, 0x35, 0xde, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x69, 0x16, 0xa7,
	0x33, 0x03, 0x0a, 0x00, 0x00,
}

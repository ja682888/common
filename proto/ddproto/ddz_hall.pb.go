// Code generated by protoc-gen-go.
// source: ddz_hall.proto
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

// Ignoring public import of ddz_base_roomTypeInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_commonRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_timerInfo from ddz_base.proto

// Ignoring public import of ddz_base_deskInfo from ddz_base.proto

// Ignoring public import of ddz_enum_protoId from ddz_base.proto

// Ignoring public import of ddz_enum_errorCode from ddz_base.proto

// Ignoring public import of ddz_enum_paiType from ddz_base.proto

// Ignoring public import of ddz_enum_actType from ddz_base.proto

// Ignoring public import of ddz_enum_gameStatus from ddz_base.proto

// Ignoring public import of ddz_enum_playerStatus from ddz_base.proto

// Ignoring public import of ddz_enum_roomType from ddz_base.proto

// Ignoring public import of ddz_enum_enterType from ddz_base.proto

// Ignoring public import of ddz_enum_coinRoomLevel from ddz_base.proto

// Ignoring public import of ddz_enum_deskGameStatus from ddz_base.proto

// 创建房间
type DdzReqCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *DdzBaseRoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzReqCreateDesk) Reset()                    { *m = DdzReqCreateDesk{} }
func (m *DdzReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzReqCreateDesk) ProtoMessage()               {}
func (*DdzReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *DdzReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqCreateDesk) GetRoomTypeInfo() *DdzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 创建房间回复的信息
type DdzAckCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32               `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string              `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserBalance      *int64               `protobuf:"varint,4,opt,name=userBalance" json:"userBalance,omitempty"`
	CreateFee        *int64               `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *DdzBaseRoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzAckCreateDesk) Reset()                    { *m = DdzAckCreateDesk{} }
func (m *DdzAckCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzAckCreateDesk) ProtoMessage()               {}
func (*DdzAckCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *DdzAckCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckCreateDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzAckCreateDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *DdzAckCreateDesk) GetUserBalance() int64 {
	if m != nil && m.UserBalance != nil {
		return *m.UserBalance
	}
	return 0
}

func (m *DdzAckCreateDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *DdzAckCreateDesk) GetRoomTypeInfo() *DdzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 游戏战绩
type DdzReqGameRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqGameRecord) Reset()                    { *m = DdzReqGameRecord{} }
func (m *DdzReqGameRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzReqGameRecord) ProtoMessage()               {}
func (*DdzReqGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

func (m *DdzReqGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DdzReqGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type DdzBaseUserRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=NickName" json:"NickName,omitempty"`
	WinAmount        *int64       `protobuf:"varint,4,opt,name=WinAmount" json:"WinAmount,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzBaseUserRecord) Reset()                    { *m = DdzBaseUserRecord{} }
func (m *DdzBaseUserRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzBaseUserRecord) ProtoMessage()               {}
func (*DdzBaseUserRecord) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{3} }

func (m *DdzBaseUserRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBaseUserRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzBaseUserRecord) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *DdzBaseUserRecord) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

//
type DdzBaseGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32               `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	DeskId           *int32               `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	BeginTime        *string              `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	Users            []*DdzBaseUserRecord `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzBaseGameRecord) Reset()                    { *m = DdzBaseGameRecord{} }
func (m *DdzBaseGameRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzBaseGameRecord) ProtoMessage()               {}
func (*DdzBaseGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{4} }

func (m *DdzBaseGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBaseGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DdzBaseGameRecord) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzBaseGameRecord) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *DdzBaseGameRecord) GetUsers() []*DdzBaseUserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

//
type DdzAckGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Records          []*DdzBaseGameRecord `protobuf:"bytes,3,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzAckGameRecord) Reset()                    { *m = DdzAckGameRecord{} }
func (m *DdzAckGameRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzAckGameRecord) ProtoMessage()               {}
func (*DdzAckGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{5} }

func (m *DdzAckGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckGameRecord) GetRecords() []*DdzBaseGameRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

// 客户端请求进入 room, 服务器返回DdzSendGameInfo
type DdzReqEnterDesk struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MatchId          *int32                `protobuf:"varint,2,opt,name=matchId" json:"matchId,omitempty"`
	TableId          *int32                `protobuf:"varint,3,opt,name=tableId" json:"tableId,omitempty"`
	PassWord         *string               `protobuf:"bytes,4,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32               `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	EnterType        *DdzEnumEnterType     `protobuf:"varint,6,opt,name=enterType,enum=ddproto.DdzEnumEnterType" json:"enterType,omitempty"`
	CoinRoomLevel    *DdzEnumCoinRoomLevel `protobuf:"varint,7,opt,name=coinRoomLevel,enum=ddproto.DdzEnumCoinRoomLevel" json:"coinRoomLevel,omitempty"`
	IsContinuous     *bool                 `protobuf:"varint,8,opt,name=isContinuous" json:"isContinuous,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzReqEnterDesk) Reset()                    { *m = DdzReqEnterDesk{} }
func (m *DdzReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzReqEnterDesk) ProtoMessage()               {}
func (*DdzReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{6} }

func (m *DdzReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqEnterDesk) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *DdzReqEnterDesk) GetTableId() int32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *DdzReqEnterDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *DdzReqEnterDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzReqEnterDesk) GetEnterType() DdzEnumEnterType {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return DdzEnumEnterType_FRIEND
}

func (m *DdzReqEnterDesk) GetCoinRoomLevel() DdzEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLevel != nil {
		return *m.CoinRoomLevel
	}
	return DdzEnumCoinRoomLevel_LV1
}

func (m *DdzReqEnterDesk) GetIsContinuous() bool {
	if m != nil && m.IsContinuous != nil {
		return *m.IsContinuous
	}
	return false
}

type DdzAckEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsContinuous     *bool        `protobuf:"varint,2,opt,name=isContinuous" json:"isContinuous,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzAckEnterDesk) Reset()                    { *m = DdzAckEnterDesk{} }
func (m *DdzAckEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzAckEnterDesk) ProtoMessage()               {}
func (*DdzAckEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{7} }

func (m *DdzAckEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckEnterDesk) GetIsContinuous() bool {
	if m != nil && m.IsContinuous != nil {
		return *m.IsContinuous
	}
	return false
}

func init() {
	proto.RegisterType((*DdzReqCreateDesk)(nil), "ddproto.ddz_req_createDesk")
	proto.RegisterType((*DdzAckCreateDesk)(nil), "ddproto.ddz_ack_createDesk")
	proto.RegisterType((*DdzReqGameRecord)(nil), "ddproto.ddz_req_gameRecord")
	proto.RegisterType((*DdzBaseUserRecord)(nil), "ddproto.ddz_base_userRecord")
	proto.RegisterType((*DdzBaseGameRecord)(nil), "ddproto.ddz_base_gameRecord")
	proto.RegisterType((*DdzAckGameRecord)(nil), "ddproto.ddz_ack_gameRecord")
	proto.RegisterType((*DdzReqEnterDesk)(nil), "ddproto.ddz_req_enterDesk")
	proto.RegisterType((*DdzAckEnterDesk)(nil), "ddproto.ddz_ack_enterDesk")
}

var fileDescriptor15 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x31, 0xf9, 0x9b, 0x36, 0x81, 0x6e, 0x7a, 0xb0, 0x0a, 0x02, 0x64, 0x71, 0x40, 0x42,
	0xe4, 0x50, 0x21, 0xee, 0x14, 0x84, 0x1a, 0x09, 0x95, 0xa8, 0x2a, 0xf4, 0x68, 0x6d, 0xec, 0xa1,
	0x59, 0xc5, 0xde, 0x0d, 0xbb, 0x36, 0xa8, 0xbc, 0x08, 0xef, 0xc2, 0xab, 0x71, 0x61, 0x76, 0x37,
	0xb1, 0xfb, 0x23, 0xa4, 0x5a, 0x5c, 0xa2, 0xec, 0xec, 0xec, 0x7c, 0x3f, 0xf3, 0x19, 0xc6, 0x59,
	0xf6, 0x33, 0x59, 0xf2, 0x3c, 0x9f, 0xae, 0xb5, 0x2a, 0x15, 0xeb, 0x67, 0x99, 0xfb, 0x73, 0x30,
	0x49, 0x55, 0x51, 0x28, 0x99, 0xa4, 0xb9, 0x40, 0x59, 0xfa, 0xdb, 0x03, 0xd7, 0xbd, 0xe0, 0x06,
	0xfd, 0x39, 0x5e, 0x03, 0xb3, 0x15, 0x8d, 0xdf, 0x92, 0x54, 0x23, 0x2f, 0xf1, 0x3d, 0x9a, 0x15,
	0x7b, 0x0e, 0xbd, 0x25, 0xf2, 0x0c, 0x75, 0x14, 0x3c, 0x0b, 0x5e, 0xec, 0x1c, 0xee, 0x4f, 0x37,
	0x43, 0xa7, 0x73, 0xfb, 0x7b, 0xec, 0xee, 0xd8, 0x6b, 0xd8, 0xd5, 0x4a, 0x15, 0x67, 0x97, 0x6b,
	0x9c, 0xc9, 0xaf, 0x2a, 0xea, 0xb8, 0xde, 0x27, 0x75, 0xef, 0x16, 0x2a, 0xb9, 0xda, 0x15, 0xff,
	0x0e, 0x3c, 0x24, 0x4f, 0x57, 0xed, 0x21, 0xc7, 0xd0, 0xcb, 0xa8, 0x7b, 0x96, 0x39, 0xb0, 0x2e,
	0x7b, 0x08, 0x83, 0x35, 0x37, 0xe6, 0x87, 0xd2, 0x59, 0x14, 0x52, 0x65, 0xc8, 0x26, 0xb0, 0x53,
	0x19, 0xd4, 0x47, 0x3c, 0xe7, 0x32, 0xc5, 0xe8, 0x3e, 0x15, 0x43, 0xb6, 0x07, 0x43, 0x0f, 0xf5,
	0x01, 0x31, 0xea, 0xba, 0xd2, 0x4d, 0xf2, 0xbd, 0x3b, 0x91, 0xff, 0xd2, 0xd8, 0x75, 0xc1, 0x0b,
	0x3c, 0xc5, 0x94, 0x90, 0xef, 0xc8, 0x1d, 0xa0, 0x23, 0xb6, 0xbc, 0x49, 0x87, 0x65, 0x39, 0xf3,
	0xac, 0x47, 0xb4, 0x86, 0x49, 0x0d, 0x68, 0x2f, 0x5a, 0x0d, 0xa6, 0x61, 0x9f, 0xfd, 0x30, 0x3b,
	0x7c, 0x64, 0x4d, 0x39, 0x11, 0xe9, 0xea, 0x84, 0x08, 0x6e, 0x4c, 0x21, 0xfd, 0xe7, 0x42, 0xbe,
	0x2d, 0x54, 0x25, 0x4b, 0x6f, 0x49, 0xfc, 0x2b, 0xb8, 0x02, 0xf9, 0xbf, 0x5a, 0x36, 0x3b, 0x09,
	0xdd, 0x99, 0xc0, 0x16, 0x78, 0x21, 0xe4, 0x99, 0x28, 0xbc, 0xff, 0x43, 0xf6, 0x12, 0xba, 0x56,
	0x95, 0x21, 0xef, 0x43, 0x9a, 0xf9, 0xf8, 0xb6, 0xcb, 0x8d, 0xe8, 0xf8, 0xb2, 0xc9, 0x47, 0x6b,
	0x5e, 0x8d, 0xaf, 0xde, 0x8a, 0x57, 0xd0, 0xd7, 0xee, 0xbd, 0x21, 0x72, 0xff, 0x80, 0x6e, 0x40,
	0xe2, 0x3f, 0x01, 0xec, 0x6d, 0xf7, 0x4b, 0xdf, 0x0c, 0xea, 0x16, 0xd1, 0x7c, 0x00, 0xfd, 0x82,
	0x97, 0xe9, 0xb2, 0xce, 0x26, 0x15, 0x4a, 0xbe, 0xc8, 0xb1, 0x36, 0x86, 0xf6, 0x32, 0xa7, 0xb0,
	0x9e, 0xdb, 0xb0, 0x7a, 0x5f, 0x1a, 0xba, 0x5d, 0x47, 0x77, 0x0a, 0x43, 0x07, 0x6b, 0xf3, 0xe6,
	0x12, 0x39, 0x3e, 0x7c, 0x74, 0x8d, 0x30, 0xca, 0xaa, 0x48, 0xea, 0x16, 0xf6, 0x06, 0x46, 0xa9,
	0x12, 0xf2, 0x94, 0x22, 0xfa, 0x11, 0xbf, 0x63, 0x1e, 0xf5, 0xdd, 0x9b, 0xa7, 0xb7, 0xdf, 0x5c,
	0x6b, 0x63, 0xfb, 0xb0, 0x2b, 0xcc, 0x3b, 0x25, 0x4b, 0x21, 0x2b, 0x55, 0x99, 0x68, 0x40, 0xcf,
	0x06, 0xf1, 0x27, 0x2f, 0xde, 0x1a, 0xdf, 0x56, 0xfc, 0xcd, 0x81, 0xd6, 0x81, 0xc1, 0x51, 0xe7,
	0x38, 0x9c, 0xdf, 0x9b, 0x07, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x45, 0xdf, 0x08, 0x4a, 0xa3,
	0x04, 0x00, 0x00,
}

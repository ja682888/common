// Code generated by protoc-gen-go.
// source: erddz_hall.proto
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

// Ignoring public import of erddz_base_roomTypeInfo from erddz_base.proto

// Ignoring public import of erddz_base_playerInfo from erddz_base.proto

// Ignoring public import of erddz_base_playerRateInfo from erddz_base.proto

// Ignoring public import of erddz_base_commonRateInfo from erddz_base.proto

// Ignoring public import of erddz_base_timerInfo from erddz_base.proto

// Ignoring public import of erddz_base_deskInfo from erddz_base.proto

// Ignoring public import of erddz_enum_protoId from erddz_base.proto

// Ignoring public import of erddz_enum_errorCode from erddz_base.proto

// Ignoring public import of erddz_enum_paiType from erddz_base.proto

// Ignoring public import of erddz_enum_actType from erddz_base.proto

// Ignoring public import of erddz_enum_deskStatus from erddz_base.proto

// Ignoring public import of erddz_enum_playerStatus from erddz_base.proto

// Ignoring public import of erddz_enum_roomType from erddz_base.proto

// Ignoring public import of erddz_enum_coinRoomLevel from erddz_base.proto

// 创建房间
type ErddzReqCreateDesk struct {
	Header           *ProtoHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *ErddzBaseRoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ErddzReqCreateDesk) Reset()                    { *m = ErddzReqCreateDesk{} }
func (m *ErddzReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqCreateDesk) ProtoMessage()               {}
func (*ErddzReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *ErddzReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqCreateDesk) GetRoomTypeInfo() *ErddzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 创建房间回复的信息
type ErddzAckCreateDesk struct {
	Header           *ProtoHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32                 `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string                `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	CreateFee        *int64                 `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *ErddzBaseRoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ErddzAckCreateDesk) Reset()                    { *m = ErddzAckCreateDesk{} }
func (m *ErddzAckCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckCreateDesk) ProtoMessage()               {}
func (*ErddzAckCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{1} }

func (m *ErddzAckCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckCreateDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ErddzAckCreateDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *ErddzAckCreateDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *ErddzAckCreateDesk) GetRoomTypeInfo() *ErddzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 客户端请求进入 room, 服务器返回DdzSendGameInfo
type ErddzReqEnterDesk struct {
	Header           *ProtoHeader            `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PassWord         *string                 `protobuf:"bytes,2,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32                 `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	RoomType         *ErddzEnumRoomType      `protobuf:"varint,4,opt,name=roomType,enum=ddproto.ErddzEnumRoomType" json:"roomType,omitempty"`
	CoinRoomLevel    *ErddzEnumCoinRoomLevel `protobuf:"varint,5,opt,name=coinRoomLevel,enum=ddproto.ErddzEnumCoinRoomLevel" json:"coinRoomLevel,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ErddzReqEnterDesk) Reset()                    { *m = ErddzReqEnterDesk{} }
func (m *ErddzReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqEnterDesk) ProtoMessage()               {}
func (*ErddzReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{2} }

func (m *ErddzReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqEnterDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *ErddzReqEnterDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzReqEnterDesk) GetRoomType() ErddzEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return ErddzEnumRoomType_ERDDZ_FRIEND
}

func (m *ErddzReqEnterDesk) GetCoinRoomLevel() ErddzEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLevel != nil {
		return *m.CoinRoomLevel
	}
	return ErddzEnumCoinRoomLevel_ERDDZ_LV1
}

type ErddzAckEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckEnterDesk) Reset()                    { *m = ErddzAckEnterDesk{} }
func (m *ErddzAckEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckEnterDesk) ProtoMessage()               {}
func (*ErddzAckEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{3} }

func (m *ErddzAckEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ErddzReqCreateDesk)(nil), "ddproto.erddz_req_createDesk")
	proto.RegisterType((*ErddzAckCreateDesk)(nil), "ddproto.erddz_ack_createDesk")
	proto.RegisterType((*ErddzReqEnterDesk)(nil), "ddproto.erddz_req_enterDesk")
	proto.RegisterType((*ErddzAckEnterDesk)(nil), "ddproto.erddz_ack_enterDesk")
}

var fileDescriptor21 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x49, 0x43, 0x43, 0x7b, 0xd0, 0xa8, 0xb8, 0x1d, 0xa2, 0x8a, 0xa1, 0x44, 0x0c, 0x4c,
	0x19, 0x32, 0x20, 0x24, 0x36, 0x84, 0x50, 0x23, 0x31, 0x44, 0x08, 0x89, 0x31, 0x32, 0xf1, 0xa1,
	0x56, 0x4d, 0xe2, 0x60, 0xa7, 0x20, 0x78, 0x23, 0x9e, 0x84, 0xd7, 0xc2, 0xb1, 0x9b, 0x96, 0x7f,
	0x03, 0x59, 0x2c, 0xfb, 0xbb, 0xcf, 0xf7, 0xfd, 0xee, 0x60, 0x88, 0x82, 0xb1, 0xb7, 0x64, 0x4e,
	0xb3, 0x2c, 0x28, 0x05, 0xaf, 0x38, 0xd9, 0x63, 0x4c, 0x5f, 0x26, 0xa3, 0x94, 0xe7, 0x39, 0x2f,
	0x92, 0x34, 0x5b, 0x60, 0x51, 0x99, 0xea, 0x64, 0xed, 0x7f, 0xa0, 0x12, 0x8d, 0xe2, 0x57, 0x30,
	0x36, 0x9a, 0xc0, 0xa7, 0x24, 0x15, 0x48, 0x2b, 0xbc, 0x42, 0xb9, 0x24, 0x27, 0xe0, 0xcc, 0x91,
	0x32, 0x14, 0x9e, 0x35, 0xb5, 0x4e, 0xf7, 0xc3, 0x71, 0xb0, 0x6e, 0x1c, 0xc4, 0xf5, 0x39, 0xd3,
	0x35, 0x72, 0x06, 0x07, 0x82, 0xf3, 0xfc, 0xee, 0xb5, 0xc4, 0xa8, 0x78, 0xe4, 0x5e, 0x47, 0x7b,
	0xa7, 0x1b, 0xef, 0x36, 0x2e, 0xf9, 0xea, 0xf3, 0xdf, 0xad, 0x26, 0x96, 0xa6, 0xcb, 0xf6, 0xb1,
	0x2e, 0x38, 0x4c, 0xb9, 0x23, 0xa6, 0x03, 0xbb, 0x64, 0x08, 0xbd, 0x92, 0x4a, 0xf9, 0xc2, 0x05,
	0xf3, 0x6c, 0xa5, 0xf4, 0xc9, 0x21, 0xf4, 0x4d, 0xd7, 0x6b, 0x44, 0xaf, 0xab, 0x24, 0xfb, 0x17,
	0xab, 0xf3, 0x4f, 0xd6, 0x0f, 0x0b, 0x46, 0xdb, 0x15, 0xa9, 0x5d, 0xa2, 0x68, 0x81, 0xaa, 0xd0,
	0x62, 0x85, 0x76, 0x5f, 0xa3, 0x75, 0x34, 0x9a, 0x82, 0x5f, 0x49, 0x14, 0x91, 0x41, 0x1d, 0x90,
	0x00, 0x7a, 0x4d, 0x9e, 0xb7, 0xab, 0x14, 0x37, 0x3c, 0xfa, 0xc1, 0x84, 0xc5, 0x2a, 0xdf, 0x30,
	0x91, 0x73, 0x18, 0xa4, 0x7c, 0x51, 0xdc, 0xaa, 0xf7, 0x0d, 0x3e, 0x63, 0xa6, 0xc7, 0x73, 0xc3,
	0xe3, 0xbf, 0x3e, 0x7d, 0x33, 0xfa, 0x17, 0xcd, 0x20, 0xf5, 0xd2, 0x5b, 0x0e, 0x72, 0xd9, 0x99,
	0xd9, 0xf1, 0x4e, 0x6c, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xae, 0x31, 0x60, 0xd1, 0x73, 0x02,
	0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: common_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Mid                *string `protobuf:"bytes,1,opt,name=mid" json:"mid,omitempty"`
	Pwd                *string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin               *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                 *uint32 `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName           *string `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	LoginTurntable     *bool   `protobuf:"varint,6,opt,name=loginTurntable" json:"loginTurntable,omitempty"`
	LoginTurntableTime *string `protobuf:"bytes,7,opt,name=loginTurntableTime" json:"loginTurntableTime,omitempty"`
	Scores             *int32  `protobuf:"varint,8,opt,name=scores" json:"scores,omitempty"`
	LastSignTime       *string `protobuf:"bytes,9,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignCount          *int32  `protobuf:"varint,10,opt,name=signCount" json:"signCount,omitempty"`
	Diamond            *int64  `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	OpenId             *string `protobuf:"bytes,12,opt,name=openId" json:"openId,omitempty"`
	HeadUrl            *string `protobuf:"bytes,13,opt,name=headUrl" json:"headUrl,omitempty"`
	City               *string `protobuf:"bytes,14,opt,name=city" json:"city,omitempty"`
	Sex                *int32  `protobuf:"varint,15,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *User) GetMid() string {
	if m != nil && m.Mid != nil {
		return *m.Mid
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *User) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *User) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *User) GetLoginTurntable() bool {
	if m != nil && m.LoginTurntable != nil {
		return *m.LoginTurntable
	}
	return false
}

func (m *User) GetLoginTurntableTime() string {
	if m != nil && m.LoginTurntableTime != nil {
		return *m.LoginTurntableTime
	}
	return ""
}

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignCount() int32 {
	if m != nil && m.SignCount != nil {
		return *m.SignCount
	}
	return 0
}

func (m *User) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

type TNotice struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32   `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string  `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string  `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string  `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Fileds           []string `protobuf:"bytes,6,rep,name=fileds" json:"fileds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *TNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *TNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *TNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *TNotice) GetFileds() []string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

// 用户的seesion 主要是游戏的
type GameSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=DeskId" json:"DeskId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,4,opt,name=GameStatus" json:"GameStatus,omitempty"`
	GameId           *int32  `protobuf:"varint,5,opt,name=GameId" json:"GameId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=GameNumber" json:"GameNumber,omitempty"`
	GameCustomStatus *int32  `protobuf:"varint,7,opt,name=GameCustomStatus" json:"GameCustomStatus,omitempty"`
	IsBreak          *bool   `protobuf:"varint,8,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,9,opt,name=isLeave" json:"isLeave,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameSession) Reset()                    { *m = GameSession{} }
func (m *GameSession) String() string            { return proto.CompactTextString(m) }
func (*GameSession) ProtoMessage()               {}
func (*GameSession) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GameSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *GameSession) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *GameSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *GameSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *GameSession) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameSession) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *GameSession) GetGameCustomStatus() int32 {
	if m != nil && m.GameCustomStatus != nil {
		return *m.GameCustomStatus
	}
	return 0
}

func (m *GameSession) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *GameSession) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

// 服务器游戏玩家通过的协议
type CommonSrvGameUser struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Coin             *int64  `protobuf:"varint,3,opt,name=Coin" json:"Coin,omitempty"`
	RoomId           *int32  `protobuf:"varint,4,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,5,opt,name=DeskId" json:"DeskId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=gameNumber" json:"gameNumber,omitempty"`
	IsBreak          *bool   `protobuf:"varint,7,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,8,opt,name=isLeave" json:"isLeave,omitempty"`
	Status           *int32  `protobuf:"varint,9,opt,name=Status" json:"Status,omitempty"`
	WaitTime         *int64  `protobuf:"varint,10,opt,name=WaitTime" json:"WaitTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommonSrvGameUser) Reset()                    { *m = CommonSrvGameUser{} }
func (m *CommonSrvGameUser) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvGameUser) ProtoMessage()               {}
func (*CommonSrvGameUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CommonSrvGameUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CommonSrvGameUser) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *CommonSrvGameUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *CommonSrvGameUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *CommonSrvGameUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *CommonSrvGameUser) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *CommonSrvGameUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *CommonSrvGameUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *CommonSrvGameUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *CommonSrvGameUser) GetWaitTime() int64 {
	if m != nil && m.WaitTime != nil {
		return *m.WaitTime
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "ddproto.User")
	proto.RegisterType((*TNotice)(nil), "ddproto.TNotice")
	proto.RegisterType((*GameSession)(nil), "ddproto.GameSession")
	proto.RegisterType((*CommonSrvGameUser)(nil), "ddproto.common_srv_game_user")
}

var fileDescriptor1 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x52, 0x4b, 0x6e, 0xdb, 0x30,
	0x10, 0x85, 0x2d, 0x29, 0x92, 0xc6, 0x9f, 0xa4, 0x4a, 0x5a, 0x10, 0x5d, 0x15, 0x5e, 0x65, 0xd5,
	0x43, 0xd4, 0x01, 0x0a, 0x03, 0xad, 0x17, 0x8d, 0x83, 0x2e, 0x0d, 0x45, 0x9c, 0xba, 0x84, 0x25,
	0xd2, 0x20, 0x29, 0xb7, 0xbe, 0x40, 0xef, 0xd3, 0x0b, 0xf4, 0x6c, 0x1d, 0x8e, 0x64, 0xc7, 0xf1,
	0x4e, 0xef, 0xe9, 0xcd, 0xf0, 0xbd, 0x99, 0x81, 0xdb, 0xca, 0x34, 0x8d, 0xd1, 0x6b, 0x87, 0x76,
	0x8f, 0xf6, 0xe3, 0xce, 0x1a, 0x6f, 0x8a, 0x54, 0x4a, 0xfe, 0x98, 0xfd, 0x19, 0x42, 0xfc, 0x44,
	0xbf, 0x8a, 0x11, 0x44, 0x8d, 0x92, 0x62, 0xf0, 0x61, 0x70, 0x9f, 0x07, 0xb0, 0xfb, 0x25, 0xc5,
	0x90, 0xc1, 0x18, 0xe2, 0xca, 0x28, 0x2d, 0x22, 0x42, 0x51, 0x01, 0x30, 0x24, 0x59, 0x4c, 0xdf,
	0x93, 0xe2, 0x06, 0x32, 0xad, 0xaa, 0xed, 0xb2, 0x6c, 0x50, 0x24, 0xac, 0x7d, 0x07, 0xd3, 0xda,
	0x6c, 0x94, 0x5e, 0xb5, 0x56, 0xfb, 0xf2, 0xb9, 0x46, 0x71, 0x45, 0x7c, 0x56, 0xbc, 0x87, 0xe2,
	0x35, 0xbf, 0x52, 0x54, 0x93, 0x72, 0xcd, 0x14, 0xae, 0x5c, 0x65, 0x2c, 0x3a, 0x91, 0x11, 0x4e,
	0x8a, 0x3b, 0x18, 0xd7, 0xa5, 0xf3, 0x8f, 0x6a, 0xa3, 0x59, 0x95, 0xb3, 0xea, 0x0d, 0xe4, 0x8e,
	0x98, 0xb9, 0x69, 0xb5, 0x17, 0xc0, 0xc2, 0x6b, 0x48, 0x1f, 0x54, 0x49, 0xd9, 0xa4, 0x18, 0xb1,
	0x37, 0xea, 0x64, 0x76, 0xa8, 0x17, 0x52, 0x8c, 0xb9, 0x86, 0x04, 0x3f, 0xb1, 0x94, 0x4f, 0xb6,
	0x16, 0x93, 0x53, 0x14, 0xe5, 0x0f, 0x62, 0x7a, 0x4c, 0xe9, 0xf0, 0xb7, 0xb8, 0x0e, 0xcd, 0x66,
	0x07, 0x48, 0x57, 0x4b, 0xe3, 0x55, 0x85, 0x7d, 0xc4, 0x01, 0xbf, 0x41, 0x40, 0x33, 0xbb, 0x3a,
	0xec, 0x90, 0x07, 0x92, 0x14, 0xb7, 0x30, 0xea, 0x39, 0xe5, 0x29, 0x61, 0xc4, 0xcd, 0xde, 0xc2,
	0xa4, 0x23, 0xe7, 0x46, 0x7b, 0x24, 0x8f, 0x31, 0xd3, 0xa7, 0xfa, 0xaf, 0xd8, 0x98, 0x7e, 0x48,
	0x64, 0xf3, 0x87, 0xaa, 0x51, 0x3a, 0x1a, 0x4e, 0x74, 0x9f, 0xcf, 0xfe, 0x0e, 0x60, 0xf4, 0x99,
	0x66, 0xf8, 0x88, 0xce, 0x29, 0xa3, 0xc3, 0xff, 0xb0, 0x92, 0x45, 0xe7, 0x61, 0x12, 0xf0, 0x37,
	0x63, 0x9a, 0x85, 0xec, 0xdf, 0x27, 0xfc, 0x80, 0x6e, 0x4b, 0x38, 0x3a, 0x7a, 0xe4, 0x72, 0x5f,
	0xfa, 0xd6, 0xf1, 0xbb, 0xac, 0x09, 0x1c, 0x69, 0x92, 0x73, 0xcd, 0xb2, 0x6d, 0x9e, 0xd1, 0xf2,
	0x52, 0x92, 0x42, 0xc0, 0x4d, 0xe0, 0xe6, 0xad, 0xf3, 0xa6, 0xe9, 0xab, 0xd3, 0xe3, 0x64, 0x95,
	0xfb, 0x64, 0xb1, 0xdc, 0xf2, 0x4e, 0xb2, 0x8e, 0xf8, 0x82, 0xe5, 0xbe, 0x5b, 0x47, 0x36, 0xfb,
	0x37, 0x80, 0xbb, 0xe3, 0x61, 0xd9, 0xfd, 0x7a, 0x43, 0x7d, 0xd6, 0x6d, 0xb8, 0x23, 0x7a, 0xb8,
	0x3d, 0x37, 0x7f, 0x7e, 0x23, 0xa7, 0x7b, 0x9a, 0xbf, 0xdc, 0xd3, 0x4b, 0xb8, 0xf8, 0x22, 0xdc,
	0xc9, 0xf8, 0xe6, 0xd2, 0xf8, 0x99, 0xbd, 0xf4, 0xd2, 0x5e, 0xe7, 0x97, 0xba, 0xf4, 0x81, 0x72,
	0xae, 0x20, 0x17, 0xdf, 0x4b, 0xe5, 0xf9, 0x9e, 0xc2, 0xf1, 0x44, 0xff, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xd1, 0x7a, 0x81, 0xe4, 0x17, 0x03, 0x00, 0x00,
}
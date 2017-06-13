// Code generated by protoc-gen-go.
// source: hall_common_server_award.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// todo 等待增加
type Award int32

const (
	Award_taskType_time Award = 1
	Award_taskType_2    Award = 2
)

var Award_name = map[int32]string{
	1: "taskType_time",
	2: "taskType_2",
}
var Award_value = map[string]int32{
	"taskType_time": 1,
	"taskType_2":    2,
}

func (x Award) Enum() *Award {
	p := new(Award)
	*p = x
	return p
}
func (x Award) String() string {
	return proto.EnumName(Award_name, int32(x))
}
func (x *Award) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Award_value, data, "Award")
	if err != nil {
		return err
	}
	*x = Award(value)
	return nil
}
func (Award) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// 用户奖励，保存到数据库的结构
type Taward struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TaskType         *int32  `protobuf:"varint,2,opt,name=taskType" json:"taskType,omitempty"`
	AwardType        *int32  `protobuf:"varint,3,opt,name=awardType" json:"awardType,omitempty"`
	UserId           *uint32 `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	Memo             *string `protobuf:"bytes,5,opt,name=memo" json:"memo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Taward) Reset()                    { *m = Taward{} }
func (m *Taward) String() string            { return proto.CompactTextString(m) }
func (*Taward) ProtoMessage()               {}
func (*Taward) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Taward) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Taward) GetTaskType() int32 {
	if m != nil && m.TaskType != nil {
		return *m.TaskType
	}
	return 0
}

func (m *Taward) GetAwardType() int32 {
	if m != nil && m.AwardType != nil {
		return *m.AwardType
	}
	return 0
}

func (m *Taward) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Taward) GetMemo() string {
	if m != nil && m.Memo != nil {
		return *m.Memo
	}
	return ""
}

// 在线奖励
type AwardMOnline struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Capital          *int64  `protobuf:"varint,2,opt,name=capital" json:"capital,omitempty"`
	Rate             *int64  `protobuf:"varint,3,opt,name=rate" json:"rate,omitempty"`
	BeginTime        *string `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	DurationSec      *int64  `protobuf:"varint,5,opt,name=durationSec" json:"durationSec,omitempty"`
	Level            *int32  `protobuf:"varint,6,opt,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AwardMOnline) Reset()                    { *m = AwardMOnline{} }
func (m *AwardMOnline) String() string            { return proto.CompactTextString(m) }
func (*AwardMOnline) ProtoMessage()               {}
func (*AwardMOnline) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *AwardMOnline) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *AwardMOnline) GetCapital() int64 {
	if m != nil && m.Capital != nil {
		return *m.Capital
	}
	return 0
}

func (m *AwardMOnline) GetRate() int64 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *AwardMOnline) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *AwardMOnline) GetDurationSec() int64 {
	if m != nil && m.DurationSec != nil {
		return *m.DurationSec
	}
	return 0
}

func (m *AwardMOnline) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*Taward)(nil), "yjprotogo.Taward")
	proto.RegisterType((*AwardMOnline)(nil), "yjprotogo.award_m_online")
	proto.RegisterEnum("yjprotogo.Award", Award_name, Award_value)
}

var fileDescriptor4 = []byte{
<<<<<<< HEAD
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xa4, 0x09, 0xe4, 0x20, 0x25, 0x35, 0x8b, 0x27, 0x84, 0x3a, 0xa1, 0x0e, 0x0c,
	0xfc, 0x0b, 0x66, 0x22, 0x31, 0x5a, 0x26, 0x3e, 0x15, 0x83, 0xed, 0x8b, 0x5c, 0x37, 0xa8, 0xff,
	0x1e, 0xe7, 0x2a, 0x50, 0x37, 0xdf, 0xf3, 0xbb, 0xf7, 0xdd, 0x83, 0x87, 0x4f, 0xed, 0x9c, 0x1a,
	0xc9, 0x7b, 0x0a, 0xea, 0x80, 0x71, 0xc6, 0xa8, 0xf4, 0x8f, 0x8e, 0xe6, 0x79, 0x8a, 0x94, 0x48,
	0xb4, 0xa7, 0x2f, 0x7e, 0xec, 0x69, 0xfb, 0x0e, 0xcd, 0xc0, 0x5f, 0x02, 0xa0, 0xb4, 0x46, 0x16,
	0x8f, 0xc5, 0x53, 0x2d, 0x7a, 0xb8, 0x4e, 0xfa, 0xf0, 0x3d, 0x9c, 0x26, 0x94, 0x25, 0x2b, 0x1b,
	0x68, 0xd9, 0xc6, 0x52, 0xc5, 0xd2, 0x1a, 0x9a, 0x63, 0x0e, 0x7f, 0x35, 0x72, 0x95, 0xe7, 0x4e,
	0xdc, 0xc2, 0xca, 0xa3, 0x27, 0x59, 0xe7, 0xa9, 0xdd, 0xce, 0xb0, 0xe6, 0x05, 0xe5, 0x15, 0x05,
	0x67, 0x03, 0x5e, 0xf8, 0x0b, 0xf6, 0xdf, 0xc1, 0xd5, 0xa8, 0x27, 0x9b, 0xb4, 0x63, 0x46, 0xb5,
	0x04, 0x44, 0x9d, 0xce, 0xf1, 0xd5, 0x42, 0xfc, 0xc0, 0xbd, 0x0d, 0x83, 0xf5, 0xc8, 0x84, 0x56,
	0xdc, 0xc3, 0x8d, 0x39, 0x66, 0x8b, 0xa5, 0xf0, 0x86, 0x23, 0x83, 0x2a, 0xd1, 0x41, 0xed, 0x70,
	0x46, 0x27, 0x9b, 0xe5, 0xaa, 0xdd, 0x0e, 0xea, 0x73, 0x9f, 0x0d, 0x74, 0x7f, 0x1d, 0x54, 0xca,
	0x19, 0x7d, 0x91, 0x2f, 0x80, 0x7f, 0xe9, 0xa5, 0x2f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x76,
	0x5f, 0x98, 0x02, 0x28, 0x01, 0x00, 0x00,
=======
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcb, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x49, 0x3b, 0xad, 0xf6, 0xc8, 0x94, 0x31, 0x88, 0x04, 0x11, 0x29, 0xb3, 0x2a, 0xb3,
	0x70, 0xe1, 0x5b, 0xb8, 0x8d, 0xdd, 0x87, 0xd8, 0x1e, 0xc6, 0x68, 0xd2, 0x94, 0x34, 0x33, 0x32,
	0x1b, 0x9f, 0xc8, 0x87, 0x94, 0x39, 0x75, 0x6e, 0xbb, 0xff, 0xc2, 0x1f, 0xbe, 0x1c, 0x78, 0xfa,
	0xd0, 0xd6, 0xaa, 0xd6, 0x3b, 0xe7, 0x7b, 0x35, 0x62, 0xd8, 0x62, 0x50, 0xfa, 0x5b, 0x87, 0xee,
	0x79, 0x08, 0x3e, 0x7a, 0x5e, 0xec, 0x3e, 0x49, 0xac, 0xfd, 0xf2, 0x07, 0xf2, 0x86, 0x2a, 0x5e,
	0x42, 0x62, 0x3a, 0xc1, 0x2a, 0x56, 0x67, 0x32, 0x31, 0x1d, 0x7f, 0x80, 0xeb, 0xa8, 0xc7, 0xaf,
	0x66, 0x37, 0xa0, 0x48, 0x28, 0x3d, 0x7a, 0xfe, 0x08, 0x05, 0x8d, 0xa8, 0x4c, 0xa9, 0x3c, 0x05,
	0xfc, 0x1e, 0xf2, 0xcd, 0x88, 0xe1, 0xb5, 0x13, 0xb3, 0x8a, 0xd5, 0x73, 0xf9, 0xef, 0x38, 0x87,
	0x99, 0x43, 0xe7, 0x45, 0x56, 0xb1, 0xba, 0x90, 0xa4, 0x97, 0xbf, 0x0c, 0x4a, 0x5a, 0x2a, 0xa7,
	0x7c, 0x6f, 0x4d, 0x7f, 0x3e, 0x67, 0x17, 0x73, 0x01, 0x57, 0xad, 0x1e, 0x4c, 0xd4, 0x96, 0x78,
	0x52, 0x79, 0xb0, 0xfb, 0x87, 0x83, 0x8e, 0x13, 0x49, 0x2a, 0x49, 0xef, 0x11, 0xdf, 0x71, 0x6d,
	0xfa, 0xc6, 0x38, 0x24, 0x8e, 0x42, 0x9e, 0x02, 0x5e, 0xc1, 0x4d, 0xb7, 0x09, 0x3a, 0x1a, 0xdf,
	0xbf, 0x61, 0x4b, 0x44, 0xa9, 0x3c, 0x8f, 0xf8, 0x1d, 0x64, 0x16, 0xb7, 0x68, 0x45, 0x4e, 0xdf,
	0x9b, 0xcc, 0x6a, 0x05, 0xd9, 0x74, 0xad, 0x5b, 0x98, 0x1f, 0xae, 0xa1, 0xa2, 0x71, 0xb8, 0x60,
	0xbc, 0x04, 0x38, 0x46, 0x2f, 0x8b, 0xe4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x01, 0x75, 0x0b, 0xbd,
	0x86, 0x01, 0x00, 0x00,
>>>>>>> 9acc1444663f452e461ae6c033b78f4059a59180
}
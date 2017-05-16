// Code generated by protoc-gen-go.
// source: common_server_award.proto
// DO NOT EDIT!

package ddproto

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
func (Award) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

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
func (*Taward) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

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
func (*AwardMOnline) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

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
	proto.RegisterType((*Taward)(nil), "ddproto.Taward")
	proto.RegisterType((*AwardMOnline)(nil), "ddproto.award_m_online")
	proto.RegisterEnum("ddproto.Award", Award_name, Award_value)
}

var fileDescriptor10 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0xb8, 0x4e, 0xc9, 0x41, 0x4a, 0x6a, 0x16, 0xb3, 0xa1, 0x4e, 0xa8, 0x03, 0x03,
	0xdf, 0x82, 0x99, 0x48, 0x8c, 0x96, 0x89, 0x4f, 0xc8, 0xc2, 0x7f, 0x22, 0xd7, 0x0d, 0xe2, 0xdb,
	0xe3, 0x5c, 0x05, 0x62, 0xf3, 0x3d, 0xbf, 0x7b, 0xbf, 0x7b, 0x70, 0x3f, 0xa5, 0x10, 0x52, 0xd4,
	0x27, 0xcc, 0x0b, 0x66, 0x6d, 0xbe, 0x4c, 0xb6, 0x4f, 0x73, 0x4e, 0x25, 0xc9, 0xad, 0xb5, 0xf4,
	0x38, 0xbc, 0x41, 0x3b, 0xd2, 0x87, 0x04, 0x68, 0x9c, 0x55, 0xec, 0x81, 0x3d, 0x0a, 0x39, 0xc0,
	0x55, 0x31, 0xa7, 0xcf, 0xf1, 0x7b, 0x46, 0xd5, 0x90, 0xb2, 0x87, 0x8e, 0x6c, 0x24, 0x71, 0x92,
	0x76, 0xd0, 0x9e, 0x6b, 0xf4, 0x8b, 0x55, 0x9b, 0x3a, 0xf7, 0xf2, 0x06, 0x36, 0x01, 0x43, 0x52,
	0xa2, 0x4e, 0xdd, 0x61, 0x81, 0x1d, 0x2d, 0xe8, 0xa0, 0x53, 0xf4, 0x2e, 0xe2, 0x3f, 0x3f, 0x23,
	0xff, 0x2d, 0x6c, 0x27, 0x33, 0xbb, 0x62, 0x3c, 0x31, 0xf8, 0x1a, 0x90, 0x4d, 0xb9, 0xc4, 0xf3,
	0x95, 0xf8, 0x8e, 0x1f, 0x2e, 0x8e, 0x2e, 0x20, 0x11, 0x3a, 0x79, 0x07, 0xd7, 0xf6, 0x5c, 0x2d,
	0x2e, 0xc5, 0x57, 0x9c, 0x08, 0xc4, 0x65, 0x0f, 0xc2, 0xe3, 0x82, 0x5e, 0xb5, 0xeb, 0x55, 0xc7,
	0x23, 0x88, 0x4b, 0x9f, 0x3d, 0xf4, 0xbf, 0x1d, 0x74, 0xa9, 0x19, 0x03, 0xab, 0x17, 0xc0, 0x9f,
	0xf4, 0x3c, 0x34, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x73, 0xfa, 0x6c, 0x21, 0x01, 0x00,
	0x00,
}

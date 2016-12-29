// Code generated by protoc-gen-go.
// source: hall_data.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HallEnumProtoId int32

const (
	HallEnumProtoId_HALL_PID_HEARTBEAT                   HallEnumProtoId = 0
	HallEnumProtoId_HALL_PID_QUICK_CONN                  HallEnumProtoId = 1
	HallEnumProtoId_HALL_PID_QUICK_CONN_ACK              HallEnumProtoId = 2
	HallEnumProtoId_HALL_PID_GAME_LOGIN                  HallEnumProtoId = 3
	HallEnumProtoId_HALL_PID_GAME_LOGIN_ACK              HallEnumProtoId = 4
	HallEnumProtoId_HALL_PID_WXPAYUNIFIEDORDER_REQ       HallEnumProtoId = 5
	HallEnumProtoId_HALL_PID_WXPAYUNIFIEDORDER_ACK       HallEnumProtoId = 6
	HallEnumProtoId_HALL_PID_WXPAYSYNCCHECKER_REQ        HallEnumProtoId = 7
	HallEnumProtoId_HALL_PID_WXPAYSYNCCHECKER_ACK        HallEnumProtoId = 8
	HallEnumProtoId_HALL_PID_USER_DATA                   HallEnumProtoId = 9
	HallEnumProtoId_HALL_PID_USER_DATA_ACK               HallEnumProtoId = 10
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY                HallEnumProtoId = 11
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY_ACK            HallEnumProtoId = 12
	HallEnumProtoId_HALL_PID_DS_LOTTERY_INFO_ACK         HallEnumProtoId = 13
	HallEnumProtoId_HALL_PID_ONLINEWARD_REQ              HallEnumProtoId = 14
	HallEnumProtoId_HALL_PID_ONLINEWARD_ACK              HallEnumProtoId = 15
	HallEnumProtoId_HALL_PID_EVENT_REQ                   HallEnumProtoId = 16
	HallEnumProtoId_HALL_PID_EVENT_ACK                   HallEnumProtoId = 17
	HallEnumProtoId_HALL_PID_GOODS_LIST_REQ              HallEnumProtoId = 18
	HallEnumProtoId_HALL_PID_GOODS_LIST_ACK              HallEnumProtoId = 19
	HallEnumProtoId_HALL_PID_HOTUPDATEVERSIONINFO_REQ    HallEnumProtoId = 20
	HallEnumProtoId_HALL_PID_HOTUPDATEVERSIONINFO_ACK    HallEnumProtoId = 21
	HallEnumProtoId_HALL_PID_BAG_ITEMS_REQ               HallEnumProtoId = 22
	HallEnumProtoId_HALL_PID_BAG_ITEMS_ACK               HallEnumProtoId = 23
	HallEnumProtoId_HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ HallEnumProtoId = 24
	HallEnumProtoId_HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK HallEnumProtoId = 25
	HallEnumProtoId_HALL_PID_APPLE_PAY_CB_REQ            HallEnumProtoId = 26
	HallEnumProtoId_HALL_PID_HOTUPDATEASSETSINFO_REQ     HallEnumProtoId = 27
	HallEnumProtoId_HALL_PID_HOTUPDATEASSETSINFO_ACK     HallEnumProtoId = 28
	HallEnumProtoId_HALL_PID_MAIL_LIST_REQ               HallEnumProtoId = 29
	HallEnumProtoId_HALL_PID_MAIL_LIST_ACK               HallEnumProtoId = 30
	HallEnumProtoId_HALL_PID_HALL_DSLOTTERYINFO_REQ      HallEnumProtoId = 31
	HallEnumProtoId_HALL_PID_FRIENDS_LIST_REQ            HallEnumProtoId = 32
	HallEnumProtoId_HALL_PID_FRIENDS_LIST_ACK            HallEnumProtoId = 33
	HallEnumProtoId_HALL_PID_RANK_REQ                    HallEnumProtoId = 34
	HallEnumProtoId_HALL_PID_RANK_ACK                    HallEnumProtoId = 35
	HallEnumProtoId_HALL_PID_STRONGBOX_INFO_REQ          HallEnumProtoId = 36
	HallEnumProtoId_HALL_PID_STRONGBOX_INFO_ACK          HallEnumProtoId = 37
	HallEnumProtoId_HALL_PID_STRONGBOX_ACCESS_REQ        HallEnumProtoId = 38
	HallEnumProtoId_HALL_PID_STRONGBOX_ACCESS_ACK        HallEnumProtoId = 39
	HallEnumProtoId_HALL_PID_FRIEND_ADD_REQ              HallEnumProtoId = 40
	HallEnumProtoId_HALL_PID_FRIEND_ADD_ACK              HallEnumProtoId = 41
	HallEnumProtoId_HALL_PID_FRIEND_DEL_REQ              HallEnumProtoId = 42
	HallEnumProtoId_HALL_PID_FRIEND_DEL_ACK              HallEnumProtoId = 43
	HallEnumProtoId_HALL_PID_FRIENDS_SEARCH_REQ          HallEnumProtoId = 44
	HallEnumProtoId_HALL_PID_FRIENDS_SEARCH_ACK          HallEnumProtoId = 45
)

var HallEnumProtoId_name = map[int32]string{
	0:  "HALL_PID_HEARTBEAT",
	1:  "HALL_PID_QUICK_CONN",
	2:  "HALL_PID_QUICK_CONN_ACK",
	3:  "HALL_PID_GAME_LOGIN",
	4:  "HALL_PID_GAME_LOGIN_ACK",
	5:  "HALL_PID_WXPAYUNIFIEDORDER_REQ",
	6:  "HALL_PID_WXPAYUNIFIEDORDER_ACK",
	7:  "HALL_PID_WXPAYSYNCCHECKER_REQ",
	8:  "HALL_PID_WXPAYSYNCCHECKER_ACK",
	9:  "HALL_PID_USER_DATA",
	10: "HALL_PID_USER_DATA_ACK",
	11: "HALL_PID_DRAW_LOTTERY",
	12: "HALL_PID_DRAW_LOTTERY_ACK",
	13: "HALL_PID_DS_LOTTERY_INFO_ACK",
	14: "HALL_PID_ONLINEWARD_REQ",
	15: "HALL_PID_ONLINEWARD_ACK",
	16: "HALL_PID_EVENT_REQ",
	17: "HALL_PID_EVENT_ACK",
	18: "HALL_PID_GOODS_LIST_REQ",
	19: "HALL_PID_GOODS_LIST_ACK",
	20: "HALL_PID_HOTUPDATEVERSIONINFO_REQ",
	21: "HALL_PID_HOTUPDATEVERSIONINFO_ACK",
	22: "HALL_PID_BAG_ITEMS_REQ",
	23: "HALL_PID_BAG_ITEMS_ACK",
	24: "HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ",
	25: "HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK",
	26: "HALL_PID_APPLE_PAY_CB_REQ",
	27: "HALL_PID_HOTUPDATEASSETSINFO_REQ",
	28: "HALL_PID_HOTUPDATEASSETSINFO_ACK",
	29: "HALL_PID_MAIL_LIST_REQ",
	30: "HALL_PID_MAIL_LIST_ACK",
	31: "HALL_PID_HALL_DSLOTTERYINFO_REQ",
	32: "HALL_PID_FRIENDS_LIST_REQ",
	33: "HALL_PID_FRIENDS_LIST_ACK",
	34: "HALL_PID_RANK_REQ",
	35: "HALL_PID_RANK_ACK",
	36: "HALL_PID_STRONGBOX_INFO_REQ",
	37: "HALL_PID_STRONGBOX_INFO_ACK",
	38: "HALL_PID_STRONGBOX_ACCESS_REQ",
	39: "HALL_PID_STRONGBOX_ACCESS_ACK",
	40: "HALL_PID_FRIEND_ADD_REQ",
	41: "HALL_PID_FRIEND_ADD_ACK",
	42: "HALL_PID_FRIEND_DEL_REQ",
	43: "HALL_PID_FRIEND_DEL_ACK",
	44: "HALL_PID_FRIENDS_SEARCH_REQ",
	45: "HALL_PID_FRIENDS_SEARCH_ACK",
}
var HallEnumProtoId_value = map[string]int32{
	"HALL_PID_HEARTBEAT":                   0,
	"HALL_PID_QUICK_CONN":                  1,
	"HALL_PID_QUICK_CONN_ACK":              2,
	"HALL_PID_GAME_LOGIN":                  3,
	"HALL_PID_GAME_LOGIN_ACK":              4,
	"HALL_PID_WXPAYUNIFIEDORDER_REQ":       5,
	"HALL_PID_WXPAYUNIFIEDORDER_ACK":       6,
	"HALL_PID_WXPAYSYNCCHECKER_REQ":        7,
	"HALL_PID_WXPAYSYNCCHECKER_ACK":        8,
	"HALL_PID_USER_DATA":                   9,
	"HALL_PID_USER_DATA_ACK":               10,
	"HALL_PID_DRAW_LOTTERY":                11,
	"HALL_PID_DRAW_LOTTERY_ACK":            12,
	"HALL_PID_DS_LOTTERY_INFO_ACK":         13,
	"HALL_PID_ONLINEWARD_REQ":              14,
	"HALL_PID_ONLINEWARD_ACK":              15,
	"HALL_PID_EVENT_REQ":                   16,
	"HALL_PID_EVENT_ACK":                   17,
	"HALL_PID_GOODS_LIST_REQ":              18,
	"HALL_PID_GOODS_LIST_ACK":              19,
	"HALL_PID_HOTUPDATEVERSIONINFO_REQ":    20,
	"HALL_PID_HOTUPDATEVERSIONINFO_ACK":    21,
	"HALL_PID_BAG_ITEMS_REQ":               22,
	"HALL_PID_BAG_ITEMS_ACK":               23,
	"HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ": 24,
	"HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK": 25,
	"HALL_PID_APPLE_PAY_CB_REQ":            26,
	"HALL_PID_HOTUPDATEASSETSINFO_REQ":     27,
	"HALL_PID_HOTUPDATEASSETSINFO_ACK":     28,
	"HALL_PID_MAIL_LIST_REQ":               29,
	"HALL_PID_MAIL_LIST_ACK":               30,
	"HALL_PID_HALL_DSLOTTERYINFO_REQ":      31,
	"HALL_PID_FRIENDS_LIST_REQ":            32,
	"HALL_PID_FRIENDS_LIST_ACK":            33,
	"HALL_PID_RANK_REQ":                    34,
	"HALL_PID_RANK_ACK":                    35,
	"HALL_PID_STRONGBOX_INFO_REQ":          36,
	"HALL_PID_STRONGBOX_INFO_ACK":          37,
	"HALL_PID_STRONGBOX_ACCESS_REQ":        38,
	"HALL_PID_STRONGBOX_ACCESS_ACK":        39,
	"HALL_PID_FRIEND_ADD_REQ":              40,
	"HALL_PID_FRIEND_ADD_ACK":              41,
	"HALL_PID_FRIEND_DEL_REQ":              42,
	"HALL_PID_FRIEND_DEL_ACK":              43,
	"HALL_PID_FRIENDS_SEARCH_REQ":          44,
	"HALL_PID_FRIENDS_SEARCH_ACK":          45,
}

func (x HallEnumProtoId) Enum() *HallEnumProtoId {
	p := new(HallEnumProtoId)
	*p = x
	return p
}
func (x HallEnumProtoId) String() string {
	return proto.EnumName(HallEnumProtoId_name, int32(x))
}
func (x *HallEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumProtoId_value, data, "HallEnumProtoId")
	if err != nil {
		return err
	}
	*x = HallEnumProtoId(value)
	return nil
}
func (HallEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

// 活动类型
type HallEnumEvent int32

const (
	HallEnumEvent_TYPE_TIME HallEnumEvent = 1
	HallEnumEvent_TYPE_NEW  HallEnumEvent = 2
	HallEnumEvent_TYPE_NULL HallEnumEvent = 3
)

var HallEnumEvent_name = map[int32]string{
	1: "TYPE_TIME",
	2: "TYPE_NEW",
	3: "TYPE_NULL",
}
var HallEnumEvent_value = map[string]int32{
	"TYPE_TIME": 1,
	"TYPE_NEW":  2,
	"TYPE_NULL": 3,
}

func (x HallEnumEvent) Enum() *HallEnumEvent {
	p := new(HallEnumEvent)
	*p = x
	return p
}
func (x HallEnumEvent) String() string {
	return proto.EnumName(HallEnumEvent_name, int32(x))
}
func (x *HallEnumEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumEvent_value, data, "HallEnumEvent")
	if err != nil {
		return err
	}
	*x = HallEnumEvent(value)
	return nil
}
func (HallEnumEvent) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

// 活动奖品
type HallEnum_Reward int32

const (
	HallEnum_Reward_RE_EXP  HallEnum_Reward = 1
	HallEnum_Reward_RE_GIFT HallEnum_Reward = 2
)

var HallEnum_Reward_name = map[int32]string{
	1: "RE_EXP",
	2: "RE_GIFT",
}
var HallEnum_Reward_value = map[string]int32{
	"RE_EXP":  1,
	"RE_GIFT": 2,
}

func (x HallEnum_Reward) Enum() *HallEnum_Reward {
	p := new(HallEnum_Reward)
	*p = x
	return p
}
func (x HallEnum_Reward) String() string {
	return proto.EnumName(HallEnum_Reward_name, int32(x))
}
func (x *HallEnum_Reward) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnum_Reward_value, data, "HallEnum_Reward")
	if err != nil {
		return err
	}
	*x = HallEnum_Reward(value)
	return nil
}
func (HallEnum_Reward) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

// 邮件类型
type HallEnumMailType int32

const (
	HallEnumMailType_SYSTEM     HallEnumMailType = 1
	HallEnumMailType_FRIEND_ADD HallEnumMailType = 2
)

var HallEnumMailType_name = map[int32]string{
	1: "SYSTEM",
	2: "FRIEND_ADD",
}
var HallEnumMailType_value = map[string]int32{
	"SYSTEM":     1,
	"FRIEND_ADD": 2,
}

func (x HallEnumMailType) Enum() *HallEnumMailType {
	p := new(HallEnumMailType)
	*p = x
	return p
}
func (x HallEnumMailType) String() string {
	return proto.EnumName(HallEnumMailType_name, int32(x))
}
func (x *HallEnumMailType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumMailType_value, data, "HallEnumMailType")
	if err != nil {
		return err
	}
	*x = HallEnumMailType(value)
	return nil
}
func (HallEnumMailType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

// 道具类型
type HallEnumPropsType int32

const (
	HallEnumPropsType_TYPE_LABA     HallEnumPropsType = 1
	HallEnumPropsType_TYPE_JIPAIQI  HallEnumPropsType = 2
	HallEnumPropsType_TYPE_FANGKA   HallEnumPropsType = 3
	HallEnumPropsType_TYPE_EXP_3000 HallEnumPropsType = 4
)

var HallEnumPropsType_name = map[int32]string{
	1: "TYPE_LABA",
	2: "TYPE_JIPAIQI",
	3: "TYPE_FANGKA",
	4: "TYPE_EXP_3000",
}
var HallEnumPropsType_value = map[string]int32{
	"TYPE_LABA":     1,
	"TYPE_JIPAIQI":  2,
	"TYPE_FANGKA":   3,
	"TYPE_EXP_3000": 4,
}

func (x HallEnumPropsType) Enum() *HallEnumPropsType {
	p := new(HallEnumPropsType)
	*p = x
	return p
}
func (x HallEnumPropsType) String() string {
	return proto.EnumName(HallEnumPropsType_name, int32(x))
}
func (x *HallEnumPropsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumPropsType_value, data, "HallEnumPropsType")
	if err != nil {
		return err
	}
	*x = HallEnumPropsType(value)
	return nil
}
func (HallEnumPropsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

// 奖励物品类型
type HallLotteryItemType int32

const (
	HallLotteryItemType_TYPE_ROOM_TICKET HallLotteryItemType = 1
	HallLotteryItemType_TYPE_FILL_SIGN   HallLotteryItemType = 2
	HallLotteryItemType_TYEP_DIAMOND     HallLotteryItemType = 3
)

var HallLotteryItemType_name = map[int32]string{
	1: "TYPE_ROOM_TICKET",
	2: "TYPE_FILL_SIGN",
	3: "TYEP_DIAMOND",
}
var HallLotteryItemType_value = map[string]int32{
	"TYPE_ROOM_TICKET": 1,
	"TYPE_FILL_SIGN":   2,
	"TYEP_DIAMOND":     3,
}

func (x HallLotteryItemType) Enum() *HallLotteryItemType {
	p := new(HallLotteryItemType)
	*p = x
	return p
}
func (x HallLotteryItemType) String() string {
	return proto.EnumName(HallLotteryItemType_name, int32(x))
}
func (x *HallLotteryItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallLotteryItemType_value, data, "HallLotteryItemType")
	if err != nil {
		return err
	}
	*x = HallLotteryItemType(value)
	return nil
}
func (HallLotteryItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

// 任务类型
type HallEnumTaskType int32

const (
	HallEnumTaskType_TYPE_MJ  HallEnumTaskType = 1
	HallEnumTaskType_TYPE_DDZ HallEnumTaskType = 2
	HallEnumTaskType_TYPE_ZJH HallEnumTaskType = 3
)

var HallEnumTaskType_name = map[int32]string{
	1: "TYPE_MJ",
	2: "TYPE_DDZ",
	3: "TYPE_ZJH",
}
var HallEnumTaskType_value = map[string]int32{
	"TYPE_MJ":  1,
	"TYPE_DDZ": 2,
	"TYPE_ZJH": 3,
}

func (x HallEnumTaskType) Enum() *HallEnumTaskType {
	p := new(HallEnumTaskType)
	*p = x
	return p
}
func (x HallEnumTaskType) String() string {
	return proto.EnumName(HallEnumTaskType_name, int32(x))
}
func (x *HallEnumTaskType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumTaskType_value, data, "HallEnumTaskType")
	if err != nil {
		return err
	}
	*x = HallEnumTaskType(value)
	return nil
}
func (HallEnumTaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

// vip等级
type HallUser_VIP int32

const (
	HallUser_VIP_LV_1 HallUser_VIP = 1
	HallUser_VIP_LV_2 HallUser_VIP = 2
	HallUser_VIP_LV_3 HallUser_VIP = 3
	HallUser_VIP_LV_4 HallUser_VIP = 4
	HallUser_VIP_LV_5 HallUser_VIP = 5
	HallUser_VIP_LV_6 HallUser_VIP = 6
)

var HallUser_VIP_name = map[int32]string{
	1: "LV_1",
	2: "LV_2",
	3: "LV_3",
	4: "LV_4",
	5: "LV_5",
	6: "LV_6",
}
var HallUser_VIP_value = map[string]int32{
	"LV_1": 1,
	"LV_2": 2,
	"LV_3": 3,
	"LV_4": 4,
	"LV_5": 5,
	"LV_6": 6,
}

func (x HallUser_VIP) Enum() *HallUser_VIP {
	p := new(HallUser_VIP)
	*p = x
	return p
}
func (x HallUser_VIP) String() string {
	return proto.EnumName(HallUser_VIP_name, int32(x))
}
func (x *HallUser_VIP) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallUser_VIP_value, data, "HallUser_VIP")
	if err != nil {
		return err
	}
	*x = HallUser_VIP(value)
	return nil
}
func (HallUser_VIP) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

// 单个活动
type HallItemEvent struct {
	Id               *int32           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumEvent   `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumEvent" json:"type,omitempty"`
	Reward           *HallEnum_Reward `protobuf:"varint,3,opt,name=reward,enum=ddproto.HallEnum_Reward" json:"reward,omitempty"`
	RichText         []string         `protobuf:"bytes,5,rep,name=richText" json:"richText,omitempty"`
	Title            *string          `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallItemEvent) Reset()                    { *m = HallItemEvent{} }
func (m *HallItemEvent) String() string            { return proto.CompactTextString(m) }
func (*HallItemEvent) ProtoMessage()               {}
func (*HallItemEvent) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *HallItemEvent) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemEvent) GetType() HallEnumEvent {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumEvent_TYPE_TIME
}

func (m *HallItemEvent) GetReward() HallEnum_Reward {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return HallEnum_Reward_RE_EXP
}

func (m *HallItemEvent) GetRichText() []string {
	if m != nil {
		return m.RichText
	}
	return nil
}

func (m *HallItemEvent) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

// 单个邮件
type HallMailItem struct {
	Id               *string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumMailType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumMailType" json:"type,omitempty"`
	Title            *string           `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Content          *string           `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	IsWatch          *bool             `protobuf:"varint,5,opt,name=isWatch" json:"isWatch,omitempty"`
	IsCheck          *bool             `protobuf:"varint,6,opt,name=isCheck" json:"isCheck,omitempty"`
	Attach           []*HallBagItem    `protobuf:"bytes,7,rep,name=attach" json:"attach,omitempty"`
	Date             *int64            `protobuf:"varint,8,opt,name=date" json:"date,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallMailItem) Reset()                    { *m = HallMailItem{} }
func (m *HallMailItem) String() string            { return proto.CompactTextString(m) }
func (*HallMailItem) ProtoMessage()               {}
func (*HallMailItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *HallMailItem) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *HallMailItem) GetType() HallEnumMailType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumMailType_SYSTEM
}

func (m *HallMailItem) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *HallMailItem) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *HallMailItem) GetIsWatch() bool {
	if m != nil && m.IsWatch != nil {
		return *m.IsWatch
	}
	return false
}

func (m *HallMailItem) GetIsCheck() bool {
	if m != nil && m.IsCheck != nil {
		return *m.IsCheck
	}
	return false
}

func (m *HallMailItem) GetAttach() []*HallBagItem {
	if m != nil {
		return m.Attach
	}
	return nil
}

func (m *HallMailItem) GetDate() int64 {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return 0
}

// 背包单个道具
type HallBagItem struct {
	Type             *HallEnumPropsType `protobuf:"varint,1,opt,name=type,enum=ddproto.HallEnumPropsType" json:"type,omitempty"`
	Amount           *int32             `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallBagItem) Reset()                    { *m = HallBagItem{} }
func (m *HallBagItem) String() string            { return proto.CompactTextString(m) }
func (*HallBagItem) ProtoMessage()               {}
func (*HallBagItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *HallBagItem) GetType() HallEnumPropsType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumPropsType_TYPE_LABA
}

func (m *HallBagItem) GetAmount() int32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// 单个任务
type HallItemTask struct {
	Id               *int32            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumTaskType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumTaskType" json:"type,omitempty"`
	Title            *string           `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Reward           *string           `protobuf:"bytes,4,opt,name=reward" json:"reward,omitempty"`
	Num              *int32            `protobuf:"varint,5,opt,name=Num" json:"Num,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallItemTask) Reset()                    { *m = HallItemTask{} }
func (m *HallItemTask) String() string            { return proto.CompactTextString(m) }
func (*HallItemTask) ProtoMessage()               {}
func (*HallItemTask) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *HallItemTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemTask) GetType() HallEnumTaskType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumTaskType_TYPE_MJ
}

func (m *HallItemTask) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *HallItemTask) GetReward() string {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return ""
}

func (m *HallItemTask) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// 个人信息
type HallUserData struct {
	UserName          *string       `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	UserId            *int32        `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NiceValue         *int64        `protobuf:"varint,3,opt,name=niceValue" json:"niceValue,omitempty"`
	EvilValue         *int64        `protobuf:"varint,4,opt,name=evilValue" json:"evilValue,omitempty"`
	UserLevel         *int32        `protobuf:"varint,5,opt,name=userLevel" json:"userLevel,omitempty"`
	UserVIP           *bool         `protobuf:"varint,6,opt,name=userVIP" json:"userVIP,omitempty"`
	UserVIPLv         *HallUser_VIP `protobuf:"varint,7,opt,name=userVIPLv,enum=ddproto.HallUser_VIP" json:"userVIPLv,omitempty"`
	UserMoney         *int64        `protobuf:"varint,8,opt,name=userMoney" json:"userMoney,omitempty"`
	UserDiamond       *int64        `protobuf:"varint,9,opt,name=userDiamond" json:"userDiamond,omitempty"`
	UserRedBag        *string       `protobuf:"bytes,10,opt,name=userRedBag" json:"userRedBag,omitempty"`
	UserLotteryTicket *int64        `protobuf:"varint,11,opt,name=userLotteryTicket" json:"userLotteryTicket,omitempty"`
	Sex               *bool         `protobuf:"varint,12,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized  []byte        `json:"-"`
}

func (m *HallUserData) Reset()                    { *m = HallUserData{} }
func (m *HallUserData) String() string            { return proto.CompactTextString(m) }
func (*HallUserData) ProtoMessage()               {}
func (*HallUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *HallUserData) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *HallUserData) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallUserData) GetNiceValue() int64 {
	if m != nil && m.NiceValue != nil {
		return *m.NiceValue
	}
	return 0
}

func (m *HallUserData) GetEvilValue() int64 {
	if m != nil && m.EvilValue != nil {
		return *m.EvilValue
	}
	return 0
}

func (m *HallUserData) GetUserLevel() int32 {
	if m != nil && m.UserLevel != nil {
		return *m.UserLevel
	}
	return 0
}

func (m *HallUserData) GetUserVIP() bool {
	if m != nil && m.UserVIP != nil {
		return *m.UserVIP
	}
	return false
}

func (m *HallUserData) GetUserVIPLv() HallUser_VIP {
	if m != nil && m.UserVIPLv != nil {
		return *m.UserVIPLv
	}
	return HallUser_VIP_LV_1
}

func (m *HallUserData) GetUserMoney() int64 {
	if m != nil && m.UserMoney != nil {
		return *m.UserMoney
	}
	return 0
}

func (m *HallUserData) GetUserDiamond() int64 {
	if m != nil && m.UserDiamond != nil {
		return *m.UserDiamond
	}
	return 0
}

func (m *HallUserData) GetUserRedBag() string {
	if m != nil && m.UserRedBag != nil {
		return *m.UserRedBag
	}
	return ""
}

func (m *HallUserData) GetUserLotteryTicket() int64 {
	if m != nil && m.UserLotteryTicket != nil {
		return *m.UserLotteryTicket
	}
	return 0
}

func (m *HallUserData) GetSex() bool {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return false
}

// 排行信息
type HallRankItem struct {
	Placing          *int32  `protobuf:"varint,1,opt,name=placing" json:"placing,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	RankInfo         *string `protobuf:"bytes,4,opt,name=rankInfo" json:"rankInfo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallRankItem) Reset()                    { *m = HallRankItem{} }
func (m *HallRankItem) String() string            { return proto.CompactTextString(m) }
func (*HallRankItem) ProtoMessage()               {}
func (*HallRankItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *HallRankItem) GetPlacing() int32 {
	if m != nil && m.Placing != nil {
		return *m.Placing
	}
	return 0
}

func (m *HallRankItem) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallRankItem) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *HallRankItem) GetRankInfo() string {
	if m != nil && m.RankInfo != nil {
		return *m.RankInfo
	}
	return ""
}

// 金币专区
type CoinZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CoinZone) Reset()                    { *m = CoinZone{} }
func (m *CoinZone) String() string            { return proto.CompactTextString(m) }
func (*CoinZone) ProtoMessage()               {}
func (*CoinZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *CoinZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 钻石专区
type DiamondZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DiamondZone) Reset()                    { *m = DiamondZone{} }
func (m *DiamondZone) String() string            { return proto.CompactTextString(m) }
func (*DiamondZone) ProtoMessage()               {}
func (*DiamondZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *DiamondZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 兑换专区
type ExchangeZone struct {
	Money            []*GoodsItem `protobuf:"bytes,1,rep,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ExchangeZone) Reset()                    { *m = ExchangeZone{} }
func (m *ExchangeZone) String() string            { return proto.CompactTextString(m) }
func (*ExchangeZone) ProtoMessage()               {}
func (*ExchangeZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *ExchangeZone) GetMoney() []*GoodsItem {
	if m != nil {
		return m.Money
	}
	return nil
}

// 购买专区
type BuyZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BuyZone) Reset()                    { *m = BuyZone{} }
func (m *BuyZone) String() string            { return proto.CompactTextString(m) }
func (*BuyZone) ProtoMessage()               {}
func (*BuyZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *BuyZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 商品类型
type GoodsItem struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Money            *int32  `protobuf:"varint,2,opt,name=money" json:"money,omitempty"`
	Img              *string `protobuf:"bytes,3,opt,name=img" json:"img,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GoodsItem) Reset()                    { *m = GoodsItem{} }
func (m *GoodsItem) String() string            { return proto.CompactTextString(m) }
func (*GoodsItem) ProtoMessage()               {}
func (*GoodsItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *GoodsItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GoodsItem) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *GoodsItem) GetImg() string {
	if m != nil && m.Img != nil {
		return *m.Img
	}
	return ""
}

// 保险箱数据
type HallStrongboxInfo struct {
	BoxCoin          *int64 `protobuf:"varint,2,opt,name=boxCoin" json:"boxCoin,omitempty"`
	UserCoin         *int64 `protobuf:"varint,3,opt,name=userCoin" json:"userCoin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HallStrongboxInfo) Reset()                    { *m = HallStrongboxInfo{} }
func (m *HallStrongboxInfo) String() string            { return proto.CompactTextString(m) }
func (*HallStrongboxInfo) ProtoMessage()               {}
func (*HallStrongboxInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{11} }

func (m *HallStrongboxInfo) GetBoxCoin() int64 {
	if m != nil && m.BoxCoin != nil {
		return *m.BoxCoin
	}
	return 0
}

func (m *HallStrongboxInfo) GetUserCoin() int64 {
	if m != nil && m.UserCoin != nil {
		return *m.UserCoin
	}
	return 0
}

func init() {
	proto.RegisterType((*HallItemEvent)(nil), "ddproto.hall_item_event")
	proto.RegisterType((*HallMailItem)(nil), "ddproto.hall_mail_item")
	proto.RegisterType((*HallBagItem)(nil), "ddproto.hall_bag_item")
	proto.RegisterType((*HallItemTask)(nil), "ddproto.hall_item_task")
	proto.RegisterType((*HallUserData)(nil), "ddproto.hall_userData")
	proto.RegisterType((*HallRankItem)(nil), "ddproto.hall_rank_item")
	proto.RegisterType((*CoinZone)(nil), "ddproto.CoinZone")
	proto.RegisterType((*DiamondZone)(nil), "ddproto.DiamondZone")
	proto.RegisterType((*ExchangeZone)(nil), "ddproto.ExchangeZone")
	proto.RegisterType((*BuyZone)(nil), "ddproto.BuyZone")
	proto.RegisterType((*GoodsItem)(nil), "ddproto.GoodsItem")
	proto.RegisterType((*HallStrongboxInfo)(nil), "ddproto.hall_strongbox_info")
	proto.RegisterEnum("ddproto.HallEnumProtoId", HallEnumProtoId_name, HallEnumProtoId_value)
	proto.RegisterEnum("ddproto.HallEnumEvent", HallEnumEvent_name, HallEnumEvent_value)
	proto.RegisterEnum("ddproto.HallEnum_Reward", HallEnum_Reward_name, HallEnum_Reward_value)
	proto.RegisterEnum("ddproto.HallEnumMailType", HallEnumMailType_name, HallEnumMailType_value)
	proto.RegisterEnum("ddproto.HallEnumPropsType", HallEnumPropsType_name, HallEnumPropsType_value)
	proto.RegisterEnum("ddproto.HallLotteryItemType", HallLotteryItemType_name, HallLotteryItemType_value)
	proto.RegisterEnum("ddproto.HallEnumTaskType", HallEnumTaskType_name, HallEnumTaskType_value)
	proto.RegisterEnum("ddproto.HallUser_VIP", HallUser_VIP_name, HallUser_VIP_value)
}

var fileDescriptor14 = []byte{
	// 1318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x56, 0xdb, 0x46,
	0x10, 0xae, 0xb1, 0x0d, 0xf6, 0x18, 0x1b, 0xb1, 0x24, 0x44, 0x84, 0x90, 0x10, 0xe7, 0xa7, 0x60,
	0x5a, 0x4e, 0x7e, 0xda, 0x9e, 0xde, 0xe4, 0x42, 0xb6, 0x84, 0xd9, 0x20, 0x4b, 0x8e, 0x24, 0x20,
	0xa4, 0x17, 0x3a, 0x8a, 0xad, 0x82, 0x0e, 0xb6, 0xc5, 0xb1, 0x45, 0x0a, 0xaf, 0xd0, 0xbb, 0x3e,
	0x44, 0xdf, 0xa3, 0x8f, 0xd6, 0xd9, 0x5d, 0x59, 0x46, 0xc1, 0x26, 0xb9, 0x5b, 0xcd, 0xcc, 0xf7,
	0xed, 0x37, 0xdf, 0x8e, 0xa4, 0x85, 0xa5, 0x33, 0xaf, 0xd7, 0x73, 0xbb, 0x5e, 0xe4, 0xed, 0x5e,
	0x0c, 0xc3, 0x28, 0x24, 0x0b, 0xdd, 0x2e, 0x5f, 0x54, 0xff, 0xc9, 0xc4, 0xc9, 0x20, 0xf2, 0xfb,
	0xae, 0xff, 0xc5, 0x1f, 0x44, 0x04, 0x60, 0x2e, 0xe8, 0xca, 0x99, 0xcd, 0xcc, 0x56, 0x9e, 0xbc,
	0x84, 0x5c, 0x74, 0x7d, 0xe1, 0xcb, 0x73, 0xf8, 0x54, 0x79, 0x23, 0xef, 0xc6, 0xb8, 0x5d, 0x8e,
	0xf1, 0x07, 0x97, 0x63, 0xcc, 0x36, 0xcc, 0x0f, 0xfd, 0xbf, 0xbc, 0x61, 0x57, 0xce, 0xf2, 0xca,
	0xb5, 0x29, 0x95, 0x16, 0x2f, 0x20, 0x12, 0x14, 0x86, 0x41, 0xe7, 0xcc, 0xf1, 0xaf, 0x22, 0x39,
	0xbf, 0x99, 0xdd, 0x2a, 0x92, 0x32, 0xe4, 0xa3, 0x20, 0xea, 0xf9, 0xf2, 0x3c, 0x62, 0x8b, 0xd5,
	0xff, 0x32, 0x50, 0xe1, 0xa8, 0xbe, 0x17, 0x08, 0x61, 0x37, 0x24, 0x15, 0x49, 0x2d, 0x25, 0xe9,
	0xd1, 0x94, 0x8d, 0x38, 0x8e, 0xd5, 0x4c, 0x98, 0xb3, 0x1c, 0xba, 0x04, 0x0b, 0x9d, 0x70, 0x10,
	0xa1, 0x60, 0x39, 0x37, 0x0e, 0x04, 0xa3, 0x63, 0x2f, 0xea, 0x9c, 0xa1, 0x94, 0xcc, 0x56, 0x41,
	0x04, 0x1a, 0x67, 0x7e, 0xe7, 0x9c, 0x8b, 0x29, 0xa0, 0x01, 0xf3, 0x5e, 0x14, 0x79, 0x58, 0xb0,
	0x80, 0x5a, 0x4b, 0x6f, 0x56, 0xd3, 0xfb, 0x7d, 0xf6, 0x4e, 0x85, 0xc2, 0x45, 0xc8, 0xa1, 0xbf,
	0xbe, 0x5c, 0x40, 0x54, 0xb6, 0xaa, 0x43, 0x39, 0x9d, 0xde, 0x89, 0x45, 0x67, 0xb8, 0xe8, 0x8d,
	0x29, 0xa2, 0xf1, 0xf9, 0x62, 0x24, 0x54, 0x57, 0x70, 0xcf, 0x7e, 0x78, 0x89, 0x2a, 0x59, 0x8f,
	0xf9, 0xea, 0x28, 0xf6, 0x83, 0x9f, 0x51, 0xe4, 0x8d, 0xce, 0x53, 0x47, 0xb4, 0x9d, 0xf2, 0x63,
	0x7d, 0x0a, 0x35, 0x83, 0x38, 0x53, 0xec, 0xa8, 0x24, 0x87, 0x26, 0xdc, 0x28, 0x41, 0xd6, 0xb8,
	0xec, 0x73, 0x27, 0xf2, 0xd5, 0xbf, 0xe7, 0xe2, 0x1e, 0x2e, 0x47, 0xfe, 0x50, 0xc5, 0xd1, 0x61,
	0x07, 0xc7, 0xd6, 0x86, 0xd7, 0xf7, 0xe3, 0xa3, 0x40, 0x02, 0x16, 0xa1, 0x5d, 0x21, 0x94, 0x2c,
	0x43, 0x71, 0x10, 0x74, 0xfc, 0x23, 0xaf, 0x77, 0x29, 0xf6, 0xc8, 0xb2, 0x90, 0xff, 0x25, 0xe8,
	0x89, 0x50, 0x6e, 0x1c, 0x62, 0x28, 0x1d, 0x27, 0xa7, 0x27, 0x36, 0x63, 0xb6, 0xb3, 0xd0, 0x11,
	0x6d, 0xc7, 0xb6, 0x6f, 0x8b, 0x1a, 0x0c, 0xe8, 0x5f, 0xd0, 0x79, 0xd6, 0xd9, 0x57, 0xce, 0xb3,
	0xb4, 0x8b, 0xf9, 0x31, 0x5d, 0x2b, 0x1c, 0xf8, 0xd7, 0xc2, 0x7e, 0xb2, 0x02, 0x25, 0xae, 0x3a,
	0x40, 0x1b, 0x07, 0x5d, 0xb9, 0xc8, 0x83, 0x68, 0x1a, 0x0b, 0x5a, 0x7e, 0xb7, 0xee, 0x9d, 0xca,
	0xc0, 0x1b, 0x58, 0x83, 0x65, 0x2e, 0x25, 0x8c, 0x22, 0x7f, 0x78, 0xed, 0x04, 0x9d, 0x73, 0x3f,
	0x92, 0x4b, 0xbc, 0x1c, 0xcd, 0x18, 0xf9, 0x57, 0xf2, 0x22, 0x93, 0x53, 0x3d, 0x8c, 0x4f, 0x60,
	0xe8, 0x0d, 0xce, 0xc5, 0x81, 0xa2, 0xe2, 0x8b, 0x9e, 0xd7, 0x09, 0x06, 0xa7, 0xf1, 0x31, 0xa4,
	0xbd, 0x28, 0x33, 0xb7, 0xd0, 0x8b, 0x73, 0xee, 0x96, 0xb0, 0x9b, 0x0d, 0x3e, 0xe2, 0xe9, 0xe0,
	0xcf, 0x50, 0x18, 0x5e, 0xdd, 0x81, 0x42, 0x23, 0x0c, 0x06, 0x9f, 0x50, 0x3a, 0x79, 0x02, 0xd9,
	0x0b, 0xef, 0x1a, 0xc9, 0xd8, 0x94, 0x91, 0xa4, 0xd7, 0x66, 0x18, 0x76, 0x47, 0x14, 0x77, 0xac,
	0xee, 0x42, 0x29, 0x6e, 0xe8, 0xfb, 0xea, 0x5f, 0xc3, 0xa2, 0x76, 0xd5, 0x39, 0xf3, 0x06, 0xa7,
	0x3e, 0x07, 0x3c, 0x85, 0x7c, 0x9f, 0x7b, 0x34, 0x1b, 0x52, 0x83, 0x85, 0xfa, 0xe5, 0xf5, 0xf7,
	0xd1, 0xbf, 0x85, 0x62, 0xf2, 0x90, 0x9a, 0xc7, 0xf2, 0x78, 0x1f, 0x31, 0x13, 0xe8, 0x63, 0xd0,
	0x3f, 0x15, 0x16, 0x54, 0x7f, 0x87, 0x15, 0xee, 0xe3, 0x28, 0x1a, 0x86, 0x83, 0xd3, 0xcf, 0xe1,
	0x95, 0x1b, 0xa0, 0x1b, 0xcc, 0x4c, 0x5c, 0x33, 0x2b, 0x38, 0x28, 0x3b, 0x1e, 0x35, 0x1e, 0xe1,
	0x73, 0x54, 0xfb, 0xb7, 0x04, 0xcb, 0xa9, 0x97, 0x25, 0x0a, 0x69, 0x97, 0xac, 0x02, 0xd9, 0x57,
	0x74, 0xdd, 0x6d, 0x53, 0xd5, 0xdd, 0xd7, 0x14, 0xcb, 0xa9, 0x6b, 0x8a, 0x23, 0xfd, 0x40, 0x1e,
	0xc0, 0x4a, 0x12, 0xff, 0x70, 0x48, 0x1b, 0x07, 0x6e, 0xc3, 0x34, 0x0c, 0x29, 0x43, 0xd6, 0xe1,
	0xc1, 0x94, 0x84, 0xab, 0x34, 0x0e, 0xa4, 0xb9, 0x14, 0xaa, 0xa9, 0xb4, 0x34, 0x57, 0x37, 0x9b,
	0xd4, 0x90, 0xb2, 0x29, 0xd4, 0x24, 0xc1, 0x51, 0x39, 0x52, 0x85, 0xc7, 0x49, 0xf2, 0xf8, 0x63,
	0x5b, 0x39, 0x39, 0x34, 0xe8, 0x1e, 0xd5, 0x54, 0xd3, 0x52, 0x35, 0xcb, 0xb5, 0xb4, 0x0f, 0x52,
	0xfe, 0x1b, 0x35, 0x8c, 0x67, 0x1e, 0xcf, 0x67, 0x23, 0x5d, 0x63, 0x9f, 0x18, 0x8d, 0xc6, 0xbe,
	0xd6, 0x38, 0x88, 0x69, 0x16, 0xee, 0x2e, 0x61, 0x2c, 0x85, 0x94, 0x23, 0x87, 0x36, 0x86, 0x55,
	0xc5, 0x51, 0xa4, 0x22, 0x79, 0x08, 0xab, 0xb7, 0xe3, 0x1c, 0x03, 0xf8, 0x16, 0xdc, 0x4f, 0x72,
	0xaa, 0xa5, 0x1c, 0x63, 0x7b, 0x8e, 0xa3, 0x59, 0x27, 0x52, 0x89, 0x6c, 0xc0, 0xda, 0xd4, 0x14,
	0x47, 0x2e, 0x92, 0x4d, 0x78, 0x34, 0x49, 0xdb, 0x49, 0x92, 0x1a, 0x7b, 0x26, 0xaf, 0x28, 0xa7,
	0xac, 0x33, 0x0d, 0x9d, 0x1a, 0xda, 0xb1, 0x62, 0xa9, 0xbc, 0x9f, 0xca, 0xac, 0x24, 0x43, 0x2e,
	0xa5, 0x3a, 0xd1, 0x8e, 0x34, 0xc3, 0xe1, 0x20, 0x69, 0x4a, 0x9c, 0xd5, 0x2f, 0xa7, 0x0f, 0xc9,
	0x34, 0x99, 0x1c, 0x6a, 0x0b, 0x10, 0x99, 0x95, 0x64, 0xc8, 0x15, 0xf2, 0x02, 0x9e, 0x4e, 0xa6,
	0xc8, 0x74, 0x0e, 0xdb, 0xe8, 0x0d, 0x32, 0x5b, 0x36, 0x35, 0x0d, 0xde, 0x0a, 0xe3, 0xb8, 0xf7,
	0xed, 0x32, 0xc6, 0x76, 0x3f, 0xe5, 0x74, 0x5d, 0x69, 0xba, 0xd4, 0xd1, 0x5a, 0x36, 0xa7, 0x58,
	0x9d, 0x91, 0x63, 0xb8, 0x07, 0x64, 0x0b, 0x9e, 0xdf, 0xa6, 0x67, 0xd3, 0xa6, 0xd8, 0xb6, 0xe6,
	0xd8, 0x89, 0x10, 0xf9, 0xbb, 0x2a, 0x19, 0xe7, 0x5a, 0xea, 0xf8, 0x94, 0x76, 0x5b, 0xd7, 0x5c,
	0x1c, 0x1a, 0xb7, 0x51, 0xe7, 0x44, 0x0f, 0xc9, 0x73, 0xd8, 0xbc, 0x4d, 0xf4, 0xd5, 0x76, 0xeb,
	0xdf, 0xac, 0x62, 0x5b, 0x3d, 0x4a, 0xb5, 0xd6, 0x52, 0xa8, 0x3e, 0x71, 0x7f, 0x63, 0x46, 0x8e,
	0xe1, 0x1e, 0x93, 0x67, 0xf0, 0x64, 0xc2, 0xce, 0x16, 0xaa, 0x1d, 0x8f, 0x51, 0x22, 0xe1, 0x49,
	0xaa, 0x8f, 0x3d, 0x8b, 0x6a, 0xc6, 0xcd, 0xd3, 0xdd, 0x9c, 0x9d, 0x66, 0x5b, 0x3c, 0x25, 0xf7,
	0x61, 0x39, 0x49, 0x5b, 0x8a, 0x71, 0xc0, 0x51, 0xd5, 0xdb, 0x61, 0x56, 0xfd, 0x0c, 0xbf, 0x7c,
	0xeb, 0x49, 0xd8, 0x76, 0x2c, 0xd3, 0x68, 0xd6, 0xcd, 0x8f, 0x6e, 0x22, 0xe6, 0xf9, 0x5d, 0x05,
	0x8c, 0xe1, 0x45, 0xea, 0x35, 0x9d, 0x14, 0x28, 0x8d, 0x86, 0x66, 0x8b, 0x41, 0x78, 0x79, 0x77,
	0x09, 0x63, 0xf9, 0x31, 0x35, 0xb2, 0xa2, 0x29, 0x57, 0x51, 0xc5, 0x9b, 0xb3, 0x35, 0x2b, 0xc9,
	0x90, 0xdb, 0xd3, 0x92, 0xaa, 0xa6, 0x73, 0x64, 0x6d, 0x56, 0x92, 0x21, 0x77, 0x52, 0xad, 0x8d,
	0x8d, 0xb4, 0xf1, 0xbb, 0xda, 0xd8, 0xe7, 0xe8, 0x9f, 0xee, 0x2a, 0x60, 0x0c, 0x3f, 0xd7, 0xde,
	0xc5, 0xf7, 0xc9, 0x1b, 0x77, 0xc3, 0x32, 0x14, 0x9d, 0x93, 0xb6, 0xe6, 0x3a, 0xb4, 0xa5, 0xe1,
	0x27, 0x78, 0x11, 0x0a, 0xfc, 0x11, 0x5f, 0x76, 0xfc, 0xe6, 0x8e, 0x93, 0xc6, 0xa1, 0xae, 0x4b,
	0xd9, 0xda, 0x0e, 0x48, 0xb7, 0x2e, 0x8c, 0x00, 0xf3, 0x96, 0xe6, 0x6a, 0x1f, 0xdb, 0x08, 0x2e,
	0xc1, 0x02, 0xae, 0x9b, 0x74, 0xcf, 0x91, 0xe6, 0x6a, 0xaf, 0xe3, 0xbf, 0xc9, 0x57, 0x97, 0x3e,
	0xac, 0xb7, 0x4f, 0x6c, 0x7c, 0xb3, 0xb0, 0xbe, 0x02, 0x30, 0xb1, 0x07, 0x21, 0x7f, 0xc0, 0xbd,
	0xa9, 0x57, 0xae, 0xb1, 0x0c, 0x5d, 0xa9, 0x2b, 0x08, 0x93, 0x60, 0x91, 0x3f, 0xbe, 0xa7, 0x6d,
	0x85, 0x7e, 0xa0, 0xa8, 0x73, 0x09, 0x4a, 0x3c, 0xb2, 0xa7, 0x18, 0xcd, 0x03, 0x45, 0x62, 0xb7,
	0x98, 0x32, 0x0f, 0xa0, 0x2e, 0xf7, 0xed, 0xab, 0x57, 0xaf, 0xa4, 0x5c, 0xcd, 0x81, 0x55, 0x4e,
	0xde, 0x13, 0xd7, 0x89, 0xf8, 0xbe, 0xc6, 0xe8, 0xef, 0x81, 0xc4, 0x8b, 0x2d, 0xd3, 0x6c, 0xa1,
	0x0f, 0xf8, 0xbd, 0x76, 0x70, 0x17, 0x02, 0x15, 0xc1, 0x49, 0xd1, 0x51, 0x9b, 0x36, 0x0d, 0xdc,
	0x87, 0xef, 0xac, 0xb5, 0x5d, 0x95, 0x2a, 0x2d, 0xd3, 0x50, 0xd1, 0x92, 0x77, 0x40, 0xa6, 0x5c,
	0xe5, 0xd0, 0x08, 0x8e, 0x6d, 0xbd, 0xbf, 0x61, 0xa9, 0xaa, 0x7e, 0x42, 0x8a, 0xf1, 0xd3, 0xa7,
	0xf7, 0xfb, 0x08, 0x3f, 0xb8, 0x71, 0x8d, 0xe3, 0xf7, 0xa5, 0x02, 0xe4, 0xf4, 0x23, 0xf7, 0x35,
	0xc2, 0xc4, 0xea, 0x0d, 0x42, 0xc4, 0xea, 0x2d, 0xb6, 0x25, 0x56, 0xbf, 0xe0, 0x7f, 0x4d, 0xac,
	0x7e, 0xc5, 0xbf, 0x97, 0x58, 0xfd, 0x26, 0xcd, 0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x76,
	0x95, 0x7e, 0x49, 0x0c, 0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: hall.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of hall_item_event from hall_data.proto

// Ignoring public import of hall_mail_item from hall_data.proto

// Ignoring public import of hall_bag_item from hall_data.proto

// Ignoring public import of hall_item_task from hall_data.proto

// Ignoring public import of hall_userData from hall_data.proto

// Ignoring public import of hall_item_rank from hall_data.proto

// Ignoring public import of CoinZone from hall_data.proto

// Ignoring public import of DiamondZone from hall_data.proto

// Ignoring public import of ExchangeZone from hall_data.proto

// Ignoring public import of BuyZone from hall_data.proto

// Ignoring public import of GoodsItem from hall_data.proto

// Ignoring public import of hall_enum_protoId from hall_data.proto

// Ignoring public import of hall_enum_event from hall_data.proto

// Ignoring public import of hall_enum_Reward from hall_data.proto

// Ignoring public import of hall_enum_mail_type from hall_data.proto

// Ignoring public import of hall_enum_props_type from hall_data.proto

// Ignoring public import of hall_enum_taskType from hall_data.proto

// Ignoring public import of hall_user_VIP from hall_data.proto

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// ======================商城商品列表======================
// 商品列表页面类型
type HallEnumGoodsType int32

const (
	HallEnumGoodsType_TYPE_COIN     HallEnumGoodsType = 1
	HallEnumGoodsType_TYPE_DIAMOND  HallEnumGoodsType = 2
	HallEnumGoodsType_TYPE_EXCHANGE HallEnumGoodsType = 3
	HallEnumGoodsType_TYPE_BUY      HallEnumGoodsType = 4
	HallEnumGoodsType_TYPE_VIP      HallEnumGoodsType = 5
)

var HallEnumGoodsType_name = map[int32]string{
	1: "TYPE_COIN",
	2: "TYPE_DIAMOND",
	3: "TYPE_EXCHANGE",
	4: "TYPE_BUY",
	5: "TYPE_VIP",
}
var HallEnumGoodsType_value = map[string]int32{
	"TYPE_COIN":     1,
	"TYPE_DIAMOND":  2,
	"TYPE_EXCHANGE": 3,
	"TYPE_BUY":      4,
	"TYPE_VIP":      5,
}

func (x HallEnumGoodsType) Enum() *HallEnumGoodsType {
	p := new(HallEnumGoodsType)
	*p = x
	return p
}
func (x HallEnumGoodsType) String() string {
	return proto.EnumName(HallEnumGoodsType_name, int32(x))
}
func (x *HallEnumGoodsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumGoodsType_value, data, "HallEnumGoodsType")
	if err != nil {
		return err
	}
	*x = HallEnumGoodsType(value)
	return nil
}
func (HallEnumGoodsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

// 货币与商品类型
type HallEnumTradeType int32

const (
	HallEnumTradeType_TRADE_COIN    HallEnumTradeType = 1
	HallEnumTradeType_TRADE_DIAMOND HallEnumTradeType = 2
	HallEnumTradeType_TRADE_BONUS   HallEnumTradeType = 3
	HallEnumTradeType_TRADE_CODE    HallEnumTradeType = 4
	HallEnumTradeType_TRADE_RMB     HallEnumTradeType = 5
	HallEnumTradeType_TRADE_PROPS   HallEnumTradeType = 6
	HallEnumTradeType_TRADE_VIP     HallEnumTradeType = 7
)

var HallEnumTradeType_name = map[int32]string{
	1: "TRADE_COIN",
	2: "TRADE_DIAMOND",
	3: "TRADE_BONUS",
	4: "TRADE_CODE",
	5: "TRADE_RMB",
	6: "TRADE_PROPS",
	7: "TRADE_VIP",
}
var HallEnumTradeType_value = map[string]int32{
	"TRADE_COIN":    1,
	"TRADE_DIAMOND": 2,
	"TRADE_BONUS":   3,
	"TRADE_CODE":    4,
	"TRADE_RMB":     5,
	"TRADE_PROPS":   6,
	"TRADE_VIP":     7,
}

func (x HallEnumTradeType) Enum() *HallEnumTradeType {
	p := new(HallEnumTradeType)
	*p = x
	return p
}
func (x HallEnumTradeType) String() string {
	return proto.EnumName(HallEnumTradeType_name, int32(x))
}
func (x *HallEnumTradeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumTradeType_value, data, "HallEnumTradeType")
	if err != nil {
		return err
	}
	*x = HallEnumTradeType(value)
	return nil
}
func (HallEnumTradeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

// 活动列表
type HallReqEvent struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqEvent) Reset()                    { *m = HallReqEvent{} }
func (m *HallReqEvent) String() string            { return proto.CompactTextString(m) }
func (*HallReqEvent) ProtoMessage()               {}
func (*HallReqEvent) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *HallReqEvent) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 活动列表ACK
type HallAckEvent struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	EventList        []*HallItemEvent `protobuf:"bytes,2,rep,name=eventList" json:"eventList,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallAckEvent) Reset()                    { *m = HallAckEvent{} }
func (m *HallAckEvent) String() string            { return proto.CompactTextString(m) }
func (*HallAckEvent) ProtoMessage()               {}
func (*HallAckEvent) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *HallAckEvent) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckEvent) GetEventList() []*HallItemEvent {
	if m != nil {
		return m.EventList
	}
	return nil
}

// 奖励物品信息
type HallLotteryItem struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallLotteryItem) Reset()                    { *m = HallLotteryItem{} }
func (m *HallLotteryItem) String() string            { return proto.CompactTextString(m) }
func (*HallLotteryItem) ProtoMessage()               {}
func (*HallLotteryItem) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *HallLotteryItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallLotteryItem) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// 签到信息与奖励物品列表
type HallSignLotteryInfo struct {
	IsSignToday      *bool              `protobuf:"varint,1,opt,name=isSignToday" json:"isSignToday,omitempty"`
	TotalDays        *int32             `protobuf:"varint,2,opt,name=totalDays" json:"totalDays,omitempty"`
	ContinuousDays   *int32             `protobuf:"varint,3,opt,name=continuousDays" json:"continuousDays,omitempty"`
	SignItems        []*HallLotteryItem `protobuf:"bytes,4,rep,name=signItems" json:"signItems,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallSignLotteryInfo) Reset()                    { *m = HallSignLotteryInfo{} }
func (m *HallSignLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*HallSignLotteryInfo) ProtoMessage()               {}
func (*HallSignLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *HallSignLotteryInfo) GetIsSignToday() bool {
	if m != nil && m.IsSignToday != nil {
		return *m.IsSignToday
	}
	return false
}

func (m *HallSignLotteryInfo) GetTotalDays() int32 {
	if m != nil && m.TotalDays != nil {
		return *m.TotalDays
	}
	return 0
}

func (m *HallSignLotteryInfo) GetContinuousDays() int32 {
	if m != nil && m.ContinuousDays != nil {
		return *m.ContinuousDays
	}
	return 0
}

func (m *HallSignLotteryInfo) GetSignItems() []*HallLotteryItem {
	if m != nil {
		return m.SignItems
	}
	return nil
}

// 转盘信息与奖励物品列表
type HallDrawLotteryInfo struct {
	IsDrawToday      *bool              `protobuf:"varint,1,opt,name=isDrawToday" json:"isDrawToday,omitempty"`
	DialItems        []*HallLotteryItem `protobuf:"bytes,2,rep,name=dialItems" json:"dialItems,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallDrawLotteryInfo) Reset()                    { *m = HallDrawLotteryInfo{} }
func (m *HallDrawLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*HallDrawLotteryInfo) ProtoMessage()               {}
func (*HallDrawLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *HallDrawLotteryInfo) GetIsDrawToday() bool {
	if m != nil && m.IsDrawToday != nil {
		return *m.IsDrawToday
	}
	return false
}

func (m *HallDrawLotteryInfo) GetDialItems() []*HallLotteryItem {
	if m != nil {
		return m.DialItems
	}
	return nil
}

// 邮件REQ
type HallReqMailList struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqMailList) Reset()                    { *m = HallReqMailList{} }
func (m *HallReqMailList) String() string            { return proto.CompactTextString(m) }
func (*HallReqMailList) ProtoMessage()               {}
func (*HallReqMailList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *HallReqMailList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 邮件ACK
type HallAckMailList struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	List             []*HallMailItem `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallAckMailList) Reset()                    { *m = HallAckMailList{} }
func (m *HallAckMailList) String() string            { return proto.CompactTextString(m) }
func (*HallAckMailList) ProtoMessage()               {}
func (*HallAckMailList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *HallAckMailList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckMailList) GetList() []*HallMailItem {
	if m != nil {
		return m.List
	}
	return nil
}

// 任务
type HallReqTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqTask) Reset()                    { *m = HallReqTask{} }
func (m *HallReqTask) String() string            { return proto.CompactTextString(m) }
func (*HallReqTask) ProtoMessage()               {}
func (*HallReqTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

func (m *HallReqTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckTask struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	TaskList         []*HallItemTask `protobuf:"bytes,2,rep,name=taskList" json:"taskList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallAckTask) Reset()                    { *m = HallAckTask{} }
func (m *HallAckTask) String() string            { return proto.CompactTextString(m) }
func (*HallAckTask) ProtoMessage()               {}
func (*HallAckTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *HallAckTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckTask) GetTaskList() []*HallItemTask {
	if m != nil {
		return m.TaskList
	}
	return nil
}

// 领取任务
type HallReqGetTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqGetTask) Reset()                    { *m = HallReqGetTask{} }
func (m *HallReqGetTask) String() string            { return proto.CompactTextString(m) }
func (*HallReqGetTask) ProtoMessage()               {}
func (*HallReqGetTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

func (m *HallReqGetTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckGetTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	FinishedNum      *int32       `protobuf:"varint,2,opt,name=finishedNum" json:"finishedNum,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAckGetTask) Reset()                    { *m = HallAckGetTask{} }
func (m *HallAckGetTask) String() string            { return proto.CompactTextString(m) }
func (*HallAckGetTask) ProtoMessage()               {}
func (*HallAckGetTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{10} }

func (m *HallAckGetTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckGetTask) GetFinishedNum() int32 {
	if m != nil && m.FinishedNum != nil {
		return *m.FinishedNum
	}
	return 0
}

// 背包道具列表
type HallReqBagItems struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqBagItems) Reset()                    { *m = HallReqBagItems{} }
func (m *HallReqBagItems) String() string            { return proto.CompactTextString(m) }
func (*HallReqBagItems) ProtoMessage()               {}
func (*HallReqBagItems) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{11} }

func (m *HallReqBagItems) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 道具列表ACK
type HallAckBagItems struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Items            []*HallBagItem `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *HallAckBagItems) Reset()                    { *m = HallAckBagItems{} }
func (m *HallAckBagItems) String() string            { return proto.CompactTextString(m) }
func (*HallAckBagItems) ProtoMessage()               {}
func (*HallAckBagItems) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{12} }

func (m *HallAckBagItems) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckBagItems) GetItems() []*HallBagItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// 个人信息
type HallReqUserData struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqUserData) Reset()                    { *m = HallReqUserData{} }
func (m *HallReqUserData) String() string            { return proto.CompactTextString(m) }
func (*HallReqUserData) ProtoMessage()               {}
func (*HallReqUserData) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{13} }

func (m *HallReqUserData) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckUserData struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	User             *HallUserData `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *HallAckUserData) Reset()                    { *m = HallAckUserData{} }
func (m *HallAckUserData) String() string            { return proto.CompactTextString(m) }
func (*HallAckUserData) ProtoMessage()               {}
func (*HallAckUserData) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{14} }

func (m *HallAckUserData) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckUserData) GetUser() *HallUserData {
	if m != nil {
		return m.User
	}
	return nil
}

// 商品列表Req
type HallReqGoodsList struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GoodsType        *HallEnumGoodsType `protobuf:"varint,2,opt,name=goods_type,enum=ddproto.HallEnumGoodsType" json:"goods_type,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallReqGoodsList) Reset()                    { *m = HallReqGoodsList{} }
func (m *HallReqGoodsList) String() string            { return proto.CompactTextString(m) }
func (*HallReqGoodsList) ProtoMessage()               {}
func (*HallReqGoodsList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{15} }

func (m *HallReqGoodsList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallReqGoodsList) GetGoodsType() HallEnumGoodsType {
	if m != nil && m.GoodsType != nil {
		return *m.GoodsType
	}
	return HallEnumGoodsType_TYPE_COIN
}

// 商品列表Ack
type HallAckGoodsList struct {
	Header           *ProtoHeader        `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GoodsType        *HallEnumGoodsType  `protobuf:"varint,2,opt,name=goods_type,enum=ddproto.HallEnumGoodsType" json:"goods_type,omitempty"`
	Items            []*HallGoodsItemMsg `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *HallAckGoodsList) Reset()                    { *m = HallAckGoodsList{} }
func (m *HallAckGoodsList) String() string            { return proto.CompactTextString(m) }
func (*HallAckGoodsList) ProtoMessage()               {}
func (*HallAckGoodsList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{16} }

func (m *HallAckGoodsList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckGoodsList) GetGoodsType() HallEnumGoodsType {
	if m != nil && m.GoodsType != nil {
		return *m.GoodsType
	}
	return HallEnumGoodsType_TYPE_COIN
}

func (m *HallAckGoodsList) GetItems() []*HallGoodsItemMsg {
	if m != nil {
		return m.Items
	}
	return nil
}

// 单个商品信息
type HallGoodsItemMsg struct {
	GoodsId          *int32             `protobuf:"varint,1,opt,name=goods_id" json:"goods_id,omitempty"`
	Name             *string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PriceType        *HallEnumTradeType `protobuf:"varint,3,opt,name=price_type,enum=ddproto.HallEnumTradeType" json:"price_type,omitempty"`
	Price            *int32             `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
	ValueType        *HallEnumTradeType `protobuf:"varint,5,opt,name=value_type,enum=ddproto.HallEnumTradeType" json:"value_type,omitempty"`
	Value            *int32             `protobuf:"varint,6,opt,name=value" json:"value,omitempty"`
	Discount         *string            `protobuf:"bytes,7,opt,name=discount" json:"discount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallGoodsItemMsg) Reset()                    { *m = HallGoodsItemMsg{} }
func (m *HallGoodsItemMsg) String() string            { return proto.CompactTextString(m) }
func (*HallGoodsItemMsg) ProtoMessage()               {}
func (*HallGoodsItemMsg) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{17} }

func (m *HallGoodsItemMsg) GetGoodsId() int32 {
	if m != nil && m.GoodsId != nil {
		return *m.GoodsId
	}
	return 0
}

func (m *HallGoodsItemMsg) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *HallGoodsItemMsg) GetPriceType() HallEnumTradeType {
	if m != nil && m.PriceType != nil {
		return *m.PriceType
	}
	return HallEnumTradeType_TRADE_COIN
}

func (m *HallGoodsItemMsg) GetPrice() int32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *HallGoodsItemMsg) GetValueType() HallEnumTradeType {
	if m != nil && m.ValueType != nil {
		return *m.ValueType
	}
	return HallEnumTradeType_TRADE_COIN
}

func (m *HallGoodsItemMsg) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *HallGoodsItemMsg) GetDiscount() string {
	if m != nil && m.Discount != nil {
		return *m.Discount
	}
	return ""
}

// 排行
type HallReqRank struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqRank) Reset()                    { *m = HallReqRank{} }
func (m *HallReqRank) String() string            { return proto.CompactTextString(m) }
func (*HallReqRank) ProtoMessage()               {}
func (*HallReqRank) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{18} }

func (m *HallReqRank) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckRank struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RankList         []*HallItemRank `protobuf:"bytes,2,rep,name=rankList" json:"rankList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallAckRank) Reset()                    { *m = HallAckRank{} }
func (m *HallAckRank) String() string            { return proto.CompactTextString(m) }
func (*HallAckRank) ProtoMessage()               {}
func (*HallAckRank) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{19} }

func (m *HallAckRank) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckRank) GetRankList() []*HallItemRank {
	if m != nil {
		return m.RankList
	}
	return nil
}

// 抽奖动作
type HallReqDrawLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqDrawLottery) Reset()                    { *m = HallReqDrawLottery{} }
func (m *HallReqDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*HallReqDrawLottery) ProtoMessage()               {}
func (*HallReqDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{20} }

func (m *HallReqDrawLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckDrawLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	LotteryId        *int32       `protobuf:"varint,2,opt,name=lotteryId" json:"lotteryId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAckDrawLottery) Reset()                    { *m = HallAckDrawLottery{} }
func (m *HallAckDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*HallAckDrawLottery) ProtoMessage()               {}
func (*HallAckDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{21} }

func (m *HallAckDrawLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckDrawLottery) GetLotteryId() int32 {
	if m != nil && m.LotteryId != nil {
		return *m.LotteryId
	}
	return 0
}

// 请求转盘与签到信息
type HallReqDsLotteryInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqDsLotteryInfo) Reset()                    { *m = HallReqDsLotteryInfo{} }
func (m *HallReqDsLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*HallReqDsLotteryInfo) ProtoMessage()               {}
func (*HallReqDsLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{22} }

func (m *HallReqDsLotteryInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 回复转盘与签到信息
type HallAckDsLotteryInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	LotteryId        *int32               `protobuf:"varint,2,opt,name=lotteryId" json:"lotteryId,omitempty"`
	DrawLotteryInfo  *HallDrawLotteryInfo `protobuf:"bytes,12,opt,name=drawLotteryInfo" json:"drawLotteryInfo,omitempty"`
	SignLotteryInfo  *HallSignLotteryInfo `protobuf:"bytes,13,opt,name=signLotteryInfo" json:"signLotteryInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *HallAckDsLotteryInfo) Reset()                    { *m = HallAckDsLotteryInfo{} }
func (m *HallAckDsLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*HallAckDsLotteryInfo) ProtoMessage()               {}
func (*HallAckDsLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{23} }

func (m *HallAckDsLotteryInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckDsLotteryInfo) GetLotteryId() int32 {
	if m != nil && m.LotteryId != nil {
		return *m.LotteryId
	}
	return 0
}

func (m *HallAckDsLotteryInfo) GetDrawLotteryInfo() *HallDrawLotteryInfo {
	if m != nil {
		return m.DrawLotteryInfo
	}
	return nil
}

func (m *HallAckDsLotteryInfo) GetSignLotteryInfo() *HallSignLotteryInfo {
	if m != nil {
		return m.SignLotteryInfo
	}
	return nil
}

// 好友列表
type HallReqFriendsList struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqFriendsList) Reset()                    { *m = HallReqFriendsList{} }
func (m *HallReqFriendsList) String() string            { return proto.CompactTextString(m) }
func (*HallReqFriendsList) ProtoMessage()               {}
func (*HallReqFriendsList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{24} }

func (m *HallReqFriendsList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckFriendsList struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Users            []*HallFriendState `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallAckFriendsList) Reset()                    { *m = HallAckFriendsList{} }
func (m *HallAckFriendsList) String() string            { return proto.CompactTextString(m) }
func (*HallAckFriendsList) ProtoMessage()               {}
func (*HallAckFriendsList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{25} }

func (m *HallAckFriendsList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckFriendsList) GetUsers() []*HallFriendState {
	if m != nil {
		return m.Users
	}
	return nil
}

// 用户状态
type HallFriendState struct {
	IsOnline         *bool         `protobuf:"varint,1,opt,name=is_online" json:"is_online,omitempty"`
	Info             *HallUserInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *HallFriendState) Reset()                    { *m = HallFriendState{} }
func (m *HallFriendState) String() string            { return proto.CompactTextString(m) }
func (*HallFriendState) ProtoMessage()               {}
func (*HallFriendState) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{26} }

func (m *HallFriendState) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *HallFriendState) GetInfo() *HallUserInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

// 用户信息
type HallUserInfo struct {
	NickName         *string `protobuf:"bytes,1,opt,name=nick_name" json:"nick_name,omitempty"`
	HeadImg          *string `protobuf:"bytes,2,opt,name=head_img" json:"head_img,omitempty"`
	Sex              *int32  `protobuf:"varint,3,opt,name=sex" json:"sex,omitempty"`
	City             *int32  `protobuf:"varint,4,opt,name=city" json:"city,omitempty"`
	Sign             *string `protobuf:"bytes,5,opt,name=sign" json:"sign,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallUserInfo) Reset()                    { *m = HallUserInfo{} }
func (m *HallUserInfo) String() string            { return proto.CompactTextString(m) }
func (*HallUserInfo) ProtoMessage()               {}
func (*HallUserInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{27} }

func (m *HallUserInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *HallUserInfo) GetHeadImg() string {
	if m != nil && m.HeadImg != nil {
		return *m.HeadImg
	}
	return ""
}

func (m *HallUserInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *HallUserInfo) GetCity() int32 {
	if m != nil && m.City != nil {
		return *m.City
	}
	return 0
}

func (m *HallUserInfo) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

func init() {
	proto.RegisterType((*HallReqEvent)(nil), "ddproto.hall_req_event")
	proto.RegisterType((*HallAckEvent)(nil), "ddproto.hall_ack_event")
	proto.RegisterType((*HallLotteryItem)(nil), "ddproto.hall_lottery_item")
	proto.RegisterType((*HallSignLotteryInfo)(nil), "ddproto.hall_sign_lottery_info")
	proto.RegisterType((*HallDrawLotteryInfo)(nil), "ddproto.hall_draw_lottery_info")
	proto.RegisterType((*HallReqMailList)(nil), "ddproto.hall_req_mail_list")
	proto.RegisterType((*HallAckMailList)(nil), "ddproto.hall_ack_mail_list")
	proto.RegisterType((*HallReqTask)(nil), "ddproto.hall_req_task")
	proto.RegisterType((*HallAckTask)(nil), "ddproto.hall_ack_task")
	proto.RegisterType((*HallReqGetTask)(nil), "ddproto.hall_req_getTask")
	proto.RegisterType((*HallAckGetTask)(nil), "ddproto.hall_ack_getTask")
	proto.RegisterType((*HallReqBagItems)(nil), "ddproto.hall_req_bag_items")
	proto.RegisterType((*HallAckBagItems)(nil), "ddproto.hall_ack_bag_items")
	proto.RegisterType((*HallReqUserData)(nil), "ddproto.hall_req_userData")
	proto.RegisterType((*HallAckUserData)(nil), "ddproto.hall_ack_userData")
	proto.RegisterType((*HallReqGoodsList)(nil), "ddproto.hall_req_goods_list")
	proto.RegisterType((*HallAckGoodsList)(nil), "ddproto.hall_ack_goods_list")
	proto.RegisterType((*HallGoodsItemMsg)(nil), "ddproto.hall_goods_item_msg")
	proto.RegisterType((*HallReqRank)(nil), "ddproto.hall_req_rank")
	proto.RegisterType((*HallAckRank)(nil), "ddproto.hall_ack_rank")
	proto.RegisterType((*HallReqDrawLottery)(nil), "ddproto.hall_req_draw_lottery")
	proto.RegisterType((*HallAckDrawLottery)(nil), "ddproto.hall_ack_draw_lottery")
	proto.RegisterType((*HallReqDsLotteryInfo)(nil), "ddproto.hall_req_ds_lottery_info")
	proto.RegisterType((*HallAckDsLotteryInfo)(nil), "ddproto.hall_ack_ds_lottery_info")
	proto.RegisterType((*HallReqFriendsList)(nil), "ddproto.hall_req_friends_list")
	proto.RegisterType((*HallAckFriendsList)(nil), "ddproto.hall_ack_friends_list")
	proto.RegisterType((*HallFriendState)(nil), "ddproto.hall_friend_state")
	proto.RegisterType((*HallUserInfo)(nil), "ddproto.hall_user_info")
	proto.RegisterEnum("ddproto.HallEnumGoodsType", HallEnumGoodsType_name, HallEnumGoodsType_value)
	proto.RegisterEnum("ddproto.HallEnumTradeType", HallEnumTradeType_name, HallEnumTradeType_value)
}

var fileDescriptor13 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xed, 0x6a, 0x24, 0x45,
	0x14, 0xb5, 0x33, 0x33, 0x49, 0xe6, 0x26, 0x93, 0x74, 0x3a, 0x6b, 0x76, 0x08, 0x8a, 0x4b, 0xa3,
	0xb0, 0xae, 0x6c, 0xc0, 0x05, 0x65, 0x15, 0x04, 0x93, 0x9d, 0xe0, 0x0e, 0x24, 0x93, 0x26, 0xc9,
	0x8a, 0x0b, 0x62, 0x5b, 0x76, 0x57, 0x26, 0xc5, 0xf6, 0x47, 0xec, 0xaa, 0x59, 0x9d, 0x27, 0x10,
	0x9f, 0xc1, 0x07, 0xf2, 0x87, 0x2f, 0xe5, 0xad, 0xdb, 0x95, 0xfe, 0x60, 0x26, 0x6c, 0x77, 0x60,
	0xff, 0x0c, 0x53, 0xb7, 0xee, 0x39, 0xf7, 0x54, 0xdd, 0xaa, 0x53, 0x0d, 0x70, 0xcd, 0xa2, 0xe8,
	0xe0, 0x26, 0x4b, 0x55, 0xea, 0xac, 0x85, 0x21, 0xfd, 0xd9, 0xdf, 0xd6, 0x41, 0x3f, 0x64, 0x8a,
	0xe5, 0x33, 0xfb, 0xbb, 0x41, 0x1a, 0xc7, 0x69, 0xe2, 0x07, 0x91, 0xe0, 0x89, 0xca, 0x83, 0xee,
	0xd7, 0xb0, 0x45, 0x79, 0x19, 0xff, 0xdd, 0xe7, 0x6f, 0x31, 0xee, 0x7c, 0x0a, 0xab, 0xd7, 0x9c,
	0x85, 0x3c, 0x1b, 0x5a, 0x8f, 0xac, 0xc7, 0x1b, 0xcf, 0x1e, 0x1c, 0x18, 0xc6, 0x03, 0x4f, 0xff,
	0xbe, 0xa4, 0x39, 0x37, 0x30, 0x38, 0x16, 0xbc, 0x69, 0x83, 0x73, 0xbe, 0x80, 0x3e, 0xa5, 0x9f,
	0x08, 0xa9, 0x86, 0x2b, 0x8f, 0x3a, 0x98, 0x38, 0x2c, 0x12, 0x89, 0x51, 0x28, 0x1e, 0xe7, 0x94,
	0xee, 0x53, 0xd8, 0xa1, 0x50, 0x94, 0x2a, 0xc5, 0xb3, 0x39, 0x4d, 0x39, 0x00, 0x2b, 0x22, 0xa4,
	0x1a, 0x3d, 0x67, 0x13, 0xba, 0x09, 0x8b, 0x39, 0x12, 0x59, 0x8f, 0xfb, 0xee, 0x5f, 0x16, 0xec,
	0x51, 0xbe, 0x14, 0xd3, 0xa4, 0x04, 0x25, 0x57, 0xa9, 0xb3, 0x0b, 0x1b, 0x42, 0x5e, 0x60, 0xf8,
	0x32, 0x0d, 0xd9, 0x9c, 0xd0, 0xeb, 0xce, 0x0e, 0xf4, 0x55, 0xaa, 0x58, 0x34, 0x62, 0x73, 0x49,
	0x14, 0x3d, 0x67, 0x0f, 0xb6, 0x82, 0x34, 0x51, 0x22, 0x99, 0xa5, 0x33, 0x49, 0xf1, 0x0e, 0xc5,
	0x9f, 0x42, 0x5f, 0x93, 0x8e, 0x51, 0x80, 0x1c, 0x76, 0x49, 0xf6, 0x7e, 0x5d, 0x76, 0x55, 0xa3,
	0xfb, 0xb3, 0x11, 0x12, 0x66, 0xec, 0x8f, 0x25, 0x42, 0x46, 0x18, 0xae, 0x0a, 0x41, 0xf6, 0x50,
	0xb0, 0x28, 0x67, 0x5f, 0x79, 0x27, 0xfb, 0xb7, 0xe0, 0x14, 0x3d, 0x8b, 0x99, 0xc0, 0x59, 0xdc,
	0xcc, 0x86, 0x7d, 0x63, 0x06, 0xab, 0xfb, 0xd6, 0x12, 0xeb, 0x7c, 0x06, 0xdd, 0xa8, 0x6c, 0xdb,
	0xc3, 0xba, 0x42, 0x22, 0x23, 0x79, 0x5f, 0xc1, 0xa0, 0x90, 0xa7, 0x98, 0x7c, 0xd3, 0x50, 0xd9,
	0xaf, 0x06, 0xa6, 0x95, 0x35, 0x87, 0x39, 0x9f, 0xc3, 0xba, 0xce, 0x3e, 0xb9, 0x53, 0x18, 0x9d,
	0x27, 0x9d, 0xe2, 0x3e, 0x07, 0xbb, 0x10, 0x36, 0xe5, 0xea, 0xb2, 0xb9, 0xb6, 0x53, 0x83, 0xd4,
	0xda, 0x5a, 0x21, 0x75, 0xbf, 0xaf, 0x44, 0x22, 0xe4, 0x35, 0x0f, 0x27, 0xb3, 0x38, 0x3f, 0x65,
	0xb5, 0x06, 0xfe, 0xc6, 0xa6, 0x24, 0x51, 0xde, 0xa3, 0x81, 0x2d, 0xb1, 0xd8, 0xc0, 0x9e, 0xa8,
	0x9c, 0xb1, 0xbd, 0xfa, 0x46, 0xdd, 0xb2, 0xb9, 0xdf, 0x98, 0x6b, 0xa7, 0xe5, 0xcd, 0x24, 0xcf,
	0x46, 0xe8, 0x21, 0x0d, 0xd5, 0xf9, 0x06, 0xaa, 0xd5, 0xb5, 0x83, 0x62, 0x56, 0x57, 0x23, 0x68,
	0x8b, 0x16, 0xb4, 0xdd, 0x72, 0xb9, 0x09, 0xec, 0x96, 0x3d, 0x4c, 0xd3, 0x50, 0xb6, 0x39, 0xc0,
	0x5f, 0x02, 0xe4, 0x18, 0x35, 0xbf, 0xc9, 0x4d, 0x63, 0xeb, 0xd9, 0xc7, 0xf5, 0x42, 0x3c, 0x99,
	0xc5, 0x7e, 0x99, 0xe4, 0xfe, 0x63, 0x99, 0x82, 0xd4, 0xfa, 0xf7, 0x5f, 0x10, 0x0d, 0xd2, 0xf4,
	0xa8, 0x43, 0x3d, 0xfa, 0xa8, 0x9e, 0x9d, 0x27, 0xd2, 0x91, 0x8e, 0xe5, 0xd4, 0xfd, 0xef, 0x56,
	0x5d, 0x3d, 0xee, 0xd8, 0xb0, 0x6e, 0x22, 0x4b, 0x9d, 0x52, 0xeb, 0xba, 0xc9, 0x44, 0xc0, 0x73,
	0x5d, 0x9d, 0x3b, 0x75, 0xa9, 0x0c, 0x97, 0x91, 0xeb, 0x1a, 0x40, 0x8f, 0x20, 0xe8, 0x7e, 0x9a,
	0x0f, 0x19, 0xde, 0xb2, 0x68, 0x66, 0x18, 0x7a, 0x0d, 0x19, 0x08, 0x32, 0x5c, 0x25, 0x06, 0xd4,
	0x18, 0x0a, 0x19, 0xa4, 0xb3, 0x44, 0x0d, 0xd7, 0xc8, 0xbf, 0xab, 0xc6, 0x91, 0xb1, 0xe4, 0x3e,
	0xc6, 0xd1, 0x1c, 0xa6, 0x8d, 0x43, 0x67, 0xbf, 0xc3, 0x38, 0x74, 0x8a, 0xfb, 0x1d, 0x7c, 0x58,
	0x08, 0xab, 0x5a, 0x7a, 0x43, 0x81, 0x9e, 0x81, 0x6b, 0x81, 0xed, 0xe1, 0xfa, 0x99, 0x32, 0x80,
	0x71, 0x68, 0x0c, 0xe4, 0x7b, 0x18, 0x96, 0x82, 0x64, 0xfd, 0x85, 0x69, 0xa6, 0xe9, 0x5f, 0xcb,
	0x50, 0x90, 0xa8, 0xfb, 0x50, 0x2c, 0xd1, 0xe5, 0x3c, 0x87, 0x6d, 0xbd, 0xc0, 0x13, 0x13, 0x46,
	0xae, 0xe1, 0x26, 0x31, 0x7c, 0x52, 0xdf, 0xda, 0xc5, 0x77, 0x11, 0x91, 0xfa, 0x81, 0xad, 0x22,
	0x07, 0xcb, 0x90, 0x0b, 0x4f, 0x7b, 0xad, 0x39, 0x57, 0x19, 0x7e, 0xda, 0xb4, 0xba, 0xa2, 0xee,
	0x75, 0xa5, 0x39, 0xed, 0xe1, 0x78, 0x8a, 0x7a, 0xda, 0x9b, 0xee, 0x78, 0xb6, 0x73, 0x42, 0x5f,
	0x2a, 0xa6, 0x38, 0x3e, 0x22, 0x3b, 0x0b, 0x41, 0xbd, 0x89, 0x42, 0xfa, 0x69, 0x12, 0x89, 0x84,
	0x9b, 0xaf, 0x01, 0x7c, 0x66, 0xf5, 0xc2, 0x8c, 0x11, 0x3e, 0x5c, 0x34, 0xc2, 0x7c, 0xdd, 0xbf,
	0x98, 0x2f, 0xb0, 0x22, 0xa2, 0xb9, 0x12, 0x81, 0xcb, 0xa0, 0x8b, 0x6e, 0xd1, 0x45, 0xc7, 0x4b,
	0xa6, 0x17, 0xe1, 0x8b, 0x78, 0x6a, 0xae, 0xfe, 0x06, 0x74, 0x24, 0xff, 0xd3, 0x7c, 0xd6, 0xa0,
	0x2b, 0x04, 0x42, 0xcd, 0xcd, 0x9d, 0xc6, 0x91, 0xde, 0x5e, 0xba, 0xcd, 0xfd, 0x27, 0x57, 0xf0,
	0x60, 0xa9, 0x41, 0x0d, 0xa0, 0x7f, 0xf9, 0xda, 0x3b, 0xf6, 0x5f, 0x9c, 0x8d, 0x27, 0xb6, 0x85,
	0x15, 0x36, 0x69, 0x38, 0x1a, 0x1f, 0x9e, 0x9e, 0x4d, 0x46, 0xf6, 0x0a, 0xca, 0x18, 0x50, 0xe4,
	0xf8, 0xa7, 0x17, 0x2f, 0x0f, 0x27, 0x3f, 0x1c, 0xdb, 0x1d, 0x64, 0x5e, 0xa7, 0xd0, 0xd1, 0xab,
	0xd7, 0x76, 0xb7, 0x18, 0xfd, 0x38, 0xf6, 0xec, 0xde, 0x93, 0xbf, 0xad, 0x6a, 0xa1, 0x8a, 0x5f,
	0x6c, 0x01, 0x5c, 0x9e, 0x1f, 0x8e, 0x8a, 0x4a, 0x9a, 0x97, 0xc6, 0x65, 0xa9, 0x6d, 0xd8, 0xc8,
	0x43, 0x47, 0x67, 0x93, 0x57, 0x17, 0x58, 0xa8, 0x82, 0x19, 0x1d, 0x63, 0x29, 0x2d, 0x96, 0xc6,
	0xe7, 0xa7, 0x47, 0x76, 0xaf, 0xcc, 0xf7, 0xce, 0xcf, 0xbc, 0x0b, 0x7b, 0xb5, 0x9c, 0xd7, 0x5a,
	0xd6, 0xbc, 0x0f, 0x3c, 0xeb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xeb, 0x4e, 0x26, 0x4d,
	0x0b, 0x00, 0x00,
}

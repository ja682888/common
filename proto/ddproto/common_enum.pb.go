// Code generated by protoc-gen-go.
// source: common_enum.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏的id
type CommonEnumGame int32

const (
	CommonEnumGame_GID_SRC                    CommonEnumGame = 0
	CommonEnumGame_GID_TEXAS                  CommonEnumGame = 1
	CommonEnumGame_GID_MAHJONG                CommonEnumGame = 2
	CommonEnumGame_GID_DDZ                    CommonEnumGame = 3
	CommonEnumGame_GID_ZJH                    CommonEnumGame = 4
	CommonEnumGame_GID_HALL                   CommonEnumGame = 5
	CommonEnumGame_GID_PDK                    CommonEnumGame = 7
	CommonEnumGame_GID_ZXZ                    CommonEnumGame = 8
	CommonEnumGame_GID_MJBAISHAN              CommonEnumGame = 9
	CommonEnumGame_GID_MJCHANGCHUN            CommonEnumGame = 10
	CommonEnumGame_GID_NIUNIUJINGDIAN         CommonEnumGame = 11
	CommonEnumGame_GID_BANTUOZI               CommonEnumGame = 12
	CommonEnumGame_GID_DOUDIZHUERREN          CommonEnumGame = 13
	CommonEnumGame_GIDDOUDIZHUJINGDIANDONGBEI CommonEnumGame = 14
	CommonEnumGame_GID_TIANDAKENG             CommonEnumGame = 15
	CommonEnumGame_GID_MJ_CHANGCHUN           CommonEnumGame = 16
	CommonEnumGame_GID_BAIRENNIUNIU           CommonEnumGame = 17
	CommonEnumGame_GID_ZHIZUNWUZHANG          CommonEnumGame = 18
)

var CommonEnumGame_name = map[int32]string{
	0:  "GID_SRC",
	1:  "GID_TEXAS",
	2:  "GID_MAHJONG",
	3:  "GID_DDZ",
	4:  "GID_ZJH",
	5:  "GID_HALL",
	7:  "GID_PDK",
	8:  "GID_ZXZ",
	9:  "GID_MJBAISHAN",
	10: "GID_MJCHANGCHUN",
	11: "GID_NIUNIUJINGDIAN",
	12: "GID_BANTUOZI",
	13: "GID_DOUDIZHUERREN",
	14: "GIDDOUDIZHUJINGDIANDONGBEI",
	15: "GID_TIANDAKENG",
	16: "GID_MJ_CHANGCHUN",
	17: "GID_BAIRENNIUNIU",
	18: "GID_ZHIZUNWUZHANG",
}
var CommonEnumGame_value = map[string]int32{
	"GID_SRC":                    0,
	"GID_TEXAS":                  1,
	"GID_MAHJONG":                2,
	"GID_DDZ":                    3,
	"GID_ZJH":                    4,
	"GID_HALL":                   5,
	"GID_PDK":                    7,
	"GID_ZXZ":                    8,
	"GID_MJBAISHAN":              9,
	"GID_MJCHANGCHUN":            10,
	"GID_NIUNIUJINGDIAN":         11,
	"GID_BANTUOZI":               12,
	"GID_DOUDIZHUERREN":          13,
	"GIDDOUDIZHUJINGDIANDONGBEI": 14,
	"GID_TIANDAKENG":             15,
	"GID_MJ_CHANGCHUN":           16,
	"GID_BAIRENNIUNIU":           17,
	"GID_ZHIZUNWUZHANG":          18,
}

func (x CommonEnumGame) Enum() *CommonEnumGame {
	p := new(CommonEnumGame)
	*p = x
	return p
}
func (x CommonEnumGame) String() string {
	return proto.EnumName(CommonEnumGame_name, int32(x))
}
func (x *CommonEnumGame) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumGame_value, data, "CommonEnumGame")
	if err != nil {
		return err
	}
	*x = CommonEnumGame(value)
	return nil
}
func (CommonEnumGame) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// 房间的类型
type COMMON_ENUM_ROOMTYPE int32

const (
	COMMON_ENUM_ROOMTYPE_DESK_FRIEND COMMON_ENUM_ROOMTYPE = 1
	COMMON_ENUM_ROOMTYPE_DESK_COIN   COMMON_ENUM_ROOMTYPE = 2
)

var COMMON_ENUM_ROOMTYPE_name = map[int32]string{
	1: "DESK_FRIEND",
	2: "DESK_COIN",
}
var COMMON_ENUM_ROOMTYPE_value = map[string]int32{
	"DESK_FRIEND": 1,
	"DESK_COIN":   2,
}

func (x COMMON_ENUM_ROOMTYPE) Enum() *COMMON_ENUM_ROOMTYPE {
	p := new(COMMON_ENUM_ROOMTYPE)
	*p = x
	return p
}
func (x COMMON_ENUM_ROOMTYPE) String() string {
	return proto.EnumName(COMMON_ENUM_ROOMTYPE_name, int32(x))
}
func (x *COMMON_ENUM_ROOMTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_ROOMTYPE_value, data, "COMMON_ENUM_ROOMTYPE")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_ROOMTYPE(value)
	return nil
}
func (COMMON_ENUM_ROOMTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

// 游戏的状态
type COMMON_ENUM_GAMESTATUS int32

const (
	COMMON_ENUM_GAMESTATUS_NOGAME COMMON_ENUM_GAMESTATUS = 1
	COMMON_ENUM_GAMESTATUS_GAMING COMMON_ENUM_GAMESTATUS = 2
)

var COMMON_ENUM_GAMESTATUS_name = map[int32]string{
	1: "NOGAME",
	2: "GAMING",
}
var COMMON_ENUM_GAMESTATUS_value = map[string]int32{
	"NOGAME": 1,
	"GAMING": 2,
}

func (x COMMON_ENUM_GAMESTATUS) Enum() *COMMON_ENUM_GAMESTATUS {
	p := new(COMMON_ENUM_GAMESTATUS)
	*p = x
	return p
}
func (x COMMON_ENUM_GAMESTATUS) String() string {
	return proto.EnumName(COMMON_ENUM_GAMESTATUS_name, int32(x))
}
func (x *COMMON_ENUM_GAMESTATUS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_GAMESTATUS_value, data, "COMMON_ENUM_GAMESTATUS")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_GAMESTATUS(value)
	return nil
}
func (COMMON_ENUM_GAMESTATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

// 发布版本
type COMMON_ENUM_RELEASETAG int32

const (
	COMMON_ENUM_RELEASETAG_R_PRO          COMMON_ENUM_RELEASETAG = 1
	COMMON_ENUM_RELEASETAG_R_APPLE_VERIFY COMMON_ENUM_RELEASETAG = 2
)

var COMMON_ENUM_RELEASETAG_name = map[int32]string{
	1: "R_PRO",
	2: "R_APPLE_VERIFY",
}
var COMMON_ENUM_RELEASETAG_value = map[string]int32{
	"R_PRO":          1,
	"R_APPLE_VERIFY": 2,
}

func (x COMMON_ENUM_RELEASETAG) Enum() *COMMON_ENUM_RELEASETAG {
	p := new(COMMON_ENUM_RELEASETAG)
	*p = x
	return p
}
func (x COMMON_ENUM_RELEASETAG) String() string {
	return proto.EnumName(COMMON_ENUM_RELEASETAG_name, int32(x))
}
func (x *COMMON_ENUM_RELEASETAG) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_RELEASETAG_value, data, "COMMON_ENUM_RELEASETAG")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_RELEASETAG(value)
	return nil
}
func (COMMON_ENUM_RELEASETAG) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

// 踢人的踢出类型
type COMMON_ENUM_KICKOUT int32

const (
	COMMON_ENUM_KICKOUT_K_COINUNMATCHED              COMMON_ENUM_KICKOUT = 1
	COMMON_ENUM_KICKOUT_K_SYSTEM                     COMMON_ENUM_KICKOUT = 2
	COMMON_ENUM_KICKOUT_K_OFFLINE                    COMMON_ENUM_KICKOUT = 3
	COMMON_ENUM_KICKOUT_K_GAME_BLOCKED               COMMON_ENUM_KICKOUT = 4
	COMMON_ENUM_KICKOUT_K_TIMEOUT_NOTREADY_LOTTERY   COMMON_ENUM_KICKOUT = 5
	COMMON_ENUM_KICKOUT_K_TIMEOUT_NOTREADY_ENTERDESK COMMON_ENUM_KICKOUT = 6
)

var COMMON_ENUM_KICKOUT_name = map[int32]string{
	1: "K_COINUNMATCHED",
	2: "K_SYSTEM",
	3: "K_OFFLINE",
	4: "K_GAME_BLOCKED",
	5: "K_TIMEOUT_NOTREADY_LOTTERY",
	6: "K_TIMEOUT_NOTREADY_ENTERDESK",
}
var COMMON_ENUM_KICKOUT_value = map[string]int32{
	"K_COINUNMATCHED":              1,
	"K_SYSTEM":                     2,
	"K_OFFLINE":                    3,
	"K_GAME_BLOCKED":               4,
	"K_TIMEOUT_NOTREADY_LOTTERY":   5,
	"K_TIMEOUT_NOTREADY_ENTERDESK": 6,
}

func (x COMMON_ENUM_KICKOUT) Enum() *COMMON_ENUM_KICKOUT {
	p := new(COMMON_ENUM_KICKOUT)
	*p = x
	return p
}
func (x COMMON_ENUM_KICKOUT) String() string {
	return proto.EnumName(COMMON_ENUM_KICKOUT_name, int32(x))
}
func (x *COMMON_ENUM_KICKOUT) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_KICKOUT_value, data, "COMMON_ENUM_KICKOUT")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_KICKOUT(value)
	return nil
}
func (COMMON_ENUM_KICKOUT) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

// 玩家申请解散房间的操作类型
type COMMON_ENUM_APPLYDISSOLVE int32

const (
	COMMON_ENUM_APPLYDISSOLVE_AD_AGREE  COMMON_ENUM_APPLYDISSOLVE = 1
	COMMON_ENUM_APPLYDISSOLVE_AD_REFUSE COMMON_ENUM_APPLYDISSOLVE = 2
	COMMON_ENUM_APPLYDISSOLVE_AD_NOACT  COMMON_ENUM_APPLYDISSOLVE = 3
)

var COMMON_ENUM_APPLYDISSOLVE_name = map[int32]string{
	1: "AD_AGREE",
	2: "AD_REFUSE",
	3: "AD_NOACT",
}
var COMMON_ENUM_APPLYDISSOLVE_value = map[string]int32{
	"AD_AGREE":  1,
	"AD_REFUSE": 2,
	"AD_NOACT":  3,
}

func (x COMMON_ENUM_APPLYDISSOLVE) Enum() *COMMON_ENUM_APPLYDISSOLVE {
	p := new(COMMON_ENUM_APPLYDISSOLVE)
	*p = x
	return p
}
func (x COMMON_ENUM_APPLYDISSOLVE) String() string {
	return proto.EnumName(COMMON_ENUM_APPLYDISSOLVE_name, int32(x))
}
func (x *COMMON_ENUM_APPLYDISSOLVE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_APPLYDISSOLVE_value, data, "COMMON_ENUM_APPLYDISSOLVE")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_APPLYDISSOLVE(value)
	return nil
}
func (COMMON_ENUM_APPLYDISSOLVE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

// 按钮的类型
type BTN_TYPE int32

const (
	BTN_TYPE_BTNTYPE_NEWUSERBUTTON     BTN_TYPE = 1
	BTN_TYPE_BTNTYPE_SHARELINKTIMELINE BTN_TYPE = 2
	BTN_TYPE_REDBAGSHAR                BTN_TYPE = 3
)

var BTN_TYPE_name = map[int32]string{
	1: "BTNTYPE_NEWUSERBUTTON",
	2: "BTNTYPE_SHARELINKTIMELINE",
	3: "REDBAGSHAR",
}
var BTN_TYPE_value = map[string]int32{
	"BTNTYPE_NEWUSERBUTTON":     1,
	"BTNTYPE_SHARELINKTIMELINE": 2,
	"REDBAGSHAR":                3,
}

func (x BTN_TYPE) Enum() *BTN_TYPE {
	p := new(BTN_TYPE)
	*p = x
	return p
}
func (x BTN_TYPE) String() string {
	return proto.EnumName(BTN_TYPE_name, int32(x))
}
func (x *BTN_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BTN_TYPE_value, data, "BTN_TYPE")
	if err != nil {
		return err
	}
	*x = BTN_TYPE(value)
	return nil
}
func (BTN_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

// 进入金币场房间的错误类型
type COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM int32

const (
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_UNDER_MIN       COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 1
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_OVER_LIMIT      COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 2
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_UNDER_LIMIT     COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 3
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_OTHER_LV_DESK_GAMING COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 4
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_NONE_DESK_AVAILABLE  COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 5
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_OTHER                COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 6
)

var COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name = map[int32]string{
	1: "ERROR_EC_COIN_UNDER_MIN",
	2: "ERROR_EC_COIN_OVER_LIMIT",
	3: "ERROR_EC_COIN_UNDER_LIMIT",
	4: "ERROR_EC_OTHER_LV_DESK_GAMING",
	5: "ERROR_EC_NONE_DESK_AVAILABLE",
	6: "ERROR_EC_OTHER",
}
var COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value = map[string]int32{
	"ERROR_EC_COIN_UNDER_MIN":       1,
	"ERROR_EC_COIN_OVER_LIMIT":      2,
	"ERROR_EC_COIN_UNDER_LIMIT":     3,
	"ERROR_EC_OTHER_LV_DESK_GAMING": 4,
	"ERROR_EC_NONE_DESK_AVAILABLE":  5,
	"ERROR_EC_OTHER":                6,
}

func (x COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) Enum() *COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM {
	p := new(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM)
	*p = x
	return p
}
func (x COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) String() string {
	return proto.EnumName(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name, int32(x))
}
func (x *COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value, data, "COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM(value)
	return nil
}
func (COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{7}
}

func init() {
	proto.RegisterEnum("ddproto.CommonEnumGame", CommonEnumGame_name, CommonEnumGame_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ROOMTYPE", COMMON_ENUM_ROOMTYPE_name, COMMON_ENUM_ROOMTYPE_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_GAMESTATUS", COMMON_ENUM_GAMESTATUS_name, COMMON_ENUM_GAMESTATUS_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_RELEASETAG", COMMON_ENUM_RELEASETAG_name, COMMON_ENUM_RELEASETAG_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_KICKOUT", COMMON_ENUM_KICKOUT_name, COMMON_ENUM_KICKOUT_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_APPLYDISSOLVE", COMMON_ENUM_APPLYDISSOLVE_name, COMMON_ENUM_APPLYDISSOLVE_value)
	proto.RegisterEnum("ddproto.BTN_TYPE", BTN_TYPE_name, BTN_TYPE_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM", COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name, COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value)
}

<<<<<<< HEAD
var fileDescriptor6 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x73, 0xf3, 0x34,
	0x10, 0xc6, 0x89, 0xd3, 0xb4, 0xe9, 0x36, 0x6d, 0x37, 0x7a, 0xff, 0x40, 0xe1, 0x2d, 0x03, 0x33,
	0x9c, 0x7c, 0x60, 0x38, 0xc1, 0x59, 0xb6, 0x36, 0xb6, 0x62, 0x7b, 0x95, 0x91, 0xa5, 0xb4, 0xf1,
	0x45, 0xc3, 0xd0, 0x0e, 0xa7, 0x34, 0x0c, 0x03, 0x1f, 0x87, 0x0f, 0xc6, 0x91, 0x6f, 0xc2, 0xc8,
	0xae, 0xfb, 0xa6, 0x33, 0xbd, 0x69, 0xf7, 0xd1, 0xae, 0x7e, 0xfb, 0xac, 0x60, 0xf9, 0xdb, 0x61,
	0xbf, 0x3f, 0x3c, 0x85, 0xc7, 0xa7, 0xbf, 0xf7, 0x3f, 0xfe, 0xf1, 0xe7, 0xe1, 0xaf, 0x83, 0x38,
	0x7b, 0x78, 0xe8, 0x0f, 0xe9, 0x7f, 0x09, 0xe0, 0x91, 0x1c, 0x7e, 0xff, 0x75, 0xff, 0x28, 0x2e,
	0xe0, 0xac, 0xd0, 0x2a, 0xb4, 0x36, 0xc7, 0x2f, 0xc4, 0x25, 0x9c, 0xc7, 0xc0, 0xd1, 0xbd, 0x6c,
	0x71, 0x22, 0xae, 0xe1, 0x22, 0x86, 0x8d, 0x2c, 0xd7, 0x86, 0x0b, 0x4c, 0xc6, 0xcb, 0x4a, 0x75,
	0x38, 0x1d, 0x83, 0x6e, 0x5d, 0xe2, 0x89, 0x58, 0xc0, 0x3c, 0x06, 0xa5, 0xac, 0x6b, 0x9c, 0x8d,
	0xd2, 0x46, 0x55, 0x78, 0xf6, 0x72, 0xef, 0xbe, 0xc3, 0xb9, 0x58, 0xc2, 0x65, 0xdf, 0x72, 0x9d,
	0x49, 0xdd, 0x96, 0x92, 0xf1, 0x5c, 0xbc, 0x83, 0xeb, 0x21, 0x95, 0x97, 0x92, 0x8b, 0xbc, 0xf4,
	0x8c, 0x20, 0x3e, 0x82, 0x88, 0x49, 0xd6, 0x9e, 0xb5, 0x5f, 0x6b, 0x2e, 0x94, 0x96, 0x8c, 0x17,
	0x02, 0x61, 0x11, 0xf3, 0x99, 0x64, 0xe7, 0x4d, 0xa7, 0x71, 0x21, 0x3e, 0xc0, 0xb2, 0x67, 0x32,
	0x5e, 0xe9, 0xae, 0xf4, 0x64, 0x2d, 0x31, 0x5e, 0x8a, 0x6f, 0xe1, 0xeb, 0x42, 0xab, 0x31, 0x3b,
	0x76, 0x50, 0x86, 0x8b, 0x8c, 0x34, 0x5e, 0x09, 0x01, 0x57, 0xfd, 0xa8, 0x31, 0x29, 0x2b, 0xe2,
	0x02, 0xaf, 0xc5, 0x7b, 0xc0, 0x81, 0x24, 0x7c, 0x46, 0xc1, 0x31, 0x9b, 0x49, 0x6d, 0x89, 0x07,
	0x20, 0x5c, 0x8e, 0xcf, 0x76, 0xa5, 0xee, 0x3c, 0xdf, 0xf9, 0x2e, 0x16, 0xa0, 0x48, 0x7f, 0x86,
	0xf7, 0xb9, 0x69, 0x1a, 0xc3, 0x81, 0xd8, 0x37, 0xc1, 0x1a, 0xd3, 0xb8, 0xdd, 0x86, 0xa2, 0x95,
	0x8a, 0xda, 0x2a, 0xac, 0xac, 0x26, 0x56, 0x38, 0x89, 0x56, 0xf7, 0x89, 0xdc, 0x68, 0xc6, 0x24,
	0xfd, 0x09, 0x3e, 0x1e, 0xd7, 0x15, 0xb2, 0xa1, 0xd6, 0x49, 0xe7, 0x5b, 0x01, 0x70, 0xca, 0x26,
	0xc6, 0x38, 0x89, 0xe7, 0x42, 0x36, 0x3a, 0xee, 0x22, 0xfd, 0xe5, 0x75, 0x85, 0xa5, 0x9a, 0x64,
	0x4b, 0x4e, 0x16, 0xe2, 0x1c, 0x66, 0x36, 0x6c, 0xac, 0xc1, 0x49, 0x9c, 0xd2, 0x06, 0xb9, 0xd9,
	0xd4, 0x14, 0xb6, 0x64, 0xf5, 0x6a, 0x87, 0x49, 0xfa, 0xcf, 0x04, 0xde, 0x1d, 0x57, 0x56, 0x3a,
	0xaf, 0x8c, 0x77, 0x71, 0x0f, 0x03, 0x8e, 0xe7, 0x46, 0xba, 0xbc, 0xa4, 0x88, 0xb9, 0x80, 0x79,
	0x15, 0xda, 0x5d, 0xeb, 0xa8, 0xc1, 0x24, 0x42, 0x57, 0xc1, 0xac, 0x56, 0xb5, 0x66, 0xc2, 0x69,
	0xec, 0x5e, 0xf5, 0xa8, 0x21, 0xab, 0x4d, 0x5e, 0x91, 0xc2, 0x93, 0xe8, 0x7b, 0x15, 0x9c, 0x6e,
	0xc8, 0x78, 0x17, 0xd8, 0x38, 0x4b, 0x52, 0xed, 0x42, 0x6d, 0x9c, 0x23, 0xbb, 0xc3, 0x99, 0xf8,
	0x0e, 0x3e, 0xbd, 0xa1, 0x13, 0x3b, 0xb2, 0xd1, 0x0f, 0x3c, 0x4d, 0x57, 0x70, 0x73, 0x8c, 0x17,
	0xe9, 0x77, 0x4a, 0xb7, 0xad, 0xa9, 0xb7, 0x14, 0x79, 0xa4, 0x0a, 0xb2, 0xb0, 0x44, 0x83, 0x89,
	0x52, 0x05, 0x4b, 0x2b, 0xdf, 0x12, 0x26, 0xcf, 0x22, 0x1b, 0x99, 0x3b, 0x9c, 0xa6, 0x0e, 0xe6,
	0x99, 0xe3, 0xd0, 0xdb, 0x7f, 0x03, 0x1f, 0x32, 0xc7, 0xf1, 0x18, 0x98, 0xee, 0x7c, 0x4b, 0x36,
	0xf3, 0xce, 0x19, 0xc6, 0x89, 0xb8, 0x85, 0x9b, 0x51, 0x6a, 0x4b, 0x69, 0xa9, 0xd6, 0x5c, 0x45,
	0xbe, 0x7e, 0xc6, 0x44, 0x5c, 0x01, 0x58, 0x52, 0x99, 0x2c, 0xa2, 0x88, 0xd3, 0xf4, 0xdf, 0x09,
	0xfc, 0x70, 0x8c, 0x47, 0xd6, 0x1a, 0xdb, 0x3f, 0x32, 0x0c, 0x11, 0x0d, 0x8c, 0x5b, 0x17, 0xdf,
	0xc0, 0x97, 0x83, 0x46, 0x79, 0xef, 0x6a, 0xf0, 0xac, 0xc8, 0x86, 0x46, 0xc7, 0x47, 0x3f, 0xc1,
	0x57, 0xaf, 0x45, 0xb3, 0x25, 0x1b, 0x6a, 0xdd, 0x68, 0x87, 0x49, 0x44, 0x7a, 0xab, 0x74, 0x90,
	0xa7, 0xe2, 0x7b, 0xb8, 0x7d, 0x91, 0x8d, 0x2b, 0xa3, 0xb2, 0x0d, 0xfd, 0x67, 0x7a, 0xfe, 0x1c,
	0x27, 0xd1, 0xe5, 0x97, 0x2b, 0x6c, 0x98, 0x06, 0x59, 0x6e, 0xa5, 0xae, 0x65, 0x56, 0x13, 0xce,
	0xe2, 0xee, 0x5e, 0x37, 0xc1, 0xd3, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x02, 0xce, 0xf1, 0x93,
	0x3d, 0x04, 0x00, 0x00,
=======
var fileDescriptor5 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xa6, 0xe9, 0xff, 0xf4, 0xcf, 0x71, 0x7f, 0xa0, 0xd0, 0x22, 0x90, 0x38, 0xe5, 0x80, 0x38,
	0xc1, 0xd9, 0xbb, 0x9e, 0xec, 0x3a, 0xbb, 0x6b, 0x47, 0x5e, 0x3b, 0x6d, 0xf6, 0x62, 0x21, 0x5a,
	0x71, 0x4a, 0x83, 0x10, 0x3c, 0x0e, 0x0f, 0xc6, 0x91, 0x37, 0x61, 0xbc, 0xcb, 0x96, 0x44, 0xaa,
	0xe4, 0xc3, 0xcc, 0x7c, 0x33, 0xe3, 0x6f, 0xbe, 0x19, 0x18, 0x7e, 0x59, 0x2e, 0x16, 0xcb, 0x87,
	0x70, 0xff, 0xf0, 0x73, 0xf1, 0xfe, 0xdb, 0xf7, 0xe5, 0x8f, 0x25, 0xdf, 0xbd, 0xbb, 0x6b, 0x8d,
	0xd1, 0x9f, 0x01, 0xb0, 0x15, 0x38, 0x7c, 0xfd, 0xbc, 0xb8, 0xe7, 0x07, 0xb0, 0x9b, 0x29, 0x19,
	0x6a, 0x9b, 0xb2, 0x67, 0xfc, 0x08, 0xf6, 0xa3, 0xe3, 0xf0, 0x56, 0xd4, 0x6c, 0x83, 0x9f, 0xc0,
	0x41, 0x74, 0x2b, 0x91, 0x4f, 0x8c, 0xce, 0xd8, 0xa0, 0x4f, 0x96, 0xb2, 0x61, 0x9b, 0xbd, 0xd3,
	0x4c, 0x72, 0xb6, 0xc5, 0x0f, 0x61, 0x2f, 0x3a, 0xb9, 0x28, 0x4b, 0xb6, 0xdd, 0x43, 0x53, 0x59,
	0xb0, 0xdd, 0xc7, 0xbc, 0xdb, 0x86, 0xed, 0xf1, 0x21, 0x1c, 0xb5, 0x2d, 0x27, 0x89, 0x50, 0x75,
	0x2e, 0x34, 0xdb, 0xe7, 0xa7, 0x70, 0xd2, 0x85, 0x52, 0xf2, 0xb3, 0x34, 0xf7, 0x9a, 0x01, 0xbf,
	0x00, 0x1e, 0x83, 0x5a, 0x79, 0x7a, 0x13, 0xa5, 0x33, 0xa9, 0x28, 0xf9, 0x80, 0x33, 0x38, 0x8c,
	0xf1, 0x44, 0x68, 0xe7, 0x4d, 0xa3, 0xd8, 0x21, 0x3f, 0x87, 0x61, 0xcb, 0xc9, 0x78, 0xa9, 0x9a,
	0xdc, 0xa3, 0xb5, 0xa8, 0xd9, 0x11, 0x7f, 0x0d, 0x2f, 0x29, 0xdc, 0x47, 0xfb, 0x0e, 0x92, 0xe6,
	0x48, 0x50, 0xb1, 0x63, 0xce, 0xe1, 0xb8, 0x1d, 0x35, 0x06, 0x45, 0x81, 0x34, 0xde, 0x09, 0x3f,
	0x03, 0xd6, 0x31, 0x09, 0xff, 0xa9, 0xb0, 0x3e, 0x4a, 0x84, 0xa9, 0x73, 0x47, 0x88, 0x0d, 0xfb,
	0x6f, 0x9b, 0x5c, 0x35, 0x5e, 0xdf, 0xf8, 0x26, 0x16, 0x30, 0x3e, 0xfa, 0x08, 0x67, 0xa9, 0xa9,
	0x2a, 0xa3, 0x03, 0x6a, 0x5f, 0x05, 0x6b, 0x4c, 0xe5, 0xe6, 0x53, 0x8c, 0x52, 0x4a, 0xac, 0x8b,
	0x30, 0xb6, 0x0a, 0xb5, 0x24, 0x6d, 0x49, 0xea, 0x36, 0x90, 0x1a, 0xa5, 0xd9, 0x60, 0xf4, 0x01,
	0x2e, 0x56, 0xeb, 0x32, 0x51, 0x61, 0xed, 0x84, 0xf3, 0x35, 0x07, 0xd8, 0xd1, 0x26, 0xfa, 0x54,
	0x44, 0x36, 0x59, 0x2a, 0xee, 0x62, 0xf4, 0x69, 0xbd, 0xc2, 0x62, 0x89, 0xa2, 0x46, 0x27, 0x32,
	0xbe, 0x0f, 0xdb, 0x36, 0x4c, 0xad, 0xa1, 0x02, 0x9a, 0xd2, 0x06, 0x31, 0x9d, 0x96, 0x18, 0x66,
	0x68, 0xd5, 0x78, 0x4e, 0x85, 0xbf, 0x36, 0xe0, 0x74, 0xb5, 0xb2, 0x50, 0x69, 0x61, 0xbc, 0x8b,
	0x7b, 0xe8, 0xe8, 0x78, 0x5d, 0x09, 0x97, 0xe6, 0x18, 0x69, 0xd2, 0x5e, 0x8b, 0x50, 0xcf, 0x6b,
	0x87, 0x15, 0xed, 0x9f, 0x48, 0x17, 0xc1, 0x8c, 0xc7, 0xa5, 0xd2, 0x48, 0x17, 0x40, 0xdd, 0x8b,
	0x96, 0x6a, 0x48, 0x4a, 0x93, 0x16, 0x54, 0xb0, 0x15, 0x75, 0x2f, 0x48, 0xd5, 0x0a, 0xa9, 0x65,
	0xd0, 0xc6, 0x59, 0x14, 0x72, 0x1e, 0x4a, 0xe3, 0x1c, 0xda, 0x39, 0x9d, 0xc6, 0x1b, 0xb8, 0x7a,
	0x02, 0x47, 0x4d, 0x70, 0xd4, 0x83, 0xed, 0x8c, 0xc6, 0x70, 0xb9, 0x4a, 0x2f, 0xb2, 0x9f, 0x4b,
	0x55, 0xd7, 0xa6, 0x9c, 0x61, 0xe4, 0x23, 0x64, 0x10, 0x99, 0x45, 0xec, 0x44, 0x24, 0xcf, 0xe2,
	0xd8, 0xd7, 0x48, 0xf4, 0x3a, 0x50, 0x1b, 0x91, 0x3a, 0xb6, 0x39, 0x72, 0xb0, 0x97, 0x38, 0x1d,
	0x5a, 0xf9, 0x2f, 0xe1, 0x9c, 0xec, 0x68, 0x06, 0x8d, 0x37, 0x94, 0x6d, 0x13, 0xef, 0x9c, 0xd1,
	0xd4, 0xe3, 0x1a, 0x2e, 0x7b, 0x88, 0x0e, 0x92, 0x64, 0x54, 0xba, 0x88, 0xfc, 0xda, 0x19, 0x07,
	0xfc, 0x18, 0xc0, 0xa2, 0x4c, 0x44, 0x16, 0x41, 0xea, 0xfa, 0x7b, 0x03, 0xde, 0xad, 0xd2, 0xa3,
	0x7b, 0x33, 0xb6, 0xfd, 0xa4, 0x1b, 0x22, 0x0a, 0x18, 0xb7, 0xce, 0x5f, 0xc1, 0xf3, 0x0e, 0xc3,
	0xb4, 0x55, 0x35, 0x78, 0x2d, 0xd1, 0x06, 0xda, 0x1e, 0x7d, 0x7a, 0x05, 0x2f, 0xd6, 0x41, 0x43,
	0xeb, 0x09, 0xa5, 0xaa, 0x94, 0xa3, 0x3f, 0x89, 0xd2, 0x53, 0xa5, 0x1d, 0xbc, 0xc9, 0xdf, 0xc2,
	0xf5, 0x23, 0x6c, 0x5c, 0x1e, 0x91, 0x59, 0x68, 0x8f, 0xe9, 0xdf, 0x71, 0x6c, 0x45, 0x95, 0x1f,
	0x53, 0xb4, 0xd1, 0xd8, 0xc1, 0x62, 0x26, 0x54, 0x29, 0x92, 0x12, 0x69, 0x0f, 0xb4, 0xbb, 0xf5,
	0x26, 0x6c, 0xe7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xce, 0xf1, 0x93, 0x3d, 0x04, 0x00,
	0x00,
>>>>>>> 42e0f119e4612ba0e4faaf5279c670789cb90f36
}

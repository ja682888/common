package api

import (
	"casino_common/api/majiang"
	"casino_common/proto/ddproto"
	"sort"
)

//跑得快解析器
type MJParser interface {
	CanPeng(interface{}, interface{}) (bool, error)
	CanGang(...interface{}) (interface{}, error)
	CanChi(interface{}, interface{}) (interface{}, error)
	CanBu(...interface{}) (interface{}, error)
	CanTing(...interface{}) (interface{}, error)      //是否可以报听 白山麻将
	CanFly(interface{}, interface{}) (bool, error)    //是否可以飞 宜宾麻将
	CanTi(interface{}) (interface{}, error)           //是否可以提 宜宾麻将
	GetJiaoInfos(...interface{}) (interface{}, error) //判断是否有叫
	GetTingInfos(...interface{}) (interface{}, error) //获取听牌
	Parse(pids []int32) (interface{}, error)          //通过一副牌的id解析牌型
	XiPai() interface{}                               //洗牌
	Hu(...interface{}) (interface{}, error)           //胡牌的方式...
	CanHu(...interface{}) (interface{}, error)        //能否胡牌 不带番数得分计算...
	InitMjPaiByIndex(index int32) *majiang.MJPAI      //通过id得到一张麻将牌
	IsTingYongPai(pai majiang.MJPAI) bool             //是否是听用牌 宜宾麻将
	IsMengQing(g MJUserGameData) bool                 //是否是门清 没有吃碰杠 白山麻将
}

//麻将的骨架，通用的方法都在这里
type MJParserCore struct {
}

//得到叫牌的信息
type JiaoInfoBean struct {
	HuPai *majiang.MJPAI
	Fan   int32
	Count int32
}

type JiaoInfo struct {
	OutPai *majiang.MJPAI  //打出去的牌
	Jiaos  []*JiaoInfoBean //打牌之后的叫牌
}

//

//统计牌 27这个谁需要考虑 东南西北发中白的情况？
func (p *MJParserCore) CountHandPais(pais []*majiang.MJPAI) []int {
	//log.T("开始统计牌的数量:%v", utils.List2Str(pais))
	//0 ~ 45
	// 0 ~ 8 万
	// 9 ~ 17 索
	// 18 ~ 26 筒
	// 27 ~ 35 东27 南28 西29 北30 中31 白32 发33

	//counts := make([]int, 27) //0~27
	counts := make([]int, 36) //0~35
	for _, p := range pais {
		counts[p.GetCountIndex() ] ++
	}
	//log.T("统计出来的count:%v", counts)
	return counts
}

//空实现
func (p *MJParserCore) CanHu(...interface{}) (interface{}, error) {
	return nil, nil
}

//是否能碰牌
func (p *MJParserCore) CanPeng(userGameData interface{}, pengPai interface{}) (bool, error) {
	//log.T("判断是否能碰:userGameData.type%v,content:%v \n", reflect.TypeOf(userGameData), userGameData)
	gameData := userGameData.(MJUserGameData) //玩家数据
	handPais := gameData.GetHandPais()        //判断的手牌
	pengPai2 := pengPai.(*majiang.MJPAI)      //判断的碰牌

	paiCount := 0
	for _, p := range handPais {
		if p.Value == pengPai2.Value && p.Flower == pengPai2.Flower {
			paiCount++
		}
	}
	return paiCount >= 2, nil
}

//是否能杠牌
/**
	判断别人打的牌是否能杠(这里都是点杠)
 */
func (p *MJParserCore) CanGang(userGameData interface{}, gangPai interface{}) (interface{}, error) {
	gameData := userGameData.(MJUserGameData) //玩家数据
	handPais := gameData.GetHandPais()        //判断的手牌
	gangPai2 := gangPai.(*majiang.MJPAI)      //判断的碰牌

	paiCount := 0
	for _, p := range handPais {
		if p.Value == gangPai2.Value && p.Flower == gangPai2.Flower {
			paiCount++
		}
	}

	ret := &CanGangInfo{}
	if paiCount == 3 {
		ret.CanGang = true
		ret.GangInfoBean = append(ret.GangInfoBean, &CanGangInfoBean{
			GangPai:  gangPai2,
			GangType: GANGTYPE_MING,
		})
	}
	return ret, nil
}

/**
	自己摸的牌，有哪些可以杠的牌
 */

var (
	GANGTYPE_MING int32 = 1
	GANGTYPE_BA   int32 = 2
	GANGTYPE_AN   int32 = 3
	GANGTYPE_BU   int32 = 4
	GANGTYPE_FENG int32 = 5
)

type CanGangInfo struct {
	OutUserId    uint32             //打牌的人
	GangsUserId  uint32             //杠牌的人
	CanGang      bool               //是否能杠
	GangInfoBean []*CanGangInfoBean //这里可以使用 map 来存储
}

type CanGangInfoBean struct {
	GangPai  *majiang.MJPAI
	GangType int32 //杠牌的类型
}

//新的结构
type CanGangInfo2 struct {
	OutUserId   uint32      //打牌的人
	GangsUserId uint32      //杠牌的人
	CanGang     bool        //是否能杠
	GangBean    []*GangBean //这里可以使用 map 来存储
}

type GangBean struct {
	Pais []*majiang.MJPAI //能杠的4张牌
	Type int32            //杠牌的类型 普通杠1~3: 1点杠 2巴杠 3暗杠 补充: 4补杠 5风杠(旋风杠、喜杠)
}

var (
	PAI_TYPE_MENQING int32 = int32(ddproto.PaiType_H_MenQing) //门清
)
//胡牌之后的信息
type CanHuInfo struct {
	CanHu    bool
	HuDesc   []string
	Fan      int32
	Score    int64
	OutUser  uint32
	HuUser   uint32
	IsZimo   bool           //自摸
	IsBanker bool           //指胡牌的人是否是庄
	Pai      *majiang.MJPAI //胡的牌
	PaiType  []int32
}

type CanChiInfo struct {
	OutUserId uint32 //打牌的人
	ChiUserId uint32 //吃牌的人
	CanChi    bool   //是否能吃
	ChiBeans  []ChiBean
}

type ChiBean []*majiang.MJPAI

type CanBuInfo struct {
	CanBu   bool
	BuCards []*majiang.MJPAI
}

//提
type CanTiInfo struct {
	CanTi   bool
	TiCards []*majiang.MJPAI
}

func (p *MJParserCore) ZiMoGangCards(userGameData interface{}) (interface{}, error) {
	gameData := userGameData.(MJUserGameData) //玩家数据
	handPais := gameData.GetHandPais()        //判断的手牌
	//把摸的牌放进手牌里进行判断
	if gameData.GetMoPai() != nil {
		handPais = append(handPais, gameData.GetMoPai())
	}

	var listHandPais majiang.MjPAIList
	listHandPais = handPais
	sort.Sort(listHandPais)

	pengPais := gameData.GetPengPais() //得到所有的吃牌
	var ret []*CanGangInfoBean
	//首先判断明杠
	counts := p.CountHandPais(handPais) //统计牌
	for _, p := range listHandPais {
		if 4 == counts[ p.GetCountIndex()] {
			r := &CanGangInfoBean{
				GangPai:  p,
				GangType: GANGTYPE_AN,
			}
			ret = append(ret, r)
		}
	}

	//再判断巴杠
	for _, p := range listHandPais {
		for _, peng := range pengPais {
			pengPai := peng.Pais[0] //判断碰牌的第一张和手牌是否一致
			if p.Flower == pengPai.Flower &&
				p.Value == pengPai.Value {
				r := &CanGangInfoBean{
					GangPai:  p,
					GangType: GANGTYPE_BA,
				}
				ret = append(ret, r)
				break
			}
		}
	}

	//清楚重复的数据
	var ret2 *CanGangInfo = &CanGangInfo{CanGang: false}
	isexistRet2 := func(g *CanGangInfoBean) bool {
		for _, r := range ret2.GangInfoBean {
			if r.GangPai.Flower == g.GangPai.Flower &&
				r.GangPai.Value == g.GangPai.Value {
				return true
			}
		}
		return false
	}

	for _, g := range ret {
		//首先判断ret2中是否有杠牌了
		if !isexistRet2(g) {
			ret2.CanGang = true
			ret2.GangInfoBean = append(ret2.GangInfoBean, g)
		}
	}

	//返回最后的杠牌的数据
	return ret2, nil
}

//是否可以吃牌
func (p *MJParserCore) CanChi(interface{}, interface{}) (interface{}, error) {
	return nil, nil
}

//是否可以报听
func (p *MJParserCore) CanTing(...interface{}) (interface{}, error) {
	return nil, nil
}

//是否可以飞
func (p *MJParserCore) CanFly(interface{}, interface{}) (bool, error) {
	return false, nil
}

//是否可以提
func (p *MJParserCore) CanTi(interface{}) (interface{}, error) {
	return nil, nil
}

//是否能补虾子
func (p *MJParserCore) CanBu(...interface{}) (interface{}, error) {
	return nil, nil
}

//得到叫牌的信息 打出哪张牌后能胡哪些牌
func (p *MJParserCore) GetJiaoInfos(...interface{}) (interface{}, error) {
	return nil, nil
}

//获取听牌 能胡哪些牌
func (p *MJParserCore) GetTingInfos(...interface{}) (interface{}, error) {
	return nil, nil
}

//通过一副牌的id解析牌型
func (p *MJParserCore) Parse(pids []int32) (interface{}, error) {
	return nil, nil
}

//洗牌
func (p *MJParserCore) XiPai() interface{} {
	return nil
}

//todo 清一色
func (p *MJParserCore) IsQingYiSe(pais []*majiang.MJPAI) bool {
	//
	//flower := pais[0].Flower
	//for i := 1; i < len(pais); i++ {
	//	if *flower != *pais[i].Flower {
	//		return false //不是清一色
	//	}
	//}
	return false
}

/**

//判断是否是七对 返回勾数
func (p *MJParserCore) IsQiDui(handPai *MJHandPai, hupai *MJPai) (isQidui bool, gou int32) {
	joinedPais := p.JoinHandPaiPaisAndHuPai(handPai, hupai)

	//合并手牌 长度不为14即有碰杠吃 不满足七对基本要求
	if len(joinedPais) != 14 {
		return false, 0
	}

	handCounts := GettPaiStats(joinedPais)
	for i := 0; i < len(handCounts); i++ {
		switch handCounts[i] {
		case 0, 2:
			continue //牌数为0和2都忽略
		case 4:
			gou++ //统计勾数
		default:
			return false, 0 //牌数1、3不为七对
		}
	}
	return true, gou
}

//七对 龙七对牌型胡牌判断
func (p *MJParserCore) tryHU7(handPai *MJHandPai, hupai *MJPai) bool {
	isQidui, _ := p.IsQiDui(handPai, hupai)
	return isQidui
}

**/
func (p *MJParserCore) is19(val int) bool {
	v := p.GettPaiValueByCountPos(val)
	return v == 1 || v == 9
}

func (p *MJParserCore) GettPaiValueByCountPos(countPos int) int32 {
	return int32(countPos%9 + 1)
}

func (p *MJParserCore) IsMengQing(g MJUserGameData) bool {
	if len(g.GetChiPais()) > 0 {
		return false
	}
	if len(g.GetPengPais()) > 0 {
		return false
	}
	for _, gang := range g.GetGangPais() {
		if gang.GangType != GANGTYPE_AN {
			return false
		}
	}
	return true
}

//是否是听用牌的空实现 宜宾麻将
func (p *MJParserCore) IsTingYongPai(pai majiang.MJPAI) bool {
	return false
}

/**

func (p *MJParserCore) is258(val int) bool {
	v := GettPaiValueByCountPos(val)
	return v == 2 || v == 5 || v == 8
}

**/

func (p *MJParserCore) TryHU(count []int, len int) (result bool, isAll19 bool, jiang int) {
	//log.T("开始判断tryHu(%v,%v)", count, len)
	isAll19 = true //全带幺
	result = false
	jiang = -1
	//递归完所有的牌表示 胡了
	if (len == 0) {
		//log.T("len == 0")
		return true, isAll19, jiang
	}

	if (len%3 == 2) {
		//log.T("if %v 取模 3 == 2", len)
		// 说明对牌出现
		for i := 0; i < 27; i++ {
			if (count[i] >= 2) {
				count[i] -= 2

				result, isAll19, jiang = p.TryHU(count, len-2)
				if (result) {
					//log.T("i: %v, value: %v", i, count[i])
					if ! p.is19(i) {
						//不是幺九
						isAll19 = false
					}
					jiang = i
					return true, isAll19, jiang
				}
				count[i] += 2
			}
		}
	} else {
		//log.T("else %v", len)
		// 是否是顺子，这里应该分开判断
		for i := 0; i < 7; i++ {
			if (count[i] > 0 && count[i+1] > 0 && count[i+2] > 0) {
				count[i] -= 1;
				count[i+1] -= 1;
				count[i+2] -= 1;
				result, isAll19, jiang = p.TryHU(count, len-3)
				if (result) {
					//log.T("i: %v, value: %v", i, count[i])
					if !p.is19(i) && !p.is19(i + 1) && !p.is19(i + 2) {
						//不是幺九
						//log.T("branch 2 pos%v不是幺九", i)
						isAll19 = false
					}
					return true, isAll19, jiang
				}
				count[i] += 1
				count[i+1] += 1
				count[i+2] += 1
			}
		}

		for i := 9; i < 16; i++ {
			if (count[i] > 0 && count[i+1] > 0 && count[i+2] > 0) {
				count[i] -= 1
				count[i+1] -= 1
				count[i+2] -= 1
				result, isAll19, jiang = p.TryHU(count, len-3)
				if (result) {
					//log.T("i: %v, value: %v", i, count[i])
					if !p.is19(i) && !p.is19(i + 1) && !p.is19(i + 2) {
						//不是幺九
						//log.T("branch 3 pos%v不是幺九", i)
						isAll19 = false
					}
					return true, isAll19, jiang
				}
				count[i] += 1
				count[i+1] += 1
				count[i+2] += 1
			}
		}

		for i := 18; i < 25; i++ {
			if (count[i] > 0 && count[i+1] > 0 && count[i+2] > 0) {
				count[i] -= 1;
				count[i+1] -= 1;
				count[i+2] -= 1;
				result, isAll19, jiang = p.TryHU(count, len-3)
				if (result) {
					//log.T("i: %v, value: %v", i, count[i])
					if !p.is19(i) && !p.is19(i + 1) && !p.is19(i + 2) {
						//不是幺九
						//log.T("branch 4 pos%v不是幺九", i)
						isAll19 = false
					}
					return true, isAll19, jiang
				}
				count[i] += 1;
				count[i+1] += 1;
				count[i+2] += 1;
			}
		}

		// 三个一样的
		for i := 0; i < 27; i++ {
			if (count[i] >= 3) {
				count[i] -= 3
				result, isAll19, jiang = p.TryHU(count, len-3)
				if (result) {
					//log.T("i: %v, value: %v", i, count[i])
					if !p.is19(i) {
						//不是幺九
						//log.T("branch 5 pos%v不是幺九", i)
						isAll19 = false
					}
					return true, isAll19, jiang
				}
				count[i] += 3
			}
		}
	}
	return false, isAll19, jiang
}

/**

////////////////////////////////////////////////////////////////////////

//将手牌碰牌刚牌与huPai拼接成数组
func (p *MJParserCore) JoinAllHandPaiAndHuPai(handPai *MJHandPai, hupai *MJPai) []*MJPai {
	pais := []*MJPai{}
	if hupai != nil {
		pais = append(pais, hupai)
	}

	if handPai.Pais != nil {
		pais = append(pais, handPai.Pais...)
	}

	if handPai.GangPais != nil {
		pais = append(pais, handPai.GangPais...)
	}

	if handPai.PengPais != nil {
		pais = append(pais, handPai.PengPais...)
	}

	if handPai.ChiPais != nil {
		pais = append(pais, handPai.ChiPais...)
	}
	return pais
}

//将手牌与huPai拼接成数组
func (p *MJParserCore) JoinHandPaiPaisAndHuPai(handPai *MJHandPai, hupai *MJPai) []*MJPai {
	pais := []*MJPai{}
	if hupai != nil {
		pais = append(pais, hupai)
	}

	if handPai.Pais != nil {
		pais = append(pais, handPai.Pais...)
	}

	return pais
}

//判断手牌中是否有杠牌
func (p *MJParserCore) isYouGang(handPai *MJHandPai) bool {
	return handPai.GangPais != nil && len(handPai.GangPais) > 0
}

//判断手牌中是否有碰牌
func (p *MJParserCore) isYouPeng(handPai *MJHandPai) bool {
	return handPai.PengPais != nil && len(handPai.PengPais) > 0
}

//判断手牌中是否有吃牌
func (p *MJParserCore) isYouChi(handPai *MJHandPai) bool {
	return handPai.ChiPais != nil && len(handPai.ChiPais) > 0
}


**/

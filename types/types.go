package types

//DAMAGETYPE ダメージタイプ
type DAMAGETYPE int

//DAMAGETYPES
const (
	NORMALATK DAMAGETYPE = iota + 1
	THUMP
	E
	Q
	PHYSICS
	FIRE
	WATER
	ICE
	THUNDER
	WIND
	ROCK
	GRASS
	ALL
)

//DamageTypeList ダメージタイプリスト
var DamageTypeList []DAMAGETYPE = []DAMAGETYPE{
	NORMALATK,
	THUMP,
	E,
	Q,
	PHYSICS,
	FIRE,
	WATER,
	ICE,
	THUNDER,
	WIND,
	ROCK,
	GRASS,
	ALL,
}

//DAMAGEMAP ダメージタイプのマップ
var DAMAGEMAP = map[string]DAMAGETYPE{
	"普": NORMALATK,
	"重": THUMP,
	"E": E,
	"Q": Q,
	"物": PHYSICS,
	"火": FIRE,
	"水": WATER,
	"冰": ICE,
	"雷": THUNDER,
	"风": WIND,
	"岩": ROCK,
	"草": GRASS,
	"全": ALL,
}

//DAMAGEMAPNEGATIVE ダメージタイプのマップ
var DAMAGEMAPNEGATIVE = map[DAMAGETYPE]string{}

func init() {
	//逆マップの初期化
	for i, v := range DAMAGEMAP {
		DAMAGEMAPNEGATIVE[v] = i
	}
}

//BaseSkill スキルベース
type BaseSkill struct {
	Name             string   `json:"name"`
	DamageRate       float32  `json:"damageRate"`
	DamageTypes      []string `json:"damageTypes"`
	DamageBoostTypes []DAMAGETYPE
}

//GetEleReation 元素反応
func (c *BaseSkill) GetEleReation() (res map[float32]string) {
	res = map[float32]string{}
	for _, v := range c.DamageBoostTypes {
		if v == FIRE {
			res[1.5] = "蒸发"
			res[2] = "融化"
		}
		if v == WATER {
			res[2] = "蒸发"
		}
		if v == ICE {
			res[1.5] = "融化"
		}
	}
	return
}

//スキルの初期化
func (c *BaseSkill) init() {
	for _, v := range c.DamageTypes {
		c.DamageBoostTypes = append(c.DamageBoostTypes, DAMAGEMAP[v])
	}
}

//DamageBoost ダメージ増幅
type DamageBoost struct {
	DamageBoostRate float32 `json:"damageBoostRate"`
	DamageType      string  `json:"damageType"`
	DamageBoostType DAMAGETYPE
}

//DamageBoostの初期化
func (c *DamageBoost) init() {
	c.DamageBoostType = DAMAGEMAP[c.DamageType]
}

//DamageBoost かける関数
func (c *DamageBoost) add(toAdd DamageBoost) *DamageBoost {
	if c.DamageBoostType == toAdd.DamageBoostType {
		c.DamageBoostRate += toAdd.DamageBoostRate
	}
	return c
}

//DamageBoost 引く関数
func (c *DamageBoost) decrease(toAdd DamageBoost) *DamageBoost {
	if c.DamageBoostType == toAdd.DamageBoostType {
		c.DamageBoostRate -= toAdd.DamageBoostRate
	}
	return c
}

//getDamageBoost 取得
func getDamageBoost(c []DamageBoost, dType DAMAGETYPE) (re DamageBoost, has bool) {
	for i := range c {
		if c[i].DamageBoostType == dType {
			has = true
			re = c[i]
			return
		}
	}
	has = false
	re = DamageBoost{}
	re.DamageBoostType = dType
	re.DamageType = DAMAGEMAPNEGATIVE[dType]
	return
}

//DamageBoosts かける関数
func mergeDamageBoost(c []DamageBoost, toAdd ...[]DamageBoost) []DamageBoost {
	var resMap = []DamageBoost{}
	for i, v := range DamageTypeList {
		x, _ := getDamageBoost(c, v)
		resMap = append(resMap, x)
		for ii := range toAdd {
			xx, _ := getDamageBoost(toAdd[ii], v)
			resMap[i].add(xx)
		}
	}
	return resMap
}

//DamageBoosts 引く関数
func mergeDamageBoostDecrease(c []DamageBoost, toAdd ...[]DamageBoost) []DamageBoost {
	var resMap = []DamageBoost{}
	for i, v := range DamageTypeList {
		x, _ := getDamageBoost(c, v)
		resMap = append(resMap, x)
		for ii := range toAdd {
			xx, _ := getDamageBoost(toAdd[ii], v)
			resMap[i].decrease(xx)
		}
	}
	return resMap
}

//Decrease デバフ
type Decrease struct {
	LevelResisRate float32 `json:"levelResisRate"`
}

//Result 計算結果ストラクト
type Result struct {
	//skill name
	SkillName string `json:"skillName"`
	//damageResult
	FinalDamageWithoutCrit float32     `json:"finalDamageWithoutCrit"`
	FinalDamageWithCrit    float32     `json:"finalDamageWithCrit"`
	FinalDamageAverage     float32     `json:"finalDamageAverage"`
	FinalEleResult         []EleResult `json:"finalEleResult"`
}

//EleResult 元素増幅計算結果ストラクト
type EleResult struct {
	//skill name
	ReactionName string `json:"reactionName"`
	//damageResult
	FinalEleDamageWithoutCrit float32 `json:"finalEleDamageWithoutCrit"`
	FinalEleDamageWithCrit    float32 `json:"finalEleDamageWithCrit"`
	FinalEleDamageAverage     float32 `json:"finalEleDamageAverage"`
}

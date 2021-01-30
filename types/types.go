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

//BaseCharacter チャラベース
type BaseCharacter struct {
	//base
	Atk            float32       `json:"atk"`
	Def            float32       `json:"def"`
	Boold          float32       `json:"boold"`
	Skills         []BaseSkill   `json:"skills"`
	DamageBoosts   []DamageBoost `json:"damageBoosts"`
	CritRate       float32       `json:"critRate"`
	CritDamageRate float32       `json:"critDamageRate"`
	EleValue       float32       `json:"eleValue"`
	//rate buff
	AtkBuffRate   float32 `json:"atkBuffRate"`
	DefBuffRate   float32 `json:"defBuffRate"`
	BooldBuffRate float32 `json:"booldBuffRate"`
	//value buff
	DamageBoostBuffs   []DamageBoost `json:"damageBoostBuffs"`
	CritRateBuff       float32       `json:"critRateBuff"`
	CritDamageRateBuff float32       `json:"critDamageRateBuff"`
	EleValueBuff       float32       `json:"eleValueBuff"`
	//hutao buff
	UseE              bool    `json:"useE"`
	UpAtkByBooldBuff  float32 `json:"upAtkByBooldBuff"`
	UpMaxBaseAtkRate  float32 `json:"upMaxBaseAtkRate"`
	UnderHalf         bool    `json:"underHalf"`
	UpDamageBoostType string  `json:"upDamageBoostType"`
	UpDamageBoostBuff float32 `json:"upDamageBoostBuff"`
}

//チャラベースの初期化
func (c *BaseCharacter) init() {
	for i := range c.Skills {
		c.Skills[i].init()
	}
	for i := range c.DamageBoosts {
		c.DamageBoosts[i].init()
	}
	for i := range c.DamageBoostBuffs {
		c.DamageBoostBuffs[i].init()
	}
}

//BaseWeapon 武器ベース
type BaseWeapon struct {
	//base
	Atk float32 `json:"atk"`
	//rate buff
	AtkBuffRate   float32 `json:"atkBuffRate"`
	DefBuffRate   float32 `json:"defBuffRate"`
	BooldBuffRate float32 `json:"booldBuffRate"`
	//value buff
	DamageBoostBuffs   []DamageBoost `json:"damageBoostBuffs"`
	CritRateBuff       float32       `json:"critRateBuff"`
	CritDamageRateBuff float32       `json:"critDamageRateBuff"`
	EleValueBuff       float32       `json:"eleValueBuff"`
	//humo buff
	AddAtkByBoold     float32 `json:"addAtkByBoold"`
	AddAtkByHalfBoold float32 `json:"addAtkByHalfBoold"`
}

//武器の初期化
func (c *BaseWeapon) init() {
	for i := range c.DamageBoostBuffs {
		c.DamageBoostBuffs[i].init()
	}
}

//BaseArtifact アーティファクトベース
type BaseArtifact struct {
	//main
	//base
	MainAtk      float32 `json:"mainAtk"`
	MainDef      float32 `json:"mainDef"`
	MainBoold    float32 `json:"mainBoold"`
	MainEleValue float32 `json:"mainEleValue"`
	//rate buff
	MainAtkBuffRate   float32 `json:"mainAtkBuffRate"`
	MainDefBuffRate   float32 `json:"mainDefBuffRate"`
	MainBooldBuffRate float32 `json:"mainBooldBuffRate"`
	//value buff
	MainDamageBoostBuffs   []DamageBoost `json:"mainDamageBoostBuffs"`
	MainCritRateBuff       float32       `json:"mainCritRateBuff"`
	MainCritDamageRateBuff float32       `json:"mainCritDamageRateBuff"`
	MainEleValueBuff       float32       `json:"mainEleValueBuff"`
	MainChargeRateBuff     float32       `json:"mainChargeRateBuff"`
	//sub or total(ALL IN ONE)
	//base
	Artifact
	//contrl flg
	Name   string `json:"name"`
	Groups []int  `json:"groups"`
	Use    bool   `json:"use"`
}

//アーティファクトの初期化
func (c *BaseArtifact) init() {
	for i := range c.MainDamageBoostBuffs {
		c.MainDamageBoostBuffs[i].init()
	}
	for i := range c.DamageBoostBuffs {
		c.DamageBoostBuffs[i].init()
	}
}

func (c *BaseArtifact) isValidInGroup(groupIndex int) bool {
	if c.Use {
		for i := range c.Groups {
			if c.Groups[i] == groupIndex {
				return true
			}
		}
	}
	return false
}

//Artifact アーティファクトトータル
type Artifact struct {
	//base
	Atk      float32 `json:"atk"`
	Def      float32 `json:"def"`
	Boold    float32 `json:"boold"`
	EleValue float32 `json:"eleValue"`
	//rate buff
	AtkBuffRate   float32 `json:"atkBuffRate"`
	DefBuffRate   float32 `json:"defBuffRate"`
	BooldBuffRate float32 `json:"booldBuffRate"`
	//value buff
	DamageBoostBuffs   []DamageBoost `json:"damageBoostBuffs"`
	CritRateBuff       float32       `json:"critRateBuff"`
	CritDamageRateBuff float32       `json:"critDamageRateBuff"`
	EleValueBuff       float32       `json:"eleValueBuff"`
	ChargeRateBuff     float32       `json:"chargeRateBuff"`
}

//アーティファクトトータルの初期化
func (c *Artifact) init() {
	for i := range c.DamageBoostBuffs {
		c.DamageBoostBuffs[i].init()
	}
}

//Character チャラ属性
type Character struct {
	Character BaseCharacter  `json:"character"`
	Weapon    BaseWeapon     `json:"weapon"`
	Artifacts []BaseArtifact `json:"artifacts"`
	Skills    []BaseSkill    `json:"skills"`
	//Groups
	GroupsMap       map[int][]int         `json:"groupsMap"`
	ArtifactsInOne  map[int]Artifact      `json:"artifactsInOne"`
	DamageBoosts    map[int][]DamageBoost `json:"damageBoosts"`
	BaseAtk         map[int]float32       `json:"baseAtk"`
	HuTaoUpAtkMax   map[int]float32       `json:"huTaoUpAtkMax"`
	HuTaoUpAtk      map[int]float32       `json:"huTaoUpAtk"`
	Atk             map[int]float32       `json:"atk"`
	Def             map[int]float32       `json:"def"`
	BooldMax        map[int]float32       `json:"booldMax"`
	BooldRate       map[int]float32       `json:"booldRate"`
	BooldCurrent    map[int]float32       `json:"booldCurrent"`
	CritRate        map[int]float32       `json:"critRate"`
	CritDamageRate  map[int]float32       `json:"critDamageRate"`
	EleValue        map[int]float32       `json:"eleValue"`
	EleReactionRate map[int]float32       `json:"eleReactionRate"`
	//Decrease
	Decrease Decrease `json:"decrease"`
	//Results
	Results map[int][]Result `json:"results"`
}

//Decrease デバフ
type Decrease struct {
	MonstarResisRate    float32 `json:"monstarResisRate"`
	LevelResisRate      float32 `json:"levelResisRate"`
	FinalDamageDownRate float32 `json:"finalDamageDownRate"`
}

//Result 計算結果ストラクト
type Result struct {
	//skill name
	SkillName string `json:"skillName"`
	//damageResult
	FinalDamageWithoutCrit    float32 `json:"finalDamageWithoutCrit"`
	FinalDamageWithCrit       float32 `json:"finalDamageWithCrit"`
	FinalDamageAverage        float32 `json:"finalDamageAverage"`
	FinalEleDamageWithoutCrit float32 `json:"finalEleDamageWithoutCrit"`
	FinalEleDamageWithCrit    float32 `json:"finalEleDamageWithCrit"`
	FinalEleDamageAverage     float32 `json:"finalEleDamageAverage"`
}

//Init 初期化
func (c *Character) Init() {
	//初期化
	if c.DamageBoosts == nil {
		c.DamageBoosts = map[int][]DamageBoost{}
	}
	if c.BaseAtk == nil {
		c.BaseAtk = map[int]float32{}
	}
	if c.Atk == nil {
		c.Atk = map[int]float32{}
	}
	if c.HuTaoUpAtkMax == nil {
		c.HuTaoUpAtkMax = map[int]float32{}
	}
	if c.HuTaoUpAtk == nil {
		c.HuTaoUpAtk = map[int]float32{}
	}
	if c.Def == nil {
		c.Def = map[int]float32{}
	}
	if c.BooldMax == nil {
		c.BooldMax = map[int]float32{}
	}
	if c.BooldRate == nil {
		c.BooldRate = map[int]float32{}
	}
	if c.BooldCurrent == nil {
		c.BooldCurrent = map[int]float32{}
	}
	if c.CritRate == nil {
		c.CritRate = map[int]float32{}
	}
	if c.CritDamageRate == nil {
		c.CritDamageRate = map[int]float32{}
	}
	if c.EleValue == nil {
		c.EleValue = map[int]float32{}
	}
	if c.EleReactionRate == nil {
		c.EleReactionRate = map[int]float32{}
	}
	if c.Results == nil {
		c.Results = map[int][]Result{}
	}

	baseChara := &c.Character
	baseChara.init()
	weapon := &c.Weapon
	weapon.init()
	for i := range c.Artifacts {
		c.Artifacts[i].init()
	}
	artifacts := c.Artifacts

	c.Skills = baseChara.Skills

	//アーティファクト処理
	c.GroupsMap, c.ArtifactsInOne = calAllArtifacts(artifacts)

	for i := range c.GroupsMap {
		artifactsInOne := c.ArtifactsInOne[i]

		c.BaseAtk[i] = baseChara.Atk + weapon.Atk
		c.Atk[i] = c.BaseAtk[i] + c.BaseAtk[i]*(baseChara.AtkBuffRate+weapon.AtkBuffRate+artifactsInOne.AtkBuffRate) + artifactsInOne.Atk
		c.Def[i] = baseChara.Def + baseChara.Def*(baseChara.DefBuffRate+weapon.DefBuffRate+artifactsInOne.DefBuffRate) + artifactsInOne.Def
		c.BooldMax[i] = baseChara.Boold + baseChara.Boold*(baseChara.BooldBuffRate+weapon.BooldBuffRate+artifactsInOne.BooldBuffRate) + artifactsInOne.Boold
		c.DamageBoosts[i] = mergeDamageBoost(baseChara.DamageBoosts, baseChara.DamageBoostBuffs, weapon.DamageBoostBuffs, artifactsInOne.DamageBoostBuffs)
		c.CritRate[i] = baseChara.CritRate + baseChara.CritRateBuff + weapon.CritRateBuff + artifactsInOne.CritRateBuff
		c.CritDamageRate[i] = baseChara.CritDamageRate + baseChara.CritDamageRateBuff + weapon.CritDamageRateBuff + artifactsInOne.CritDamageRateBuff
		c.EleValue[i] = baseChara.EleValue + baseChara.EleValueBuff + weapon.EleValueBuff + artifactsInOne.EleValueBuff

		//huTaoAndhumoBuff
		huTaoAndhumoBuff(c, i)
	}

}

//すべての聖遺物を合計
func calAllArtifacts(s []BaseArtifact) (groups map[int][]int, res map[int]Artifact) {
	groups = map[int][]int{}
	res = map[int]Artifact{}
	//1. グループの確認
	for i, v := range s {
		//グループに追加
		for _, vv := range v.Groups {
			_, ok := groups[vv]
			if !ok {
				groups[vv] = []int{}
			}
			groups[vv] = append(groups[vv], i)
		}
	}
	//2. グループ一つづつ計算
	for i, v := range groups {
		var tempRes Artifact = Artifact{}
		//合計
		for _, vv := range v {
			tempRes.Atk += s[vv].Atk + s[vv].MainAtk
			tempRes.Def += s[vv].Def + s[vv].MainDef
			tempRes.Boold += s[vv].Boold + s[vv].MainBoold
			tempRes.EleValue += s[vv].EleValue + s[vv].MainEleValue
			tempRes.AtkBuffRate += s[vv].AtkBuffRate + s[vv].MainAtkBuffRate
			tempRes.DefBuffRate += s[vv].DefBuffRate + s[vv].MainDefBuffRate
			tempRes.BooldBuffRate += s[vv].BooldBuffRate + s[vv].MainBooldBuffRate

			tempRes.DamageBoostBuffs = mergeDamageBoost(tempRes.DamageBoostBuffs, mergeDamageBoost(s[vv].DamageBoostBuffs, s[vv].MainDamageBoostBuffs))

			tempRes.CritRateBuff += s[vv].CritRateBuff + s[vv].MainCritRateBuff
			tempRes.CritDamageRateBuff += s[vv].CritDamageRateBuff + s[vv].MainCritDamageRateBuff
			tempRes.EleValueBuff += s[vv].EleValueBuff + s[vv].MainEleValueBuff
			tempRes.ChargeRateBuff += s[vv].ChargeRateBuff + s[vv].MainChargeRateBuff
		}

		tempRes.init()

		res[i] = tempRes
	}
	return
}

//huTaoAndhumoBuff buff
func huTaoAndhumoBuff(c *Character, i int) {
	//humo
	c.Atk[i] += c.Weapon.AddAtkByBoold * c.BooldMax[i]
	if c.Character.UseE {
		//hutao
		toUp := c.Character.UpAtkByBooldBuff * c.BooldMax[i]
		toUpMax := c.BaseAtk[i] * c.Character.UpMaxBaseAtkRate
		if toUp > toUpMax {
			toUp = toUpMax
		}
		c.HuTaoUpAtk[i] = toUp
		c.HuTaoUpAtkMax[i] = toUpMax
		c.Atk[i] += toUp
	}
	if c.Character.UnderHalf {
		//hutao
		temp := DamageBoost{
			DamageBoostRate: c.Character.UpDamageBoostBuff,
			DamageType:      c.Character.UpDamageBoostType,
			DamageBoostType: DAMAGEMAP[c.Character.UpDamageBoostType],
		}
		c.DamageBoosts[i] = mergeDamageBoost(c.DamageBoosts[i], []DamageBoost{temp})
		//humo
		c.Atk[i] += c.Weapon.AddAtkByHalfBoold * c.BooldMax[i]
		//other
		c.BooldRate[i] = 0.5
		c.BooldCurrent[i] = c.BooldRate[i] * c.BooldMax[i]
	}
}

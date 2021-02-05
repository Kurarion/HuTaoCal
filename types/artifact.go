package types

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

//すべての聖遺物を合計
func calAllArtifacts(s []BaseArtifact) (groups map[int][]int, res map[int]*Artifact) {
	groups = map[int][]int{}
	res = map[int]*Artifact{}
	//1. グループの確認
	var hasGroup bool = false
	for i, v := range s {
		//グループに追加
		for _, vv := range v.Groups {
			_, ok := groups[vv]
			if !ok {
				groups[vv] = []int{}
			}
			groups[vv] = append(groups[vv], i)
			hasGroup = true
		}
	}
	if !hasGroup {
		groups[0] = []int{}
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

		res[i] = &tempRes
	}
	return
}

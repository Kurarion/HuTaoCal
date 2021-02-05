package types

//Data チャラ属性
type Data struct {
	Character BaseCharacter  `json:"character"`
	Weapon    BaseWeapon     `json:"weapon"`
	Artifacts []BaseArtifact `json:"artifacts"`
	Skills    []BaseSkill    `json:"skills"`
	//monstrt
	Monster BaseMonster `json:"monster"`
	//Groups
	GroupsMap       map[int][]int         `json:"groupsMap"`
	ArtifactsInOne  map[int]*Artifact     `json:"artifactsInOne"`
	DamageBoosts    map[int][]DamageBoost `json:"damageBoosts"`
	BaseAtk         map[int]float32       `json:"baseAtk"`
	Atk             map[int]float32       `json:"atk"`
	Def             map[int]float32       `json:"def"`
	BooldMax        map[int]float32       `json:"booldMax"`
	CritRate        map[int]float32       `json:"critRate"`
	CritDamageRate  map[int]float32       `json:"critDamageRate"`
	EleValue        map[int]float32       `json:"eleValue"`
	EleReactionRate map[int]float32       `json:"eleReactionRate"`
	//Decrease
	DamageDecrease Decrease `json:"damageDecrease"`
	//character Results
	HuTaoInfo   map[int]*HutaoResult   `json:"huTaoInfo"`
	NuoaierInfo map[int]*NuoaierResult `json:"nuoaierInfo"`
	//weapon Results
	HumoInfo map[int]*HumoResult `json:"humoInfo"`
	//Results
	Results map[int][]Result `json:"results"`
}

/*  キャラクター  */

//HutaoResult hutao info
type HutaoResult struct {
	UpAtkMax     float32 `json:"upAtkMax"`
	UpAtk        float32 `json:"upAtk"`
	BooldRate    float32 `json:"booldRate"`
	BooldCurrent float32 `json:"booldCurrent"`
}

//NuoaierResult nuoaier info
type NuoaierResult struct {
	UpAtk float32 `json:"upAtk"`
}

/*  武器  */

//HumoResult humo info
type HumoResult struct {
	UpAtk       float32 `json:"upAtk"`
	UpAtkNormal float32 `json:"upAtkNormal"`
	UpAtkHalf   float32 `json:"upAtkHalf"`
}

//Init 初期化
func (c *Data) Init() {
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
	if c.Def == nil {
		c.Def = map[int]float32{}
	}
	if c.BooldMax == nil {
		c.BooldMax = map[int]float32{}
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
	baseMonster := &c.Monster
	baseMonster.init()
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
		if c.EleValue[i] == 0 {
			c.EleReactionRate[i] = 0
		} else {
			c.EleReactionRate[i] = (6.665 - 9340/(c.EleValue[i]+1401)) / 2.4
		}

		//チャラの特別対応
		execCharacterSpecial(c, i)
		//武器の特別対応
		execWeaponSpecial(c, i)
	}

	//ダメージ軽減
	var result float32 = 1.0
	result = (float32(c.Character.Level) + 100.0) / (float32(c.Character.Level) + 100.0 + ((float32(c.Monster.Level) + 100.0) * (1.0 - c.Monster.DefDeBuffRate)))
	//レベル圧制
	if c.Character.Level-c.Monster.Level >= 70 && c.Monster.Level <= 10 {
		result *= 1.5
	} else if c.Monster.Level-c.Character.Level >= 70 && c.Character.Level <= 10 {
		result *= 0.5
	}
	c.DamageDecrease.LevelResisRate = result
}

package types

//BaseCharacter チャラベース
type BaseCharacter struct {
	//base
	Level            int           `json:"level"`
	Atk              float32       `json:"atk"`
	Def              float32       `json:"def"`
	Boold            float32       `json:"boold"`
	CurrentBooldRate float32       `json:"currentBooldRate"`
	Skills           []BaseSkill   `json:"skills"`
	DamageBoosts     []DamageBoost `json:"damageBoosts"`
	CritRate         float32       `json:"critRate"`
	CritDamageRate   float32       `json:"critDamageRate"`
	EleValue         float32       `json:"eleValue"`
	//rate buff
	AtkBuffRate   float32 `json:"atkBuffRate"`
	DefBuffRate   float32 `json:"defBuffRate"`
	BooldBuffRate float32 `json:"booldBuffRate"`
	//value buff
	DamageBoostBuffs   []DamageBoost `json:"damageBoostBuffs"`
	CritRateBuff       float32       `json:"critRateBuff"`
	CritDamageRateBuff float32       `json:"critDamageRateBuff"`
	EleValueBuff       float32       `json:"eleValueBuff"`
	//character detail
	//hutao buff
	Hutao HutaoBase `json:"hutao"`
	//Nuoaier buff
	Nuoaier NuoaierBase `json:"nuoaier"`
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

//HutaoBase チャラベース
type HutaoBase struct {
	On                              bool    `json:"on"`
	UseE                            bool    `json:"useE"`
	UpAtkByBooldBuff                float32 `json:"upAtkByBooldBuff"`
	UpMaxBaseAtkRate                float32 `json:"upMaxBaseAtkRate"`
	UnderHalfBooldUpDamageBoostType string  `json:"underHalfBooldUpDamageBoostType"`
	UnderHalfBooldUpDamageBoostBuff float32 `json:"underHalfBooldUpDamageBoostBuff"`
}

//NuoaierBase チャラベース
type NuoaierBase struct {
	On           bool    `json:"on"`
	UseQ         bool    `json:"useQ"`
	DefToAtkRate float32 `json:"defToAtkRate"`
}

//チャラの特別対応
func execCharacterSpecial(c *Data, i int) {
	//Hutao
	HutaoBuff(c, i)
	//Nuoaier
	NuoaierBuff(c, i)
}

//HutaoBuff buff
func HutaoBuff(c *Data, i int) {
	if c.Character.Hutao.On {
		if c.Character.Hutao.UseE {
			//hutao
			toUp := c.Character.Hutao.UpAtkByBooldBuff * c.BooldMax[i]
			toUpMax := c.BaseAtk[i] * c.Character.Hutao.UpMaxBaseAtkRate
			if toUp > toUpMax {
				toUp = toUpMax
			}
			//init
			if c.HuTaoInfo == nil {
				c.HuTaoInfo = map[int]*HutaoResult{}
			}
			if c.HuTaoInfo[i] == nil {
				c.HuTaoInfo[i] = &HutaoResult{}
			}
			c.HuTaoInfo[i].UpAtk = toUp
			c.HuTaoInfo[i].UpAtkMax = toUpMax
			//加算
			c.Atk[i] += toUp
		}
		//half boold rate
		if c.Character.CurrentBooldRate != 0 && c.Character.CurrentBooldRate <= 0.5 {
			//hutao
			temp := DamageBoost{
				DamageBoostRate: c.Character.Hutao.UnderHalfBooldUpDamageBoostBuff,
				DamageType:      c.Character.Hutao.UnderHalfBooldUpDamageBoostType,
				DamageBoostType: DAMAGEMAP[c.Character.Hutao.UnderHalfBooldUpDamageBoostType],
			}
			//init
			if c.HuTaoInfo == nil {
				c.HuTaoInfo = map[int]*HutaoResult{}
			}
			if c.HuTaoInfo[i] == nil {
				c.HuTaoInfo[i] = &HutaoResult{}
			}
			c.HuTaoInfo[i].BooldRate = c.Character.CurrentBooldRate
			c.HuTaoInfo[i].BooldCurrent = c.HuTaoInfo[i].BooldRate * c.BooldMax[i]
			//加算
			c.DamageBoosts[i] = mergeDamageBoost(c.DamageBoosts[i], []DamageBoost{temp})
		}
	}
}

//NuoaierBuff buff
func NuoaierBuff(c *Data, i int) {
	if c.Character.Nuoaier.On {
		if c.Character.Nuoaier.UseQ {
			//nuoaier
			toUp := c.Character.Nuoaier.DefToAtkRate * c.Def[i]
			//init
			if c.NuoaierInfo == nil {
				c.NuoaierInfo = map[int]*NuoaierResult{}
			}
			if c.NuoaierInfo[i] == nil {
				c.NuoaierInfo[i] = &NuoaierResult{}
			}
			c.NuoaierInfo[i].UpAtk = toUp
			//加算
			c.Atk[i] += toUp
		}
	}
}

package types

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
	Humo HumoBase `json:"humo"`
}

//武器の初期化
func (c *BaseWeapon) init() {
	for i := range c.DamageBoostBuffs {
		c.DamageBoostBuffs[i].init()
	}
}

//HumoBase 武器ベース
type HumoBase struct {
	On                bool    `json:"on"`
	AddAtkByBoold     float32 `json:"addAtkByBoold"`
	AddAtkByHalfBoold float32 `json:"addAtkByHalfBoold"`
}

//武器の特別対応
func execWeaponSpecial(c *Data, i int) {
	//Humo
	HumoBuff(c, i)
}

//HumoBuff buff
func HumoBuff(c *Data, i int) {
	if c.Weapon.Humo.On {
		var toUpAtk float32 = 0.0
		var toUpAtkNormal float32 = 0.0
		var toUpAtkHalf float32 = 0.0
		//humo
		toUpAtkNormal = c.Weapon.Humo.AddAtkByBoold * c.BooldMax[i]
		//half boold rate
		if c.Character.CurrentBooldRate != 0 && c.Character.CurrentBooldRate <= 0.5 {
			toUpAtkHalf = c.Weapon.Humo.AddAtkByHalfBoold * c.BooldMax[i]
		}
		toUpAtk = toUpAtkNormal + toUpAtkHalf
		//init
		if c.HumoInfo == nil {
			c.HumoInfo = map[int]*HumoResult{}
		}
		if c.HumoInfo[i] == nil {
			c.HumoInfo[i] = &HumoResult{}
		}
		c.HumoInfo[i].UpAtk = toUpAtk
		c.HumoInfo[i].UpAtkNormal = toUpAtkNormal
		c.HumoInfo[i].UpAtkHalf = toUpAtkHalf
		//加算
		c.Atk[i] += toUpAtk
	}
}

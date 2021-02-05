package types

//BaseMonster モンスターベース
type BaseMonster struct {
	//base
	Level              int           `json:"level"`
	EleResisRates      []DamageBoost `json:"eleResisRates"`
	EleDecreaseRates   []DamageBoost `json:"eleDecreaseRates"`
	FinalEleResisRates []DamageBoost `json:"FinalEleResisRates"`
	DefDeBuffRate      float32       `json:"defDeBuffRate"`
}

//モンスターベースの初期化
func (c *BaseMonster) init() {
	for i := range c.EleResisRates {
		c.EleResisRates[i].init()
	}
	for i := range c.EleDecreaseRates {
		c.EleDecreaseRates[i].init()
	}
	//最終耐性計算
	c.FinalEleResisRates = mergeDamageBoostDecrease(c.EleResisRates, c.EleDecreaseRates)
	//結果修正
	for i := range c.FinalEleResisRates {
		if c.FinalEleResisRates[i].DamageBoostRate < 0 {
			c.FinalEleResisRates[i].DamageBoostRate *= 0.5
		}
	}
}

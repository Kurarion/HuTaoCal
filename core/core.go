package core

import (
	json "huTao/readJSON"
	types "huTao/types"
)

//GenerateJSON JSONを生成
func GenerateJSON(toGen interface{}, err error) string {

	if err == nil {
		content, err := json.GenerateJSON(toGen)
		if err != nil {
			return ""
		}
		return string(content)
	}
	return ""
}

//CalDamage 計算最終ダメージ
func CalDamage(path string) (types.Character, error) {
	//JSONからデータを読み込み
	character, err := json.ReadJSON(path)
	if err != nil {
		return types.Character{}, err
	}
	character.Init()

	//環境デバフ等
	var decrese = getDecrease(&character)
	decrese.MonstarResisRate = getMonstarResisRate()
	decrese.LevelResisRate = getLevelResisRate()
	decrese.FinalDamageDownRate = decrese.LevelResisRate * decrese.MonstarResisRate
	//相対固定
	skills := getSkills(&character)
	groupMap := getGroupMap(&character)
	//ResultMap
	resultMap := getResultMap(&character)
	for i := range groupMap {
		//配置順で計算
		//計算用
		atk := getAtk(&character, i)
		// def := getDef(&character, i)
		// boold := getBoold(&character, i)
		damageBoostRates := getDamageBoostRate(&character, i)
		critRate := getCritRate(&character, i)
		critDamageRate := getCritDamageRate(&character, i)
		eleReactionRate := getEleReactionRate(&character, i)
		//Result
		var tempResults []types.Result = []types.Result{}
		//スキル順で計算
		for ii := range skills {
			//適用
			skillPowerRate, damageBoostRate := calDamageBoostRate(skills[ii], damageBoostRates)
			//結果
			var tempResult types.Result = types.Result{}
			tempResult.FinalDamageWithoutCrit = atk * skillPowerRate * (1 + damageBoostRate) * decrese.FinalDamageDownRate
			tempResult.FinalDamageWithCrit = tempResult.FinalDamageWithoutCrit * (1 + critDamageRate)
			if critRate > 1 {
				tempResult.FinalDamageAverage = tempResult.FinalDamageWithCrit
			} else {
				tempResult.FinalDamageAverage = tempResult.FinalDamageWithCrit*critRate + tempResult.FinalDamageWithoutCrit*(1-critRate)
			}
			tempResult.FinalEleDamageWithoutCrit = atk * skillPowerRate * (1 + damageBoostRate) * decrese.FinalDamageDownRate * (1 + eleReactionRate)
			tempResult.FinalEleDamageWithCrit = tempResult.FinalEleDamageWithoutCrit * (1 + critDamageRate)
			if critRate > 1 {
				tempResult.FinalEleDamageAverage = tempResult.FinalEleDamageWithCrit
			} else {
				tempResult.FinalEleDamageAverage = tempResult.FinalEleDamageWithCrit*critRate + tempResult.FinalEleDamageWithoutCrit*(1-critRate)
			}

			tempResults = append(tempResults, tempResult)
		}
		resultMap[i] = tempResults
	}

	return character, nil
}

func calDamageBoostRate(skill types.BaseSkill, boosts []types.DamageBoost) (skillPowerRate, damageBoostRate float32) {
	skillPowerRate = skill.DamageRate
	damageBoostRate = 0
	for _, v := range skill.DamageBoostTypes {

		for ii := range boosts {
			if boosts[ii].DamageBoostType == v {
				damageBoostRate += boosts[ii].DamageBoostRate
			}
		}
	}
	return
}

func getGroupMap(c *types.Character) map[int][]int {
	return c.GroupsMap
}

func getAtk(c *types.Character, i int) float32 {
	return c.Atk[i]
}

func getDef(c *types.Character, i int) float32 {
	return c.Def[i]
}

func getBoold(c *types.Character, i int) float32 {
	return c.BooldMax[i]
}

func getSkills(c *types.Character) []types.BaseSkill {
	return c.Skills
}

func getDecrease(c *types.Character) types.Decrease {
	return c.Decrease
}

func getResultMap(c *types.Character) map[int][]types.Result {
	return c.Results
}

func getDamageBoostRate(c *types.Character, i int) []types.DamageBoost {
	return c.DamageBoosts[i]
}

func getCritRate(c *types.Character, i int) float32 {
	return c.CritRate[i]
}

func getCritDamageRate(c *types.Character, i int) float32 {
	return c.CritDamageRate[i]
}

func getEleReactionRate(c *types.Character, i int) float32 {
	return c.EleReactionRate[i]
}

func getMonstarResisRate() float32 {
	return 0.9
}

func getLevelResisRate() float32 {
	return 0.503
}

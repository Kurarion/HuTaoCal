package core

import (
	json "huTao/readJSON"
	types "huTao/types"
)

//GenerateJSON JSONを生成
func GenerateJSON() string {
	var test types.BaseArtifact = types.BaseArtifact{}
	content, err := json.GenerateJSON(test)
	if err != nil {
		return ""
	}
	return string(content)
}

//CalDamage 計算最終ダメージ
func CalDamage(path string) {
	//JSONからデータを読み込み
	character, err := json.ReadJSON(path)
	if err != nil {
		return
	}

	//環境デバフ等
	var decrese types.Decrease = types.Decrease{}
	decrese.MonstarResisRate = getMonstarResisRate()
	decrese.LevelResisRate = getLevelResisRate()
	decrese.FinalDamageDownRate = decrese.LevelResisRate * decrese.MonstarResisRate
	//相対固定
	skills := getSkills(&character)
	groupMap := getGroupMap(&character)
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
		//スキル順で計算
		for ii := range skills {

			//適用
			damageBoostRate, skillPowerRate := calDamageBoostRate(skills[ii], damageBoostRates)
			//結果
			var tempResult types.Result = types.Result{}
			tempResult.FinalDamageWithoutCrit = atk * skillPowerRate * (1 + damageBoostRate) * decrese.FinalDamageDownRate
			tempResult.FinalDamageWithCrit = tempResult.FinalDamageWithoutCrit * (1 + critDamageRate)
			tempResult.FinalDamageAverage = tempResult.FinalDamageWithCrit * critRate
			tempResult.FinalEleDamageWithoutCrit = atk * skillPowerRate * (1 + damageBoostRate) * decrese.FinalDamageDownRate * (1 + eleReactionRate)
			tempResult.FinalEleDamageWithCrit = tempResult.FinalEleDamageWithoutCrit * (1 + critDamageRate)
			tempResult.FinalEleDamageAverage = tempResult.FinalEleDamageWithCrit * critRate
		}
	}

	return
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

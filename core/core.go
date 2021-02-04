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
	decrese := getDecrease(&character)
	decrese.LevelResisRate = getLevelResisRate(&character)
	//モンスター
	monsterEleResisRates := getMonsterEleResisRates(&character)
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
			currentMonsterResisRate := calDamageResisRate(skills[ii], monsterEleResisRates)
			currentDamageDownRate := decrese.LevelResisRate * currentMonsterResisRate
			//結果
			var tempResult types.Result = types.Result{}
			//name
			tempResult.SkillName = skills[ii].Name
			//result
			tempResult.FinalDamageWithoutCrit = atk * skillPowerRate * (1 + damageBoostRate) * currentDamageDownRate
			tempResult.FinalDamageWithCrit = tempResult.FinalDamageWithoutCrit * (1 + critDamageRate)
			if critRate > 1 {
				tempResult.FinalDamageAverage = tempResult.FinalDamageWithCrit
			} else {
				tempResult.FinalDamageAverage = tempResult.FinalDamageWithCrit*critRate + tempResult.FinalDamageWithoutCrit*(1-critRate)
			}
			//判断元素増幅
			reactionMaps := skills[ii].GetEleReation()
			for iii, vvv := range reactionMaps {
				var tempEleResult = types.EleResult{}
				tempEleResult.ReactionName = vvv
				tempEleResult.FinalEleDamageWithoutCrit = tempResult.FinalDamageWithoutCrit * iii * (1 + eleReactionRate)
				tempEleResult.FinalEleDamageWithCrit = tempResult.FinalDamageWithCrit * iii * (1 + eleReactionRate)
				tempEleResult.FinalEleDamageAverage = tempResult.FinalDamageAverage * iii * (1 + eleReactionRate)
				tempResult.FinalEleResult = append(tempResult.FinalEleResult, tempEleResult)
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

func calDamageResisRate(skill types.BaseSkill, ResisRates []types.DamageBoost) (damageResisRate float32) {
	damageResisRate = 1
	for _, v := range skill.DamageBoostTypes {

		for ii := range ResisRates {
			if ResisRates[ii].DamageBoostType == v {
				damageResisRate = (1 - ResisRates[ii].DamageBoostRate)
				return
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
	return c.DamageDecrease
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

func getMonsterEleResisRates(c *types.Character) []types.DamageBoost {
	return c.Monster.FinalEleResisRates
}

func getLevelResisRate(c *types.Character) float32 {
	return c.DamageDecrease.LevelResisRate
}

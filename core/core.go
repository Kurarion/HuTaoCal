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
func CalDamage(path string) (types.Data, error) {
	//JSONからデータを読み込み
	data, err := json.ReadJSON(path)
	if err != nil {
		return types.Data{}, err
	}
	data.Init()

	//環境デバフ等
	decrese := getDecrease(&data)
	decrese.LevelResisRate = getLevelResisRate(&data)
	//モンスター
	monsterEleResisRates := getMonsterEleResisRates(&data)
	//相対固定
	skills := getSkills(&data)
	groupMap := getGroupMap(&data)
	//ResultMap
	resultMap := getResultMap(&data)
	for i := range groupMap {
		//配置順で計算
		//計算用
		atk := getAtk(&data, i)
		// def := getDef(&data, i)
		// boold := getBoold(&data, i)
		damageBoostRates := getDamageBoostRate(&data, i)
		critRate := getCritRate(&data, i)
		critDamageRate := getCritDamageRate(&data, i)
		eleReactionRate := getEleReactionRate(&data, i)
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

	return data, nil
}

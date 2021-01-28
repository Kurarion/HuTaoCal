package core

import (
	types "huTao/types"
)

//CalDamage 計算最終ダメージ
func CalDamage() (result types.Result) {
	result.Atk = getAtk()
	result.Def = getDef()
	result.Boold = getBoold()
	result.SkillPowerRate = getSkillPowerRate()
	result.DamageBoostRate = getDamageBoostRate()
	result.CritRate = getCritRate()
	result.CritDamageRate = getCritDamageRate()
	result.EleReactionRate = getEleReactionRate()

	result.MonstarResisRate = getMonstarResisRate()
	result.LevelResisRate = getLevelResisRate()
	result.FinalDamageDownRate = result.LevelResisRate * result.MonstarResisRate

	result.FinalDamageWithoutCirt = atk * skillPower * (1 + damageBoost) * damageDownRate
	result.FinalDamageWithCirt = finalDamageWithoutCirt * critDamage
	result.FinalDamageAverage = finalDamageWithCirt * critRate

	return
}

func getAtk() {

}

func getDef() {

}

func getBoold() {

}

func getSkillPowerRate() {

}

func getDamageBoostRate() {

}

func getCritRate() {

}

func getCritDamageRate() {

}

func getEleReactionRate() {

}

func getMonstarResisRate() {
	return 0.9
}

func getLevelResisRate() {
	return 0.503
}

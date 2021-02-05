package core

import types "huTao/types"

func calDamageBoostRate(skill types.BaseSkill, boosts []types.DamageBoost) (skillPowerRate, damageBoostRate float32) {
	skillPowerRate = skill.DamageRate
	damageBoostRate = 0
	var hasAddAll = false
	for _, v := range skill.DamageBoostTypes {

		if v == types.ALL {
			continue
		}
		for ii := range boosts {
			if boosts[ii].DamageBoostType == v {
				damageBoostRate += boosts[ii].DamageBoostRate
			}
			if boosts[ii].DamageBoostType == types.ALL && !hasAddAll {
				damageBoostRate += boosts[ii].DamageBoostRate
				hasAddAll = true
			}
		}
	}
	return
}

func calDamageResisRate(skill types.BaseSkill, ResisRates []types.DamageBoost) (damageResisRate float32) {
	damageResisRate = 1
	for _, v := range skill.DamageBoostTypes {
		if v == types.ALL ||
			v == types.NORMALATK ||
			v == types.THUMP ||
			v == types.E ||
			v == types.Q {
			continue
		}
		for ii := range ResisRates {
			if ResisRates[ii].DamageBoostType == v {
				damageResisRate = (1 - ResisRates[ii].DamageBoostRate)
				return
			}
		}
	}
	return
}

func getGroupMap(c *types.Data) map[int][]int {
	return c.GroupsMap
}

func getAtk(c *types.Data, i int) float32 {
	return c.Atk[i]
}

func getDef(c *types.Data, i int) float32 {
	return c.Def[i]
}

func getBoold(c *types.Data, i int) float32 {
	return c.BooldMax[i]
}

func getSkills(c *types.Data) []types.BaseSkill {
	return c.Skills
}

func getDecrease(c *types.Data) types.Decrease {
	return c.DamageDecrease
}

func getResultMap(c *types.Data) map[int][]types.Result {
	return c.Results
}

func getDamageBoostRate(c *types.Data, i int) []types.DamageBoost {
	return c.DamageBoosts[i]
}

func getCritRate(c *types.Data, i int) float32 {
	return c.CritRate[i]
}

func getCritDamageRate(c *types.Data, i int) float32 {
	return c.CritDamageRate[i]
}

func getEleReactionRate(c *types.Data, i int) float32 {
	return c.EleReactionRate[i]
}

func getMonsterEleResisRates(c *types.Data) []types.DamageBoost {
	return c.Monster.FinalEleResisRates
}

func getLevelResisRate(c *types.Data) float32 {
	return c.DamageDecrease.LevelResisRate
}

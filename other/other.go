package other

import "fmt"

var (
	baseBlood                 float64 = 15552
	baseAtk                   float64 = 714
	baseCritRate              float64 = 0.05
	baseCritDamage            float64 = 0.884
	hutaoChangeRate           float64 = 0.0715
	hutaoChangeMaxBaseAtkRate float64 = 4
	humoChangeRate            float64 = 0.016 + 0.018
	humoBooldBuff             float64 = 0.4
	humoCritDamage            float64 = 0.662
)

var (
	mainBlood      float64 = 4780
	mainAtk        float64 = 311
	mainBloodRate  float64 = 0.466
	mainAtkRate    float64 = 0
	mainCriteRate  float64 = 0
	mainCritDamage float64 = 0.622
)

var (
	maxSub = 4*5 + 20/4*5
	minSub = 3*5 + 20/4*5
)

type subType = int

const (
	start subType = iota
	// atk
	// boold
	// def
	// ele
	// charge
	booldBuff
	atkBuff
	// defBuff
	critRate
	critDamage
	end
)

var subValues map[subType]float64
var subNames map[subType]string

func init() {
	subValues = make(map[subType]float64)
	// subValues[boold] = 299
	// subValues[atk] = 19
	// subValues[def] = 23
	// subValues[ele] = 23
	// subValues[charge] = 0.065
	subValues[booldBuff] = 0.058
	subValues[atkBuff] = 0.058
	// subValues[defBuff] = 0.073
	subValues[critRate] = 0.039
	subValues[critDamage] = 0.078

	subNames = make(map[subType]string)
	subNames[booldBuff] = "血量百分比"
	subNames[atkBuff] = "攻击百分比"
	subNames[critRate] = "暴击率"
	subNames[critDamage] = "暴击伤害"
}

type result struct {
	Subs       map[subType]int
	DamageRate float64
	UpRate     float64
}

func (c *result) init() {
	c.Subs = make(map[subType]int)
}

func (c *result) clone() (res result) {
	res = result{}
	res.init()
	for k, v := range c.Subs {
		res.Subs[k] = v
	}
	res.DamageRate = c.DamageRate
	res.UpRate = c.UpRate

	return
}

func getDamageRate(lastRes result) float64 {
	var subBooldBuff, subAtkBuff, subCritRateBuff, subCirtDamageBuff float64
	subBooldBuff = float64(lastRes.Subs[booldBuff]) * subValues[booldBuff]
	subAtkBuff = float64(lastRes.Subs[atkBuff]) * subValues[atkBuff]
	subCritRateBuff = float64(lastRes.Subs[critRate]) * subValues[critRate]
	subCirtDamageBuff = float64(lastRes.Subs[critDamage]) * subValues[critDamage]
	//Atk
	var totalBlood float64 = baseBlood + baseBlood*(humoBooldBuff+mainBloodRate+subBooldBuff) + (mainBlood)
	//hutao
	var maxBloodToAtk float64 = baseAtk * hutaoChangeMaxBaseAtkRate
	var currentBloodToAtk float64 = hutaoChangeRate * totalBlood
	var toAddAtk float64
	if maxBloodToAtk > currentBloodToAtk {
		toAddAtk = currentBloodToAtk
	} else {
		toAddAtk = maxBloodToAtk
	}
	//humo
	toAddAtk += humoChangeRate * totalBlood
	var totalAtk float64 = baseAtk + baseAtk*(mainAtkRate+subAtkBuff) + (mainAtk) + toAddAtk
	//Damage
	var damageRate float64
	var totalCriteDamage float64 = 1 + baseCritDamage + humoCritDamage + mainCritDamage + subCirtDamageBuff
	var currentCriteRate float64 = baseCritRate + mainCriteRate + subCritRateBuff
	var maxCriteRate float64 = 1
	var totalCriteRate float64
	if maxCriteRate > currentCriteRate {
		totalCriteRate = currentCriteRate
	} else {
		totalCriteRate = maxCriteRate
	}
	damageRate = totalAtk * (totalCriteRate) * (totalCriteDamage)
	return damageRate
}

//TestArtifact the best value of artifact
func TestArtifact() {
	var data map[int]result = make(map[int]result)
	//init
	var zeroResult = result{}
	zeroResult.init()
	for j := int(start) + 1; j < int(end); j++ {
		zeroResult.Subs[subType(j)] = 0
	}
	data[0] = zeroResult
	//exec
	for i := 1; i <= maxSub; i++ {
		var tempResult = data[i-1]
		var thisResult = tempResult.clone()
		var currentMaxDamage float64
		var currentSelectType subType
		for j := int(start) + 1; j < int(end); j++ {
			thisResult.Subs[subType(j)]++
			var damage = getDamageRate(thisResult)
			if damage > currentMaxDamage {
				currentSelectType = subType(j)
				currentMaxDamage = damage
			}
			thisResult.Subs[subType(j)]--
		}
		thisResult.Subs[currentSelectType]++
		thisResult.DamageRate = currentMaxDamage
		thisResult.UpRate = thisResult.DamageRate / tempResult.DamageRate
		data[i] = thisResult
	}
	//output
	for i := 1; i <= maxSub; i++ {
		k := i
		v := data[i]
		fmt.Printf("[%2v]->\n", k)
		for j := int(start) + 1; j < int(end); j++ {
			kk := j
			vv := v.Subs[j]
			fmt.Printf("       %s: %v", subNames[kk], float64(vv)*subValues[kk])
			switch kk {
			case booldBuff:
				fmt.Printf(" (%v,%v)\n", humoBooldBuff+mainBloodRate+float64(vv)*subValues[kk], baseBlood+baseBlood*(humoBooldBuff+mainBloodRate+float64(vv)*subValues[kk])+(mainBlood))
			case atkBuff:
				fmt.Print("\n")
			case critRate:
				fmt.Printf(" (%v)\n", baseCritRate+mainCriteRate+float64(vv)*subValues[kk])
			case critDamage:
				fmt.Printf(" (%v)\n", baseCritDamage+humoCritDamage+mainCritDamage+float64(vv)*subValues[kk])
			}
		}
		fmt.Printf("       <伤害倍率: %v>\n", v.DamageRate)
		fmt.Printf("       <提升: %v>\n", v.UpRate)
		fmt.Print("\n")
	}

}

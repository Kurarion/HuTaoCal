package types

//BaseCharacter チャラベース
type BaseCharacter struct {
	//base
	SelfAtk             int
	SelfDef             int
	SelfBoold           int
	SelfSkillPowerRates int
	SelfDamageBoostRate int
	SelfCritRate        int
	SelfCritDamageRate  int
	SelfEleValue        int
	//rate buff
	SelfAtkBuffRate   int
	SelfDefBuffRate   int
	SelfBooldBuffRate int
	//value buff
	SelfDamageBoostRateBuff int
	SelfCritRateBuff        int
	SelfCritDamageRateBuff  int
	SelfEleValueBuff        int
}

//BaseWeapon 武器ベース
type BaseWeapon struct {
	//base
	WeaponAtk int
	//rate buff
	WeaponAtkBuffRate   int
	WeaponDefBuffRate   int
	WeaponBooldBuffRate int
	//value buff
	WeaponDamageBoostRateBuff int
	WeaponCritRateBuff        int
	WeaponCritDamageRateBuff  int
	WeaponEleValueBuff        int
}

//BaseArtifact アーティファクトベース
type BaseArtifact struct {
	//base
	SelfAtk      int
	SelfDef      int
	SelfBoold    int
	SelfEleValue int
	//rate buff
	SelfAtkBuffRate   int
	SelfDefBuffRate   int
	SelfBooldBuffRate int
	//value buff
	SelfDamageBoostRateBuff int
	SelfCritRateBuff        int
	SelfCritDamageRateBuff  int
}

//Character チャラ属性
type Character struct {
	BaseCharacter
	BaseWeapon
	List            []BaseArtifact
	Atk             int
	Def             int
	BooldMax        int
	BooldRate       int
	BooldCurrent    int
	SkillPowerRates int
	DamageBoostRate int
	CritRate        int
	CritDamageRate  int
	EleValue        int
	EleReactionRate int
}

//Result 計算結果ストラクト
type Result struct {
	//character base
	Character
	//decrease
	MonstarResisRate
	LevelResisRate
	FinalDamageDownRate int
	//damageResult
	FinalDamageWithoutCirt
	FinalDamageWithCrit
	FinalDamageAverage int
	FinalEleDamageWithoutCirt
	FinalEleDamageWithCrit
	FinalEleDamageAverage int
}

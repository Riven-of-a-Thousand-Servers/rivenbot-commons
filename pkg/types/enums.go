package types

type CharacterClass string

const (
	TITAN   CharacterClass = "TITAN"
	WARLOCK CharacterClass = "WARLOCK"
	HUNTER  CharacterClass = "HUNTER"
)

type RaidName string

const (
	SALVATIONS_EDGE     RaidName = "SALVATIONS_EDGE"
	CROTAS_END          RaidName = "CROTAS_END"
	ROOT_OF_NIGHTMARES  RaidName = "ROOT_OF_NIGHTMARES"
	KINGS_FALL          RaidName = "KINGS_FALL"
	VOW_OF_THE_DISCIPLE RaidName = "VOW_OF_THE_DISCIPLE"
	VAULT_OF_GLASS      RaidName = "VAULT_OF_GLASS"
	DEEP_STONE_CRYPT    RaidName = "DEEP_STONE_CRYPT"
	GARDEN_OF_SALVATION RaidName = "GARDEN_OF_SALVATION"
	CROWN_OF_SORROW     RaidName = "CROWN_OF_SORROW"
	LAST_WISH           RaidName = "LAST_WISH"
	SPIRE_OF_STARS      RaidName = "SPIRE_OF_STARS"
	EATER_OF_WORLDS     RaidName = "EATER_OF_WORLDS"
	LEVIATHAN           RaidName = "LEVIATHAN"
	SCOURGE_OF_THE_PAST RaidName = "SCOURGE_OF_THE_PAST"
)

type RaidDifficulty string

const (
	NORMAL         RaidDifficulty = "NORMAL"
	PRESTIGE       RaidDifficulty = "PRESTIGE"
	MASTER         RaidDifficulty = "MASTER"
	GUIDED_GAMES   RaidDifficulty = "GUIDED_GAMES"
	CHALLENGE_MODE RaidDifficulty = "CHALLENGE_MODE"
)

type DamageType string

const (
	KINETIC DamageType = "KINETIC"
	ARC     DamageType = "ARC"
	VOID    DamageType = "VOID"
	SOLAR   DamageType = "SOLAR"
	STASIS  DamageType = "STASIS"
	STRAND  DamageType = "STRAND"
)

type EquipmentSlot string

const (
	PRIMARY EquipmentSlot = "PRIMARY"
	SPECIAL EquipmentSlot = "SPECIAL"
	HEAVY   EquipmentSlot = "HEAVY"
)

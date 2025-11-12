package commons

import "time"

type InstanceActivityEntity struct {
	InstanceId         int64
	PlayerMembershipId int64
	PlayerCharacterId  int64
	CharacterEmblem    int64
	IsCompleted        bool
	Kills              int32
	Deaths             int32
	Assists            int32
	MeleeKills         int
	SuperKills         int
	GrenadeKills       int
	KillsDeathsAssists float32
	KillsDeathsRatio   float32
	DurationSeconds    int64
	TimeplayedSeconds  int64
}

type InstanceWeaponStats struct {
	InstanceId          int64
	PlayerCharacterId   int64
	WeaponId            int64
	TotalKills          int
	TotalPrecisionKills int
	PrecisionRatio      float32
}

type PlayerEntity struct {
	MembershipId    int64
	DisplayName     string
	DisplayNameCode int32
	MembershipType  int32
	LastSeen        time.Time
	Characters      []PlayerCharacterEntity
}

type PlayerCharacterEntity struct {
	CharacterId        int64
	CharacterClass     string
	CharacterEmblem    int64
	PlayerMembershipId int64
}

type PlayerRaidStatsEntity struct {
	RaidName           RaidName
	RaidDifficulty     RaidDifficulty
	PlayerMembershipId int64
	Kills              int
	Deaths             int
	Assists            int
	HoursPlayed        int
	Clears             int
	FullClears         int
	Flawless           bool
	ContestClear       bool
	DayOne             bool
	Solo               bool
	Duo                bool
	Trio               bool
	SoloFlawless       bool
	DuoFlawless        bool
	TrioFlawless       bool
}

type RaidEntity struct {
	RaidName       string
	RaidDifficulty string
	IsActive       bool
	ReleaseDate    time.Time
	RaidHash       int64
}

type RaidPgcr struct {
	InstanceId int64
	Blob       []byte
}

type WeaponEntity struct {
	WeaponHash          int64
	WeaponIcon          string
	WeaponName          string
	WeaponDamageType    DamageType
	WeaponEquipmentSlot EquipmentSlot
}

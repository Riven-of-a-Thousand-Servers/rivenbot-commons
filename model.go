package commons

import "time"

type Activity struct {
	Hash          int64
	Label         string
	NameId        int64
	DifficultyId  int64
	IsWorldsFirst bool
}

type ActivityDifficulty struct {
	Id    int64
	Label string
}

type ActivityName struct {
	Id          int64
	Label       string
	IsActive    bool
	ReleaseDate time.Time
}

type DestinyPlayer struct {
	MembershipId          int64
	MembershipType        int
	IconPath              string
	DisplayName           string
	GlobalDisplayName     string
	GlobalDisplayNameCode int
	TotalClears           int
	TotalFullClears       int
	IsPrivate             bool
	LastCrawled           time.Time
	LastSeen              time.Time
}

type Instance struct {
	Id           int64
	ActivityHash int64
	IsFresh      bool
	Flawless     bool
	Completed    bool
	PlayerCount  int
	Duration     int
	StartTime    time.Time
	EndTime      time.Time
}

type InstancePlayer struct {
	InstanceId   int64
	MembershipId int64
	Completed    bool
	TimePlayed   int
}

type InstanceCharacter struct {
	InstanceId        int64
	MembershipId      int64
	CharacterId       int64
	ClassHash         int64
	EmblemHash        int64
	Completed         bool
	Kills             int
	Deaths            int
	Assists           int
	Kda               float32
	Kdr               float32
	SuperKills        int
	MeleeKills        int
	GrenadeKills      int
	Efficiency        int
	TimePlayedSeconds int
}

type InstanceCharacterWeapon struct {
	InstanceId         int64
	PlayerMembershipId int64
	PlayerCharacterId  int64
	WeaponId           int64
	Kills              int64
	PrecisionKills     int64
	PrecisionRatio     float32
}

type Weapon struct {
	Hash          int64
	IconUrl       string
	Name          string
	EquipmentSlot string
	DamageType    string
}

package utils

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/Riven-of-a-Thousand-Servers/rivenbot-commons"
)

var reverseClassLabels map[string]commons.CharacterClass
var characterClassLabels = map[commons.CharacterClass]string{
	commons.TITAN:   "Titan",
	commons.WARLOCK: "Warlock",
	commons.HUNTER:  "Hunter",
}

func ClassLabel(c commons.CharacterClass) string {
	if label, exists := characterClassLabels[c]; exists {
		return label
	} else {
		return ""
	}
}

var reverseRaidLabels map[string]commons.RaidName
var raidNameLabels = map[commons.RaidName]string{
	commons.SALVATIONS_EDGE:     "Salvation's Edge",
	commons.CROTAS_END:          "Crota's End",
	commons.ROOT_OF_NIGHTMARES:  "Root of Nightmares",
	commons.KINGS_FALL:          "King's Fall",
	commons.VOW_OF_THE_DISCIPLE: "Vow of the Disciple",
	commons.VAULT_OF_GLASS:      "Vault of Glass",
	commons.DEEP_STONE_CRYPT:    "Deep Stone Crypt",
	commons.GARDEN_OF_SALVATION: "Garden of Salvation",
	commons.CROWN_OF_SORROW:     "Crown of Sorrow",
	commons.LAST_WISH:           "Last Wish",
	commons.SPIRE_OF_STARS:      "Leviathan, Spire of Stars",
	commons.EATER_OF_WORLDS:     "Leviathan, Eater of Worlds",
	commons.LEVIATHAN:           "Leviathan",
	commons.SCOURGE_OF_THE_PAST: "Scourge of the Past",
}

func RaidLabel(rn commons.RaidName) string {
	if label, exists := raidNameLabels[rn]; exists {
		return label
	} else {
		return ""
	}
}

var reverseDifficultyLabels map[string]commons.RaidDifficulty
var raidDifficultyLabels = map[commons.RaidDifficulty]string{
	commons.NORMAL:         "Normal",
	commons.PRESTIGE:       "Prestige",
	commons.MASTER:         "Master",
	commons.GUIDED_GAMES:   "Guided Games",
	commons.CHALLENGE_MODE: "Challenge Mode",
}

var once sync.Once

// Initializes the reverse look up maps only once
func initReverseMaps() {
	once.Do(func() {
		reverseDifficultyLabels = make(map[string]commons.RaidDifficulty)
		for raidDifficulty, label := range raidDifficultyLabels {
			reverseDifficultyLabels[label] = raidDifficulty
		}

		reverseRaidLabels = make(map[string]commons.RaidName)
		for raidName, label := range raidNameLabels {
			reverseRaidLabels[label] = raidName
		}

		reverseClassLabels = make(map[string]commons.CharacterClass)
		for characterClass, label := range characterClassLabels {
			reverseClassLabels[label] = characterClass
		}
	})
}

// This method returns an instance of a Raid Name and RaidDifficulty
// given a string, e.g., Last Wish should yield both the RaidName LAST_WISH
// and the RaidDiffculty NORMAL. On the other hand, Salvation's Edge: Master
// should yield RaidName SALVATIONS_EDGE and RaidDifficulty MASTER
func GetRaidAndDifficulty(label string) (commons.RaidName, commons.RaidDifficulty, error) {
	initReverseMaps()
	tokens := strings.Split(label, ":")

	if len(tokens) <= 0 {
		log.Panicf("Unable to tokenize raid Manifest Display Name [%s]", label)
		return "", "", errors.New("Unable to tokenize raid Manifest Display Name")
	}
	name := strings.TrimSpace(tokens[0])
	raidName, nameExists := reverseRaidLabels[name]

	if !nameExists {
		return "", "", fmt.Errorf("Raid name [%s] has no match", name)
	}

	if len(tokens) <= 1 {
		return raidName, commons.NORMAL, nil
	}

	difficulty := strings.TrimSpace(tokens[1]) // Default difficulty
	raidDifficulty, difficultyExists := reverseDifficultyLabels[difficulty]
	log.Printf("Difficulty exists for raid [%s]: %v", label, difficultyExists)
	if !difficultyExists {
		switch {
		case strings.EqualFold(difficulty, "Standard"):
			raidDifficulty = reverseDifficultyLabels["Normal"]
		case strings.EqualFold(difficulty, "Expert") || strings.EqualFold(difficulty, "Legend"):
			raidDifficulty = reverseDifficultyLabels["Challenge Mode"]
		default:
			return "", "", fmt.Errorf("Raid difficulty [%s] has no match", difficulty)
		}
	}

	return raidName, raidDifficulty, nil
}

func GetDamageType(enumValue int) commons.DamageType {
	switch enumValue {
	case 1:
		return commons.KINETIC
	case 2:
		return commons.ARC
	case 3:
		return commons.SOLAR
	case 4:
		return commons.VOID
	case 6:
		return commons.STASIS
	case 7:
		return commons.STRAND
	default:
		return ""
	}
}

type EquippingBlockcommons interface {
	~int64 | ~string
}

func GetEquippingSlot[T EquippingBlockcommons](enumValue T) commons.EquipmentSlot {
	value := any(enumValue)
	if v, ok := value.(int64); ok {
		switch v {
		case 1498876634:
			return commons.PRIMARY
		case 2465295065:
			return commons.SPECIAL
		case 953998645:
			return commons.HEAVY
		}
	} else if s, ok := value.(string); ok {
		switch strings.ToLower(s) {
		case "kinetic weapons":
			return commons.PRIMARY
		case "energy weapons":
			return commons.SPECIAL
		case "power weapons":
			return commons.HEAVY
		}
	}
	return ""
}

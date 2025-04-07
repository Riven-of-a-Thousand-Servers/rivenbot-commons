package utils

import (
	"fmt"
	"testing"

	"github.com/Riven-of-a-Thousand-Servers/rivenbot-commons/pkg/types"
	"github.com/stretchr/testify/assert"
)

var enumUtilTests = map[string]struct {
	input      string
	raid       types.RaidName
	difficulty types.RaidDifficulty
}{
	"leviathan_normal": {
		input:      "Leviathan: Normal",
		raid:       types.LEVIATHAN,
		difficulty: types.NORMAL,
	},
	"leviathan_prestige": {
		input:      "Leviathan: Prestige",
		raid:       types.LEVIATHAN,
		difficulty: types.PRESTIGE,
	},
	"leviathan_guided_games": {
		input:      "Leviathan: Guided Games",
		raid:       types.LEVIATHAN,
		difficulty: types.GUIDED_GAMES,
	},
	"eater_of_worlds_normal": {
		input:      "Leviathan, Eater of Worlds: Normal",
		raid:       types.EATER_OF_WORLDS,
		difficulty: types.NORMAL,
	},
	"eater_of_worlds_prestige": {
		input:      "Leviathan, Eater of Worlds: Prestige",
		raid:       types.EATER_OF_WORLDS,
		difficulty: types.PRESTIGE,
	},
	"eater_of_worlds_guided_games": {
		input:      "Leviathan, Eater of Worlds: Guided Games",
		raid:       types.EATER_OF_WORLDS,
		difficulty: types.GUIDED_GAMES,
	},
	"spire_of_stars_normal": {
		input:      "Leviathan, Spire of Stars: Normal",
		raid:       types.SPIRE_OF_STARS,
		difficulty: types.NORMAL,
	},
	"spire_of_stars_prestige": {
		input:      "Leviathan, Spire of Stars: Prestige",
		raid:       types.SPIRE_OF_STARS,
		difficulty: types.PRESTIGE,
	},
	"spire_of_stars_guided_games": {
		input:      "Leviathan, Spire of Stars: Guided Games",
		raid:       types.SPIRE_OF_STARS,
		difficulty: types.GUIDED_GAMES,
	},
	"scourge_of_the_past": {
		input:      "Scourge of the Past",
		raid:       types.SCOURGE_OF_THE_PAST,
		difficulty: types.NORMAL,
	},
	"last_wish": {
		input:      "Last Wish",
		raid:       types.LAST_WISH,
		difficulty: types.NORMAL,
	},
	"crown_of_sorrow": {
		input:      "Crown of Sorrow",
		raid:       types.CROWN_OF_SORROW,
		difficulty: types.NORMAL,
	},
	"garden_of_salvation": {
		input:      "Garden of Salvation",
		raid:       types.GARDEN_OF_SALVATION,
		difficulty: types.NORMAL,
	},
	"deep_stone_crypt": {
		input:      "Deep Stone Crypt",
		raid:       types.DEEP_STONE_CRYPT,
		difficulty: types.NORMAL,
	},
	"vault_of_glass_normal": {
		input:      "Vault of Glass: Standard",
		raid:       types.VAULT_OF_GLASS,
		difficulty: types.NORMAL,
	},
	"vault_of_glass_master": {
		input:      "Vault of Glass: Master",
		raid:       types.VAULT_OF_GLASS,
		difficulty: types.MASTER,
	},
	"vault_of_glass_challenge": {
		input:      "Vault of Glass: Challenge Mode",
		raid:       types.VAULT_OF_GLASS,
		difficulty: types.CHALLENGE_MODE,
	},
	"vow_of_the_disciple_normal": {
		input:      "Vow of the Disciple: Standard",
		raid:       types.VOW_OF_THE_DISCIPLE,
		difficulty: types.NORMAL,
	},
	"vow_of_the_disciple_master": {
		input:      "Vow of the Disciple: Master",
		raid:       types.VOW_OF_THE_DISCIPLE,
		difficulty: types.MASTER,
	},
	"kings_fall_challenge_mode": {
		input:      "King's Fall: Expert",
		raid:       types.KINGS_FALL,
		difficulty: types.CHALLENGE_MODE,
	},
	"kings_fall_normal": {
		input:      "King's Fall: Standard",
		raid:       types.KINGS_FALL,
		difficulty: types.NORMAL,
	},
	"kings_fall_master": {
		input:      "King's Fall: Master",
		raid:       types.KINGS_FALL,
		difficulty: types.MASTER,
	},
	"root_of_nightmares_normal": {
		input:      "Root of Nightmares: Standard",
		raid:       types.ROOT_OF_NIGHTMARES,
		difficulty: types.NORMAL,
	},
	"root_of_nightmares_master": {
		input:      "Root of Nightmares: Master",
		raid:       types.ROOT_OF_NIGHTMARES,
		difficulty: types.MASTER,
	},
	"crotas_end_challenge_mode": {
		input:      "Crota's End: Legend",
		raid:       types.CROTAS_END,
		difficulty: types.CHALLENGE_MODE,
	},
	"crotas_end_normal": {
		input:      "Crota's End: Normal",
		raid:       types.CROTAS_END,
		difficulty: types.NORMAL,
	},
	"crotas_end_master": {
		input:      "Crota's End: Master",
		raid:       types.CROTAS_END,
		difficulty: types.MASTER,
	},
	"salvations_edge_normal": {
		input:      "Salvation's Edge: Standard",
		raid:       types.SALVATIONS_EDGE,
		difficulty: types.NORMAL,
	},
	// This one is a contest mode clear
	"salvations_edge_normal_2.0": {
		input:      "Salvation's Edge",
		raid:       types.SALVATIONS_EDGE,
		difficulty: types.NORMAL,
	},
	"salvations_edge_master": {
		input:      "Salvation's Edge: Master",
		raid:       types.SALVATIONS_EDGE,
		difficulty: types.MASTER,
	},
}

func TestGetRaidAndDifficulty_Success(t *testing.T) {
	// given: A string representing a raid
	for test, params := range enumUtilTests {
		t.Run(test, func(t *testing.T) {
			// when: GetRaidAndDifficulty is called
			raid, difficulty, err := GetRaidAndDifficulty(params.input)

			if err != nil {
				t.Fatalf("Expected no errors, found one: %v", err)
			}

			assert := assert.New(t)
			// then: Raid and Difficulty values are correct
			assert.Equal(params.raid, raid, fmt.Sprintf("Raid should be [%s]", params.raid))
			assert.Equal(params.difficulty, difficulty, fmt.Sprintf("Difficulty should be [%s]", params.difficulty))
		})
	}
}

type DamageTypeTest struct {
	Value    int
	Expected types.DamageType
}

var damageTypeInputs = map[string]DamageTypeTest{
	"kinetic damage test": {
		Value:    1,
		Expected: types.KINETIC,
	},
	"arc damage test": {
		Value:    2,
		Expected: types.ARC,
	},
	"solar damage test": {
		Value:    3,
		Expected: types.SOLAR,
	},
	"void damage test": {
		Value:    4,
		Expected: types.VOID,
	},
	"stasis damage test": {
		Value:    6,
		Expected: types.STASIS,
	},
	"strand damage test": {
		Value:    7,
		Expected: types.STRAND,
	},
	"zero damage test": {
		Value:    0,
		Expected: "",
	},
	"raids? damage test": {
		Value:    5,
		Expected: "",
	},
}

func TestGetDamageType(t *testing.T) {
	for test, input := range damageTypeInputs {
		t.Run(test, func(t *testing.T) {
			// when: GetDamageType is called
			result := GetDamageType(input.Value)

			if result != input.Expected {
				t.Fatalf("Expected: %s. Got: %s", input.Expected, result)
			}
		})
	}
}

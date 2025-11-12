package utils

import (
	"fmt"
	"testing"

	"github.com/Riven-of-a-Thousand-Servers/rivenbot-commons"
	"github.com/stretchr/testify/assert"
)

var enumUtilTests = map[string]struct {
	input      string
	raid       commons.RaidName
	difficulty commons.RaidDifficulty
}{
	"leviathan_normal": {
		input:      "Leviathan: Normal",
		raid:       commons.LEVIATHAN,
		difficulty: commons.NORMAL,
	},
	"leviathan_prestige": {
		input:      "Leviathan: Prestige",
		raid:       commons.LEVIATHAN,
		difficulty: commons.PRESTIGE,
	},
	"leviathan_guided_games": {
		input:      "Leviathan: Guided Games",
		raid:       commons.LEVIATHAN,
		difficulty: commons.GUIDED_GAMES,
	},
	"eater_of_worlds_normal": {
		input:      "Leviathan, Eater of Worlds: Normal",
		raid:       commons.EATER_OF_WORLDS,
		difficulty: commons.NORMAL,
	},
	"eater_of_worlds_prestige": {
		input:      "Leviathan, Eater of Worlds: Prestige",
		raid:       commons.EATER_OF_WORLDS,
		difficulty: commons.PRESTIGE,
	},
	"eater_of_worlds_guided_games": {
		input:      "Leviathan, Eater of Worlds: Guided Games",
		raid:       commons.EATER_OF_WORLDS,
		difficulty: commons.GUIDED_GAMES,
	},
	"spire_of_stars_normal": {
		input:      "Leviathan, Spire of Stars: Normal",
		raid:       commons.SPIRE_OF_STARS,
		difficulty: commons.NORMAL,
	},
	"spire_of_stars_prestige": {
		input:      "Leviathan, Spire of Stars: Prestige",
		raid:       commons.SPIRE_OF_STARS,
		difficulty: commons.PRESTIGE,
	},
	"spire_of_stars_guided_games": {
		input:      "Leviathan, Spire of Stars: Guided Games",
		raid:       commons.SPIRE_OF_STARS,
		difficulty: commons.GUIDED_GAMES,
	},
	"scourge_of_the_past": {
		input:      "Scourge of the Past",
		raid:       commons.SCOURGE_OF_THE_PAST,
		difficulty: commons.NORMAL,
	},
	"last_wish": {
		input:      "Last Wish",
		raid:       commons.LAST_WISH,
		difficulty: commons.NORMAL,
	},
	"crown_of_sorrow": {
		input:      "Crown of Sorrow",
		raid:       commons.CROWN_OF_SORROW,
		difficulty: commons.NORMAL,
	},
	"garden_of_salvation": {
		input:      "Garden of Salvation",
		raid:       commons.GARDEN_OF_SALVATION,
		difficulty: commons.NORMAL,
	},
	"deep_stone_crypt": {
		input:      "Deep Stone Crypt",
		raid:       commons.DEEP_STONE_CRYPT,
		difficulty: commons.NORMAL,
	},
	"vault_of_glass_normal": {
		input:      "Vault of Glass: Standard",
		raid:       commons.VAULT_OF_GLASS,
		difficulty: commons.NORMAL,
	},
	"vault_of_glass_master": {
		input:      "Vault of Glass: Master",
		raid:       commons.VAULT_OF_GLASS,
		difficulty: commons.MASTER,
	},
	"vault_of_glass_challenge": {
		input:      "Vault of Glass: Challenge Mode",
		raid:       commons.VAULT_OF_GLASS,
		difficulty: commons.CHALLENGE_MODE,
	},
	"vow_of_the_disciple_normal": {
		input:      "Vow of the Disciple: Standard",
		raid:       commons.VOW_OF_THE_DISCIPLE,
		difficulty: commons.NORMAL,
	},
	"vow_of_the_disciple_master": {
		input:      "Vow of the Disciple: Master",
		raid:       commons.VOW_OF_THE_DISCIPLE,
		difficulty: commons.MASTER,
	},
	"kings_fall_challenge_mode": {
		input:      "King's Fall: Expert",
		raid:       commons.KINGS_FALL,
		difficulty: commons.CHALLENGE_MODE,
	},
	"kings_fall_normal": {
		input:      "King's Fall: Standard",
		raid:       commons.KINGS_FALL,
		difficulty: commons.NORMAL,
	},
	"kings_fall_master": {
		input:      "King's Fall: Master",
		raid:       commons.KINGS_FALL,
		difficulty: commons.MASTER,
	},
	"root_of_nightmares_normal": {
		input:      "Root of Nightmares: Standard",
		raid:       commons.ROOT_OF_NIGHTMARES,
		difficulty: commons.NORMAL,
	},
	"root_of_nightmares_master": {
		input:      "Root of Nightmares: Master",
		raid:       commons.ROOT_OF_NIGHTMARES,
		difficulty: commons.MASTER,
	},
	"crotas_end_challenge_mode": {
		input:      "Crota's End: Legend",
		raid:       commons.CROTAS_END,
		difficulty: commons.CHALLENGE_MODE,
	},
	"crotas_end_normal": {
		input:      "Crota's End: Normal",
		raid:       commons.CROTAS_END,
		difficulty: commons.NORMAL,
	},
	"crotas_end_master": {
		input:      "Crota's End: Master",
		raid:       commons.CROTAS_END,
		difficulty: commons.MASTER,
	},
	"salvations_edge_normal": {
		input:      "Salvation's Edge: Standard",
		raid:       commons.SALVATIONS_EDGE,
		difficulty: commons.NORMAL,
	},
	// This one is a contest mode clear
	"salvations_edge_normal_2.0": {
		input:      "Salvation's Edge",
		raid:       commons.SALVATIONS_EDGE,
		difficulty: commons.NORMAL,
	},
	"salvations_edge_master": {
		input:      "Salvation's Edge: Master",
		raid:       commons.SALVATIONS_EDGE,
		difficulty: commons.MASTER,
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
	Expected commons.DamageType
}

var damageTypeInputs = map[string]DamageTypeTest{
	"kinetic damage test": {
		Value:    1,
		Expected: commons.KINETIC,
	},
	"arc damage test": {
		Value:    2,
		Expected: commons.ARC,
	},
	"solar damage test": {
		Value:    3,
		Expected: commons.SOLAR,
	},
	"void damage test": {
		Value:    4,
		Expected: commons.VOID,
	},
	"stasis damage test": {
		Value:    6,
		Expected: commons.STASIS,
	},
	"strand damage test": {
		Value:    7,
		Expected: commons.STRAND,
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

			// then: the result should match the expected output
			if result != input.Expected {
				t.Fatalf("Expected: %s. Got: %s", input.Expected, result)
			}
		})
	}
}

type EquipmentSlotTestCase[T EquippingBlockcommons] struct {
	Input    T
	Expected commons.EquipmentSlot
}

func TestGetEquippingSlot(t *testing.T) {
	stringTests := []EquipmentSlotTestCase[string]{
		{
			Input:    "KinetiC weapOns",
			Expected: commons.PRIMARY,
		},
		{
			Input:    "EneRgY WeapOns",
			Expected: commons.SPECIAL,
		},
		{
			Input:    "PowEr WeaPONs",
			Expected: commons.HEAVY,
		},
		{
			Input:    "ESome Typo",
			Expected: "",
		},
	}

	for _, testcase := range stringTests {
		t.Run(fmt.Sprintf("Input: %v", testcase.Input), func(t *testing.T) {
			// when: GetEquippingSlot is called
			result := GetEquippingSlot(testcase.Input)

			// then: the result should match the expected output
			if result != testcase.Expected {
				t.Fatalf("Expected: %v. Actual: %v", testcase.Expected, result)
			}
		})
	}

	intTests := []EquipmentSlotTestCase[int64]{
		{
			Input:    1498876634,
			Expected: commons.PRIMARY,
		},
		{
			Input:    2465295065,
			Expected: commons.SPECIAL,
		},
		{
			Input:    953998645,
			Expected: commons.HEAVY,
		},
	}

	for _, testcase := range intTests {
		t.Run(fmt.Sprintf("Input: %v", testcase.Input), func(t *testing.T) {
			// when: GetEquippingSlot is called
			result := GetEquippingSlot(testcase.Input)

			// then: the result should match the expected output
			if result != testcase.Expected {
				t.Fatalf("Expected: %v. Actual: %v", testcase.Expected, result)
			}
		})
	}
}

package test

import "testing"

func TestNewDice(t *testing.T)  {
	dice := setupDice()

	if dice.GetNumberOfFace() != 3 {
		t.Errorf("Expected numberOfFace to be %d, but got %d", 3, dice.GetNumberOfFace())
	}
}

func TestRoll(t *testing.T)  {
	dice := setupDice()

	for i := 0; i < 10; i++ {
		result := dice.Roll()
		if result < 1 || result > dice.GetNumberOfFace() {
			t.Errorf("Roll() returned %d, expected value between 1 and %d", result, dice.GetNumberOfFace())
		}
	}
}
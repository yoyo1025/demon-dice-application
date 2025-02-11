package test

import (
	"testing"
)

func TestNewTurn(t *testing.T)  {
	turn := setupTurn(15)
	if turn.GetMaxTurn() != 15 {
		t.Errorf("Expected maxTurn to be %d, but got %d", 15, turn.GetMaxTurn())
	}
	if turn.GetCurrentTurn() != 1 {
		t.Errorf("Expected currentTurn to be %d, but got %d", 1, turn.GetCurrentTurn())
	}
	if turn.GetCurrentPlayerIndex() != 1 {
		t.Errorf("Expected currentPlayerIndex to be %d, but got %d", 1, turn.GetCurrentPlayerIndex())
	}
	if turn.GetMaxPlayerIndex() != 4 {
		t.Errorf("Expected maxPlayerIndex to be %d, but got %d", 4, turn.GetMaxPlayerIndex())
	}
}

func TestNextTurn(t *testing.T)  {
	turn := setupTurn(15)
	if turn.GetCurrentTurn() != 1 {
		t.Errorf("Expected currentTurn to be %d, but got %d", 1, turn.GetCurrentTurn())
	}
	turn.NextTurn()
	if turn.GetCurrentTurn() != 2 {
		t.Errorf("Expected currentTurn to be %d, but got %d", 2, turn.GetCurrentTurn())
	}
}

func TestIsMaxTurnReached(t *testing.T)  {
	turn := setupTurn(15)
	if turn.IsMaxTurnReached() != false {
		t.Errorf("Expected isMaxTurnReached to be %v, but got %v", false, turn.IsMaxTurnReached())
	}
	for i := 0; i < int(turn.GetMaxTurn()) - 1; i++ {
		turn.NextTurn()
	}
	if turn.IsMaxTurnReached() != true {
		t.Errorf("Expected isMaxTurnReached to be %v, but got %v", true, turn.IsMaxTurnReached())
	}
}

func TestNextPlayerIndex(t *testing.T)  {
	turn := setupTurn(15)	
	if turn.GetCurrentPlayerIndex() != 1 {
		t.Errorf("Expected currentPlayerIndex to be %d, but got %d", 1, turn.GetCurrentPlayerIndex())
	}
	turn.NextPlayerIndex()
	if turn.GetCurrentPlayerIndex() != 2 {
		t.Errorf("Expected currentPlayerIndex to be %d, but got %d", 2, turn.GetCurrentPlayerIndex())
	}
}
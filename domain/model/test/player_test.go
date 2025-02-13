package test

import (
	"demon-dice-application/domain/model"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	player := model.NewPlayer(1, 1001, "TestPlayer", true, false)

	if player.GetId() != 1 {
		t.Errorf("expected ID to be %d, got %d", 1, player.GetId())
	}
	if player.GetUserId() != 1001 {
		t.Errorf("expected UserID to be %d, got %d", 1001, player.GetUserId())
	}
	if player.GetName() != "TestPlayer" {
		t.Errorf("expected Name to be %s, got %s", "TestPlayer", player.GetName())
	}
	if player.GetIsConnected() != true {
		t.Errorf("expected IsConnected to be %v, got %v", true, player.GetIsOnBreak())
	}
	if player.GetIsOnBreak() != false {
		t.Errorf("expected IsOnBreak to be %v, got %v", false, player.GetIsConnected())
	}
}
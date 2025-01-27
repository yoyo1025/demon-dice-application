package test

import (
	"demon-dice-application/domain/model"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	id := int64(1)
	userId := int64(1001)
	name := "TestPlayer"
	isConnected := true

	player := model.NewPlayer(id, userId, name, isConnected)

	if player.ID != id {
		t.Errorf("expected ID to be %d, got %d", id, player.ID)
	}
	if player.UserID != userId {
		t.Errorf("expected UserID to be %d, got %d", userId, player.UserID)
	}
	if player.Name != name {
		t.Errorf("expected Name to be %s, got %s", name, player.Name)
	}
	if player.IsConnected != isConnected {
		t.Errorf("expected IsConnected to be %v, got %v", isConnected, player.IsConnected)
	}
	if player.IsOnBreak != false {
		t.Errorf("expected IsOnBreak to be false, got %v", player.IsOnBreak)
	}
}

func TestPlayerSettersAndGetters(t *testing.T) {
	player := model.NewPlayer(1, 1001, "InitialPlayer", true)

	player.ID = 2
	player.UserID = 2002
	player.Name = "UpdatedPlayer"
	player.IsConnected = false
	player.IsOnBreak = true

	if player.ID != 2 {
		t.Errorf("expected ID to be %d, got %d", 2, player.ID)
	}
	if player.UserID != 2002 {
		t.Errorf("expected UserID to be %d, got %d", 2002, player.UserID)
	}
	if player.Name != "UpdatedPlayer" {
		t.Errorf("expected Name to be %s, got %s", "UpdatedPlayer", player.Name)
	}
	if player.IsConnected != false {
		t.Errorf("expected IsConnected to be false, got %v", player.IsConnected)
	}
	if player.IsOnBreak != true {
		t.Errorf("expected IsOnBreak to be true, got %v", player.IsOnBreak)
	}
}
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

	if player.GetId() != id {
		t.Errorf("expected ID to be %d, got %d", id, player.GetId())
	}
	if player.GetUserId() != userId {
		t.Errorf("expected UserID to be %d, got %d", userId, player.GetUserId())
	}
	if player.GetName() != name {
		t.Errorf("expected Name to be %s, got %s", name, player.GetName())
	}
	if player.GetIsConnected() != isConnected {
		t.Errorf("expected IsConnected to be %v, got %v", isConnected, player.GetIsOnBreak())
	}
	if player.GetIsOnBreak() != false {
		t.Errorf("expected IsOnBreak to be false, got %v", player.GetIsConnected())
	}
}
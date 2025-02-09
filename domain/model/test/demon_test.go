package test

import (
	"testing"
)

func TestNewDemon(t *testing.T) {
	demon := setupDemon()

	if demon.GetId() != 2 {
		t.Errorf("Expected id to be %d, but got %d", 2, demon.GetId())
	}
	if demon.GetUserId() != 1002 {
		t.Errorf("Expected userId to be %d, but got %d", 1002, demon.GetUserId())
	}
	if demon.GetName() != "TestDemon" {
		t.Errorf("Expected name to be %s, but got %s", "TestDemon", demon.GetName())
	}
	if demon.GetIsConnected() != true {
		t.Errorf("Expected isConnected to be %v, but got %v", true, demon.GetIsConnected())
	}
	if demon.GetIsOnBreak() != false {
		t.Errorf("Expected isOnBreak to be %v, but got %v", false, demon.GetIsOnBreak())
	}
}

func TestCapture(t *testing.T) {
	demon := setupDemon()
	villager := setupVillager()
	demon.Capture(villager)
	if villager.GetIsAlive() != false {
		t.Errorf("Expected isAlive to be %v, but got %v", false, villager.GetIsAlive())
	}
}
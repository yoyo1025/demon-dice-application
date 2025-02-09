package test

import (
	"testing"
)

func TestNewVillager(t *testing.T) {
	villager := setupVillager()

	if villager.GetId() != 3 {
		t.Errorf("Expected id to be %d, but got %d", 3, villager.GetId())
	}
	if villager.GetUserId() != 1003 {
		t.Errorf("Expected userId to be %d, but got %d", 1002, villager.GetUserId())
	}
	if villager.GetName() != "TestVillager" {
		t.Errorf("Expected name to be %s, but got %s", "TestVillager", villager.GetName())
	}
	if villager.GetIsConnected() != true {
		t.Errorf("Expected isConnected to be %v, but got %v", true, villager.GetIsConnected())
	}
	if villager.GetIsOnBreak() != false {
		t.Errorf("Expected isOnBreak to be %v, but got %v", false, villager.GetIsOnBreak())
	}
	if villager.GetIsAlive() != true {
		t.Errorf("Expected isAlive to be %v, but got %v", true, villager.GetIsAlive())
	}
	if villager.GetRank() != 0 {
		t.Errorf("Expected points to be %d, but got %d", 0, villager.GetRank())
	}
	if villager.GetPoints() != 0 {
		t.Errorf("Expected points to be %d, but got %d", 0, villager.GetRank())
	}
}

func TestAddPoint(t *testing.T)  {
	villager := setupVillager()
	if villager.GetPoints() != 0 {
		t.Errorf("Expected points to be %d, but got %d", 0, villager.GetPoints())
	}
	villager.AddPoint()
	if villager.GetPoints() != 1 {
		t.Errorf("Expected points to be %d, but got %d", 1, villager.GetPoints())
	}
	villager.AddPoint()
	if villager.GetPoints() != 2 {
		t.Errorf("Expected points to be %d, but got %d", 2, villager.GetPoints())
	}
}

func TestCaptured(t *testing.T) {
	villager := setupVillager()
	if villager.GetIsAlive() != true {
		t.Errorf("Expected isAlive to be %v, but got %v", true, villager.GetIsAlive())
	}
	villager.Captured()
	if villager.GetIsAlive() != false {
		t.Errorf("Expected isAlive to be %v, but got %v", false, villager.GetIsAlive())
	}
}
package test

import (
	"reflect"
	"testing"
)

func TestNewPosition(t *testing.T)  {
	position := setupPosition(5, 5)
	tmpPosition := setupPosition(5, 5)
	if position.GetX() != 5 {
		t.Errorf("Expected x to be %d, but got %d", 5, position.GetX())
	}
	if position.GetY() != 5 {
		t.Errorf("Expected y to be %d, but got %d", 5, position.GetY())
	}
	// Position構造体の中身を比較
	if !reflect.DeepEqual(position, tmpPosition) {
		t.Errorf("Expected position to be %+v, but got %+v", *tmpPosition, *position.GetPosition())
	}
}

func TestChangePosition(t *testing.T)  {
	position := setupPosition(5, 5)
	tmpPosition := setupPosition(6, 6)
	position.ChangePosition(6, 6)
	if !reflect.DeepEqual(position, tmpPosition) {
		t.Errorf("Expected position to be %+v, but got %+v", *tmpPosition, *position.GetPosition())
	}
}
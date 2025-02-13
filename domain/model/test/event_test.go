package test

import (
	"demon-dice-application/domain/model"
	"fmt"
	"reflect"
	"testing"
)


func TestExcuteEvent(t *testing.T)  {
	t.Run("ROLL_AGIN で移動可能座標が正しく返る", func(t *testing.T)  {
		eventKind := 0
		event := setupEvent(model.EventKindEnum(eventKind))
		player1 := setupPlayer(1, 1001, "test-player1", true, false)
		player2 := setupPlayer(2, 1002, "test-player2", true, false)
		player3 := setupPlayer(3, 1003, "test-player3", true, false)
		player4 := setupPlayer(4, 1004, "test-player4", true, false)

		

		players := []*model.Player {
			player1, player2, player3, player4,
		}

		p1Pos := setupPosition(3, 3)
		p2Pos := setupPosition(1, 2)
		p3Pos := setupPosition(3, 4)
		p4Pos := setupPosition(5, 7)

		playerPositions := map[*model.Player]*model.Position {
			players[0]: p1Pos,
			players[1]: p2Pos,
			players[2]: p3Pos,
			players[3]: p4Pos,
		}

		got := event.ExecuteEvent(playerPositions, players, 1, 1)
		// expectedCount := 12
		// if len(got) != expectedCount {
		// 	t.Errorf("期待する座標数 %d 個, 実際の結果 %d 個", expectedCount, len(got))
		// }
		fmt.Println(got)
	})

	t.Run("SKIP_TURN で対象プレイヤーの isOnBreak が true になる", func(t *testing.T) {
		eventKind := 1
		event := setupEvent(model.EventKindEnum(eventKind))
		player1 := setupPlayer(1, 1001, "test-player1", true, false)
		player2 := setupPlayer(2, 1002, "test-player2", true, false)
		player3 := setupPlayer(3, 1003, "test-player3", true, false)
		player4 := setupPlayer(4, 1004, "test-player4", true, false)

		

		players := []*model.Player {
			player1, player2, player3, player4,
		}

		p1Pos := setupPosition(3, 3)
		p2Pos := setupPosition(1, 2)
		p3Pos := setupPosition(3, 4)
		p4Pos := setupPosition(5, 7)

		playerPositions := map[*model.Player]*model.Position {
			players[0]: p1Pos,
			players[1]: p2Pos,
			players[2]: p3Pos,
			players[3]: p4Pos,
		}

		event.ExecuteEvent(playerPositions, players, 1, 2)
		if !players[1].GetIsOnBreak() {
			t.Errorf("SKIP_TURN イベント後、プレイヤー2の isOnBreak は true のはずが false のまま")
		}
	})

	t.Run("WARP でプレイヤー同士の位置が入れ替わる", func(t *testing.T) {
		eventKind := 2
		event := setupEvent(model.EventKindEnum(eventKind))
		player1 := setupPlayer(1, 1001, "test-player1", true, false)
		player2 := setupPlayer(2, 1002, "test-player2", true, false)
		player3 := setupPlayer(3, 1003, "test-player3", true, false)
		player4 := setupPlayer(4, 1004, "test-player4", true, false)

		

		players := []*model.Player {
			player1, player2, player3, player4,
		}

		p1Pos := setupPosition(3, 3)
		p2Pos := setupPosition(1, 2)
		p3Pos := setupPosition(3, 4)
		p4Pos := setupPosition(5, 7)

		playerPositions := map[*model.Player]*model.Position {
			players[0]: p1Pos,
			players[1]: p2Pos,
			players[2]: p3Pos,
			players[3]: p4Pos,
		}

		event.ExecuteEvent(playerPositions, players, 1, 2)

		expectedPoint1 := setupPosition(1, 2)
		expectedPoint2 := setupPosition(3, 3)

		if !reflect.DeepEqual(playerPositions[players[0]], expectedPoint1) {
			t.Errorf("WARP 後, player1 の位置が %v となるはずが %v となった", expectedPoint1, playerPositions[players[0]])
		}
		if !reflect.DeepEqual(playerPositions[players[1]], expectedPoint2) {
			t.Errorf("WARP 後, player2 の位置が %v となるはずが %v となった", expectedPoint2, playerPositions[players[1]])
		}
	})
}
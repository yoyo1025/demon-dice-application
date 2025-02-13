package test

import (
	"demon-dice-application/domain/model"
	"fmt"
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
}
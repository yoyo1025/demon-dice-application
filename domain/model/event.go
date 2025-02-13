package model

import (
	"fmt"
	"log"
)


type Event struct {
	kind EventKindEnum
	dice Dice
}

func NewEvent(kind EventKindEnum) *Event {
	return &Event{
		kind: kind,
		dice: *NewDice(),
	}
}

func (e *Event) GetKind() EventKindEnum {
	return e.kind
}

func (e *Event) SetKind(kind EventKindEnum) {
	e.kind = kind
}

// ExecuteEvent は、ROLL_AGAINの場合は移動可能なPositionスライスを返し、
// SKIP_TURNやWARPの場合は状態変更処理を行い、空のスライスを返します。
func (e *Event) ExecuteEvent(playerPositions map[*Player]*Position, players []*Player, eventPlayerId, targetPlayerId int) []Position {
	// IDを0ベースに変換する
	eventPlayerIndex := eventPlayerId - 1
	targetPlayerIndex := targetPlayerId - 1

	if e.kind == roll_again {
		return e.ExecuteRollAgain(playerPositions, players, eventPlayerIndex)
	}

	// ROLL_AGAIN以外は状態変更処理を行う
	switch e.kind {
	case skip_turn:
		e.ExecuteSkipTurn(players, targetPlayerIndex)
	case warp:
		e.ExecuteWarp(playerPositions, players, eventPlayerIndex, targetPlayerIndex)
	default:
		log.Fatalf("Unknown EventKind: %v", e.kind)
	}

	// SKIP_TURN, WARPの場合は空のスライスを返す
	return []Position{}
}

// ExecuteRollAgain は、サイコロの結果に応じて移動可能な座標を計算して返します。
func (e *Event) ExecuteRollAgain(playerPositions map[*Player]*Position, players []*Player, eventPlayerIndex int) []Position {
	// サイコロを振る
	diceResult := e.dice.Roll()

	// イベント発生プレイヤーの取得
	if eventPlayerIndex < 0 || eventPlayerIndex >= len(players) {
		panic("Invalid eventPlayerIndex")
	}
	eventPlayer := players[eventPlayerIndex]

	// 現在の位置を取得
	currentPositionPtr, ok := playerPositions[eventPlayer]
	if !ok || currentPositionPtr == nil {
		panic(fmt.Sprintf("Position not found for player with ID: %d", eventPlayer.GetId()))
	}
	currentPosition := *currentPositionPtr

	// マップから全プレイヤーの座標（occupiedPositions）を作成
	occupiedPositions := []Position{}
	for _, pos := range playerPositions {
		if pos != nil {
			occupiedPositions = append(occupiedPositions, *pos)
		}
	}

	possiblePositionsSet := make(map[Position]struct{})
	fieldSize := 8 // フィールドサイズ（例：8×8）

	// サイコロの出目に応じた移動範囲を計算
	for dx := 0; dx <= int(diceResult); dx++ {
		for dy := 0; dy <= int(diceResult)-dx; dy++ {
			// 現在位置は除外
			if dx == 0 && dy == 0 {
				continue
			}

			x := int(currentPosition.GetX())
			y := int(currentPosition.GetY())

			// 第1象限: (x+dx, y+dy)
			addPositionIfValid(possiblePositionsSet, occupiedPositions, x+dx, y+dy, fieldSize)
			// 第2象限: dx != 0 の場合 (x-dx, y+dy)
			if dx != 0 {
				addPositionIfValid(possiblePositionsSet, occupiedPositions, x-dx, y+dy, fieldSize)
			}
			// 第3象限: dx, dy 共に0でない場合 (x-dx, y-dy)
			if dx != 0 && dy != 0 {
				addPositionIfValid(possiblePositionsSet, occupiedPositions, x-dx, y-dy, fieldSize)
			}
			// 第4象限: dy != 0 の場合 (x+dx, y-dy)
			if dy != 0 {
				addPositionIfValid(possiblePositionsSet, occupiedPositions, x+dx, y-dy, fieldSize)
			}
		}
	}

	// マップからスライスに変換して返す
	possiblePositions := []Position{}
	for pos := range possiblePositionsSet {
		possiblePositions = append(possiblePositions, pos)
	}
	return possiblePositions
}

// addPositionIfValid は、指定された座標 (x, y) がフィールド内にあり、かつ他のプレイヤーが占有していなければ
// possiblePositions に追加します。
func addPositionIfValid(possiblePositions map[Position]struct{}, occupiedPositions []Position, x, y, fieldSize int) {
	if x >= 0 && x < fieldSize && y >= 0 && y < fieldSize {
		newPos := Position{
			x: int64(x),
			y: int64(y),
		}
		if !containsPosition(occupiedPositions, newPos) {
			possiblePositions[newPos] = struct{}{}
		}
	}
}

// containsPosition は、スライス内に同じ座標が存在するかをチェックします。
func containsPosition(positions []Position, pos Position) bool {
	for _, p := range positions {
		if p == pos { // Position型はフィールドが同じであれば等価とする
			return true
		}
	}
	return false
}

// ExecuteSkipTurn は、対象プレイヤーの IsOnBreak フラグを true にします。
func (e *Event) ExecuteSkipTurn(players []*Player, targetPlayerIndex int) {
	if targetPlayerIndex < 0 || targetPlayerIndex >= len(players) {
		panic("Invalid targetPlayerIndex")
	}
	players[targetPlayerIndex].isOnBreak = true
}

// ExecuteWarp は、イベント発生プレイヤーと対象プレイヤーの位置を入れ替えます。
func (e *Event) ExecuteWarp(playerPositions map[*Player]*Position, players []*Player, eventPlayerIndex, targetPlayerIndex int) {
	if eventPlayerIndex < 0 || eventPlayerIndex >= len(players) || targetPlayerIndex < 0 || targetPlayerIndex >= len(players) {
		panic("Invalid player index")
	}
	eventPlayer := players[eventPlayerIndex]
	targetPlayer := players[targetPlayerIndex]

	eventPos, ok1 := playerPositions[eventPlayer]
	targetPos, ok2 := playerPositions[targetPlayer]
	if !ok1 || !ok2 || eventPos == nil || targetPos == nil {
		panic("Position not found for one or both players")
	}

	// 位置の入れ替え
	playerPositions[eventPlayer], playerPositions[targetPlayer] = targetPos, eventPos
}
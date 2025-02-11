package model

type Turn struct {
	maxTurn int64
	currentTurn int64
	currentPlayerIndex int64
	maxPlayerIndex int64 
}

func NewTurn(maxTurn int64) *Turn {
	return &Turn{
		maxTurn: maxTurn,
		currentTurn: 1,
		currentPlayerIndex: 1,
		maxPlayerIndex: 4,
	}
}

func (t *Turn) GetMaxTurn() int64 {
	return t.maxTurn
}

func (t *Turn) GetCurrentTurn() int64 {
	return t.currentTurn
}

func (t *Turn) GetCurrentPlayerIndex() int64  {
	return t.currentPlayerIndex
}

func (t *Turn) GetMaxPlayerIndex() int64 {
	return t.maxPlayerIndex
}

func (t *Turn) NextTurn() {
	t.currentTurn += 1
}

func (t *Turn) IsMaxTurnReached() bool {
	if t.currentTurn == t.maxTurn {
		return true
	} else {
		return false
	}
}

func (t *Turn) NextPlayerIndex() {
	t.currentPlayerIndex += 1
	if t.currentPlayerIndex > 4 {
		t.currentPlayerIndex = 1
	}
}

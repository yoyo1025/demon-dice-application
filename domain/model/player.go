package model

type Player struct {
	ID          int64
	UserID      int64
	Name        string
	IsConnected bool
	IsOnBreak   bool
}

func NewPlayer(id, userId int64, name string, isConnected bool) *Player {
	return &Player{
		ID:          id,
		UserID:      userId,
		Name:        name,
		IsConnected: isConnected,
		IsOnBreak:   false,
	}
}
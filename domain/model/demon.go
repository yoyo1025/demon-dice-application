package model

type Demon struct {
	Player
}

func NewDemon(id, userId int64, name string, isConnected bool) *Demon {
	return &Demon{
		Player: Player{
			ID: id,
			UserID: userId,
			Name: name,
			IsConnected: isConnected,
			IsOnBreak: false,
		},
	}
}

func Capture(v *Villager) {
	v.IsAlive = false
}
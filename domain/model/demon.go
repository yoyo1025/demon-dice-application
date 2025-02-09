package model

type Demon struct {
	Player
}

func NewDemon(id, userId int64, name string, isConnected bool) *Demon {
	return &Demon{
		Player: Player{
			id: id,
			userId: userId,
			name: name,
			isConnected: isConnected,
			isOnBreak: false,
		},
	}
}

func (d *Demon) Capture(v *Villager) {
	v.Captured()
}
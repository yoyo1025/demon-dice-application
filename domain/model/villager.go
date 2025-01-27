package model

type Villager struct {
	Player
	Points int64
	IsAlive bool
	Rank int64
}

func NewVillager(id, userId int64, name string, isConnected bool) (*Villager){
	return &Villager{
		Player: Player{
			ID: id,
			UserID: userId,
			Name: name,
			IsConnected: isConnected,
			IsOnBreak: false,
		},
		Points: 0,
		IsAlive: true,
		Rank: 0,
	}
}

func (v *Villager) AddPoint() {
	v.Points += 1
}
package model

type Villager struct {
	Player
	points int64
	isAlive bool
	rank int64
}

func NewVillager(id, userId int64, name string, isConnected bool) (*Villager){
	return &Villager{
		Player: Player{
			id: id,
			userId: userId,
			name: name,
			isConnected: isConnected,
			isOnBreak: false,
		},
		points: 0,
		isAlive: true,
		rank: 0,
	}
}

func (v *Villager) AddPoint() {
	v.points += 1
}

func (v *Villager) Captured() {
	v.isAlive = false
}
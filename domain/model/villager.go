package model

type Villager struct {
	Player
	points int64
	isAlive bool
	rank int64
}

func NewVillager(id, userId int64, name string, isConnected, isOnBreak bool) (*Villager){
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

func (v *Villager) GetIsAlive() bool {
	return v.isAlive
}

func (v *Villager) GetRank() int64 {
	return v.rank
}

func (v *Villager) GetPoints() int64 {
	return v.points
}

func (v *Villager) AddPoint() {
	v.points += 1
}

func (v *Villager) Captured() {
	v.isAlive = false
}
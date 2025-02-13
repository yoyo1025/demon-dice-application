package model

type Player struct {
	id          int64
	userId      int64
	name        string
	isConnected bool
	isOnBreak   bool
}

func NewPlayer(id, userId int64, name string, isConnected, isOnBreak bool) *Player {
	return &Player{
		id:          id,
		userId:      userId,
		name:        name,
		isConnected: isConnected,
		isOnBreak:   isOnBreak,
	}
}

func (p *Player) GetId() int64 {
	return p.id;
}

func (p *Player) GetUserId() int64 {
	return p.userId;
}

func (p *Player) GetName() string {
	return p.name;
}

func (p *Player) GetIsConnected() bool {
	return p.isConnected;
}

func (p *Player) GetIsOnBreak() bool {
	return p.isOnBreak;
}

func (p *Player) SetIsConnected(isConnected bool) {
	p.isConnected = isConnected
}

func (p *Player) SetIsOnBreak(isOnBreak bool) {
	p.isOnBreak = isOnBreak
}
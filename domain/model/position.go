package model

type Position struct {
	x int64
	y int64
}

func NewPosition(x, y int64) *Position {
	return &Position{
		x: x,
		y: y,
	}
}

func (p *Position) GetX() int64{
	return p.x
}

func (p *Position) GetY() int64 {
	return p.y
}

func (p *Position) GetPosition() *Position {
	return p
}

func (p *Position) ChangePosition(x, y int64) {
	p.x = x
	p.y = y
}
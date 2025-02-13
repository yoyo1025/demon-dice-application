package model

type EventKindEnum int64

const (
	roll_again EventKindEnum = iota
	skip_turn
	warp
)

package model

import (
	"crypto/rand"
	"math/big"
)

type Dice struct {
	numberOfFace int64
}

func NewDice() *Dice {
	return &Dice {
		numberOfFace: 3,
	}
}

func (d *Dice) GetNumberOfFace() int64 {
	return d.numberOfFace
}

func (d *Dice) Roll() int64 {
	result, err := rand.Int(rand.Reader, big.NewInt(d.numberOfFace))
	if err != nil {
		panic(err)
	}
	return (result.Int64() + 1)
}

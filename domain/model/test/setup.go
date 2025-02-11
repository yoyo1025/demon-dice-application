package test

import "demon-dice-application/domain/model"

func setupDemon() *model.Demon {
	return model.NewDemon(2, 1002, "TestDemon", true, false)
}

func setupVillager() *model.Villager {
	return model.NewVillager(3, 1003, "TestVillager", true, false)
}

func setupDice() *model.Dice {
	return model.NewDice()
}

func setupTurn(maxTurn int64) *model.Turn  {
	return model.NewTurn(maxTurn)
}

func setupPosition(x, y int64) *model.Position {
	return model.NewPosition(x, y)
}
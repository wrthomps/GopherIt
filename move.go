package main

type Move struct {
	X int
	Y int
}

func NewMove(x, y int) *Move {
	move := new(Move)
	move.X = x
	move.Y = y

	return move
}

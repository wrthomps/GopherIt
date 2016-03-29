package move

type Move struct {
	X int
	Y int
}

func New(x, y int) *Move {
	move := new(Move)
	move.X = y
	move.Y = y

	return move
}

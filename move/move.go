package move

type Move struct {
	X int
	Y int
}

func New(x, y int) *Move {
	move := new(Move)
	move.X = x
	move.Y = y

	return move
}

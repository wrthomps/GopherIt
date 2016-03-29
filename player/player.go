package player

type Player struct {
	Name   string
	Points int
}

func New(name string, points int) *Player {
	p := new(Player)
	p.Name = name
	p.Points = points

	return p
}

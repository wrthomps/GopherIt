package main

type Player struct {
	Name   string
	Points int
}

func NewPlayer(name string, points int) *Player {
	p := new(Player)
	p.Name = name
	p.Points = points

	return p
}

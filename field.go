package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SEMI_COLON = ";"
	COMMA      = ","

	REPLACE_ALL = -1
)

type Field struct {
	MyId       int
	OpponentId int
	Rows       int
	Cols       int
	field      [][]int
}

func (f *Field) Init() error {
	if f.Rows <= 0 || f.Cols <= 0 {
		return fmt.Errorf("Invalid settings for field, cannot initialize")
	}

	f.field = make([][]int, f.Cols)
	for col := range f.field {
		f.field[col] = make([]int, f.Rows)
	}

	f.clearField()
	return nil
}

func (f *Field) ParseFromString(s string) {
	replaced := strings.Replace(s, SEMI_COLON, COMMA, REPLACE_ALL)
	r := strings.Split(replaced, COMMA)
	counter := 0
	for y := 0; y < f.Rows; y++ {
		for x := 0; x < f.Cols; x++ {
			f.field[x][y], _ = strconv.Atoi(r[counter])
			counter++
		}
	}
}

func (f *Field) clearField() {
	for y := 0; y < f.Rows; y++ {
		for x := 0; x < f.Cols; x++ {
			f.field[x][y] = 0
		}
	}
}

func (f *Field) AvailableMoves() []*Move {
	moves := make([]*Move, 0)
	for y := 0; y < f.Rows; y++ {
		for x := 0; x < f.Cols; x++ {
			if f.isEmptyPoint(x, y) && !f.isSuicideMove(x, y) {
				moves = append(moves, NewMove(x, y))
			}
		}
	}

	return moves
}

func (f *Field) isEmptyPoint(x, y int) bool {
	return f.field[x][y] == 0
}

func (f *Field) isSuicideMove(x, y int) bool {
	mark := make([][]bool, f.Rows)
	for tx := 0; tx < len(mark); tx++ {
		mark[tx] = make([]bool, f.Cols)
	}

	f.field[x][y] = f.MyId
	liberties := f.flood(mark, x, y, f.MyId, 0)
	f.field[x][y] = 0

	return liberties <= 0
}

func (f *Field) flood(mark [][]bool, x, y, player, stackCounter int) int {
	if x < 0 {
		return 0
	}
	if y < 0 {
		return 0
	}
	if x >= f.Rows {
		return 0
	}
	if y >= f.Cols {
		return 0
	}
	if mark[x][y] {
		return 0
	}

	if f.field[x][y] != player {
		if f.field[x][y] == 0 {
			return 1
		}
		return 0
	}

	mark[x][y] = true

	neighborLiberties := 0
	if stackCounter < f.Rows*f.Cols {
		neighborLiberties += f.flood(mark, x-1, y, player, stackCounter+1)
		neighborLiberties += f.flood(mark, x+1, y, player, stackCounter+1)
		neighborLiberties += f.flood(mark, x, y-1, player, stackCounter+1)
		neighborLiberties += f.flood(mark, x, y+1, player, stackCounter+1)
	}
	return neighborLiberties
}

func (f *Field) String() string {
	s := ""
	counter := 0
	for y := 0; y < f.Rows; y++ {
		for x := 0; x < f.Cols; x++ {
			if counter > 0 {
				s += COMMA
			}
			s += string(f.field[x][y])
			counter++
		}
	}

	return s
}

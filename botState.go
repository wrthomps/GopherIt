package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PLAYER_DELIMITER = ","
	MAX_PLAYERS      = 2
)

// Bot commands
const (
	TIMEBANK      = "timebank"
	TIME_PER_MOVE = "time_per_move"
	PLAYER_NAMES  = "player_names"
	YOUR_BOT      = "your_bot"
	YOUR_BOTID    = "your_botid"
	FIELD_WIDTH   = "field_width"
	FIELD_HEIGHT  = "field_height"

	ROUND = "round"
	FIELD = "field"

	POINTS = "points"
)

type BotState struct {
	maxTimeBank int
	timePerMove int
	roundNumber int
	moveNumber  int
	Timebank    int
	myName      string
	players     map[string]*Player
	field       *Field
	history     map[uint32]bool
}

func (bs *BotState) Init() {
	bs.players = make(map[string]*Player, 0)
	bs.history = make(map[uint32]bool)
	bs.field = new(Field)
	bs.field.Init()
	bs.history[bs.field.zhash] = true
}

func (bs *BotState) ParseSettings(key, value string) error {
	switch key {
	case TIMEBANK:
		time, _ := strconv.Atoi(value)
		bs.maxTimeBank = time
		bs.Timebank = time
		break
	case TIME_PER_MOVE:
		bs.timePerMove, _ = strconv.Atoi(value)
		break
	case PLAYER_NAMES:
		playerNames := strings.Split(value, PLAYER_DELIMITER)
		for _, playerName := range playerNames {
			bs.players[playerName] = NewPlayer(playerName, 0)
		}
		break
	case YOUR_BOT:
		bs.myName = value
		break
	case YOUR_BOTID:
		myId, _ := strconv.Atoi(value)
		opponentId := MAX_PLAYERS - myId + 1
		bs.field.MyId = myId
		bs.field.OpponentId = opponentId
		break
	case FIELD_WIDTH:
		bs.field.Cols, _ = strconv.Atoi(value)
		break
	case FIELD_HEIGHT:
		bs.field.Rows, _ = strconv.Atoi(value)
		break
	default:
		return fmt.Errorf("Cannot parse settings input with key '%s'", key)
	}

	return nil
}

func (bs *BotState) ParseGameData(key, value string) error {
	switch key {
	case ROUND:
		bs.roundNumber, _ = strconv.Atoi(value)
		break
	case MOVE:
		bs.moveNumber, _ = strconv.Atoi(value)
		break
	case FIELD:
		bs.field.Init()
		bs.field.ParseFromString(value)
	default:
		return fmt.Errorf("Cannot parse game data input with key '%s'", key)
	}

	return nil
}

func (bs *BotState) ParsePlayerData(playerName, key, value string) error {
	switch key {
	case POINTS:
		bs.players[playerName].Points, _ = strconv.Atoi(value)
		break
	default:
		return fmt.Errorf("Cannot parse %s data input with key '%s'", playerName, key)
	}

	return nil
}

func (bs *BotState) AvailableMoves() []*Move {
	var movesWithKo []*Move = bs.field.AvailableMoves()
	var moves []*Move = make([]*Move, 0)
	for _, move := range movesWithKo {
		var currentHashForMove uint32 = bs.field.HashAt(move.X, move.Y)
		_, has := bs.history[bs.field.zhash^currentHashForMove^zhashes[move.X][move.Y][1]]
		if !has {
			moves = append(moves, move)
		}
	}
	return bs.field.AvailableMoves()
}

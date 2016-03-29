package bot

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	SPACE_DELIMITER = " "

	SETTINGS   = "settings"
	UPDATE     = "update"
	GAME       = "game"
	ACTION     = "action"
	MOVE       = "move"
	PLACE_MOVE = "place_move"
	PASS       = "pass"
)

type BotParser struct {
	scanner      *bufio.Scanner
	bot          *BotStarter
	currentState *BotState
}

func (bp *BotParser) Init(bot *BotStarter) {
	bp.scanner = bufio.NewScanner(os.Stdin)
	bp.bot = bot
	bp.currentState = new(BotState)
	bp.currentState.Init()
}

func (bp *BotParser) Run() {
	for bp.scanner.Scan() {
		line := bp.scanner.Text()

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, SPACE_DELIMITER)
		if parts[0] == SETTINGS {
			bp.currentState.ParseSettings(parts[1], parts[2])
		} else if parts[0] == UPDATE {
			if parts[1] == GAME {
				bp.currentState.ParseGameData(parts[2], parts[3])
			} else {
				bp.currentState.ParsePlayerData(parts[1], parts[2], parts[3])
			}
		} else if parts[0] == ACTION {
			if parts[1] == MOVE {
				time, _ := strconv.Atoi(parts[2])
				move := bp.bot.GetMove(bp.currentState, time)
				if move != nil {
					fmt.Printf(PLACE_MOVE+"%s"+SPACE_DELIMITER+"%s\n", move.X, move.Y)
				} else {
					fmt.Println(PASS + "\n")
				}
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

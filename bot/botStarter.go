package bot

import (
	"github.com/wrthomps/GopherIt/move"
	"math/rand"
)

type BotStarter struct{}

func (bs *BotStarter) GetMove(state *BotState, timeout int) *move.Move {
	state.Timebank = timeout

	moves := state.AvailableMoves()
	moveCount := len(moves)

	if moveCount <= 0 {
		return nil
	} else {
		return moves[rand.Intn(moveCount)]
	}
}

package main

import (
	"github.com/wrthomps/GopherIt/bot"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	starter := new(bot.BotStarter)
	parser := new(bot.BotParser)
	parser.Init(starter)
	parser.Run()
}

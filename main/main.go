package main

import (
	"github.com/wrthomps/GopherIt/bot"
)

func main() {
	starter := new(bot.BotStarter)
	parser := new(bot.BotParser)
	parser.Init(starter)
	parser.Run()
}

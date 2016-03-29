package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	starter := new(BotStarter)
	parser := new(BotParser)
	parser.Init(starter)
	parser.Run()
}

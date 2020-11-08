package main

import (
	"fmt"
	"greedyGamePalash/processors/requestProcessors"
	"greedyGamePalash/server"
	"greedyGamePalash/state"
)

func main() {
	fmt.Print("hello")
	state := state.GetInitialState()
	requestProcessors.InitializeRequestProcessors(state)
	server.Start()
}

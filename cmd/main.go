package main

import (
	"github.com/MatteoZacca/Rock-Paper-Scissors/internal/game"
)

const (
	FirstPlayerId  = 1
	SecondPlayerId = 2
)

func main() {
	p1 := game.NewPlayer(FirstPlayerId)
	p2 := game.NewPlayer(SecondPlayerId)

	go p1.Play()
	go p2.Play()

	ref := game.NewReferee(p1, p2)
	ref.StartMatch()
}

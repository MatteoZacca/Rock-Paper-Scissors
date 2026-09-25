package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	FirstPlayerId  = 1
	SecondPlayerId = 2
)

type Move int

const (
	Rock     Move = iota // 0
	Paper                // 1
	Scissors             // 2
)

type Player struct {
	id   int
	move Move
}

func translateMove(move int) Move {
	switch move {
	case 0:
		return Rock
	case 1:
		return Paper
	case 2:
		return Scissors
	}

	return Rock
}

func playerMove(player Player, ref chan Move) {
	move := rand.Intn(3)

	// func for mapping int to Move
	player.move = translateMove(move)

	ref <- player.move
}

func main() {
	ref := make(chan Move, 2)
	firstPlayer := Player{
		id: FirstPlayerId,
	}
	secondPlayer := Player{
		id: SecondPlayerId,
	}

	var wg sync.WaitGroup

	for {
		wg.Add(2)

		go func() {
			defer wg.Done()
			playerMove(firstPlayer, ref)
			firstPlayer.move = <-ref
		}()

		go func() {
			defer wg.Done()
			playerMove(secondPlayer, ref)
			secondPlayer.move = <-ref
		}()

		if firstPlayer.move == secondPlayer.move {
			fmt.Println("It's a tie!")
		} else if (firstPlayer.move == Rock && secondPlayer.move == Scissors) ||
			(firstPlayer.move == Paper && secondPlayer.move == Rock) ||
			(firstPlayer.move == Scissors && secondPlayer.move == Paper) {
			fmt.Printf("Player %d wins!\n", firstPlayer.id)
		} else {
			fmt.Printf("Player %d wins!\n", secondPlayer.id)
		}

		wg.Wait()
	}

}

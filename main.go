package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
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
	id    int
	score int
}

var availableMoves = [3]string{"Rock", "Paper", "Scissors"}

func (m Move) String() string {
	return availableMoves[m]
}

func playerRoutine(p Player, moveChan chan<- Move, resChan <-chan int, wg *sync.WaitGroup) {

	for {
		move := Move(rand.IntN(3))

		moveChan <- move

		result := <-resChan

		switch result {
		case 1:
			p.score++
			fmt.Printf("Player %d WON with %s! (Score: %d)\n", p.id, move, p.score)
		case -1:
			fmt.Printf("Player %d LOST with %s... (Score: %d)\n", p.id, move, p.score)
		default:
			fmt.Printf("Player %d TIED with %s. (Score: %d)\n", p.id, move, p.score)
		}

		wg.Done()
	}
}

func main() {
	p1MoveChan, p1ResChan := make(chan Move), make(chan int)
	p2MoveChan, p2ResChan := make(chan Move), make(chan int)

	firstPlayer := Player{id: FirstPlayerId}
	secondPlayer := Player{id: SecondPlayerId}

	var wg sync.WaitGroup

	go playerRoutine(firstPlayer, p1MoveChan, p1ResChan, &wg)
	go playerRoutine(secondPlayer, p2MoveChan, p2ResChan, &wg)

	turn := 1

	for {
		fmt.Printf("\n--- Turn %d ---\n", turn)

		wg.Add(2)

		p1Move := <-p1MoveChan
		p2Move := <-p2MoveChan

		outcome := (p1Move - p2Move + 3) % 3

		switch outcome {
		case 0:
			p1ResChan <- 0
			p2ResChan <- 0
		case 1:
			p1ResChan <- 1
			p2ResChan <- 0
		default: // outcome == 2 means Player 2 wins
			p1ResChan <- 0
			p2ResChan <- 1
		}

		wg.Wait()

		time.Sleep(1 * time.Second) // Pause so the terminal is readable
		turn++
	}

}

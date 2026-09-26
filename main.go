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
	id       int
	score    int
	moveChan chan Move
	resChan  chan int
}

var availableMoves = [3]string{"Rock", "Paper", "Scissors"}

func (m Move) String() string {
	return availableMoves[m]
}

func NewPlayer(id int) *Player {
	return &Player{
		id:       id,
		moveChan: make(chan Move),
		resChan:  make(chan int),
	}
}

func (p *Player) Play(wg *sync.WaitGroup) {
	for {
		move := Move(rand.IntN(3))

		p.moveChan <- move

		result := <-p.resChan

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
	p1 := NewPlayer(FirstPlayerId)
	p2 := NewPlayer(SecondPlayerId)

	var wg sync.WaitGroup

	go p1.Play(&wg)
	go p2.Play(&wg)

	turn := 1

	for {
		fmt.Printf("\n--- Turn %d ---\n", turn)

		wg.Add(2)

		p1Move := <-p1.moveChan
		p2Move := <-p2.moveChan

		outcome := (p1Move - p2Move + 3) % 3

		switch outcome {
		case 0:
			p1.resChan <- 0
			p2.resChan <- 0
		case 1:
			p1.resChan <- 1
			p2.resChan <- -1
		default: // outcome == 2 means Player 2 wins
			p1.resChan <- -1
			p2.resChan <- 1
		}

		// Wait for both players to finish printing their scores
		wg.Wait()

		time.Sleep(1 * time.Second) // Pause so the terminal is readable
		turn++
	}

}

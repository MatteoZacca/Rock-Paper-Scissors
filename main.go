package main

import (
	"fmt"
	"math/rand/v2"
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

var availableMoves = [3]string{"Rock", "Paper", "Scissors"}

func (m Move) String() string {
	return availableMoves[m]
}

type Player struct {
	id          int
	score       int
	askMoveChan chan struct{}
	moveChan    chan Move
	resChan     chan int
	doneChan    chan struct{}
}

func NewPlayer(id int) *Player {
	return &Player{
		id:          id,
		askMoveChan: make(chan struct{}),
		moveChan:    make(chan Move),
		resChan:     make(chan int),
		doneChan:    make(chan struct{}),
	}
}

func (p *Player) Play() {
	for {
		// Wait for the referee to ask for the move
		<-p.askMoveChan

		// Randomly select the move
		move := Move(rand.IntN(3))

		// Communicate the move to the referee
		p.moveChan <- move

		// Wait for the result
		result := <-p.resChan

		// Update score and print
		switch result {
		case 1:
			p.score++
			fmt.Printf("Player %d WON with %s! (Score: %d)\n", p.id, move, p.score)
		case -1:
			fmt.Printf("Player %d LOST with %s... (Score: %d)\n", p.id, move, p.score)
		default:
			fmt.Printf("Player %d TIED with %s. (Score: %d)\n", p.id, move, p.score)
		}

		p.doneChan <- struct{}{}
	}
}

func main() {
	p1 := NewPlayer(FirstPlayerId)
	p2 := NewPlayer(SecondPlayerId)

	go p1.Play()
	go p2.Play()

	turn := 1

	for {
		fmt.Printf("\n--- Turn %d ---\n", turn)

		// Referee asks both players for moves
		p1.askMoveChan <- struct{}{}
		p2.askMoveChan <- struct{}{}

		// Referee receives the moves
		p1Move := <-p1.moveChan
		p2Move := <-p2.moveChan

		fmt.Printf("Referee: Player 1 played %s | Player 2 played %s\n", p1Move, p2Move)

		// Referee calculates the winner
		outcome := (p1Move - p2Move + 3) % 3

		// Referee communicates outcomes
		switch outcome {
		case 0:
			fmt.Println("Referee: It's a Tie!")
			p1.resChan <- 0
			p2.resChan <- 0
		case 1:
			fmt.Println("Referee: Player 1 takes the round!")
			p1.resChan <- 1
			p2.resChan <- -1
		default: // outcome == 2
			fmt.Println("Referee: Player 2 takes the round!")
			p1.resChan <- -1
			p2.resChan <- 1
		}

		<-p1.doneChan
		<-p2.doneChan

		time.Sleep(1 * time.Second) // Pause so the terminal is readable
		turn++
	}

}

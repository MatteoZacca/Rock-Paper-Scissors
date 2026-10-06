package game

import (
	"fmt"
	"math/rand/v2"
)

type Player struct {
	id          int
	score       int
	askMoveChan chan struct{} // Channel to signal the player to make a move
	moveChan    chan Move     // Channel to send the player's move to the referee
	resChan     chan Result   // Channel to receive the result of the round from the referee
	doneChan    chan struct{} // Channel to signal the referee that the player has finished processing the result
}

func NewPlayer(id int) *Player {
	return &Player{
		id:          id,
		askMoveChan: make(chan struct{}),
		moveChan:    make(chan Move),
		resChan:     make(chan Result),
		doneChan:    make(chan struct{}),
	}
}

func (p *Player) Play() {
	for {
		<-p.askMoveChan

		move := Move(rand.IntN(len(availableMoves)))
		p.moveChan <- move

		result := <-p.resChan

		switch result {
		case Win:
			p.score++
			fmt.Printf("Player %d: I WON with %s! (Score: %d)\n", p.id, move, p.score)
		case Lose:
			fmt.Printf("Player %d: I LOST with %s... (Score: %d)\n", p.id, move, p.score)
		case Draw:
			fmt.Printf("Player %d: DRAW with %s. (Score: %d)\n", p.id, move, p.score)
		}

		p.doneChan <- struct{}{}
	}
}

package game

import (
	"fmt"
	"math/rand/v2"
)

type Player struct {
	id          int
	score       int
	askMoveChan chan struct{}
	moveChan    chan Move
	resChan     chan Result
	doneChan    chan struct{}
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

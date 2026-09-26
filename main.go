package main

import (
	"fmt"
	"math/rand"
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

func (m Move) String() string {
	switch m {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	}
	return "Unknown"
}

func playerRoutine(p Player, moveChan chan<- Move, resChan <-chan int, wg *sync.WaitGroup) {

	for {
		move := Move(rand.Intn(3))

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

		if p1Move == p2Move {
			p1ResChan <- 0
			p2ResChan <- 0
		} else if (p1Move == Rock && p2Move == Scissors) ||
			(p1Move == Paper && p2Move == Rock) ||
			(p1Move == Scissors && p2Move == Paper) {

			p1ResChan <- 1
			p2ResChan <- 0
		} else {
			p1ResChan <- 0
			p2ResChan <- 1
		}

		wg.Wait()

		time.Sleep(1 * time.Second) // Pause so the terminal is readable
		turn++
	}

}

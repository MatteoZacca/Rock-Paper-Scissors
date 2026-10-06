package game

import (
	"fmt"
	"time"
)

type Referee struct {
	p1 *Player
	p2 *Player
}

func NewReferee(p1, p2 *Player) *Referee {
	return &Referee{
		p1: p1,
		p2: p2,
	}
}

func (r *Referee) StartMatch() {
	turn := 1

	for {
		fmt.Printf("\n--- Turn %d ---\n", turn)

		r.p1.askMoveChan <- struct{}{}
		r.p2.askMoveChan <- struct{}{}

		p1Move := <-r.p1.moveChan
		p2Move := <-r.p2.moveChan

		fmt.Printf("Referee: Player %d played %s | Player %d played %s\n", r.p1.id, p1Move, r.p2.id, p2Move)

		// Logic to determine the winner based on the moves
		outcome := (int(p1Move) - int(p2Move) + len(availableMoves)) % len(availableMoves)

		switch outcome {
		case 0:
			fmt.Println("Referee: It's a Tie!")
			r.p1.resChan <- Draw
			r.p2.resChan <- Draw
		case 1:
			fmt.Println("Referee: Player 1 takes the round!")
			r.p1.resChan <- Win
			r.p2.resChan <- Lose
		default: // 2
			fmt.Println("Referee: Player 2 takes the round!")
			r.p1.resChan <- Lose
			r.p2.resChan <- Win
		}

		<-r.p1.doneChan
		<-r.p2.doneChan

		time.Sleep(1 * time.Second) // Optional: Add a delay between turns for better readability
		turn++
	}
}

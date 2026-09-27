// This file holds the game domain language
package game

// Move logic
type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

var availableMoves = [3]string{"Rock", "Paper", "Scissors"}

func (m Move) String() string {
	return availableMoves[m]
}

// Result logic
type Result int

const (
	Draw Result = iota
	Win
	Lose
)

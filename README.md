# Rock-Paper-Scissors in Go

A simple, concurrent implementation of the classic Rock-Paper-Scissors game written in Go. 

This project serves as a practical demonstration of **synchronous message communication** in Go. It heavily leverages Goroutines and unbuffered Channels to coordinate state and actions between different concurrent entities without using explicit locks.

## How It Works

The game simulation consists of two main entities:
*   **Players (Goroutines):** Two independent players run concurrently. They wait for a signal to make a move, randomly select Rock, Paper, or Scissors, and wait to receive the outcome of the round.
*   **Referee:** The central coordinator. The referee signals both players to make a move, collects their choices, calculates the outcome based on the classic game rules, and sends the result back to each player.

### Synchronous Channel Communication

The coordination is entirely handled by unbuffered channels, ensuring strict synchronization during each turn:
1.  `askMoveChan`: Referee signals players that a new turn has started.
2.  `moveChan`: Players send their chosen `Move` to the referee.
3.  `resChan`: Referee sends the `Result` (Win, Lose, Draw) back to the players.
4.  `doneChan`: Players signal to the referee that they have logged their result and are ready for the next turn.

## Example Output

```text
--- Turn 1 ---
Referee: Player 1 played Scissors | Player 2 played Rock
Referee: Player 2 takes the round!
Player 1: I LOST with Scissors... (Score: 0)
Player 2: I WON with Rock! (Score: 1)

--- Turn 2 ---
Referee: Player 1 played Paper | Player 2 played Paper
Referee: It's a Tie!
Player 1: DRAW with Paper. (Score: 0)
Player 2: DRAW with Paper. (Score: 1)
```
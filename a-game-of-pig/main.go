package main

import (
	"fmt"
	"math/rand"
)

const winningScore = 100

func main() {

	var playerScores [2]int
	var currentPlayer int
	var currentTurnScore int

	fmt.Println("Welcome to the game of Pig!")

	for !isGameOver(playerScores) {
		fmt.Printf("\nPlayer %d's turn (current score: %d)\n", currentPlayer+1, playerScores[currentPlayer])
		roll := rollDie()

		if roll == 1 {
			fmt.Println("You rolled a 1. Turn over!")
			currentTurnScore = 0
			switchPlayer(&currentPlayer)
		} else {
			fmt.Printf("You rolled a %d. Current turn score: %d\n", roll, currentTurnScore)
			currentTurnScore += roll
			fmt.Print("Enter 'r' to roll again or 'h' to hold: ")
			var input string
			fmt.Scanln(&input)

			if input == "h" {
				playerScores[currentPlayer] += currentTurnScore
				currentTurnScore = 0
				switchPlayer(&currentPlayer)
			}
		}
	}

	winner := currentPlayer + 1
	fmt.Printf("\nPlayer %d wins with a score of %d!\n", winner, playerScores[currentPlayer])
}

func rollDie() int {
	return rand.Intn(6) + 1
}

func switchPlayer(currentPlayer *int) {
	*currentPlayer = 1 - *currentPlayer
}

func isGameOver(playerScores [2]int) bool {
	return playerScores[0] >= winningScore || playerScores[1] >= winningScore
}

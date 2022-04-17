package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurn = 4 // change the maxTurn if you wanna to 

// this a game where you have to enter a number and in line number 16 loop will run(you can change the number of repetation if you like ) and if your guess number = random number you won
func main() {
	rand.Seed(time.Now().UnixNano())
	guess, _ := strconv.Atoi(os.Args[1])
	for turn := 0; turn < maxTurn; turn++ {
		n := rand.Intn(guess + 1)
		if turn == 1 && n == guess {
			fmt.Println("first shot you won ")
			return
		}
		if n == guess {
			fmt.Println("You won the game")
			return
		}
	}
	fmt.Println("sorry better luck next time ")

}

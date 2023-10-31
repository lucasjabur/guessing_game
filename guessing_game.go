package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func archive(trys int) {

	file, _ := os.Create("highscore.txt")
	defer file.Close()
	dat, _ := os.ReadFile("highscore.txt")

	if dat != nil {
		data, _ := strconv.Atoi(string(dat))
		if data > trys {
			_, err3 := file.WriteString(fmt.Sprintf("%d", trys))
			if err3 != nil {
				fmt.Println("Error: ", err3)
				return
			}
		} else if data == 0 {
			_, err4 := file.WriteString(fmt.Sprintf("%d", trys))
			if err4 != nil {
				fmt.Println("Error: ", err4)
				return
			}
		} else if data < trys {
			fmt.Println("Uh... that's close. You almost beat the record!")
		}
	}
}

func game(numRandom int) {

	for trys := 1; trys > 0; trys++ {
		fmt.Printf("\nTry number: %d\n", trys)
		fmt.Printf("Take a guess: ")
		var input int

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		if input > numRandom {
			fmt.Println("The value you put in is higher!")
		} else if input < numRandom {
			fmt.Println("The value you put in is lower!")
		} else {
			fmt.Printf("\nCongratulations! You guessed it in %d trys!\n", trys)
			archive(trys)

			fmt.Println("\nWanna keep playing? ('1' to play again, '0' to give up)")
			var play_again int

			_, err2 := fmt.Scan(&play_again)

			if err2 != nil {
				fmt.Println("Error: ", err2)
				return
			}
			if play_again == 0 {
				os.Exit(0)
			} else {
				main()
			}

			break
		}
	}
}

func main() {
	var numRandom = rand.Intn(100)

	fmt.Println("\nWelcome to the Guessing Game!")
	fmt.Println("I'm thinking on a number from 1 to 100. I dare you to guess it!")

	game(numRandom)
}

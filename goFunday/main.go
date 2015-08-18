package main

import "fmt"
import "math/rand"

var mainAI [500]int
var secondAI [500]int
var thirdAI [500]int
var fourthAI [500]int
var fifthAI [500]int

var AIArrays [500][500]int

var board [9]int
var moves = 0

func main() {
	var done = false

	for i := 0; i < 500; i++ {
		mainAI[i] = rand.Intn(2)
		secondAI[i] = rand.Intn(2)
		thirdAI[i] = rand.Intn(2)
		fourthAI[i] = rand.Intn(2)
		fifthAI[i] = rand.Intn(2)
		for j := 0; j < 500; j++ {
			AIArrays[j][i] = rand.Intn(2)
		}
	}

	for !done {
		for i := 0; i < 9; i++ {
			board[i] = 0
		}
		fmt.Println("Enter a 1 to play tic tac toe, a 2 to train, or a 3 to quit")
		input := ""
		fmt.Scanln(&input)
		if input == "1" {
			PlayGame()
		} else if input == "2" {
			Train()
		} else if input == "3" {
			done = true
			fmt.Println("Bye!")
		} else {
			fmt.Println("Unknown input")
		}

	}

}

func getAIMove(ai [500]int, value int) {
	var val = 0

	var total = board[0]*1 + board[1]*2 + board[2]*3 + board[3]*10 + board[4]*17 + board[5]*23 + board[6]*47 + board[7]*53 + board[8]*61
	val = ai[total]*4 + ai[total+1]*2 + ai[total+2]

	for board[val] != 0 {
		val = (val + 1) % 9
	}

	board[val] = value
}

func isWinner() int {
	for i := 0; i < 3; i++ {
		if board[i+0] == board[i+3] && board[i+3] == board[i+6] && board[i+6] != 0 {
			return board[i]
		} else if board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] && board[i*3] != 0 {
			return board[i*3]
		}
	}

	if board[0] == board[4] && board[4] == board[8] && board[8] != 0 {
		return board[0]
	} else if board[2] == board[4] && board[4] == board[6] && board[6] != 0 {
		return board[2]
	}

	for i := 0; i < 9; i++ {
		if board[i] == 0 {
			return 0
		}
	}
	return -1
}

func PlayGame() {
	fmt.Println("Time to play tic tac toe")
	var winner = 0

	for winner == 0 {
		getAIMove(mainAI, 1)
		winner = isWinner()
		if winner == 0 {
			printBoard()
			fmt.Println("enter move 0-8")
			getUserInput()
			winner = isWinner()
		}
	}

	printBoard()
	fmt.Println("The winner is ", winner)
}

func getUserInput() {
	for {
		var input = 0
		fmt.Scanln(&input)
		if input >= 0 && input <= 8 && board[input] == 0 {
			board[input] = 2
			return
		} else {
			fmt.Println("Invalid move, try again")
		}
	}
}

func printBoard() {
	var val = 0
	fmt.Println()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val = board[i*3+j]

			if val == 0 {
				fmt.Print("-")
			} else if val == 1 {
				fmt.Print("X")
			} else if val == 2 {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func Train() {
	fmt.Println("Training time")
	var arrays [500][500]int
	var newArrays [500][500]int
	arrays = AIArrays

	for k := 0; k < 50; k++ {
		var scores [500]int

		for i := 0; i < 500; i++ {
			for j := 0; j < 500; j++ {
				outcome := playAIs(arrays[i], arrays[j])
				if outcome == -1 {
				} else if outcome == 1 {
					scores[i] += 1
				}
			}
		}
		var best = 0
		for i := 0; i < 500; i++ {
			if scores[i] > scores[best] {
				best = i
			}
		}

		newArrays[0] = arrays[best]

		var counter = 1
		i := 0

		fmt.Println(best)

		for counter < 20 {
			var bestScore = scores[best]
			if bestScore > 99 {
				bestScore = 99
			}
			if (rand.Intn(bestScore) * 2) < scores[i] {
				newArrays[counter] = arrays[i]
				counter++
			}
			i++
		}

		// "breeding" of the best and random string
		for counter < 100 {
			newArrays[counter] = spliceArrays(newArrays[rand.Intn(20)], newArrays[rand.Intn(20)])
			counter++
		}

		for counter < 200 {
			newArrays[counter] = newArrays[counter-100]
			newArrays[counter+100] = newArrays[counter]
			newArrays[counter+200] = newArrays[counter+100]
			newArrays[counter+300] = newArrays[counter+200]
			counter++
		}

		// mutation of all strings
		for i := 1; i < 500; i++ {
			for j := 0; j < 500; j++ {
				if rand.Intn(100) == 0 {
					newArrays[i][j] = (newArrays[i][j] + 1) % 2
				}
			}
		}

		arrays = newArrays
	}

	mainAI = arrays[0]
	AIArrays = arrays
}

func spliceArrays(a, b [500]int) [500]int {
	var splicePoint = rand.Intn(400)
	var newArray [500]int

	for i := 0; i < 500; i++ {
		if i < splicePoint {
			newArray[i] = a[i]
		} else {
			newArray[i] = b[i]
		}
	}
	return newArray
}

func playAIs(a, b [500]int) int {
	var winner = 0

	for i := 0; i < 9; i++ {
		board[i] = 0
	}

	for winner == 0 {
		getAIMove(a, 1)
		winner = isWinner()
		if winner == 0 {
			getAIMove(b, 2)
			winner = isWinner()
		}
	}
	return winner
}

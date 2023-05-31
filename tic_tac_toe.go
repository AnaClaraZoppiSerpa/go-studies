package main

import (
	"fmt"
)

func printGame(game [3][3]string) {
	fmt.Println("   0 1 2 ")
	for i, _ := range game {
		fmt.Println(i, game[i])
	}
}

func playerInput(playerSymbol string, game [3][3]string) [3][3]string {
	fmt.Println("Where to mark? Enter row and column.")
	var i,j int

	markedValid := false

	for markedValid == false {
		fmt.Scanf("%d %d", &i, &j)
		if i >= 3 || j >= 3 || i < 0 || j < 0 || game[i][j] != "-" {
			fmt.Println("This is in invalid position. Choose another.")
		} else {
			game[i][j] = playerSymbol
			markedValid = true
		}
	}
	return game
}

func win(playerSymbol string, game [3][3]string) bool {
	// Rows
	for row := 0; row < 3; row++ {
		same_symbol_on_row := true
		for iterate_on_row := 0; iterate_on_row < 3; iterate_on_row++ {
			if game[row][iterate_on_row] != playerSymbol {
				same_symbol_on_row = false
			}
		}
		if same_symbol_on_row {
			return true
		}
	}
	// Columns
	for col := 0; col < 3; col++ {
		same_symbol_on_col := true
		for iterate_on_col := 0; iterate_on_col < 3; iterate_on_col++ {
			if game[iterate_on_col][col] != playerSymbol {
				same_symbol_on_col = false
			}
		}
		if same_symbol_on_col {
			return true
		}
	}
	// Diagonal 1
	same_on_all := true
	for index := 0; index < 3; index++ {
		if game[index][index] != playerSymbol {
			same_on_all = false
		}
	}
	if same_on_all {
		return true
	}
	// Diagonal 2
	same_on_all = true
	for index := 0; index < 3; index++ {
		if game[2-index][index] != playerSymbol {
			same_on_all = false
		}
	}
	if same_on_all {
		return true
	}
	return false
}

func main() {
	markedPlaces := 0

	game := [3][3]string {
		[3]string {"-", "-", "-"},
		[3]string {"-", "-", "-"},
		[3]string {"-", "-", "-"},
	}
	for {
		printGame(game)

		if markedPlaces >= 9 {
			fmt.Println("The game resulted in a stalemate.")
			break
		}

		game = playerInput("x", game)

		if win("x", game) {
			fmt.Println("x won the game.")
			break
		}

		markedPlaces++
		printGame(game)
		game = playerInput("o", game)

		if win("o", game) {
			fmt.Println("o won the game.")
			break
		}

		markedPlaces++
	}
}

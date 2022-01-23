package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type TicTacGame struct {
	board      [9]string
	player     string
	turnNumber int
}

func main() {
	var game TicTacGame
	game.player = "O"

	gameOver := false
	var winner string

	for gameOver != true {
		PrintBoard(game.board)
		move := askforplay()
		err := game.play(move)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(game.turnNumber)
		gameOver, winner = CheckForWinner(game.board, game.turnNumber)
	}
	PrintBoard(game.board)
	if winner == "" {
		fmt.Println("it's a draw ")
	} else {
		fmt.Printf("YaaY %s is winner ", winner)
	}
}

func CheckForWinner(b [9]string, n int) (bool, string) {

	check := false
	i := 0

	//horizantel
	for i < 9 {
		check = b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != ""
		if !check {
			i += 3
		} else {
			return true, b[i]
		}
	}

	i = 0
	//vertical
	for i < 3 {
		check = b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != ""
		if !check {
			i += 1
		} else {
			return true, b[i]
		}
	}

	//diagonal 1-9
	check = b[0] == b[4] && b[4] == b[8] && b[0] != ""
	if check {
		return true, b[0]
	}

	//diagonal 3-7
	check = b[2] == b[4] && b[4] == b[6] && b[2] != ""
	if check {
		return true, b[2]
	}

	if n == 9 {
		return true, ""
	}
	return false, ""
}

func ClearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func (game *TicTacGame) SwitchPlayers() {
	if game.player == "O" {
		game.player = "X"
		return
	}
	game.player = "O"
}

func (game *TicTacGame) play(pos int) error {
	if game.board[pos-1] == "" {
		game.board[pos-1] = game.player
		game.SwitchPlayers()
		game.turnNumber += 1
		return nil
	}
	return errors.New("try another move")
}

func askforplay() int {
	var moveInt int
	fmt.Println("Enter Position to play: ")
	fmt.Scan(&moveInt)
	return moveInt
}

func PrintBoard(b [9]string) {
	ClearScreen()
	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}

		if i > 0 && (i+1)%3 == 0 {

			fmt.Printf("\n")
		} else {
			fmt.Printf("|")

		}

	}
}

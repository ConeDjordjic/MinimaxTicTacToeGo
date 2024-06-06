package main

import (
	"fmt"
	"math"
)

type TicTacToe struct {
	Board   [3][3]rune
	Players [2]Player
}

type Player struct {
	Sign    rune
	Moves   []int
	IsHuman bool
}

func (p *Player) GetMove(board [3][3]rune) int {
	if p.IsHuman {
		return p.getHumanMove(board)
	}
	return p.getAIMove(board)
}

func (p *Player) getHumanMove(board [3][3]rune) int {
	var move int
	for {
		fmt.Printf("Player %c enter your move (0-8): ", p.Sign)
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 0 and 8.")
			continue
		}
		if move < 0 || move > 8 {
			fmt.Println("Invalid move. Please enter a number between 0 and 8.")
			continue
		}

		row := move / 3
		col := move % 3

		if board[row][col] != ' ' {
			fmt.Println("Cell is already occupied. Please choose another cell.")
			continue
		}

		break
	}
	p.Moves = append(p.Moves, move)
	return move
}

func (p *Player) getAIMove(board [3][3]rune) int {
	bestScore := math.Inf(-1)
	var bestMove int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == ' ' {
				board[i][j] = p.Sign
				score := minimax(board, 0, false, p.Sign)
				board[i][j] = ' '
				if score > bestScore {
					bestScore = score
					bestMove = i*3 + j
				}
			}
		}
	}
	p.Moves = append(p.Moves, bestMove)
	return bestMove
}

func getOpponent(sign rune) rune {
	if sign == 'X' {
		return 'O'
	}
	return 'X'
}

func (t *TicTacToe) NewGame() {
	println("Welcome to Tic-Tac-Toe!")
	t.Init()
	t.SelectPlayer()
	t.GameLoop()
}

func (t *TicTacToe) GameLoop() {
	for {
		t.Print()
		for _, p := range t.Players {
			move := p.GetMove(t.Board)
			row := move / 3
			col := move % 3

			t.Board[row][col] = p.Sign
			if t.CheckWin(p.Sign) {
				t.Print()
				fmt.Printf("Player %c wins!\n", p.Sign)
				return
			}
			if t.IsBoardFull() {
				t.Print()
				fmt.Println("The game is a draw!")
				return
			}
		}
	}
}

func (t *TicTacToe) CheckWin(sign rune) bool {
	return checkWin(t.Board, sign)
}

func checkWin(board [3][3]rune, sign rune) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if (board[i][0] == sign && board[i][1] == sign && board[i][2] == sign) ||
			(board[0][i] == sign && board[1][i] == sign && board[2][i] == sign) {
			return true
		}
	}
	// Check diagonals
	if (board[0][0] == sign && board[1][1] == sign && board[2][2] == sign) ||
		(board[0][2] == sign && board[1][1] == sign && board[2][0] == sign) {
		return true
	}
	return false
}

func (t *TicTacToe) IsBoardFull() bool {
	return isBoardFull(t.Board)
}

func isBoardFull(board [3][3]rune) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == ' ' {
				return false
			}
		}
	}
	return true
}

func (t *TicTacToe) Print() {
	b := t.Board

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == ' ' {
				fmt.Print(" ")
			} else {
				fmt.Print(string(b[i][j]))
			}
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		if i < 2 {
			fmt.Println("\n---------")
		}
	}
	fmt.Println()
}

func (t *TicTacToe) Init() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t.Board[i][j] = ' '
		}
	}
}

func (t *TicTacToe) SelectPlayer() {
	fmt.Println("Choose player: 1 - X, 2 - O")

	var player int
	_, err := fmt.Scan(&player)
	if err != nil {
		panic(err)
	}

	if player == 1 {
		t.Players[0].Sign = 'X'
		t.Players[0].IsHuman = true
		t.Players[1].Sign = 'O'
		t.Players[1].IsHuman = false
	} else if player == 2 {
		t.Players[0].Sign = 'O'
		t.Players[0].IsHuman = false
		t.Players[1].Sign = 'X'
		t.Players[1].IsHuman = true
	} else {
		panic("Invalid input")
	}
}

package main

import "math"

func minimax(board [3][3]rune, depth int, isMaximizing bool, aiSign rune) float64 {
	if checkWin(board, aiSign) {
		return 10
	}
	if checkWin(board, getOpponent(aiSign)) {
		return -10
	}
	if isBoardFull(board) {
		return 0
	}

	if isMaximizing {
		bestScore := math.Inf(-1)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == ' ' {
					board[i][j] = aiSign
					score := minimax(board, depth+1, false, aiSign)
					board[i][j] = ' '
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == ' ' {
					board[i][j] = getOpponent(aiSign)
					score := minimax(board, depth+1, true, aiSign)
					board[i][j] = ' '
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
}

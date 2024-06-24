package main

import (
	"strings"
)

func checkWinCons(board [][]string)(gameActive bool){
	gameActive = true
	mines := getMines()

	uncoveredSpaces := 0
	totalSpaces := (len(board) * len(board[0]))

	for y := range len(board){
		for x := range len(board[0]){
			if strings.HasPrefix(board[y][x], "U"){
				uncoveredSpaces++
			}
		}
	}

	if (totalSpaces - mines) == uncoveredSpaces{
		setMessages("~~ YOU WIN ~~")
		board = revealMines(board)
		displayBoard(board)
		gameActive = false
	}

	return gameActive
}


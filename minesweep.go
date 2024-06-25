package main

import (
	"regexp"
	"strings"
)

func minesweep(coord []int, oldBoard [][]string, firstTurn bool)(gameActive bool, board [][]string){

	gameActive = true
	if !(coord[0] >= 0) || !(coord[0] >= len(oldBoard) || !(coord[1] >= 0) || !(coord[1] >= len(oldBoard[0]))){
		setMessages("co-ordinate out of bounds")
		return true, oldBoard
	}

	var digitCheck = regexp.MustCompile(`[0-9]`)

	switch {
		case oldBoard[coord[0]][coord[1]] == "!":
			if (!firstTurn){
				gameActive = false
				oldBoard[coord[0]][coord[1]] = "E"
				board = oldBoard
				board = revealMines(board)
				setMessages("~~ game over! ~~")
				displayBoard(board)
			} else{
				firstTurn = false
				oldBoard = firstTurnFailsafe(coord, oldBoard)
				gameActive, oldBoard = minesweep(coord, oldBoard, gameActive)
			}
		case digitCheck.MatchString(oldBoard[coord[0]][coord[1]]):
			if !strings.HasPrefix(oldBoard[coord[0]][coord[1]], "U"){
				oldBoard[coord[0]][coord[1]] = "U" + oldBoard[coord[0]][coord[1]]
			}
		case oldBoard[coord[0]][coord[1]] == "X":
			oldBoard[coord[0]][coord[1]] = "U"
			oldBoard = uncoverSpace(oldBoard, coord)
	}

	board = oldBoard
	return gameActive, board
}

func uncoverSpace(oldBoard [][]string, coord []int)(board [][]string){

	var digitCheck = regexp.MustCompile(`[0-9]`)

	xRows := []int{(coord[1])-1,coord[1],(coord[1])+1}
	yRows := []int{(coord[0])-1,coord[0],(coord[0])+1}

	for i := range yRows{
		if yRows[i] >= 0 && yRows[i] < len(oldBoard){
			for x := range xRows{
				if xRows[x] >= 0 && xRows[x] < len(oldBoard[0]){
					switch {
						case oldBoard[yRows[i]][xRows[x]] == "X":
							oldBoard[yRows[i]][xRows[x]] = "U"
							coord := []int{yRows[i], xRows[x]} 
							uncoverSpace(oldBoard, coord)
						
						case digitCheck.MatchString(oldBoard[yRows[i]][xRows[x]]):
							if !strings.HasPrefix(oldBoard[yRows[i]][xRows[x]], "U"){
								oldBoard[yRows[i]][xRows[x]] = "U" + oldBoard[yRows[i]][xRows[x]]
							}
						}
				}
			}
		}
	}
	board = oldBoard
	return board
}
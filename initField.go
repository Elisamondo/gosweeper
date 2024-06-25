package main

import (
	"fmt"
	"math/rand"
)
var mines int

func initField(difficulty int) (board [][]string){
	fmt.Printf("difficulty selected - %d\n", difficulty)

	board = make([][]string, 0)
	var boardSize [2] int

	switch difficulty{
		case 0:
			mines = 10
			boardSize[0] = 9
			boardSize[1] = 9
		case 1:
			mines = 40
			boardSize[0] = 16
			boardSize[1] = 16
		case 2:
			mines = 80
			boardSize[0] = 16
			boardSize[1] = 26
	} 
	

	for range boardSize[0]{
		row := []string{}
		for range boardSize[1]{
			row = append(row, "X")
		}
		board = append(board, row)
	}

	for range mines {
		board = placeMine(board)
	}

	return board
}


func placeMine(oldBoard [][]string)(board [][]string){

	var coord [2]int
	coord[0],coord[1] = rand.Intn(len(oldBoard)), rand.Intn(len(oldBoard[0]))

	if (oldBoard[coord[0]][coord[1]] != "!"){
		oldBoard[coord[0]][coord[1]] = "!"
		board = placeNumbers(oldBoard, coord)
	} else
	{
		board = placeMine(oldBoard)
	}

	return board
} 

func placeNumbers(oldBoard [][]string, coord [2]int)(board [][]string){

	xRows := []int{(coord[1])-1,coord[1],(coord[1])+1}
	yRows := []int{(coord[0])-1,coord[0],(coord[0])+1}

	for y := range yRows{
		if yRows[y] >= 0 && yRows[y] < len(oldBoard){
			for x := range xRows{
				if xRows[x] >= 0 && xRows[x] < len(oldBoard[0]){
					switch oldBoard[yRows[y]][xRows[x]]{
						case "!":
							continue
						case "X":
							oldBoard[yRows[y]][xRows[x]] = "1"
						case "1":
							oldBoard[yRows[y]][xRows[x]] = "2"
						case "2":
							oldBoard[yRows[y]][xRows[x]] = "3"
						case "3":
							oldBoard[yRows[y]][xRows[x]] = "4"
						case "4":
							oldBoard[yRows[y]][xRows[x]] = "5"
						case "5":
							oldBoard[yRows[y]][xRows[x]] = "6"
						case "6":
							oldBoard[yRows[y]][xRows[x]] = "7"
						case "7":
							oldBoard[yRows[y]][xRows[x]] = "8"
						} 
					}
				}
			}
		}

	board = oldBoard
	return board
}


func getMines()(sendMines int){
	sendMines = mines
	return sendMines
}
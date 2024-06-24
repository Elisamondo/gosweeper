package main

import ("fmt"
"strings")

func main(){
	
	var board = make([][]string, 0)

	fmt.Println("GoSweeper \n By Elisa Edson \n \n ")
	board = initField(setDifficulty())

	gameController(board)
}

func setDifficulty() (difficulty int){
	var input string
	fmt.Println("Choose difficulty: Easy | Medium | Hard")
	fmt.Scan(&input)
	switch strings.ToUpper(input) {
	case "EASY", "E", "0":
		difficulty = 0
	case "MEDIUM", "M", "1":
		difficulty = 1
	case "HARD", "H", "2":
		difficulty = 2
	default: 
		difficulty = setDifficulty()
	}
	return difficulty
}
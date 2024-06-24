package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"bufio"
	"os"

)
func gameController(board [][]string){
	gameActive := true
	for (gameActive){
		displayBoard(board)
		gameActive, board = playerInteract(board)
		if (gameActive){
			gameActive = checkWinCons(board)
		}
	}
}

func displayBoard(board [][]string){	
	clear()
	for y := range len(board){
		row := []string{}
		for x := range len(board[y]){
			row = append(row, getCharacter(board[y][x]))
		}

		yLabel := (len(board) - y)
		var labelWhitespace string
		switch len(fmt.Sprint(yLabel)){
			case 1:
				labelWhitespace = "  "
			case 2: 
				labelWhitespace = " "
		}
		fmt.Printf("%v%v", yLabel, labelWhitespace)
		fmt.Printf("%v\n", row)
	}

	var Xlabel string
	for i := range len(board[0]){
		concat := rune('A' + i)
		Xlabel = Xlabel + " " + string(concat)
	}
	fmt.Printf("   %v\n", Xlabel)
	messages = getMessages()
	if (len(messages) != 0){
		for i := range messages{
			fmt.Printf("%v\n", messages[i])
		}
	}
	clearMessages()
}

func getCharacter(coord string)(char string){
	switch {
		case coord == "U":
			char = " "
		case strings.HasPrefix(coord, "f"), strings.HasPrefix(coord, "F"):
			char = "⚑"
		case coord == "E":
			char = "*"
		case strings.HasPrefix(coord, "U"):
			char = strings.Trim(coord, "U")
		default:
			char = "#"
	}
	return char
}

func playerInteract(oldBoard [][]string)(gameActive bool, board [][]string){
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	flag, _ := regexp.MatchString(`-F`, input)
	if (flag){
		input = strings.Replace(input, "-F", "", 1)
	}

	validCoord, _ := regexp.MatchString("([A-Z]{1})(\\d{2}|\\d{1})", input)
	board = oldBoard
	if (!validCoord ){
		setMessages("not a valid target - Choose a valid co-ordinate (I.E D10)")
		gameActive = true
		//gameActive, board = playerInteract(oldBoard)
		return true, oldBoard
	}

	xCoord, _ := regexp.Compile("([A-Z]{1})")
	yCoord, _ := regexp.Compile(`(\d{2}|\d{1})`)
	var targetCoord [2]string
	targetCoord[1] = xCoord.FindString(input)
	targetCoord[0] = yCoord.FindString(input)


	var convertedCoord [2]int
	temp := rune(targetCoord[1][0])
	convertedCoord[1] = int(temp - 'A')
	convertedCoord[0], _ = strconv.Atoi(targetCoord[0])

	if flag{
		board = placeFlag(convertedCoord[:], board)
		gameActive = true
	} else {
		gameActive, board = minesweep(convertedCoord[:], board)
	}

	return gameActive, board
}

func placeFlag(coord []int, oldBoard [][]string)(board [][]string){
	board = oldBoard

	coord[0] = len(board) - coord[0]

	switch board[coord[0]][coord[1]]{
		case "!":
			board[coord[0]][coord[1]] = "F"
		case "X":
			board[coord[0]][coord[1]] = "f"
		case "F":
			board[coord[0]][coord[1]] = "!"
		case "f":
			board[coord[0]][coord[1]] = "X"
		case "1":
			board[coord[0]][coord[1]] = "f1"
		case "2":
			board[coord[0]][coord[1]] = "f2"
		case "3":
			board[coord[0]][coord[1]] = "f3"
		case "4":
			board[coord[0]][coord[1]] = "f4"
		case "5":
			board[coord[0]][coord[1]] = "f5"
		case "6":
			board[coord[0]][coord[1]] = "f6"
		case "7":
			board[coord[0]][coord[1]] = "f7"
		case "8":
			board[coord[0]][coord[1]] = "f8"
		case "f1":
			board[coord[0]][coord[1]] = "1"
		case "f2":
			board[coord[0]][coord[1]] = "2"
		case "f3":
			board[coord[0]][coord[1]] = "3"
		case "f4":
			board[coord[0]][coord[1]] = "4"
		case "f5":
			board[coord[0]][coord[1]] = "5"
		case "f6":
			board[coord[0]][coord[1]] = "6"
		case "f7":
			board[coord[0]][coord[1]] = "7"
		case "f8":
			board[coord[0]][coord[1]] = "8"
		default:
			setMessages("You can not place a flag on this tile")

	}
	return board
}
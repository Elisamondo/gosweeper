package main

func revealMines(oldBoard [][]string)(board [][]string){
	board = oldBoard
	for y := range len(board){
		for x := range len(board[0]){
			switch board[y][x]{
				case "!":
					board[y][x] = "E"
				case "F":
					board[y][x] = "E"
			}
		}
	}

	return board
}
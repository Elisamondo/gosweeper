package main

func firstTurnFailsafe(coord []int, oldBoard [][]string)(board [][]string){
	board = placeMine(oldBoard)

	board = removeNumbers(board, coord)
	board = removeMine(board, coord)
	return board
}

func removeNumbers(oldBoard [][]string, coord []int)(board [][]string){

	xRows := []int{(coord[1])-1,coord[1],(coord[1])+1}
	yRows := []int{(coord[0])-1,coord[0],(coord[0])+1}

	for i := range yRows{
		if yRows[i] >= 0 && yRows[i] < len(oldBoard){
			for x := range xRows{
				if xRows[x] >= 0 && xRows[x] < len(oldBoard[0]){
					switch oldBoard[yRows[i]][xRows[x]]{
						case "!":
							continue
						case "1":
							oldBoard[yRows[i]][xRows[x]] = "X"
						case "2":
							oldBoard[yRows[i]][xRows[x]] = "1"
						case "3":
							oldBoard[yRows[i]][xRows[x]] = "2"
						case "4":
							oldBoard[yRows[i]][xRows[x]] = "3"
						case "5":
							oldBoard[yRows[i]][xRows[x]] = "4"
						case "6":
							oldBoard[yRows[i]][xRows[x]] = "5"
						case "7":
							oldBoard[yRows[i]][xRows[x]] = "6"
						case "8":
							oldBoard[yRows[i]][xRows[x]] = "7"
						}
					}
				}
			}
		}
		board = oldBoard
	return board
}

func removeMine(oldBoard [][]string, coord []int)(board [][]string){
	board = oldBoard
	xRows := []int{(coord[1])-1,coord[1],(coord[1])+1}
	yRows := []int{(coord[0])-1,coord[0],(coord[0])+1}

	board[coord[0]][coord[1]] = "X"

	for y := range yRows{
		if yRows[y] >= 0 && yRows[y] < len(oldBoard){
			for x := range xRows{
				if xRows[x] >= 0 && xRows[x] < len(oldBoard[0]){
				if (board[yRows[y]][xRows[x]] == "!"){
					switch board[coord[0]][coord[1]]{
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
			}}
	}}

	return board
}

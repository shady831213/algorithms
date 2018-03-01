package dp

import (
	"math"
	"strings"
	"fmt"
)

//for debug, show the board
func printBoard(board [][]int) {
	line := strings.Repeat(" ", 5)+strings.Repeat(" -----", len(board[0]))

	for i := range board {
		if i == 0 {
			fmt.Println(line)
		}
		content := " "+fmt.Sprintf("%3d", len(board)-1-i)+" |"
		for _,v := range board[len(board)-1-i] {
			content += " "+fmt.Sprintf("%3d",v)+" |"
		}
		fmt.Println(content)
		fmt.Println(line)
		if i == len(board) - 1 {
			xIdxs:= strings.Repeat(" ", 5)
			for j := range board[0] {
				xIdxs += " "+fmt.Sprintf("%3d", j)+"  "
			}
			fmt.Println(xIdxs)
		}
	}
}


func chessGame(board [][]int, start, end int) (totalScore int, path [][]int) {
	ySize, xSize := len(board), len(board[0])
	//must be n*n size board
	if xSize != ySize {
		panic("xSize must be equal to ySize")
	}
	//side data store score and path
	scoreboard := make([][]int, ySize, ySize)
	pathboard := make([][]int, ySize, ySize)
	for i := range scoreboard {
		//deal with right and left boundary, all filled Min value
		scoreboard[i] = make([]int, xSize+2, xSize+2)
		pathboard[i] = make([]int, xSize, xSize)
		scoreboard[i][0], scoreboard[i][xSize+1] = math.MinInt32, math.MinInt32
	}
	for j := 1; j < xSize+1; j++ {
		//mask all other position, use min value
		if j == start + 1 {
			scoreboard[0][j] = board[0][start]
		} else {
			scoreboard[0][j] = math.MinInt32
		}
	}


	//max(scoreboard[i-1][j-1], scoreboard[i-1][j], scoreboard[i-1][j+1]) + score[i][j]
	//store the previous position with -1, 0, 1, which mean left-down, down, right-down
	for i := 1; i < ySize; i ++ {
		for j := 1; j < xSize+1; j++ {
			scoreboard[i][j] = math.MinInt32
			for k:= j-1;k < j+2;k ++ {
				temp := scoreboard[i-1][k] + board[i][j-1]
				if temp > scoreboard[i][j] {
					scoreboard[i][j] = temp
					pathboard[i][j-1] = k - j
				}
			}
		}
	}

	totalScore = scoreboard[ySize-1][end+1]

	//track the path
	path = make([][]int,ySize,ySize)
	path[ySize-1] = []int{ySize-1, end}
	for i,j := ySize -1, end; i > 0; i--{
		j = pathboard[i][j]+j
		path[i-1] = []int{i-1, j}
	}
	return
}

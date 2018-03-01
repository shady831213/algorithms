package dp

import (
	"testing"
	"math/rand"
	"fmt"
	"reflect"
)

func buildBoard(size int)([][]int) {
	board := make([][]int, size, size)
	for i := range board {
		board[i] = make([]int, size, size)
		for j := range board[i] {
			board[i][j] = rand.Intn(20) - 10
		}
	}
	return board
}


func TestChessGame(t *testing.T)  {
	board := buildBoard(5)
	score, path := chessGame(board, 0, 3)
	if score != 3 {
		t.Log(fmt.Sprintf("score expect 3, but get %d", score))
		printBoard(board)
		t.Fail()
	}
	if !reflect.DeepEqual(path, [][]int{[]int{0,0}, []int{1,0}, []int{2,1}, []int{3,2}, []int{4,3}}) {
		t.Log("path wrong!\n")
		printBoard(board)
		t.Log(path)
		t.Fail()
	}

}
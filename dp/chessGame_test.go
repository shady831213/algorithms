package dp

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func buildBoard(size int) [][]int {
	_rand := rand.New(rand.NewSource(1))
	board := make([][]int, size, size)
	for i := range board {
		board[i] = make([]int, size, size)
		for j := range board[i] {
			board[i][j] = _rand.Intn(20) - 10
		}
	}
	return board
}

func TestChessGame(t *testing.T) {
	board := buildBoard(5)
	score, path := chessGame(board, 0, 3)
	if score != 3 {
		t.Log(fmt.Sprintf("score expect 3, but get %d", score))
		printBoard(board)
		t.Fail()
	}
	if !reflect.DeepEqual(path, [][]int{{0, 0}, {1, 0}, {2, 1}, {3, 2}, {4, 3}}) {
		t.Log("path wrong!\n")
		printBoard(board)
		t.Log(path)
		t.Fail()
	}

}

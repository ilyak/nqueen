package main

import (
	"fmt"
	"os"
	"strconv"
)

func checkij(board [][]bool, i, j int) bool {
	n := len(board)
	for d := 1; i+d < n; d++ {
		if board[i+d][j] {
			return false
		}
	}
	for d := 1; j+d < n; d++ {
		if board[i][j+d] {
			return false
		}
	}
	for d := 1; i+d < n && j+d < n; d++ {
		if board[i+d][j+d] {
			return false
		}
	}
	for d := 1; i-d >= 0 && j+d < n; d++ {
		if board[i-d][j+d] {
			return false
		}
	}
	return true
}

func checkboard(board [][]bool) bool {
	n := len(board)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] && !checkij(board, row, col) {
				return false
			}
		}
	}
	return true
}

func printboard(board [][]bool) {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] {
				fmt.Print("q")
			} else {
				fmt.Print("x")
			}
		}
		fmt.Println()
	}
}

func solve(board [][]bool, row int) int {
	n := len(board)
	if row == n {
		printboard(board)
		fmt.Println()
		return 1
	}
	ns := 0
	for col := 0; col < n; col++ {
		board[row][col] = true
		if checkboard(board) {
			ns += solve(board, row+1)
		}
		board[row][col] = false
	}
	return ns
}

func nqueen(n int) {
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}
	ns := solve(board, 0)
	if ns == 1 {
		fmt.Println(ns, "solution found")
	} else {
		fmt.Println(ns, "solutions found")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("specify n")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 1 {
		fmt.Println("n must be positive")
		return
	}
	nqueen(n)
}

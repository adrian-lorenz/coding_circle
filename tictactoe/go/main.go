package main

import "fmt"

func checkWin(n int, coordinates [][]int) bool {
	
	rowCount := make([]int, n)
	colCount := make([]int, n)
	diagonalCount1 := 0
	diagonalCount2 := 0

	for _, coord := range coordinates {
		x, y := coord[0], coord[1]
		//sum rows and cols
		rowCount[x]++
		colCount[y]++
		//check rows and cols
		if rowCount[x] == n || colCount[y] == n {
			return true
		}
        //sum / check diag1
		if x == y {
			diagonalCount1++
			if diagonalCount1 == n {
				return true
			}
		}
		//sum / check diag2
		if x+y == n-1 {
			diagonalCount2++
			if diagonalCount2== n {
				return true
			}
		}
	}

	return false
}


func drawBoard(n int, coordinates [][]int) {
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "_"
		}
	}
	for _, coord := range coordinates {
		x, y := coord[0], coord[1]
		board[x][y] = "X"
	}

	for _, row := range board {
	fmt.Println(row)
	}
}

func main() {
	n := 3
	coordinates := [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 2}}
	drawBoard(n, coordinates)
	result := checkWin(n, coordinates)
	fmt.Println(result)
}

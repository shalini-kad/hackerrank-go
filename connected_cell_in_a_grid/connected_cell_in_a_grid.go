package main

import "fmt"

func main() {
	var numRows, numCols int
	fmt.Scan(&numRows, &numCols)

	matrix := make([][]int, numRows)
	memoMatrix := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		matrix[i] = make([]int, numCols)
		memoMatrix[i] = make([]int, numCols)

		for j := 0; j < numCols; j++ {
			fmt.Scan(&matrix[i][j])
			memoMatrix[i][j] = 0
		}
	}

	largestArea := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			area := fillSearch(matrix, memoMatrix, i, j)
			if area > largestArea {
				largestArea = area
			}
		}
	}

	fmt.Println(largestArea)
}

func fillSearch(matrix, memoMatrix [][]int, row, col int) int {
	// base case: if position is not valid
	//            or if position has already been considered
	//            or if position is not filled
	if row >= len(matrix) || row < 0 || col >= len(matrix[0]) || col < 0 {
		return 0
	} else if memoMatrix[row][col] == 1 {
		return 0
	} else if matrix[row][col] == 0 {
		return 0
	}
	
	// start from current position
	area := 1
	memoMatrix[row][col] = 1
	
	// fill search from every adjacent position
	area += fillSearch(matrix, memoMatrix, row+1, col)   +
		fillSearch(matrix, memoMatrix, row+1, col+1) +
		fillSearch(matrix, memoMatrix, row, col+1)   +
		fillSearch(matrix, memoMatrix, row-1, col+1) +
		fillSearch(matrix, memoMatrix, row-1, col)   +
		fillSearch(matrix, memoMatrix, row-1, col-1) +
		fillSearch(matrix, memoMatrix, row, col-1)   +
		fillSearch(matrix, memoMatrix, row+1, col-1)

	return area
}

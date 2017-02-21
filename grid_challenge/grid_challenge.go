package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var numTests int
	line, _ := r.ReadString('\n')
	fmt.Sscan(line, &numTests)

	for i := 0; i < numTests; i++ {
		runTest(r)
	}
}

func runTest(r *bufio.Reader) {
	var size int
	line, _ := r.ReadString('\n')
	fmt.Sscan(line, &size)

	matrix := make([][]int, size)
	for row := range matrix {
		matrix[row] = make([]int, size)

		line, _ := r.ReadString('\n')
		for col := range matrix[row] {
			matrix[row][col] = int(line[col])
		}

		sort.Ints(matrix[row])
	}

	sorted := true
	for col := range matrix[0] {
		colSlice := make([]int, size)

		for row := range matrix {
			colSlice[row] = matrix[row][col]
		}

		if !sort.IntsAreSorted(colSlice) {
			sorted = false
			break
		}
	}

	if sorted {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

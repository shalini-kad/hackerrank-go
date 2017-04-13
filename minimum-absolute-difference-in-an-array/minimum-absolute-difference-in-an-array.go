package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var arrSize int
	fmt.Scan(&arrSize)

	arr := make([]int, arrSize)
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)

	minDiff := math.MaxInt32
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if diff := abs(arr[i] - prev); diff < minDiff {
			minDiff = diff
		}
		prev = arr[i]
	}

	fmt.Print(minDiff)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

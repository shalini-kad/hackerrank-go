package main

import (
	"fmt"
	"sort"
)

func main() {
	var numTests int
	fmt.Scan(&numTests)
	for i := 0; i < numTests; i++ {
		runTest()
	}
}

func runTest() {
	var arrLen, min int
	fmt.Scan(&arrLen, &min)

	a1 := make([]int, arrLen)
	a2 := make([]int, arrLen)

	for i, _ := range a1 {
		fmt.Scan(&a1[i])
	}

	for i, _ := range a2 {
		fmt.Scan(&a2[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a1)))
	sort.Ints(a2)

	for i, _ := range a1 {
		if a1[i]+a2[i] < min {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

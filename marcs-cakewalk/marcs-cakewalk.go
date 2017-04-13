package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var numCupcakes int
	fmt.Scan(&numCupcakes)

	cupcakes := make([]int, numCupcakes)
	for i, _ := range cupcakes {
		fmt.Scan(&cupcakes[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cupcakes)))

	var miles int64
	for i, cals := range cupcakes {
		miles += int64(cals) * int64(math.Pow(2, float64(i)))
	}

	fmt.Print(miles)
}

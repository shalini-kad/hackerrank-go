package main

import (
	"fmt"
	"sort"
)

func main() {
	var numToys int
	fmt.Scan(&numToys)

	toys := make([]int, numToys)
	for i := range toys {
		fmt.Scan(&toys[i])
	}

	sort.Ints(toys)

	numPurchases := 0
	for i := 0; i < len(toys); i++ {

		startWeight := toys[i]
		for ; i < len(toys); i++ {
			if toys[i] - startWeight > 4 {
				break
			}
		}
		i--

		numPurchases++
	}

	fmt.Println(numPurchases)
}

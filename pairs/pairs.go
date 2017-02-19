package main

import (
	"fmt"
)

func main() {

	var numNums, goalDifference int
	fmt.Scan(&numNums, &goalDifference)

	numbers := make([]int, numNums)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}
	
	numMap := make(map[int]int)
	for i := range numbers {
		num := numbers[i]
		numMap[num] = 0
		if _, ok := numMap[num + goalDifference]; ok {
			numMap[num + goalDifference]++
			numMap[num]++
		}
		if _, ok := numMap[num - goalDifference]; ok {
			numMap[num - goalDifference]++
			numMap[num]++
		}
	}

	numFound := 0
	for _, v := range numMap {
		numFound += v
	}
	numFound /= 2
	
	fmt.Println(numFound)
}

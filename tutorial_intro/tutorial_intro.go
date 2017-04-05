package main

import "fmt"

func main() {
	var targetNum, arrayLength int
	fmt.Scan(&targetNum, &arrayLength)

	for i := 0; i < arrayLength; i++ {
		var num int
		fmt.Scan(&num)
		if num == targetNum {
			fmt.Printf("%d", i)
			break
		}
	}
}

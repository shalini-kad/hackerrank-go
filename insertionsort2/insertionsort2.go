package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)

	arr := make([]int, size)
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	for i := 1; i < len(arr); i++ {
		insertElement(arr, i)
		printArr(arr)
	}
}

func insertElement(arr []int, pos int) {
	elementVal := arr[pos]
	elementPlaced := false

	for i := pos - 1; i >= 0; i-- {
		if arr[i] < elementVal {
			arr[i+1] = elementVal
			elementPlaced = true
			break
		} else {
			arr[i+1] = arr[i]
		}
	}

	if !elementPlaced {
		arr[0] = elementVal
	}

}

func printArr(arr []int) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

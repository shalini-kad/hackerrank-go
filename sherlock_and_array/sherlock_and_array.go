package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var numTests int
	fmt.Fscan(in, &numTests)

	for i := 0; i < numTests; i++ {
		runTest(in)
	}

}

func runTest(in *bufio.Reader) {
	var len int
	fmt.Fscan(in, &len)

	sum := 0
	arr := make([]int, len)
	for i := range arr {
		fmt.Fscan(in, &arr[i])
		sum += arr[i]
	}

	success := false
	tmpSum := 0
	for _, val := range arr {
		if sum-tmpSum-val == tmpSum {
			success = true
		}
		tmpSum += val
	}

	if success {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

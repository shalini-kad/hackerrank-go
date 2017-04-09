package main

import (
	"fmt"
	"math"
)

func main() {
	var targetNum, power int
	fmt.Scan(&targetNum, &power)

	currentAddendBase := computeMaxAddendBase(targetNum, power)
	numFound := numPowerSums(targetNum, power, currentAddendBase)
	fmt.Println(numFound)
}

func numPowerSums(targetNum, power, currentAddendBase int) int {

	// base cases
	if targetNum == 0 {
		return 1
	} else if currentAddendBase == 0 || targetNum < 0 {
		return 0
	}

	// try using current addend base
	addend := int(math.Pow(float64(currentAddendBase), float64(power)))
	numFound := numPowerSums(targetNum-addend, power, currentAddendBase-1)
	// try without using current addend base
	numFound += numPowerSums(targetNum, power, currentAddendBase-1)

	return numFound
}

func computeMaxAddendBase(num, b int) int {
	// x ^ b = num
	// Log(x ^ b) = Log(num)
	// b Log(x) = Log(num)
	// Log(x) = Log(num) / b
	// x = e ^ (Log(num) / b)
	log := math.Pow(math.E, math.Log(float64(num))/float64(b))
	log = math.Floor(log)
	// sometimes floating point math is a little off
	if math.Pow(float64(log+1), float64(b)) <= float64(num) {
		log = log + 1
	}
	return log
}

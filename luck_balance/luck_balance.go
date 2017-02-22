package main

import (
	"fmt"
	"sort"
)

type competition []contest

type contest struct {
	luckValue int
	important bool
}

func main() {
	var numContests, numAllowableLosses int
	fmt.Scan(&numContests, &numAllowableLosses)

	contests := make([]contest, numContests)
	numImportant := 0

	for i := range contests {
		fmt.Scan(&contests[i].luckValue, &contests[i].important)
		if contests[i].important {
			numImportant++
		}
	}

	sort.Sort(competition(contests))

	luck := 0

	numWins := numImportant - numAllowableLosses
	if numWins < 0 {
		numWins = 0
	}
	
	for i := 0; i < numWins; i++ {
		// win the first ones
		luck -= contests[i].luckValue
	}

	for i := numWins; i < len(contests); i++ {
		// lose the rest (for extra luck??)
		luck += contests[i].luckValue
	}

	fmt.Println(luck)
}

func (contests competition) Len() int {
	return len(contests)
}

func (contests competition) Less(i, j int) bool {
	switch {
	case contests[i].important && !contests[j].important:
		return true
	case !contests[i].important && contests[j].important:
		return false
	case contests[i].luckValue < contests[j].luckValue:
		return true
	default:
		return false
	}
}

func (contests competition) Swap(i, j int) {
	contests[i], contests[j] = contests[j], contests[i]
}

package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	var numToys, cash int
	fmt.Scan(&numToys, &cash)

	toys := &IntHeap{}
	heap.Init(toys)
	
	for i := 0; i < numToys; i++ {
		var toy int
		fmt.Scan(&toy)
		heap.Push(toys, toy)
	}

	numToysBought := 0
	cashSpent := 0
	for toys.Len() > 0 {
		toy := heap.Pop(toys)
		if cashSpent += toy.(int); cashSpent > cash {
			break
		}
		numToysBought++
	}

	fmt.Println(numToysBought)
}

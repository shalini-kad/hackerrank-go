package main

import "fmt"

// numRods is the number of rods in a (potentially) modified Hanoi game
const NUM_RODS = 4

// gameState represents a state in the modified Hanoi game
type gameState struct {
	// each rod contains a slice of ints representing the
	// diameters of the disks on the rod from bottom to top
	rods [][]int 
}

func main() {
	var numDisks int
	fmt.Scan(&numDisks)
	
	diskInput := make([]int, numDisks)
	for i := 0; i < numDisks; i++ {
		fmt.Scan(&diskInput[i])
	}

	// parse disk input into a gameState
	goalState := newGameState(NUM_RODS)
	for i := numDisks - 1; i >= 0; i-- {
		rodNum := diskInput[i] - 1
		diskDiameter := i + 1
		goalState.rods[rodNum] = append(goalState.rods[rodNum], diskDiameter)
	}

	// create initial state
	rootState := newGameState(NUM_RODS)
	for i := numDisks; i > 0; i-- {
		rootState.rods[0] = append(rootState.rods[0], i)
	}

	toSearch := getNextStates(rootState)
	found := make([]gameState, 0)

	depth := 0
SearchLoop:
	for len(toSearch) != 0 {
		depth++
		for _, state := range toSearch {
			if gameStatesEqual(state, goalState) {
				break SearchLoop
			}

			// add to found the new states
			nextStates := getNextStates(state)
			for _, nextState := range nextStates {
				isRepeatState := false
				for _, existingState := range found {
					if gameStatesEqual(nextState, existingState) {
						isRepeatState = true
						break
					}
				}
				if !isRepeatState {
					found = append(found, nextState)
				}
			}
		}

		toSearch = found
		found = make([]gameState, 0)
	}

	fmt.Println(depth)	
}

func getNextStates(state gameState) []gameState {
	nextStates := make([]gameState, 0)

	// loop over all rods
	for i := 0; i < NUM_RODS; i++ {
		// if rod contains a disk
		if rod := state.rods[i]; len(rod) > 0 {
			diskDiameter := rod[len(rod)-1]
			// loop over all other rods and check
			// if disk can be moved there
			for j := 0; j < NUM_RODS; j++ {
				// skip current rod
				if i == j {
					continue
				}

				if targetRod := state.rods[j]; len(targetRod) == 0 || targetRod[len(targetRod)-1] > diskDiameter {
					nextState := copyGameState(state)
					nextState.rods[i] = rod[:len(rod)-1]
					nextState.rods[j] = append(nextState.rods[j], diskDiameter)

					nextStates = append(nextStates, nextState)
				}
			}
		}
	}

	return nextStates
}

func newGameState(numRods int) gameState {
	var state gameState
	state.rods = make([][]int, numRods)
	for i := range state.rods {
		state.rods[i] = make([]int, 0)
	}
	return state
}

func copyGameState(state gameState) gameState {
	var newState gameState
	newState.rods = make([][]int, len(state.rods))
	for i := range newState.rods {
		newState.rods[i] = make([]int, len(state.rods[i]))
		copy(newState.rods[i], state.rods[i])
	}
	return newState
}

func gameStatesEqual(s1, s2 gameState) bool {
	if len(s1.rods) != len(s2.rods) {
		return false
	}

	for i := 0; i < len(s1.rods); i++ {
		r1, r2 := s1.rods[i], s2.rods[i]

		if len(r1) != len(r2) {
			return false
		}

		for j := 0; j < len(r1); j++ {
			d1, d2 := r1[j], r2[j]
			if d1 != d2 {
				return false
			}
		}
	}

	return true
}

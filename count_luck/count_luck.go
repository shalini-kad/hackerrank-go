package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	var numTests int
	line, _ := reader.ReadString('\n')
	fmt.Sscan(line, &numTests)
	for i := 0; i < numTests; i++ {
		runTest(reader)
	}
}

// position is a 2d coordinate
type position struct {
	x, y int
}

// treeNode is a node representing a position within the maze
type treeNode struct {
	pos position
	lastPos position
	children []*treeNode
}

func runTest(r *bufio.Reader) {
	var height, width int
	line, _ := r.ReadString('\n')
	fmt.Sscan(line, &height, &width)

	forest := make([][]byte, height)
	start := position{}
	end := position{}

	// save maze and start/end points
	for i := 0; i < height; i++ {
		forest[i] = make([]byte, width)
		line, _ := r.ReadString('\n')

		for j := 0; j < width; j++ {
			forest[i][j] = line[j]

			if forest[i][j] == 'M' {
				start.x = j
				start.y = i
			} else if forest[i][j] == '*' {
				end.x = j
				end.y = i
			}
		}
	}

	// save guess as to how many choices it will take to reach end
	// (aka depth of goal node)
	var guess int
	line, _ = r.ReadString('\n')
	fmt.Sscan(line, &guess)

	// initialize root node as starting position
	root := new(treeNode)
	root.pos = start
	root.lastPos = start
	root.children = make([]*treeNode, 0)

	// build a tree of important positions
	buildPositionTree(forest, root)

	// if root has only one child, skip to child
	if len(root.children) == 1 {
		root = root.children[0]
	}

	// perform breadth first search on tree to find goal node
	toSearch := make([]*treeNode, 0)
	found := make([]*treeNode, 0)

	for _, node := range root.children {
		toSearch = append(toSearch, node)
	}

	depth := 0
SearchLoop:
	for len(toSearch) != 0 {
		depth++
		for _, node := range toSearch {
			if node.pos == end {
				break SearchLoop
			}
			found = append(found, node.children...)
		}
		toSearch = found
		found = make([]*treeNode, 0)
	}

	// output result
	if depth == guess {
		fmt.Println("Impressed")
	} else {
		fmt.Println("Oops!")
	}
}

// buildPositionTree builds a tree of the important positions in a 2d slice representing the maze.
// Important positions are places where the path forks, dead ends, the start position, and the goal
func buildPositionTree(forest [][]byte, node *treeNode) {
	adjacentPositions := getAdjacentPositions(forest, node.pos, node.lastPos)
	for _, pos := range adjacentPositions {
		nextPos, lastPos := getNextPosition(forest, pos, node.pos)

		childNode := new(treeNode)
		childNode.pos = nextPos
		childNode.lastPos = lastPos
		childNode.children = make([]*treeNode, 0)

 		node.children = append(node.children, childNode)

		buildPositionTree(forest, childNode)
	}
}

// getNextPosition returns the next position of interest after currentPos. Positions of interest
// are either places where the path forks, dead ends, or the goal position
func getNextPosition(forest [][]byte, currentPos, lastPos position) (position, position) {
	for {
		adjacentPositions := getAdjacentPositions(forest, currentPos, lastPos)
		if len(adjacentPositions) != 1 {
			break
		}

		lastPos = currentPos
		currentPos = adjacentPositions[0]

		// check for final position
		if forest[currentPos.y][currentPos.x] == '*' {
			break
		}
	}

	return currentPos, lastPos
}

// getAdjacentPositions returns a slice of possible positions to move to after currentPos, not
// counting lastPos (to prevent moving backwards)
func getAdjacentPositions(forest [][]byte, currentPos, lastPos position) []position {
	adjacentPositions := make([]position, 0)

	leftPos := position{x:currentPos.x - 1, y:currentPos.y}
	rightPos := position{x:currentPos.x + 1, y:currentPos.y}
	upPos := position{x:currentPos.x, y:currentPos.y + 1}
	downPos := position{x:currentPos.x, y:currentPos.y - 1}
	
	if isValidPosition(forest, leftPos) && leftPos != lastPos {
		adjacentPositions = append(adjacentPositions, leftPos)
	}
	if isValidPosition(forest, rightPos) && rightPos != lastPos {
		adjacentPositions = append(adjacentPositions, rightPos)
	}
	if isValidPosition(forest, upPos) && upPos != lastPos {
		adjacentPositions = append(adjacentPositions, upPos)
	}
	if isValidPosition(forest, downPos) && downPos != lastPos {
		adjacentPositions = append(adjacentPositions, downPos)
	}

	return adjacentPositions
}

// isValidPosition returns true if pos is a valid position to move within the forest (. or *)
func isValidPosition(forest [][]byte, pos position) bool {
	if pos.x >= 0 && pos.x < len(forest[0]) && pos.y >= 0 && pos.y < len(forest) {
		if forest[pos.y][pos.x] == '.' || forest[pos.y][pos.x] == '*' {
			return true
		}
	}
	return false
}

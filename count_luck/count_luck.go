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

type position struct {
	x, y int
}

type treeNode struct {
	pos position
	lastPos position
	depth int
	children []*treeNode
}

func runTest(r *bufio.Reader) {
	var height, width int
	line, _ := r.ReadString('\n')
	fmt.Sscan(line, &height, &width)

	forest := make([][]byte, height)
	start := position{}
	end := position{}

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

	var guess int
	line, _ = r.ReadString('\n')
	fmt.Sscan(line, &guess)

	root := new(treeNode)
	root.pos = start
	root.lastPos = start
	root.children = make([]*treeNode, 0)

	buildPositionTree(forest, root)
}

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

func isValidPosition(forest [][]byte, pos position) bool {
	if pos.x >= 0 && pos.x < len(forest[0]) && pos.y >= 0 && pos.y < len(forest) {
		if forest[pos.y][pos.x] == '.' || forest[pos.y][pos.x] == '*' {
			return true
		}
	}
	return false
}

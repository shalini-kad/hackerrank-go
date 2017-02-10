package main

import "fmt"

type treeNode struct {
	id int
	value int
	edges []*treeEdge
	parentNode *treeNode
}

type treeEdge struct {
	n1, n2 *treeNode
	subtreeValue int
}

func main() {
	var numNodes int
	fmt.Scan(&numNodes)

	// read nodes
	totalValue := 0
	nodes := make([]*treeNode, numNodes)
	for i := 0; i < numNodes; i++ {
		var value int
		fmt.Scan(&value)

		nodes[i] = new(treeNode)
		nodes[i].id = i
		nodes[i].value = value
		nodes[i].edges = make([]*treeEdge, 0)
		nodes[i].parentNode = nil

		totalValue += value
	}

	// read edges
	edges := make([]*treeEdge, numNodes-1)
	for i := 0; i < len(edges); i++ {
		var id1, id2 int
		fmt.Scan(&id1, &id2)
		id1 -= 1
		id2 -= 1
		
		edge := new(treeEdge)
		edge.n1 = nodes[id1]
		edge.n2 = nodes[id2]
		edge.subtreeValue = -1

		nodes[id1].edges = append(nodes[id1].edges, edge)
		nodes[id2].edges = append(nodes[id2].edges, edge)

		edges[i] = edge
	}

	findSubtreeValues(nodes[0])

	edgeDifferenceValues := make([]int, numNodes-1)
	for i := 0; i < len(edges); i++ {
		subtreeValue1 := edges[i].subtreeValue
		subtreeValue2 := totalValue - subtreeValue1
		differenceValue := abs(subtreeValue1 - subtreeValue2)
		edgeDifferenceValues[i] = differenceValue
	}

	minimumDifference := edgeDifferenceValues[0]
	for i := 1; i < len(edgeDifferenceValues); i++ {
		if edgeDifferenceValues[i] < minimumDifference {
			minimumDifference = edgeDifferenceValues[i]
		}
	}

	fmt.Println(minimumDifference)
}

func findSubtreeValues(root *treeNode) {

	children := getChildren(root)

	// base case: no child nodes
	if len(children) == 0 {
		return
	}

	// recursive case: loop through children
	for i := 0; i < len(children); i++ {
		childNode := children[i]
		childNode.parentNode = root

		findSubtreeValues(childNode)

		subtreeValue := getSubtreeValue(childNode)
		edgeToChild := getEdgeToNode(root, childNode)
		edgeToChild.subtreeValue = subtreeValue
	}
}

func getSubtreeValue(node *treeNode) int {
	totalValue := 0
	for i := 0; i < len(node.edges); i++ {
		if otherNode := getOtherNode(node, node.edges[i]); otherNode != node.parentNode {
			totalValue += node.edges[i].subtreeValue
		}
	}
	totalValue += node.value
	return totalValue
}

func getChildren(node *treeNode) []*treeNode {
	children := make([]*treeNode, 0)
	for i := 0; i < len(node.edges); i++ {
		if otherNode := getOtherNode(node, node.edges[i]); otherNode != node.parentNode {
			children = append(children, otherNode)
		}
	}
	return children
}

func getEdgeToNode(n1, n2 *treeNode) *treeEdge {
	for i := 0; i < len(n1.edges); i++ {
		if otherNode := getOtherNode(n1, n1.edges[i]); otherNode == n2 {
			return n1.edges[i]
		}
	}
	return nil
}

func getOtherNode(node *treeNode, edge *treeEdge) *treeNode {
	if edge.n1 != node {
		return edge.n1
	} else {
		return edge.n2
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
func (n *treeNode) String() string {
	return fmt.Sprintf("{%v:%v, %v}\n", n.id, n.value, n.edges)
}

func (e *treeEdge) String() string {
	return fmt.Sprintf("{%v %v->%v}", e.subtreeValue, e.n1.id, e.n2.id)
}
*/

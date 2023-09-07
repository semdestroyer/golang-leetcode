package main

type Node struct {
	Val      int
	Children []*Node
}

var rootNode Node = Node{3, []*Node{&Node1, &Node2, &Node3, &Node4}}
var Node1 Node = Node{9, nil}
var Node2 Node = Node{20, nil}
var Node3 Node = Node{15, nil}
var Node4 Node = Node{7, []*Node{&Node5}}
var Node5 Node = Node{7, nil}

func main() {
	println(maxDepth(&rootNode))
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	childHeights := make([]int, 0)
	for _, child := range root.Children {
		childHeights = append(childHeights, maxDepth(child))
	}

	max := 0

	for _, c := range childHeights {
		if c > max {
			max = c
		}
	}

	return max + 1
}

//TODO: оптимизировать(оставить один цикл)

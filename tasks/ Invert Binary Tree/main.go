package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var rootNode TreeNode = TreeNode{4, &Node1L, &Node1R}
var Node1L TreeNode = TreeNode{2, &Node2L, &Node2R}
var Node1R TreeNode = TreeNode{7, &Node3L, &Node3R}
var Node2L TreeNode = TreeNode{1, nil, nil}
var Node2R TreeNode = TreeNode{3, nil, nil}
var Node3L TreeNode = TreeNode{6, nil, nil}
var Node3R TreeNode = TreeNode{9, nil, nil}

func main() {
	invertTree(&rootNode)
}

func invertTree(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		println(root.Val)
		return root
	}
	tmpRight := root.Right
	root.Right = root.Left
	root.Left = tmpRight

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

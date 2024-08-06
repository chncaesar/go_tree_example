package main

import "go_tree_example/internal"

func main() {
	names := [][]string{{"A", "A", "C"}, {"A", "B", "D"}, {"A", "B", "E"}, {"A", "F", "G"}, {"A", "F", "H"}}
	var tree *internal.TreeNode
	for _, nameArr := range names {
		tree = internal.AddToTree(tree, nameArr)
	}
	if tree == nil || 3 != len(tree.Children) {
		panic("tree should have 3 children")
	}
}

package internal

type TreeNode struct {
	Label    string      `json:"label"`
	IsLeaf   bool        `json:"isLeaf"`
	Children []*TreeNode `json:"children"`
}

func AddToTree(tree *TreeNode, labels []string) *TreeNode {
	if tree == nil && len(labels) > 0 {
		tree = &TreeNode{Label: labels[0], Children: make([]*TreeNode, 0)}
	}
	parent := &TreeNode{Children: []*TreeNode{tree}}
	for i, mqlName := range labels {
		var find = false
		for _, child := range parent.Children {
			if child.Label == mqlName {
				parent = child
				find = true
				break
			}
		}
		if !find {
			node := &TreeNode{Label: mqlName}
			if i == len(labels)-1 {
				node.IsLeaf = true
			}
			parent.Children = append(parent.Children, node)
			parent = parent.Children[len(parent.Children)-1]
		}
	}
	return tree
}

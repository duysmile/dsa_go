package main

type TreeNode struct {
	Value        int
	Left         *TreeNode
	Right        *TreeNode
	Parent       *TreeNode
	IsParentLeft bool
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value:  value,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}

func (t *TreeNode) SetLeft(node *TreeNode) {
	t.Left = node
	if node != nil {
		node.Parent = t
		node.IsParentLeft = true
	}
}

func (t *TreeNode) SetRight(node *TreeNode) {
	t.Right = node
	if node != nil {
		node.Parent = t
		node.IsParentLeft = false
	}
}

func (t *TreeNode) GetHeight() int {
	return findMax(t.GetLeftSubtreeHeight(), t.GetRightSubtreeHeight())
}

func (t *TreeNode) GetLeftSubtreeHeight() int {
	if t.Left != nil {
		return t.Left.GetHeight() + 1
	}

	return 0
}

func (t *TreeNode) GetRightSubtreeHeight() int {
	if t.Right != nil {
		return t.Right.GetHeight() + 1
	}

	return 0
}

func (t *TreeNode) GetBalanceFactor() int {
	return t.GetLeftSubtreeHeight() - t.GetRightSubtreeHeight()
}

func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

package main

type AVLTree struct {
	BSTree
}

func (t *AVLTree) LeftRotation(node *TreeNode) *TreeNode {
	newParent := node.Right
	grandParent := node.Parent
	previousLeft := newParent.Left

	t.SwapParentChild(node, newParent, grandParent)

	newParent.SetLeft(node)
	node.SetRight(previousLeft)

	return newParent
}

func (t *AVLTree) RightRotation(node *TreeNode) *TreeNode {
	newParent := node.Left
	grandParent := node.Parent
	previousRight := newParent.Right

	t.SwapParentChild(node, newParent, grandParent)

	newParent.SetRight(node)
	node.SetLeft(previousRight)

	return newParent
}

func (t *AVLTree) LeftRightRotation(node *TreeNode) *TreeNode {
	t.LeftRotation(node.Left)
	return t.RightRotation(node)
}

func (t *AVLTree) RightLeftRotation(node *TreeNode) *TreeNode {
	t.RightRotation(node.Right)
	return t.LeftRotation(node)
}

func (t *AVLTree) SwapParentChild(oldChild, newChild, parent *TreeNode) {
	if parent != nil {
		if oldChild.IsParentLeft {
			parent.SetLeft(newChild)
		} else {
			parent.SetRight(newChild)
		}
	} else {
		newChild.Parent = nil
	}
}

func (t *AVLTree) Balance(node *TreeNode) *TreeNode {
	if node.GetBalanceFactor() > 1 {
		if node.Left.GetBalanceFactor() < 0 {
			return t.LeftRightRotation(node)
		}

		return t.RightRotation(node)
	} else if node.GetBalanceFactor() < -1 {
		if node.Right.GetBalanceFactor() > 0 {
			return t.RightLeftRotation(node)
		}

		return t.LeftRotation(node)
	}

	return node
}

func (t *AVLTree) BalanceUpstream(node *TreeNode) *TreeNode {
	current := node
	var newParent *TreeNode
	for current != nil {
		newParent = t.Balance(current)
		current = current.Parent
	}

	return newParent
}

func (t *AVLTree) Add(value int) *TreeNode {
	newNode := t.BSTree.Add(value)
	t.BSTree.Root = t.BalanceUpstream(newNode)
	return newNode
}

func (t *AVLTree) Remove(value int) bool {
	removeNode := t.BSTree.Find(value)
	if removeNode != nil {
		found := t.BSTree.Remove(value)
		t.BalanceUpstream(removeNode)
		return found
	}

	return false
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func main() {
	bst := NewAVLTree()

	bst.Add(1)
	bst.Add(3)
	bst.Add(2)
	bst.Add(5)
	bst.Add(4)
	bst.Traversal()
}

package main

import "fmt"

type BSTree struct {
	Root *TreeNode
	Size int
}

func (t *BSTree) Add(value int) *TreeNode {
	newNode := NewTreeNode(value)
	if t.Root == nil {
		t.Root = newNode
	} else {
		found, parent := t.FindNodeAndParent(value, t.Root, nil)
		if found != nil {
			fmt.Println("Value exists in this tree")
			return nil
		}

		if value < parent.Value {
			parent.SetLeft(newNode)
		} else {
			parent.SetRight(newNode)
		}
	}

	t.Size += 1
	return newNode
}

func (t *BSTree) FindNodeAndParent(value int, node, parent *TreeNode) (found, returningParent *TreeNode) {
	if node == nil || node.Value == value {
		found = node
		returningParent = parent
		return
	} else if value < node.Value {
		return t.FindNodeAndParent(value, node.Left, node)
	}
	return t.FindNodeAndParent(value, node.Right, node)
}

func (t *BSTree) Remove(value int) bool {
	nodeToRemove, parent := t.FindNodeAndParent(value, t.Root, nil)
	if nodeToRemove == nil {
		return false
	}

	removedNode := t.CombineLeftIntoRightSubtree(nodeToRemove)

	if nodeToRemove == t.Root {
		t.Root = nodeToRemove
		if nodeToRemove != nil {
			nodeToRemove.Parent = nil
		}
	} else if nodeToRemove.IsParentLeft {
		parent.SetLeft(removedNode)
	} else if !nodeToRemove.IsParentLeft {
		parent.SetRight(removedNode)
	}

	t.Size -= 1
	return true
}

func (t *BSTree) CombineLeftIntoRightSubtree(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}

	if node.Right != nil {
		leftMostNode := t.GetLeftMostNode(node.Right)
		leftMostNode.SetLeft(node.Left)
		return node.Right
	}

	return node.Left
}

func (t *BSTree) GetLeftMostNode(node *TreeNode) *TreeNode {
	tmp := node
	for tmp != nil && tmp.Left != nil {
		tmp = tmp.Left
	}

	return tmp
}

func (t *BSTree) Find(value int) *TreeNode {
	node := t.Root

	for node != nil {
		if node.Value == value {
			return node
		} else if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return nil
}

func (t *BSTree) Traversal() {
	arrNode := make([]*TreeNode, 0)
	arrNode = append(arrNode, t.Root)

	for len(arrNode) > 0 {
		tmp := make([]*TreeNode, 0)
		for _, n := range arrNode {
			var left, right int
			if n.Left != nil {
				tmp = append(tmp, n.Left)
				left = n.Left.Value
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
				right = n.Right.Value
			}
			fmt.Printf("%v has left %v and right %v\n", n.Value, left, right)
		}
		fmt.Println("===========================")
		arrNode = tmp
	}
}

func NewBSTree() *BSTree {
	return &BSTree{
		Root: nil,
	}
}

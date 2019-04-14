package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(values []int) *TreeNode {
	return MakeBinaryTree(values)
}

func MakeBinaryTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{values[0], nil, nil}

	length := len(values)
	for i := 1; i < length; i++ {
		node := &TreeNode{values[i], nil, nil}
		root.insertNode(node)
	}

	return root
}

func (root *TreeNode) insertNode(node *TreeNode) {
	if node == nil {
		return
	}

	if root.Val >= node.Val {
		if root.Left == nil {
			root.Left = node
		} else {
			root.Left.insertNode(node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			root.Right.insertNode(node)
		}
	}
}

func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{root.Val}

	if root.Left != nil {
		l := PreOrderTraversal(root.Left)
		result = append(result, l...)
	}
	if root.Right != nil {
		r := PreOrderTraversal(root.Right)
		result = append(result, r...)
	}

	return result
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int

	if root.Left != nil {
		l := PreOrderTraversal(root.Left)
		result = append(result, l...)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		r := PreOrderTraversal(root.Right)
		result = append(result, r...)
	}

	return result
}

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int

	if root.Left != nil {
		l := PreOrderTraversal(root.Left)
		result = append(result, l...)
	}

	if root.Right != nil {
		r := PreOrderTraversal(root.Right)
		result = append(result, r...)
	}

	result = append(result, root.Val)

	return result
}

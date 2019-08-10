package _108_convert_sorted_array_to_binary_search_tree

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func SortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	middle := length / 2

	return &TreeNode{
		Val:   nums[middle],
		Left:  SortedArrayToBST(nums[:middle]),
		Right: SortedArrayToBST(nums[middle+1:]),
	}
}

package ex00

import "day5/tree"

func AreToysBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return getCount(root.Left) == getCount(root.Right)
}

func getCount(root *tree.TreeNode) (ret int) {
	if root == nil {
		return
	}
	ret = getCount(root.Left) + getCount(root.Right)
	if root.HasToy {
		ret++
	}
	return
}

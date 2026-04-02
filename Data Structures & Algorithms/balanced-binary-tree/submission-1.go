/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}

	var checkHeight func(node *TreeNode) int
	checkHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := checkHeight(node.Left)
		rightHeight := checkHeight(node.Right)

		if leftHeight == -1 || rightHeight == -1 {
			return -1
		}
		if leftHeight - rightHeight > 1 || rightHeight - leftHeight > 1 {
			return -1
		}

		return max(leftHeight, rightHeight) + 1
	}

	diff := checkHeight(root)
	if diff == -1 {
		return false
	}
	return true
}

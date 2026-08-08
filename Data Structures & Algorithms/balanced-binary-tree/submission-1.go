/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    var dfs func(node *TreeNode) int

	balanced := true

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		diff := rightHeight-leftHeight

		if diff > 1 || diff < -1 {
			balanced = false
		}

		return 1 + max(leftHeight,rightHeight)
	}

	dfs(root)

	return balanced
}

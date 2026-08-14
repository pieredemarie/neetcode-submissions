/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, root.Val)

}

func dfs(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		count = 1
	}

	maxSoFar = max(maxSoFar, node.Val)

	count += dfs(node.Right, maxSoFar)
	count += dfs(node.Left, maxSoFar)

	return count
}

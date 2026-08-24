/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	var stack []*TreeNode
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		idx := len(stack) - 1
		current = stack[idx]
		stack = stack[:idx]

		result = append(result, current.Val)

		current = current.Right
	}

	return result
}

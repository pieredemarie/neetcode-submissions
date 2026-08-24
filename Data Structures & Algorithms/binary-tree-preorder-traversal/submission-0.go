/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}

	var result []int
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Val)

		
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}

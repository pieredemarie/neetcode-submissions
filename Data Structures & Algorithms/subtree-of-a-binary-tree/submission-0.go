/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
   if root == nil {
		return false
   }


   var dfs func(node1, node2 *TreeNode) bool
   dfs = func(node1,node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		return dfs(node1.Left,node2.Left) && dfs(node1.Right,node2.Right)
   }

   if dfs(root,subRoot) {
		return true
   }

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

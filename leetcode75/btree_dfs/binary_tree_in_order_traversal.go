package btree_dfs

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
	output := make([]int, 0)
	if root.Left != nil {
		output = append(output, inorderTraversal(root.Left)...)
	}
	output = append(output, root.Val)
	if root.Right != nil {
		output = append(output, inorderTraversal(root.Right)...)
	}
	return output
}

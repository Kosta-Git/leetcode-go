package btree_dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}

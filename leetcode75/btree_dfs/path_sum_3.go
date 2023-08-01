package btree_dfs

func pathSum(root *TreeNode, targetSum int) int {
    return executeOnEveryNode(root, func(node *TreeNode) int {
        return findTargetSumForNode(node, targetSum)
    })
}

func executeOnEveryNode(root *TreeNode, f func(*TreeNode) int) int {
    if root == nil {
        return 0
    }
    return f(root) + executeOnEveryNode(root.Left, f) + executeOnEveryNode(root.Right, f)
}

func findTargetSumForNode(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    if root.Val == targetSum {
        return 1
    }

    return findTargetSumForNode(root.Left, targetSum-root.Val) + findTargetSumForNode(root.Right, targetSum-root.Val)
}

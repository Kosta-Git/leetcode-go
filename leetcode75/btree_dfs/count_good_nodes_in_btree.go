package btree_dfs

func goodNodes(root *TreeNode) int {
    return goodNodesHelper(root, root.Val)
}

func goodNodesHelper(root *TreeNode, highestHistory int) int {
    if root == nil {
        return 0
    }

    isGood := 0
    if highestHistory <= root.Val {
        isGood = 1
        highestHistory = root.Val
    }

    return isGood + goodNodesHelper(root.Left, highestHistory) + goodNodesHelper(root.Right, highestHistory)
}

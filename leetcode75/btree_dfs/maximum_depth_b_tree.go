package btree_dfs

import (
    "leetcode75/leetcode75/array"
)

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }

    return 1 + array.Max(maxDepth(root.Left), maxDepth(root.Right))
}

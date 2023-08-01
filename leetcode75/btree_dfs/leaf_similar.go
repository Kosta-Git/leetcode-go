package btree_dfs

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    output1 := make([]int, 0)
    output2 := make([]int, 0)
    leafSimilarHelper(root1, &output1)
    leafSimilarHelper(root2, &output2)
    if len(output1) != len(output2) {
        return false
    }
    for i := 0; i < len(output1); i++ {
        if output1[i] != output2[i] {
            return false
        }
    }
    return true
}

func leafSimilarHelper(root *TreeNode, output *[]int) {
    if root.Left == nil && root.Right == nil {
        *output = append(*output, root.Val)
        return
    }

    if root.Left != nil {
        leafSimilarHelper(root.Left, output)
    }

    if root.Right != nil {
        leafSimilarHelper(root.Right, output)
    }
}

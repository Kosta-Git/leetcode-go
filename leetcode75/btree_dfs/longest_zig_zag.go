package btree_dfs

func longestZigZag(root *TreeNode) int {
	longest := 0
	longestZigZagDFS(root, true, 0, &longest)
	longestZigZagDFS(root, false, 0, &longest)
	return longest
}

func longestZigZagDFS(root *TreeNode, left bool, pathLength int, maxLength *int) {
	if root == nil {
		return
	}

	if pathLength > *maxLength {
		*maxLength = pathLength
	}

	if left {
		longestZigZagDFS(root.Left, false, pathLength+1, maxLength)
		longestZigZagDFS(root.Right, true, 1, maxLength)
	} else {
		longestZigZagDFS(root.Left, false, 1, maxLength)
		longestZigZagDFS(root.Right, true, pathLength+1, maxLength)
	}
}

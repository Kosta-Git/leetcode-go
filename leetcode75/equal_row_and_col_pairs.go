package leetcode75

func equalPairs(grid [][]int) int {
	sumRows := make(map[int]int, len(grid))
	for _, row := range grid {
		sumRows[hashRow(row)]++
	}
	count := 0
	for i := 0; i < len(grid[0]); i++ {
		colHash := hashCol(grid, i)
		if v, ok := sumRows[colHash]; ok {
			count += v
		}
	}
	return count
}

func hashCol(grid [][]int, col int) int {
	sum := 0
	for i := 0; i < len(grid[0]); i++ {
		sum ^= grid[i][col]
		sum *= 2166136261
	}
	return sum
}

func hashRow(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum ^= arr[i]
		sum *= 2166136261
	}
	return sum
}

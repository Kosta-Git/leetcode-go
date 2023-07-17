package leetcode75

func maxArea(height []int) int {
	maxAreaFound := 0
	l, r := 0, len(height)-1
	for width := r; width > 0; width-- {
		area := computeArea(minHeight(height[l], height[r]), width)
		if maxAreaFound < area {
			maxAreaFound = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxAreaFound
}

func minHeight(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func computeArea(height int, width int) int {
	return height * width
}

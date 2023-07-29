package leetcode75

func largestAltitude(gain []int) int {
	max := 0
	current := 0
	for _, val := range gain {
		current += val
		if current > max {
			max = current
		}
	}
	return max
}

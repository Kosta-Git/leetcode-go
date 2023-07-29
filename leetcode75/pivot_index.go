package leetcode75

func pivotIndex(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	currentSum := 0

	for i := 0; i < len(nums); i++ {
		if (sum - nums[i]) == currentSum {
			return i
		}

		currentSum += nums[i]
		sum -= nums[i]
	}

	return -1
}

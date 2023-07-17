package leetcode75

func maxOperations(nums []int, k int) int {
	count := 0
	values := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := k - nums[i]

		if complement <= 0 {
			continue
		}

		if v, ok := values[complement]; ok {
			count++
			if v > 1 {
				values[complement]--
			} else {
				delete(values, complement)
			}
		} else {
			values[nums[i]]++
		}
	}

	return count
}

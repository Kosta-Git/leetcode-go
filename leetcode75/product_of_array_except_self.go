package leetcode75

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	output[0] = 1
	for i := 0; i < len(output); i++ {
		output[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				output[j] *= nums[i]
			}
		}
	}

	return output
}

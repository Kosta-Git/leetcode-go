package leetcode75

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var maxAverage float64 = math.Inf(-1)
	for i := 0; i < len(nums)-k+1; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += nums[j]
		}
		avg := float64(sum) / float64(k)
		if maxAverage < avg {
			maxAverage = avg
		}
	}
	return maxAverage
}

package windows

import (
    "leetcode75/leetcode75/array"
    "math"
)

func findMaxAverage(nums []int, k int) float64 {
    maxAverage := 0
    for i := 0; i < k; i++ {
        maxAverage += nums[i]
    }

    windowSum := maxAverage
    for i := k; i < len(nums); i++ {
        windowSum += nums[i] - nums[i-k]
        maxAverage = array.Max(maxAverage, windowSum)
    }
    return float64(maxAverage) / float64(k)
}

func findMaxAverageBruteForce(nums []int, k int) float64 {
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

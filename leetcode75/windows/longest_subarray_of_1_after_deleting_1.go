package windows

// same as leetcode75\maxConsecutiveOne.go with k=1
// not the best solution but it works with 1 line of code
func longestSubarray(nums []int) int {
    k := 1
    max, windowStartIdx, currentZeroCount := 0, 0, 0
    for i, val := range nums {
        if val == 0 {
            currentZeroCount++
        }

        if currentZeroCount > k {
            if nums[windowStartIdx] == 0 {
                currentZeroCount--
            }
            windowStartIdx++
        }
        if i+1-windowStartIdx > max {
            max = i + 1 - windowStartIdx
        }
    }
    return max - 1
}

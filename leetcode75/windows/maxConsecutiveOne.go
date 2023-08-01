package windows

func longestOnes(nums []int, k int) int {
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
    return max
}

package array

import "math"

func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }

    minimums := make([]int, 2)
    minimums[0] = math.MaxInt
    minimums[1] = math.MaxInt
    for i := 0; i < len(nums); i++ {
        if nums[i] <= minimums[0] {
            minimums[0] = nums[i]
        } else if nums[i] <= minimums[1] {
            minimums[1] = nums[i]
        } else {
            return true
        }
    }
    return false
}

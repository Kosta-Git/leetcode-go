package two_pointers

func moveZeroes(nums []int) {
    if len(nums) == 1 {
        return
    }

    indexedZeros := indexAllZeros(nums)

    if len(indexedZeros) == len(nums) {
        return
    }

    shiftRemoveAllZeros(nums, indexedZeros)

    for i := len(nums) - len(indexedZeros); i < len(nums); i++ {
        nums[i] = 0
    }
}

func indexAllZeros(nums []int) []int {
    zeros := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            zeros = append(zeros, i)
        }
    }
    return zeros
}

func shiftRemoveAllZeros(nums []int, indexedZeros []int) {
    for i := 0; i < len(indexedZeros); i++ {
        end := len(nums) - len(indexedZeros)
        if i+1 < len(indexedZeros) {
            end = indexedZeros[i+1]
        }
        for j := indexedZeros[i] - i; j < end; j++ {
            if j+i+1 >= len(nums) {
                continue
            }
            nums[j] = nums[j+i+1]
        }
    }
}

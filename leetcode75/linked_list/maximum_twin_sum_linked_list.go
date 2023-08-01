package linked_list

import (
    "leetcode75/intro"
)

func pairSum(head *intro.ListNode) int {
    it := head
    array := toArray(it)

    max := 0
    currentMax := 0
    for i := 0; i < len(array)/2; i++ {
        currentMax = array[i] + array[len(array)-1-i]
        if max < currentMax {
            max = currentMax
        }
    }
    return max
}

func toArray(head *intro.ListNode) []int {
    var result []int
    for head != nil {
        result = append(result, head.Val)
        head = head.Next
    }
    return result
}

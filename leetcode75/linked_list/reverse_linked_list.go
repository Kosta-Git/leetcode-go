package linked_list

import (
    "leetcode75/intro"
)

func reverseList(head *intro.ListNode) *intro.ListNode {
    var prev *intro.ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev

        prev = curr
        curr = next
    }

    return prev
}

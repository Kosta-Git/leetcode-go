package linked_list

import (
    "leetcode75/intro"
)

func oddEvenList(head *intro.ListNode) *intro.ListNode {
    if head == nil {
        return nil
    }

    var oddNode, evenNode, oddTail, evenTail *intro.ListNode
    isOdd := true

    for head != nil {
        next := head.Next
        head.Next = nil

        if isOdd {
            if oddNode == nil {
                oddNode = head
                oddTail = head
            } else {
                oddTail.Next = head
                oddTail = head
            }
        } else {
            if evenNode == nil {
                evenNode = head
                evenTail = head
            } else {
                evenTail.Next = head
                evenTail = head
            }
        }

        head = next
        isOdd = !isOdd
    }
    oddTail.Next = evenNode
    return oddNode
}

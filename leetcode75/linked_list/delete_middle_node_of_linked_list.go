package linked_list

import (
    "leetcode75/intro"
)

func deleteMiddle(head *intro.ListNode) *intro.ListNode {
    it := head
    counter := 1
    for it.Next != nil {
        counter++
        it = it.Next
    }

    if counter == 1 {
        return nil
    }

    if counter == 2 {
        head.Next = nil
        return head
    }

    it = head
    for i := 0; i < (counter/2)-1; i++ {
        it = it.Next
    }

    it.Next = it.Next.Next

    return head
}

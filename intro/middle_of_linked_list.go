package intro

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val  int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    it := head
    counter := 1
    for it.Next != nil {
        counter++
        it = it.Next
    }

    it = head
    for i := 0; i < counter/2; i++ {
        it = it.Next
    }
    return it
}

package linked_list

import (
    "leetcode75/intro"
    "testing"
)

func buildNodesFromArray(array []int) *intro.ListNode {
    var head *intro.ListNode
    var tail *intro.ListNode

    for _, v := range array {
        node := &intro.ListNode{
            Val: v,
        }

        if head == nil {
            head = node
            tail = node
        } else {
            tail.Next = node
            tail = node
        }
    }

    return head
}

func Test_pairSum(t *testing.T) {
    type args struct {
        head *intro.ListNode
    }
    tests := []struct {
        name string
        args args
        want int
    }{

        {
            name: "1",
            args: args{
                head: buildNodesFromArray([]int{1, 2, 3, 4, 5}),
            },
            want: 6,
        },
        {
            name: "2",
            args: args{
                head: buildNodesFromArray([]int{4, 2, 2, 3}),
            },
            want: 7,
        },
        {
            name: "3",
            args: args{
                head: buildNodesFromArray([]int{47, 22, 81, 46, 94, 95, 90, 22, 55, 91, 6, 83, 49, 65, 10, 32, 41, 26, 83, 99, 14, 85, 42, 99, 89, 69, 30, 92, 32, 74, 9, 81, 5, 9}),
            },
            want: 182,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := pairSum(tt.args.head); got != tt.want {
                t.Errorf("pairSum() = %v, want %v", got, tt.want)
            }
        })
    }
}

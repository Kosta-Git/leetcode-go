package linked_list

import (
    "leetcode75/intro"
    "reflect"
    "testing"
)

func Test_deleteMiddle(t *testing.T) {
    type args struct {
        head *intro.ListNode
    }
    tests := []struct {
        name string
        args args
        want *intro.ListNode
    }{
        {
            name: "test1",
            args: args{
                head: &intro.ListNode{
                    Val: 1,
                    Next: &intro.ListNode{
                        Val: 2,
                        Next: &intro.ListNode{
                            Val: 3,
                            Next: &intro.ListNode{
                                Val: 4,
                                Next: &intro.ListNode{
                                    Val: 5,
                                },
                            },
                        },
                    },
                },
            },
            want: &intro.ListNode{
                Val: 1,
                Next: &intro.ListNode{
                    Val: 2,
                    Next: &intro.ListNode{
                        Val: 4,
                        Next: &intro.ListNode{
                            Val: 5,
                        },
                    },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
            }
        })
    }
}

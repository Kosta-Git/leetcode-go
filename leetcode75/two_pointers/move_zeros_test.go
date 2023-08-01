package two_pointers

import "testing"

func Test_moveZeroes(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "Example 1",
            args: args{
                nums: []int{0, 1, 0, 3, 12},
            },
            want: []int{1, 3, 12, 0, 0},
        },
        {
            name: "Example 2",
            args: args{
                nums: []int{0},
            },
            want: []int{0},
        },
        {
            name: "Example 3",
            args: args{
                nums: []int{0, 0, 0, 1, 5, 0, 0, 0, 0, 0, 0, 9},
            },
            want: []int{1, 5, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            moveZeroes(tt.args.nums)
            if len(tt.args.nums) != len(tt.want) {
                t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
            }
            for i := 0; i < len(tt.args.nums); i++ {
                if tt.args.nums[i] != tt.want[i] {
                    t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
                    break
                }
            }
        })
    }
}

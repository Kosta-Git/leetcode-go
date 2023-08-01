package two_pointers

import (
    "math/rand"
    "testing"
    "time"
)

func Test_maxArea(t *testing.T) {
    type args struct {
        height []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "Example 1",
            args: args{
                height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
            },
            want: 49,
        },
        {
            name: "Example 2",
            args: args{
                height: []int{1, 1},
            },
            want: 1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxArea(tt.args.height); got != tt.want {
                t.Errorf("maxArea() = %v, want %v", got, tt.want)
            }
        })
    }
}

var arr = generateRandomArray(10000, 0, 1000)

func Benchmark_maxArea(b *testing.B) {
    for i := 0; i < b.N; i++ {
        maxArea(arr)
    }
}

func generateRandomArray(length, min, max int) []int {
    rand.Seed(time.Now().UnixNano())

    arr := make([]int, length)
    for i := 0; i < length; i++ {
        arr[i] = rand.Intn(max-min+1) + min
    }

    return arr
}

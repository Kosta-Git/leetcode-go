package windows

import "testing"

func Test_findMaxAverage(t *testing.T) {
    type args struct {
        nums []int
        k    int
    }
    tests := []struct {
        name string
        args args
        want float64
    }{
        {
            name: "test 1",
            args: args{
                nums: []int{1, 12, -5, -6, 50, 3},
                k:    4,
            },
            want: 12.75,
        },
        {
            name: "test 2",
            args: args{
                nums: []int{5},
                k:    1,
            },
            want: 5,
        },
        {
            name: "test 3",
            args: args{
                nums: []int{-1},
                k:    1,
            },
            want: -1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
                t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Benchmark(b *testing.B) {
    type args struct {
        nums []int
        k    int
    }
    benchmarks := []struct {
        name string
        args args
        want float64
    }{
        {
            name: "test 1",
            args: args{
                nums: []int{1, 12, -5, -6, 50, 3},
                k:    4,
            },
            want: 12.75,
        },
    }
    for _, bm := range benchmarks {
        b.Run(bm.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                findMaxAverage(bm.args.nums, bm.args.k)
            }
        })
    }
}

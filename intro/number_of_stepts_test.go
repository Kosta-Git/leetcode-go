package intro

import "testing"

func Test_numberOfSteps(t *testing.T) {
    type args struct {
        num int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "test 1",
            args: args{
                num: 14,
            },
            want: 6,
        },
        {
            name: "test 2",
            args: args{
                num: 8,
            },
            want: 4,
        },
        {
            name: "test 3",
            args: args{
                num: 123,
            },
            want: 12,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := numberOfSteps(tt.args.num); got != tt.want {
                t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
            }
        })
    }
}

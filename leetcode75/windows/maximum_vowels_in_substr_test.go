package windows

import "testing"

func Test_maxVowels(t *testing.T) {
    type args struct {
        s string
        k int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "3 in a row",
            args: args{
                s: "abciiidef",
                k: 3,
            },
            want: 3,
        },
        {
            name: "2 in a row",
            args: args{
                s: "aeiou",
                k: 2,
            },
            want: 2,
        },
        {
            name: "Less than k",
            args: args{
                s: "leetcode",
                k: 3,
            },
            want: 2,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
                t.Errorf("maxVowels() = %v, want %v", got, tt.want)
            }
        })
    }
}

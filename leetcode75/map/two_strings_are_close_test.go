package _map

import "testing"

func Test_closeStrings(t *testing.T) {
    type args struct {
        word1 string
        word2 string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "Example 1:",
            args: args{
                word1: "abc",
                word2: "bca",
            },
            want: true,
        },
        {
            name: "Example 2:",
            args: args{
                word1: "a",
                word2: "aa",
            },
            want: false,
        },
        {
            name: "Example 3:",
            args: args{
                word1: "cabbba",
                word2: "abbccc",
            },
            want: true,
        },
        {
            name: "Example 4:",
            args: args{
                word1: "aabbcccddd",
                word2: "abccccdddd",
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := closeStrings(tt.args.word1, tt.args.word2); got != tt.want {
                t.Errorf("closeStrings() = %v, want %v", got, tt.want)
            }
        })
    }
}

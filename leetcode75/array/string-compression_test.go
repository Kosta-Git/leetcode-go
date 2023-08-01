package array

import "testing"

func Test_compress(t *testing.T) {
    type args struct {
        chars []byte
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "test1",
            args: args{
                chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
            },
            want: 6,
        },
        {
            name: "test2",
            args: args{
                chars: []byte{'a'},
            },
            want: 1,
        },
        {
            name: "test3",
            args: args{
                chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
            },
            want: 4,
        },
        {
            name: "test4",
            args: args{
                chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
            },
            want: 6,
        },
        {
            name: "test5",
            args: args{
                chars: []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
            },
            want: 6,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := compress(tt.args.chars); got != tt.want {
                t.Errorf("compress() = %v, want %v", got, tt.want)
            }
        })
    }
}

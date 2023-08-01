package queue

import "testing"

func Test_predictPartyVictory(t *testing.T) {
    type args struct {
        senate string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "test1",
            args: args{
                senate: "RD",
            },
            want: "Radiant",
        },
        {
            name: "test2",
            args: args{
                senate: "RDD",
            },
            want: "Dire",
        },
        {
            name: "test3",
            args: args{
                senate: "DDRRR",
            },
            want: "Dire",
        },
        {
            name: "test4",
            args: args{
                senate: "DDDRDRRDRRDRDRRRDDRRDDDRDRDDDRRRRDDDDRDRRRRDRRRDRDRDDRDRRRRDRDRRRDRDDDRRDDDRDRDRDRRDRDDRDDRDDDDRDRRR",
            },
            want: "Radiant",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := predictPartyVictory(tt.args.senate); got != tt.want {
                t.Errorf("predictPartyVictory() = %v, want %v", got, tt.want)
            }
        })
    }
}

package btree_dfs

import (
    "testing"
)

func Test_goodNodes(t *testing.T) {
    type args struct {
        root *TreeNode
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "1",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := goodNodes(tt.args.root); got != tt.want {
                t.Errorf("goodNodes() = %v, want %v", got, tt.want)
            }
        })
    }
}

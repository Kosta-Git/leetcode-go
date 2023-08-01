package btree_dfs

import (
    "testing"
)

func Test_pathSum(t *testing.T) {
    type args struct {
        root      *TreeNode
        targetSum int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
                t.Errorf("pathSum() = %v, want %v", got, tt.want)
            }
        })
    }
}

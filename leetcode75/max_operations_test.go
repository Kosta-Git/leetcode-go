package leetcode75

import "testing"

func Test_maxOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    5,
			},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{3, 1, 3, 4, 3},
				k:    6,
			},
			want: 1,
		},
		{
			name: "test 3",
			args: args{
				nums: []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2},
				k:    3,
			},
			want: 4,
		},
		{
			name: "test 4",
			args: args{
				nums: []int{64, 35, 69, 79, 76, 60, 58, 38, 3, 81, 74, 9, 77, 21, 54, 54, 14, 72},
				k:    47,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

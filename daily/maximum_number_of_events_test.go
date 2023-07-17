package daily

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		events [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				events: [][]int{
					{1, 2, 4},
					{3, 4, 3},
					{2, 3, 1},
				},
				k: 2,
			},
			want: 7,
		},
		{
			name: "test 2",
			args: args{
				events: [][]int{
					{1, 2, 4},
					{3, 4, 3},
					{2, 3, 10},
				},
				k: 2,
			},
			want: 10,
		},
		{
			name: "test 3",
			args: args{
				events: [][]int{
					{1, 1, 1},
					{2, 2, 2},
					{3, 3, 3},
					{4, 4, 4},
				},
				k: 3,
			},
			want: 9,
		},
		{
			name: "test 4",
			args: args{
				events: [][]int{
					{31, 57, 53},
					{5, 63, 29},
					{54, 57, 32},
					{55, 83, 28},
					{25, 64, 5},
					{5, 33, 33},
					{32, 68, 27},
					{30, 99, 54},
				},
				k: 4,
			},
			want: 65,
		},
		{
			name: "test 5",
			args: args{
				events: [][]int{
					{30, 40, 34},
					{6, 11, 6},
					{60, 81, 36},
				},
				k: 1,
			},
			want: 36,
		},
		{
			name: "test 6",
			args: args{
				events: [][]int{
					{11, 17, 56},
					{24, 40, 53},
					{5, 62, 67},
					{66, 69, 84},
					{56, 89, 15},
				},
				k: 2,
			},
			want: 151,
		},
		{
			name: "test 7",
			args: args{
				events: [][]int{
					{21, 77, 43},
					{2, 74, 47},
					{6, 59, 22},
					{47, 47, 38},
					{13, 74, 57},
					{27, 55, 27},
					{8, 15, 8},
				},
				k: 4,
			},
			want: 57,
		},
		{
			name: "test 8",
			args: args{
				events: [][]int{
					{40, 73, 65},
					{57, 79, 80},
					{51, 57, 19},
					{33, 99, 77},
					{47, 85, 44},
					{18, 56, 11},
					{19, 78, 87},
					{23, 80, 37},
					{26, 48, 74},
					{45, 87, 80},
					{10, 12, 52},
					{23, 83, 43},
					{48, 84, 42},
					{23, 39, 3},
				},
				k: 13,
			},
			want: 206,
		}, {
			name: "test 9",
			args: args{
				events: [][]int{
					{3, 68, 97},
					{12, 46, 13},
					{21, 24, 75},
					{64, 85, 74},
					{10, 98, 15},
					{23, 84, 62},
					{87, 96, 29},
					{80, 85, 39},
					{52, 89, 77},
					{31, 63, 91},
					{29, 40, 48},
					{30, 96, 42},
					{69, 81, 68},
					{52, 58, 65},
					{41, 52, 37},
				},
				k: 10,
			},
			want: 291,
		},
		{
			name: "test 10",
			args: args{
				events: [][]int{
					{44, 81, 11},
					{31, 41, 14},
					{85, 90, 5},
					{42, 78, 68},
					{46, 61, 11},
					{23, 92, 6},
					{64, 67, 70},
					{77, 92, 32},
					{5, 26, 87},
					{10, 44, 78},
					{18, 90, 40},
					{15, 65, 66},
					{1, 10, 29},
					{36, 80, 80},
					{3, 94, 5},
					{2, 62, 90},
					{19, 28, 47},
				},
				k: 16,
			},
			want: 214,
		},
		{
			name: "test 11",
			args: args{
				events: [][]int{
					{42, 66, 44},
					{36, 86, 18},
					{75, 76, 16},
					{75, 83, 80},
					{10, 78, 1},
					{6, 21, 99},
					{35, 70, 83},
					{55, 100, 39},
					{46, 76, 98},
					{20, 68, 47},
					{59, 69, 32},
					{58, 80, 22},
					{62, 84, 75},
					{4, 70, 75},
					{72, 73, 25},
					{34, 76, 88},
					{74, 93, 77},
					{58, 90, 43},
					{21, 39, 21},
					{39, 41, 41},
				},
				k: 17,
			},
			want: 289,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.events, tt.args.k); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	type args struct {
		events [][]int
		k      int
	}
	benchmarks := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 11",
			args: args{
				events: [][]int{
					{42, 66, 44},
					{36, 86, 18},
					{75, 76, 16},
					{75, 83, 80},
					{10, 78, 1},
					{6, 21, 99},
					{35, 70, 83},
					{55, 100, 39},
					{46, 76, 98},
					{20, 68, 47},
					{59, 69, 32},
					{58, 80, 22},
					{62, 84, 75},
					{4, 70, 75},
					{72, 73, 25},
					{34, 76, 88},
					{74, 93, 77},
					{58, 90, 43},
					{21, 39, 21},
					{39, 41, 41},
				},
				k: 17,
			},
			want: 289,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				maxValue(bm.args.events, bm.args.k)
			}
		})
	}
}

package daily

import (
	"sort"
)

func find(events [][]int, start int) int {
	l, r := 0, len(events)
	for l < r {
		mid := (l + r) / 2
		if events[mid][1] < start {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}

// Chat GPT solution, this solution works using DP
// Used for reference purpose
func maxValueGTP(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	dp := make([][]int, len(events)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= len(events); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = Max(dp[i][j], dp[i-1][j])
			idx := find(events, events[i-1][0])
			if idx != -1 {
				dp[i][j] = Max(dp[i][j], dp[idx+1][j-1]+events[i-1][2])
			} else {
				dp[i][j] = Max(dp[i][j], events[i-1][2])
			}
		}
	}
	return dp[len(events)][k]
}

package leetcode75

func numberOfSteps(num int) int {
	step := 0
	for num > 0 {
		step++
		if num&1 == 0 {
			num >>= 1
		} else {
			num ^= 1
		}
	}
	return step
}

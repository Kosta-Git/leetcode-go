package leetcode75

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	for _, val := range arr {
		occurrences[val]++
	}
	foundOccurrences := make(map[int]bool)
	for _, val := range occurrences {
		if foundOccurrences[val] {
			return false
		}
		foundOccurrences[val] = true
	}
	return true
}

package leetcode75

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	output := strings.Builder{}
	for i := 0; i < Max(len(word1), len(word2)); i++ {
		if i < len(word1) {
			output.WriteByte(word1[i])
		}
		if i < len(word2) {
			output.WriteByte(word2[i])
		}
	}
	return output.String()
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

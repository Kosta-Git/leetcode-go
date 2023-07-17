package leetcode75

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	minimumLength := Min(len(str1), len(str2))
	var smallestString string
	if len(str1) < len(str2) {
		smallestString = str1
	} else {
		smallestString = str2
	}

	for i := 0; i < minimumLength; i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}

	patternBuilder := strings.Builder{}
	patternFound := ""
	for i := 0; i < minimumLength; i++ {
		patternBuilder.WriteByte(smallestString[i])
		if IsPatternOf(patternBuilder.String(), str1) && IsPatternOf(patternBuilder.String(), str2) {
			patternFound = patternBuilder.String()
		}
	}
	return patternFound
}

func IsPatternOf(pattern string, str string) bool {
	if len(str)%len(pattern) != 0 {
		return false
	}

	for i := 0; i < len(str); i += len(pattern) {
		if str[i:i+len(pattern)] != pattern {
			return false
		}
	}

	return true
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

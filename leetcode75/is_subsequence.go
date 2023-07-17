package leetcode75

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	sIndex := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[sIndex] {
			sIndex++
		}

		if sIndex == len(s) {
			return true
		}
	}

	return false
}

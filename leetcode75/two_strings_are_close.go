package leetcode75

func closeStrings(word1 string, word2 string) bool {
	freq1, freq2 := freq(word1), freq(word2)

	if len(freq1) != len(freq2) {
		return false
	}

	for v := range freq1 {
		if _, ok := freq2[v]; !ok {
			return false
		}
	}

	values := make(map[int]int, len(freq1))
	for v := range freq1 {
		if v, ok := freq1[v]; ok {
			values[v]++
		}
		if v, ok := freq2[v]; ok {
			values[v]--
		}
	}

	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

func freq(s string) map[rune]int {
	f := make(map[rune]int)
	for _, ch := range s {
		f[ch]++
	}
	return f
}

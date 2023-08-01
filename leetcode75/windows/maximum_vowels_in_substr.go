package windows

import "strings"

func maxVowels(s string, k int) int {
    max := 0
    currentVowels := 0
    const vowels = "aeiou"
    for i := 0; i < len(s); i++ {
        currentIsVowel := strings.Contains(vowels, string(s[i]))

        if currentIsVowel {
            currentVowels++
        }

        if i >= k {
            if strings.Contains(vowels, string(s[i-k])) {
                currentVowels--
            }
        }

        if currentVowels > max {
            max = currentVowels
        }
    }

    return max
}

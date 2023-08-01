package array

import (
    "strings"
)

func reverseVowels(s string) string {
    vowels := "aeiouAEIOU"
    output := s
    indices := make([]int, 0)

    for i := 0; i < len(s); i++ {
        if strings.Contains(vowels, s[i:i+1]) {
            indices = append(indices, i)
        }
    }

    if len(indices) > 1 {
        for i := 0; i < len(indices)/2; i++ {
            output = output[:indices[i]] + s[indices[len(indices)-1-i]:indices[len(indices)-1-i]+1] + output[indices[i]+1:]
            output = output[:indices[len(indices)-1-i]] + s[indices[i]:indices[i]+1] + output[indices[len(indices)-1-i]+1:]
        }
    }

    return output
}

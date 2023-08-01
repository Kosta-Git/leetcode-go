package array

import "strings"

func reverseWords(s string) string {
    str := strings.TrimSpace(s)
    words := strings.Split(str, " ")
    builder := strings.Builder{}
    for i := len(words) - 1; i >= 0; i-- {
        if strings.TrimSpace(words[i]) == "" {
            continue
        }
        builder.WriteString(strings.TrimSpace(words[i]) + " ")
    }
    return strings.TrimSpace(builder.String())
}

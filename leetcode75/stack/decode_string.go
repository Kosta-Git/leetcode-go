package stack

import (
    "strings"
)

type encoded struct {
    Times int
    Str   string
}

func decodeString(s string) string {
    builder := strings.Builder{}
    var stack []encoded
    lastWasDigit := false
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            if lastWasDigit {
                stack[len(stack)-1].Times = stack[len(stack)-1].Times*10 + int(s[i]-'0')
            } else {
                stack = append(stack, encoded{
                    Times: int(s[i] - '0'),
                    Str:   "",
                })
                lastWasDigit = true
            }
            continue

        } else if s[i] == '[' {
            // Do nothing
        } else if s[i] == ']' {
            if len(stack) == 1 {
                builder.WriteString(strings.Repeat(stack[len(stack)-1].Str, stack[len(stack)-1].Times))
                stack = stack[:len(stack)-1]
            } else {
                stack[len(stack)-2].Str += strings.Repeat(stack[len(stack)-1].Str, stack[len(stack)-1].Times)
                stack = stack[:len(stack)-1]
            }
        } else {
            if len(stack) == 0 {
                builder.WriteByte(s[i])
            } else {
                stack[len(stack)-1].Str += string(s[i])
            }
        }
        lastWasDigit = false
    }

    return builder.String()
}

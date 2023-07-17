package leetcode75

import (
	"strconv"
	"strings"
)

func fizzBuzz(n int) []string {
	builder := strings.Builder{}
	output := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			builder.WriteString("Fizz")
		}
		if i%5 == 0 {
			builder.WriteString("Buzz")
		}
		if builder.Len() == 0 {
			builder.WriteString(strconv.Itoa(i))
		}
		output = append(output, builder.String())
		builder.Reset()
	}

	return output
}

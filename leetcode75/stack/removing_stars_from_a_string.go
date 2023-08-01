package stack

func removeStars(s string) string {
    var b []byte
    for c := range s {
        if s[c] != '*' {
            b = append(b, s[c])
            continue
        }

        if len(b) == 0 {
            continue
        }

        b = b[:len(b)-1]
    }
    return string(b)
}

package array

func compress(chars []byte) int {
    currentChar := chars[0]
    currentCount := 0
    startIndex := 0
    outputSize := 0
    for i := 0; i < len(chars); i++ {
        if chars[i] == currentChar {
            currentCount++
            continue
        }

        chars[startIndex] = currentChar
        outputSize++
        if currentCount > 1 {
            outputSize += numberToBytes(currentCount, startIndex+1, chars)
        }

        currentChar = chars[i]
        currentCount = 1
        startIndex = outputSize
    }

    chars[startIndex] = currentChar
    outputSize++
    if currentCount > 1 {
        outputSize += numberToBytes(currentCount, startIndex+1, chars)
    }

    chars = chars[:outputSize]
    return outputSize
}

func numberToBytes(num int, index int, chars []byte) int {
    bytes := []byte{}
    for num > 0 {
        bytes = append(bytes, byte(rune(num%10+48)))
        num /= 10
    }
    for i := 0; i < len(bytes); i++ {
        chars[index+i] = bytes[len(bytes)-1-i]
    }
    return len(bytes)
}

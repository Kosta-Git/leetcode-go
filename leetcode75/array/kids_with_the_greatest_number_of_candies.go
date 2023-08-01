package array

func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := maxInArray(candies)
    output := make([]bool, len(candies))

    for i := 0; i < len(candies); i++ {
        output[i] = candies[i]+extraCandies >= maxCandies
    }

    return output
}

func maxInArray(array []int) int {
    max := 0
    for _, value := range array {
        if value > max {
            max = value
        }
    }
    return max
}

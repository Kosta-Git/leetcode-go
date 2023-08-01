package array

func canPlaceFlowers(flowerbed []int, n int) bool {
    canBePlaced := 0

    if len(flowerbed) == 1 && flowerbed[0] == 0 {
        canBePlaced++
    }

    if len(flowerbed) > 1 && flowerbed[0] == 0 && flowerbed[1] == 0 {
        canBePlaced++
    }

    if len(flowerbed) > 2 && flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
        canBePlaced++
    }

    if len(flowerbed) > 2 {
        zeroStreakCounter := 0
        for i := 1; i < len(flowerbed)-1; i++ {
            if flowerbed[i] == 0 {
                zeroStreakCounter += 1
            } else {
                canBePlaced += (zeroStreakCounter - 1) / 2
                zeroStreakCounter = 0
            }
        }
        canBePlaced += (zeroStreakCounter - 1) / 2
    }

    return canBePlaced >= n
}

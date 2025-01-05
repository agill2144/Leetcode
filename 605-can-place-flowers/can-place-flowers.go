func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := 0; i < len(flowerbed) && n > 0; i++ {
        prev := 0
        next := 0
        if i+1 < len(flowerbed) {next = flowerbed[i+1]}
        if i-1 >= 0 {prev = flowerbed[i-1]}

        if prev == 0 && next == 0 && flowerbed[i] != 1 {
            flowerbed[i] = 1
            n--
        }
    }
    return n == 0
}
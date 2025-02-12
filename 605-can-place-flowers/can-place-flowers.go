func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {return true}
    for i := 0; i < len(flowerbed); i++ {
        prev := 0
        next := 0
        if i-1 >= 0 {prev = flowerbed[i-1]}
        if i+1 < len(flowerbed) {next = flowerbed[i+1]}
        if prev == 1 || next == 1 {continue}
        if flowerbed[i] == 0 && n > 0 {
            n--
            flowerbed[i] = 1
        }
    }
    return n==0
}
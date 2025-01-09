func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {return true}
    for i := 0; i < len(flowerbed); i++ {
        curr := flowerbed[i]
        if curr == 1 {continue}
        prev := 0
        if i-1 >= 0 {prev = flowerbed[i-1]}
        next := 0
        if i+1 < len(flowerbed) {next = flowerbed[i+1]}
        if prev == 0 && next == 0 {
            flowerbed[i] = 1
            n--
            if n == 0 {break}
        }
    }
    return n == 0
}
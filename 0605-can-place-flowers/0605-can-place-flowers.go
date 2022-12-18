func canPlaceFlowers(flowerbed []int, n int) bool {
    if n <= 0 {return true}
    if len(flowerbed) == 1 && flowerbed[0] != 1 && n == 1 {return true}

    for i := 0; i < len(flowerbed); i++ {
        if n == 0 {return true}
        curr := flowerbed[i]        
        prevIsEmpty := (i == 0) || (flowerbed[i-1] == 0)
        nextIsEmpty := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)
        if curr != 1 && prevIsEmpty && nextIsEmpty {
            flowerbed[i] = 1
            n--
        }
    }
    return n==0
}


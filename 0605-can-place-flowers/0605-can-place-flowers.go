func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {return true}
    if flowerbed == nil || len(flowerbed) == 0 {return false}
    
    for i := 0; i < len(flowerbed); i++ {
        if n == 0 {return true}
        prevEmpty := (i == 0)|| (flowerbed[i-1] == 0)
        nextEmpty := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)
        currEmpty := flowerbed[i] == 0
        
        if currEmpty && prevEmpty && nextEmpty {
            flowerbed[i] = 1
            n--
        }
    }
    return n == 0
}
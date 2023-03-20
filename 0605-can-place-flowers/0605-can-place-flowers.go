func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {return true}
    
    for i := 0; i < len(flowerbed); i++ {
        isEmpty := flowerbed[i] == 0
        leftEmpty := i == 0 || flowerbed[i-1] == 0
        rightEmpty := i == len(flowerbed)-1 || flowerbed[i+1] == 0
        
        if isEmpty && leftEmpty && rightEmpty {
            flowerbed[i] = 1
            n--
            if n == 0 {return true}
        }
    }
    return n == 0
}
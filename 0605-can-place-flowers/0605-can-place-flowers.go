func canPlaceFlowers(flowerbed []int, n int) bool {
    if n <= 0 {return true}
    if len(flowerbed) == 1 && flowerbed[0] != 1 && n == 1 {return true}

    for i := 0; i < len(flowerbed); i++ {
        if n == 0 {return true}
        curr := flowerbed[i]
        
        if i == 0 || i == len(flowerbed)-1 {
            if (i == 0 && curr != 1 && flowerbed[i+1] != 1) || (i == len(flowerbed)-1 && curr != 1 && flowerbed[i-1] != 1) {
                flowerbed[i] = 1
                n--
            }
            continue
        }
                
        next := flowerbed[i+1]
        prev := flowerbed[i-1]
        if curr != 1 && prev != 1 && next != 1 {
            flowerbed[i] = 1
            n--
        }
    }
    return n==0
}


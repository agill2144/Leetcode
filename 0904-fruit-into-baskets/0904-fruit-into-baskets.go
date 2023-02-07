func totalFruit(fruits []int) int {
    basket := map[int]int{}
    left := 0
    count := 0
    maxCount := 0
    for i := 0; i < len(fruits); i++ {
        basket[fruits[i]]++
        count++
        for len(basket) > 2 {
            leftFruit := fruits[left]
            basket[leftFruit]--
            if val := basket[leftFruit]; val == 0 {
                delete(basket,leftFruit)
            }
            count--
            left++
        }
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}
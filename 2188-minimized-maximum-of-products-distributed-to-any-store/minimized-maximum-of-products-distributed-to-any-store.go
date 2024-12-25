func minimizedMaximum(n int, quantities []int) int {
    left := 1
    right := -1
    for i := 0; i < len(quantities); i++ {
        right= max(right, quantities[i])
    }
    ans := -1
    for left <= right {
        mid := left +(right-left)/2
        countPerStore := mid
        storesNeeded := 0
        for j := 0; j < len(quantities); j++ {
            storesNeeded += int(math.Ceil(float64(quantities[j])/float64(countPerStore)))
        }
        // when does mid not work?
        // when we need more stores than we have (n)
        if storesNeeded > n {
            // increase count per product type
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}
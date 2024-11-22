func minimizedMaximum(n int, quantities []int) int {
    m := len(quantities)
    if n < m {return -1}
    left := 1
    right := -1
    for i := 0; i < m; i++ {right = max(right, quantities[i])}
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < m; i++ {
            count += int(math.Ceil( float64(quantities[i])/float64(mid)))
        }
        // when does mid not work?
        // when count num of stores needed is more than num of stores we have
        if count > n {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}
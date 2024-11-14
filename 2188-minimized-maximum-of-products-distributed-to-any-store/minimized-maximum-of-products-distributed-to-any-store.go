func minimizedMaximum(n int, quantities []int) int {
    left := 1
    right := -1
    for i := 0; i < len(quantities); i++ {right = max(right, quantities[i])}
    ans := -1
    // minimize the maximum number ( binary search on answers hint )
    // max = mid
    // find the smallest such max
    // mid works when stores needed <= n
    // if mid works, we want smallest such mid value
    // therefore take the search on left side
    // otherwise, we needed more store for our current mid (atMax)
    // which means, our atMax (or mid) was too small that more stores were needed to distribute all products
    // therefore search right
    // n = len(quantities)
    // q = max(quantities)
    // time = o(n) + o(log q * n)
    // space = o(1)
    for left <= right {
        mid := left + (right-left)/2       
        atMax := mid
        count := 0
        for i := 0; i < len(quantities); i++ {
            count += int(math.Ceil(float64(quantities[i])/float64(atMax)))
            if count > n {break}
        }
        if count <= n {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
func minimizedMaximum(n int, quantities []int) int {
    left := 1
    right := slices.Max(quantities)
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        stores := 0
        for i := 0; i < len(quantities); i++ {
            stores += int(math.Ceil(float64(quantities[i])/float64(mid)))
            if stores > n {break}
        }
        if stores > n {
            left = mid+1
        } else {
            res = mid
            right = mid-1
        }
    }
    return res
}
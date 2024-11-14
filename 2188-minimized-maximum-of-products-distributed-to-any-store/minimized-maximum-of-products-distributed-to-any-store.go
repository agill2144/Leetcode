func minimizedMaximum(n int, quantities []int) int {
    left := 1
    right := -1
    for i := 0; i < len(quantities); i++ {right = max(right, quantities[i])}
    ans := -1
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
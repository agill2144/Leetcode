func maxLength(ribbons []int, k int) int {
    left := 1
    right := -1

    for i := 0; i < len(ribbons); i++ {right = max(right, ribbons[i])}
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < len(ribbons); i++ {
            count += (ribbons[i]/mid)
            if count >= k {break}
        }
        if count >= k {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}
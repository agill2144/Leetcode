func maxLength(ribbons []int, k int) int {
    left := 1
    right := slices.Max(ribbons)
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < len(ribbons); i++ {
            count += int(math.Floor(float64(ribbons[i])/float64(mid)))
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
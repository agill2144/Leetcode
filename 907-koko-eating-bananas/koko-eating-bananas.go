func minEatingSpeed(piles []int, h int) int {
    left := 1
    right := slices.Max(piles)
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        hours := 0
        for i := 0; i < len(piles); i++ {
            hours += int(math.Ceil(float64(piles[i])/float64(mid)))
        }
        if hours > h {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}
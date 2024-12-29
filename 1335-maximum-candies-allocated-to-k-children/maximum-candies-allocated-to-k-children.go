func maximumCandies(candies []int, k int64) int {
    left := 1
    right := slices.Max(candies)
    ans := 0
    for left <= right {
        mid := left+(right-left)/2
        piles := 0
        for i := 0; i < len(candies); i++ {
            piles += (candies[i]/mid)
        }
        if int64(piles) >= k {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}
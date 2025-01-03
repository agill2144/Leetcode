func maximumCandies(candies []int, k int64) int {
    left := 1
    right := slices.Max(candies)
    ans := 0
    for left <= right {
        // num candies allowed per pile
        mid := left + (right-left)/2
        pileCount := 0
        for i := 0; i < len(candies); i++ {
            pileCount += (candies[i]/mid)
        }
        // when does mid not work?
        // when we could not create k piles ( one for each child )
        if int64(pileCount) < k {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
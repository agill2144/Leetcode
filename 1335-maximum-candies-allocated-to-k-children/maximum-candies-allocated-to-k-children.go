func maximumCandies(candies []int, k int64) int {
    left := 1
    right := 0
    for i := 0; i < len(candies); i++ {
        right += candies[i]
    }
    if k > int64(right) {return 0}
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < len(candies); i++ {
            count += (candies[i] / mid)
        }
        if int64(count) >= k {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}
func maximumCandies(candies []int, k int64) int {
    total := 0
    left := 1
    right := math.MinInt64
    for i := 0; i < len(candies); i++ {
        right = max(right, candies[i])
        total += candies[i]
    }
    if int64(total) < k {return 0}

    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < len(candies); i++ {
            count += int(math.Floor(float64(candies[i]) / float64(mid)))
        }
        if int64(count) < k {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
func maximumCandies(candies []int, k int64) int {
    left := 1
    right := 0
    total := 0
    for i := 0; i < len(candies); i++ {
        total += candies[i]
        right = max(right, candies[i])
    }
    if k > int64(total) {return 0}
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < len(candies); i++ {
            count += (candies[i]/mid)
        }
        // when does mid not work?
        // when num of subpiles of equal sizes are < k
        // if we reduce the pile size, we can more piles
        if int64(count) < k {
            right = mid-1
        } else {
            ans = mid
            left = mid+1 // we want max such ans
        }
    }
    return ans
}
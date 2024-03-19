func shipWithinDays(weights []int, days int) int {
    left := math.MinInt64
    right := 0
    for i := 0; i < len(weights); i++ {
        right += weights[i]
        left = max(left, weights[i])
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        // if at max, we are allowed to ship $mid weights in a day
        // how many days will it take to ship all packages
        // remember; packages must be shipped in the same order as they are given in
        runningWeight := 0
        count := 1
        for j := 0; j < len(weights); j++ {
            runningWeight += weights[j]
            if runningWeight > mid {
                runningWeight = weights[j]
                count++
            }
        }

        // when does mid work?
        // mid works, if we were able to ship within $days
        if count <= days {
            // save this as a potential ans
            ans = mid
            // keep searching left, since we want smallest such ans
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
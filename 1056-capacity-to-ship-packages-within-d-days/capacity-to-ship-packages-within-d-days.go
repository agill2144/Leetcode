func shipWithinDays(weights []int, days int) int {
    left := -1
    right := 0
    for i := 0; i < len(weights); i++ {
        left = max(left, weights[i])
        right += weights[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        curr := 0
        numDays := 1
        for i := 0; i < len(weights); i++ {
            curr += weights[i]
            if curr > mid {
                curr = weights[i]
                numDays++
            }
        }
        if numDays <= days {
            ans = mid
            right = mid -1
        } else {
            left = mid+1
        }
    }
    return ans
}
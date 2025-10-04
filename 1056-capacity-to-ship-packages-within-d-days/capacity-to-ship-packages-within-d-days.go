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
        rSum := 0
        d := 1
        for i := 0; i < len(weights); i++ {
            rSum += weights[i]
            if rSum > mid {
                rSum = weights[i]
                d++
            }
        }
        if d <= days {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
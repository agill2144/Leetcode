func shipWithinDays(weights []int, days int) int {
    left := slices.Max(weights)
    right := 0
    for i := 0; i < len(weights); i++ {right += weights[i]}

    res := -1
    for left <= right {
        mid := left + (right-left)/2
        sum := 0
        count := 1
        for i := 0; i < len(weights); i++ {
            sum += weights[i]
            if sum > mid {
                sum = weights[i]
                count++
            }
        }
        if count > days {
            left = mid+1
        } else {
            res = mid
            right = mid-1
        }
    }
    return res
}
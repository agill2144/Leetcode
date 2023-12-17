func shipWithinDays(weights []int, days int) int {
    left := math.MinInt64
    right := 0
    for i := 0; i < len(weights); i++ {
        if weights[i] > left {left = weights[i]}
        right += weights[i]
    }

    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        
        // how many day will we need if our capacity is $mid
        total := 0
        daysTaken := 1
        for i := 0; i < len(weights); i++ {
            total += weights[i]
            if total > mid {
                daysTaken++
                total = weights[i]
            }
        }
        
        if daysTaken <= days {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
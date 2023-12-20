func shipWithinDays(weights []int, days int) int {
    start := math.MinInt64
    end := 0
    for i := 0; i < len(weights); i++ {
        if weights[i] > start {start = weights[i]}
        end += weights[i]
    }
    left := start
    right := end
    ans := -1
    for left <= right {
        // mid is acting as our capacity to evaluate
        mid := left + (right-left)/2        
        numDays := 0        
        total := 0
        for j := 0; j < len(weights); j++ {
            total += weights[j]
            if total == mid {
                numDays++
                total = 0
            } else if total > mid{
                numDays++
                total = weights[j]
            }
            if j == len(weights)-1 && total != 0 {numDays++; total=0}
        }
        
        if numDays > days {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
        
    }
    return ans
}
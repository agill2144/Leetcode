func minDays(bloomDay []int, m int, k int) int {
    left := math.MaxInt64
    right := math.MinInt64
    for i := 0; i < len(bloomDay); i++ {
        day := bloomDay[i]
        if day < left {
            left = day
        }
        if day > right {
            right = day
        }
    }
    
    
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        
        count := 0
        bouquetsCreated := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                count++
                if count == k {
                    bouquetsCreated++
                    count = 0
                }
            } else {
                count = 0
            }
        }
        if bouquetsCreated >= m {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
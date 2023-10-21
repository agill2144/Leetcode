func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {return -1}
    left := math.MaxInt64
    right := math.MinInt64
    for i := 0; i < len(bloomDay); i++ {
        if bloomDay[i] < left {
            left = bloomDay[i]
        }
        if bloomDay[i] > right {
            right = bloomDay[i]
        }
    }
    ans := -1    
    for left <= right {
        mid := left + (right-left)/2
        // check whether we can make m bouquets of k adjacent BLOOMED flowers
        numBouquets := 0
        count := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                count++
                if count == k {
                    numBouquets++
                    count = 0
                }
            } else {
                count = 0
            }
        }
        
        if numBouquets >= m {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
func minDays(bloomDay []int, m int, k int) int {
    left := math.MaxInt64
    right := math.MinInt64
    
    // time = o(n)
    for i := 0; i < len(bloomDay); i++ {
        if bloomDay[i] < left {
            left = bloomDay[i]
        }
        if bloomDay[i] > right {
            right = bloomDay[i]
        }
    }
    
    ans := -1
    // time = o(log$maxBloomDay * n)
    for left <= right {
        mid := left + (right-left) / 2
        
        // count adjacent flowers that have bloomed on/before this day(mid)
        // and keep track of num of bouquets we were able to make on this day(mid)
        count := 0
        numBouquets := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                // this flow is bloomed
                count++
                if count == k {
                    count = 0
                    numBouquets++
                }
            } else {
                count = 0
            }
        }
        
        if numBouquets < m {
            // if we could not even make m bouquets,
            // it means we were too early on guessing which day was the correct day
            // as in, the day we picked, didnt have enough adjacent flowers blooming
            // therefore we need to pick a later day
            // therefore left = mid+1
            left = mid+1
        } else {
            // if we were able to make atleast m bouquets,
            // going to a later day after $midth day does not make any sense, because for sure they will be bloomed
            // because we are binary searching on a sorted range of days ( minVal -> maxVal )
            // save this as a potential ans, and see if we can find an earlier day instead
            // therefore right = mid-1
            ans = mid
            right = mid-1
        }
    }
    return ans
}
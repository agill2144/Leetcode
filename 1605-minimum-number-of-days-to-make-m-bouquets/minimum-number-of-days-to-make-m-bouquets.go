// time = o(n) + o((maxDay-minDay+1) * n)
func minDays(bloomDay []int, m int, k int) int {
    left := math.MaxInt64
    right := math.MinInt64
    for i := 0; i < len(bloomDay); i++ {
        left = min(left, bloomDay[i])
        right = max(right, bloomDay[i])
    }

    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        flowers := 0
        bouquets := 0
        for j := 0; j < len(bloomDay); j++ {
            if bloomDay[j] <= mid {
                flowers++
                if flowers == k {
                    bouquets++
                    flowers = 0
                }
            } else {
                flowers = 0
            }
        }
        
        // when does mid (day) not work ?
        // when num of bouquets < m
        // this means, we need to try for a future day ( left = mid+1 )
        if bouquets < m {
            left = mid+1
        } else {
            // it worked, save this day as a potential ans
            // keep searching left, since we want the earliest such day
            ans = mid
            right = mid-1
        }        
    }
    return ans
}
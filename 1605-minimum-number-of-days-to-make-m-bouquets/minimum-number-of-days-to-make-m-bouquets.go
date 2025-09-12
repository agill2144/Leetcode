func minDays(bloomDay []int, m int, k int) int {
    left := slices.Min(bloomDay)
    right := slices.Max(bloomDay)
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        bq := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                count++
                if count == k {
                    count = 0
                    bq++
                }
            } else {
                count = 0
            }
        }
        if bq >= m {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
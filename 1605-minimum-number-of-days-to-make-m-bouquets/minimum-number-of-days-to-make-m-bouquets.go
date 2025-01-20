func minDays(bloomDay []int, m int, k int) int {
    left := 1
    right := slices.Max(bloomDay)
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        boqs := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                count++
                if count == k {
                    count = 0
                    boqs++
                }
            } else {
                count = 0
            }
        }
        if boqs >= m {
            res = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return res
}
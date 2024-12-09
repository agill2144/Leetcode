func minDays(bloomDay []int, m int, k int) int {
    if k > len(bloomDay) || m*k > len(bloomDay) {return -1}
    left := 1
    right := -1
    for i := 0; i < len(bloomDay); i++ {
        right = max(right, bloomDay[i])
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        boqs := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                count++
                if count == k {
                    boqs++
                    count = 0
                }
            } else { 
                count = 0
            }
        }
        if boqs < m {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}
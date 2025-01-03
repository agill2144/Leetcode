func maxCount(banned []int, n int, maxSum int64) int {
    set := map[int]bool{}
    for i := 0; i < len(banned); i++ {
        set[banned[i]] = true
    }
    left := 1
    right := n
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        ap := (mid*(mid+1))/2
        count := mid
        for k, _ := range set {
            if k <= mid {
                count--
                ap -= k
            }
        }
        // when does mid not work ?
        if int64(ap) > maxSum {
            right = mid-1
        } else {
            ans = count
            left = mid+1
        }
    }
    return ans
}
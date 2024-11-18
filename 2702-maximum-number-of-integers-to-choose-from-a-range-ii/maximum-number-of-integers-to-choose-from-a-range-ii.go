func maxCount(banned []int, n int, maxSum int64) int {
    set := map[int64]bool{}
    for i := 0; i < len(banned); i++ {set[int64(banned[i])] = true}
    left := 1
    right := n
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        sum := int64(mid * (mid+1)/2)
        count := mid
        for k,_ := range set {if k <= int64(mid) {sum -= k; count--}}
        if  sum <= maxSum {
            ans = count
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}
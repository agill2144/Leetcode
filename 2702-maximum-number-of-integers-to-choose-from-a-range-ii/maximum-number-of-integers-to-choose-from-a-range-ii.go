func maxCount(banned []int, n int, maxSum int64) int {
    set := map[int]bool{}
    for i := 0; i < len(banned); i++ {set[banned[i]] = true}
    left := 1
    right := n
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        // choosing integers from 1 to mid
        count := mid
        sum := mid*(mid+1)/2
        // remove banned integers from the sum / count
        for k, _ := range set {
            if k <= mid {sum -= k; count--}
        }
        if int64(sum) <= maxSum {
            ans = count
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}
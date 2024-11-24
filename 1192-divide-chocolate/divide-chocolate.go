func maximizeSweetness(sweetness []int, k int) int {
    n := len(sweetness)
    left := 1
    right := 0
    for i := 0; i < n; i++ {right += sweetness[i]}
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        rSum := 0
        chunks := 0
        for j := 0; j < n; j++ {
            rSum += sweetness[j]
            if rSum >= mid {
                chunks++
                rSum = 0
            }
        }
        if chunks < k+1 {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
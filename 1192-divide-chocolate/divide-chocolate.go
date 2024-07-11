func maximizeSweetness(sweetness []int, k int) int {
    left := math.MaxInt64
    right := 0
    for i := 0; i < len(sweetness); i++ {
        left = min(left, sweetness[i])
        right += sweetness[i]
    }
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        
        cuts := 0
        rSum := 0
        for i := 0; i < len(sweetness); i++ {
            rSum += sweetness[i]
            if rSum >= mid {
                rSum = 0
                cuts++
            }
        }
        if cuts < k+1 {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
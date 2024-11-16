func maximizeSweetness(sweetness []int, k int) int {
    start := math.MaxInt64
    end := 0
    for i := 0; i < len(sweetness); i++ {
        start = min(start, sweetness[i])
        end += sweetness[i]
    }
    ans := 0
    left := start
    right := end
    for left <= right {
        mid := left + (right-left)/2
        chunks := 0
        rSum := 0
        for j := 0; j < len(sweetness); j++ {
            rSum += sweetness[j]
            if rSum >= mid {
                chunks++
                rSum = 0
            }
        }
        // when does mid not work?
        // when chunks < k+1
        if chunks < k+1 {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
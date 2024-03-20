func maximizeSweetness(sweetness []int, k int) int {
    start := math.MaxInt64
    end := 0
    for i := 0; i < len(sweetness); i++ {
        end += sweetness[i]
        start = min(start, sweetness[i])
    }
    ans := -1
    left := start
    right := end
    for left <= right {
        // atMin sweetness level to be met
        mid := left + (right-left)/2
        count := 0
        sum := 0
        for j := 0; j < len(sweetness); j++ {
            sum += sweetness[j]
            if sum >= mid {
                count++
                sum = 0
            }
        }
        // when does mid not work ?
        // when there are less than k+1 chunks
        if count < k+1 {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
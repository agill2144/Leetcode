func maximizeSweetness(sweetness []int, k int) int {
    left := 1
    right := 0
    for i := 0; i < len(sweetness); i++ {
        right += sweetness[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        sum := 0
        for i := 0; i < len(sweetness); i++ {
            sum += sweetness[i]
            if sum >= mid {
                count++
                sum = 0
            }
        }
        if count < k+1 {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
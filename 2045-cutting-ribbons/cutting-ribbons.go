func maxLength(ribbons []int, k int) int {
    left := 1
    right := 0
    for i := 0; i < len(ribbons); i++ {
        right = max(right, ribbons[i])
    }

    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        cutSize := mid
        count := 0
        for i := 0 ; i < len(ribbons); i++ {
            count += int(math.Floor(float64(ribbons[i])/float64(cutSize)))
        }

        if count < k {
            right = mid-1
        } else {
            ans = cutSize
            left = mid+1
        }
    }
    return ans
}
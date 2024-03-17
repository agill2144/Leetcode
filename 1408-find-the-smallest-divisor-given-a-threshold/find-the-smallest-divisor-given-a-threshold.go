func smallestDivisor(nums []int, threshold int) int {
    left := 1
    right := math.MinInt64
    for i := 0; i < len(nums); i++ { right = max(right, nums[i]) }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        sum := 0
        for j := 0; j < len(nums); j++ {
            sum = sum + (int(math.Ceil(float64(nums[j]) / float64(mid))))
        }

        // when does mid not work?
        if sum > threshold {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }

    }
    return ans
}
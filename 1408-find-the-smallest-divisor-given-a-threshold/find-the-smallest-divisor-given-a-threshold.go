func smallestDivisor(nums []int, threshold int) int {
    left := 1
    right := -1
    for i := 0; i < len(nums); i++ {
        right = max(nums[i], right)
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        sum := 0
        for i := 0; i < len(nums); i++ {
            sum += int(math.Ceil(float64(nums[i])/float64(mid)))
        }
        if sum > threshold {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}
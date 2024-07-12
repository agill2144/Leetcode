func splitArray(nums []int, k int) int {
    start := math.MinInt64
    end := 0
    for i := 0; i < len(nums); i++ {
        start = max(start, nums[i])
        end += nums[i]
    }
    ans := 0
    left := start
    right := end
    for left <= right {
        mid := left + (right-left)/2

        count := 1
        rSum := 0
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > mid {
                rSum = nums[i]
                count++
            }
        }

        if count > k {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}
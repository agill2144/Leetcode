func splitArray(nums []int, k int) int {
    left := 0
    right := 0
    for i := 0; i < len(nums); i++ {
        left = max(left, nums[i])
        right += nums[i]
    }
    ans := 0
    for left <= right {
        mid := left + (right-left)/2

        rSum := 0
        cuts := 1
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > mid {
                cuts++
                rSum = nums[i]
            }
        }
        if cuts > k {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans

}
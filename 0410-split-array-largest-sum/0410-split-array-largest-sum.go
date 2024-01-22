func splitArray(nums []int, k int) int {
    if k > len(nums) {return -1}
    left := math.MinInt64
    right := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > left {left = nums[i]}
        right += nums[i]
    }
    
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        // if atMax, subarray sum cannot exceed $mid
        // how many subarrays would we have?
        count := 1
        rSum := 0
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > mid {
                count++
                rSum = nums[i]
            }
        }
        
        // does $mid work?
        if count > k {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}
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
        // largest sum of any subarray is minimized.
        // max ans is minimized
        // so at max if our split sum is mid
        // how many subarrays would we get?
        
        count := 1
        rSum := 0
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > mid {
                count++
                rSum = nums[i]
            }
        }
        if count <= k {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
    
}
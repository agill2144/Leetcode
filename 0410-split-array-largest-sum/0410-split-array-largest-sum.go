func splitArray(nums []int, k int) int {
    left := math.MinInt64
    right := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > left {left = nums[i]}
        right += nums[i]
    }
    
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        splits := totalSplits(nums, mid)

        if splits > k {
            left = mid + 1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}

func totalSplits(nums []int, maxSum int) int {
    rSum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if rSum + nums[i] <= maxSum {
            rSum += nums[i]
        } else {
            rSum = nums[i]
            count++
        }
    }
    return count+1
}
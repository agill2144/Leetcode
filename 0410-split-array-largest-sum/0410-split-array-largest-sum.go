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
        partitionsPlaced := totalSplits(nums, mid)

        if partitionsPlaced > k-1 {
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
    partitions := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        if rSum > maxSum {
            partitions++
            rSum = nums[i]
        }
    }
    // return the number of paritions placed
    return partitions
}
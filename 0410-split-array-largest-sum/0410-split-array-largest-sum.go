// time = o(n) + o( (log$totalSum-$maxNum+1)*n )
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
        partitionsPlaced := totalPartitionsPlaced(nums, mid)
        
        // if k = 2; this means, create 2 subarrays with 1 partition
        // if k = 3; this means, create 3 subarrays with 2 partition
        // therefore number of partitions placed must be < k-1 
        // if we have placed more than k-1 partition, our partition sum is too small,
        // we need to increase our allowed partition sum, therefore left = mid+1
        if partitionsPlaced > k-1 {
            left = mid + 1
        } else {
            // we have a potential ans, but since we need the smallest possible ans
            // save this ans, and continue searching for a smaller ans
            // i.e search left; i.e move left
            ans = mid
            right = mid-1
        }
    }
    return ans
}

func totalPartitionsPlaced(nums []int, maxSum int) int {
    rSum := 0
    partitions := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        if rSum > maxSum {
            partitions++
            rSum = nums[i]
        }
    }
    return partitions
}
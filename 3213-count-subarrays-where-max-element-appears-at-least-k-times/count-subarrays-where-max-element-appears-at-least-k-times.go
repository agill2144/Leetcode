func countSubarrays(nums []int, k int) int64 {
    if k > len(nums) {return 0}
    maxNum := math.MinInt64
    for i := 0; i < len(nums); i++ {maxNum = max(maxNum, nums[i])}

    total := 0
    maxCount := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == maxNum {
            maxCount++
        }
        for maxCount >= k {
            total += len(nums)-i
            if nums[left] == maxNum {maxCount--}
            left++
        }
    }

    return int64(total)
}   
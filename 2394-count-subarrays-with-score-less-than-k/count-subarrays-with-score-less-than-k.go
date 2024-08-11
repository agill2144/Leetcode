func countSubarrays(nums []int, k int64) int64 {
    count := 0
    left := 0
    rSum := 0 
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        score := rSum * (i-left+1)
        for int64(score) >= k {
            rSum -= nums[left]
            left++
            score = rSum * (i-left+1)            
        }
        count += (i-left+1)
    }
    return int64(count)
}


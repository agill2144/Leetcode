func numOfSubarrays(nums []int, k int, threshold int) int {
    left := 0
    rSum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        if i-left+1 == k {
            yes := (rSum/k) >= threshold
            if yes {
                count++
            }
            rSum -= nums[left]
            left++
        }
    }
    return count
}
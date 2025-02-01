func numSubarraysWithSum(nums []int, goal int) int {
    if goal < 0 {return 0}
    return atMostK(nums,goal) - atMostK(nums, goal-1)
}

func atMostK(nums []int, k int) int {
    left := 0
    n := len(nums)
    rSum := 0
    total := 0
    for i := 0; i < n; i++ {
        rSum += nums[i]
        for rSum > k && left < i {
            rSum -= nums[left]
            left++
        }
        if rSum <= k {
            total += (i-left+1)
        }
    }
    return total
}
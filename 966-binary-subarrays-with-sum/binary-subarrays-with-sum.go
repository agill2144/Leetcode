func numSubarraysWithSum(nums []int, goal int) int {
    if goal < 0 {return 0}
    return atMostK(nums, goal) - atMostK(nums, goal-1)

}

func atMostK(nums []int, k int) int {
    n := len(nums)
    count := 0
    left := 0
    sum := 0
    for i := 0; i < n; i++ {
        sum += nums[i]
        for sum > k && left <= i {
            sum -= nums[left]
            left++
        }
        count += (i-left+1)
    }
    return count
}
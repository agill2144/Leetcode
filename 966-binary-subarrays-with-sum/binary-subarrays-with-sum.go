// prefixSum / runningSum approach
func numSubarraysWithSum(nums []int, goal int) int {
    freq := map[int]int{0:1}
    rSum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        toRemove := rSum - goal
        count += freq[toRemove]
        freq[rSum]++
    }
    return count
}
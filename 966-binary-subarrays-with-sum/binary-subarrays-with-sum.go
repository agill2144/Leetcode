func numSubarraysWithSum(nums []int, goal int) int {
    freq := map[int]int{0:1}
    rSum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        count += freq[rSum-goal]
        freq[rSum]++
    }
    return count
}
func numSubarraysWithSum(nums []int, goal int) int {
    rSum := 0
    freq := map[int]int{0:1}
    count := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        count += freq[rSum-goal]
        freq[rSum]++
    }
    return count
}

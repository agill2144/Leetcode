func numSubarraysWithSum(nums []int, goal int) int {
    freq := map[int]int{0:1}
    rSum := 0
    total := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        total += freq[rSum - goal]
        freq[rSum]++
    }
    return total
}


func subarraySum(nums []int, k int) int {
    rSum := 0
    freq := map[int]int{0:1}
    total := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        total += freq[rSum - k]
        freq[rSum]++
    }
    return total
}

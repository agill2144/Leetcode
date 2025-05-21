func subarraySum(nums []int, k int) int {
    freq := map[int]int{0:1}
    rSum := 0
    total := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        diff := rSum-k
        total += freq[diff]
        freq[rSum]++
    }
    return total
}
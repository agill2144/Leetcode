func subarraySum(nums []int, k int) int {
    count := 0
    freq := map[int]int{0:1}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        count += freq[rSum - k]
        freq[rSum]++
    }
    return count
}
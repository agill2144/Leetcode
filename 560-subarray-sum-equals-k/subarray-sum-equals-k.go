func subarraySum(nums []int, k int) int {
    freq := map[int]int{}
    rSum := 0
    total := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        if rSum == k {total++}
        total += freq[rSum - k]
        freq[rSum]++
    }
    return total
}
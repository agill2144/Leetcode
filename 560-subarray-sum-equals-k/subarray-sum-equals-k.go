func subarraySum(nums []int, k int) int {
    count := 0
    freq := map[int]int{}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        if rSum == k {
            count++
        }
        diff := rSum - k
        count += freq[diff]
        freq[rSum]++
    }
    return count
}
// sliding window temptations
// but sliding window only works when window sum is either increasing or descreasing
// its always monotonic in nature
// with negatives, the window sum will be everywhere, hence we cannot use sliding window
// runningSum pattern.
func subarraySum(nums []int, k int) int {
    sum := 0
    freq := map[int]int{0:1}
    count := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        diff := sum-k
        count += freq[diff]
        freq[sum]++
    }
    return count
}

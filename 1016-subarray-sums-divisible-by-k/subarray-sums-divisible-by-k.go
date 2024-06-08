// sliding window temptations
// but sliding window only works when window sum is either increasing or descreasing
// its always monotonic in nature
// with negatives, the window sum will be everywhere, hence we cannot use sliding window
// runningSum pattern.
func subarraysDivByK(nums []int, k int) int {
    remainders := map[int]int{0:1}
    sum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        if rem < 0 {rem+=k}
        count += remainders[rem]
        remainders[rem]++
    }
    return count
}
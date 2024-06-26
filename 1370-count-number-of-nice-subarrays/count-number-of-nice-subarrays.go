
// time = o(n)
// space = o(n)
// similar approach as runningSum pattern;
// https://leetcode.com/problems/subarray-sum-equals-k/
func numberOfSubarrays(nums []int, k int) int {
    oddCounts := map[int]int{0:1}
    res := 0
    runningCount := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {
            runningCount++
        }
        res += oddCounts[runningCount-k]
        oddCounts[runningCount]++
    }
    return res
}
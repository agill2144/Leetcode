// duplicate of: https://leetcode.com/problems/subarray-sum-equals-k/
func numSubarraysWithSum(nums []int, goal int) int {
    count := 0
    rSum := 0
    sumCount := map[int]int{0:1}
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        diff := rSum - goal
        val, ok := sumCount[diff]
        if ok {
            count += val
        }
        sumCount[rSum]++
    }
    return count
}
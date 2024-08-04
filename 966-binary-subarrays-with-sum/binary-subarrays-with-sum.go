func numSubarraysWithSum(nums []int, goal int) int {
    sumCount := map[int]int{0:1}
    rSum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        diff := rSum-goal
        count += sumCount[diff]
        sumCount[rSum]++
    }
    return count
}
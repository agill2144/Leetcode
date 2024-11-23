func numSubarraysWithSum(nums []int, goal int) int {
    if goal < 0 {return 0}
    return atMostK(nums, goal) - atMostK(nums, goal-1)
}

func atMostK(nums []int, target int) int {
    n := len(nums)
    left := 0
    count := 0
    rSum := 0
    for i := 0; i < n; i++ {
        rSum += nums[i]
        for rSum > target && left < i {
            rSum -= nums[left]
            left++
        }
        if rSum <= target {count += (i-left+1)}
    }
    return count
}
// func numSubarraysWithSum(nums []int, goal int) int {
//     freq := map[int]int{0:1}
//     rSum := 0
//     count := 0
//     for i := 0; i < len(nums); i++ {
//         rSum += nums[i]
//         count += freq[rSum-goal]
//         freq[rSum]++
//     }
//     return count
// }
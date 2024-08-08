// sliding window atMost k pattern
func numSubarraysWithSum(nums []int, goal int) int {
    if goal < 0 {return 0}
    return countAtMost(nums, goal) - countAtMost(nums, goal-1)
}

func countAtMost(nums []int, k int)int {
    count := 0
    left := 0
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        for rSum > k && left <= i {
            rSum -= nums[left]
            left++
        }
        count += (i-left+1)
    }
    return count
}

// prefixSum / runningSum approach
// func numSubarraysWithSum(nums []int, goal int) int {
//     freq := map[int]int{0:1}
//     rSum := 0
//     count := 0
//     for i := 0; i < len(nums); i++ {
//         rSum += nums[i]
//         toRemove := rSum - goal
//         count += freq[toRemove]
//         freq[rSum]++
//     }
//     return count
// }
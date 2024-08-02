func numSubarraysWithSum(nums []int, goal int) int {
    if goal < 0 {return 0}
    return count(nums, goal) - count(nums, goal-1)
}

func count(nums []int, goal int) int {
    c := 0
    left := 0
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        for sum > goal && left <= i {
            sum -= nums[left]
            left++
        }
        c += (i-left+1)
    }
    return c
}

// duplicate of: https://leetcode.com/problems/subarray-sum-equals-k/
// func numSubarraysWithSum(nums []int, goal int) int {
//     count := 0
//     rSum := 0
//     sumFreq := map[int]int{0:1}
//     for i := 0; i < len(nums); i++ {
//         rSum += nums[i]
//         diff := rSum - goal
//         count += sumFreq[diff]
//         sumFreq[rSum]++
//     }
//     return count
// }
func numSubarraysWithSum(nums []int, goal int) int {
    if goal < 0 {return 0}    
    return countLessThanOrEqualTo(nums, goal) - countLessThanOrEqualTo(nums, goal-1) 
}

func countLessThanOrEqualTo(nums []int, target int) int {
    count := 0
    left := 0
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        for sum > target && left <= i {
            sum -= nums[left]
            left++
        }
        count += (i-left+1)
    }
    return count
}

/*
    approach: runningSum pattern
    - num of subarrays + some mathematical = runningSum
    - identical to https://leetcode.com/problems/subarray-sum-equals-k/description/
    time = o(n)
    space = o(n)
*/

// func numSubarraysWithSum(nums []int, goal int) int {
//     sumCount := map[int]int{0:1}
//     rSum := 0
//     count := 0
//     for i := 0; i < len(nums); i++ {
//         rSum += nums[i]
//         diff := rSum-goal
//         count += sumCount[diff]
//         sumCount[rSum]++
//     }
//     return count
// }
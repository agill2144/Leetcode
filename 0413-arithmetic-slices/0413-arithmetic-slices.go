func numberOfArithmeticSlices(nums []int) int {
    dp := make([]int, len(nums))
    count := 0
    for i := 2; i < len(nums); i++ {
        if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
            dp[i] = dp[i-1]+1
            count += dp[i]
        }
    }
    return count
}

/*
    approach: brute force
    - form each subarray
    time: o(n^2)
    space: o(1)
*/
// func numberOfArithmeticSlices(nums []int) int {
//     count := 0
//     for i := 0; i < len(nums)-1; i++ {
//         diff := nums[i+1]-nums[i]
//         for j := i+2; j < len(nums); j++ {
//             if nums[j] - nums[j-1] != diff {break}
//             if j-i+1 >= 3 {count++}
//         }
//     }
//     return count
// }
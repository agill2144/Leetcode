// brute force dfs - no need to backtrack, just maintain the path len and prev element that was added to path
// time: o(2^n)
// space: o(n)()
// func lengthOfLIS(nums []int) int {
//     if nums == nil || len(nums) == 0 {
//         return 0
//     }
//     maxLen := 0
//     var dfs func(start int, count int, prev int)
//     dfs = func(start int, count, prev int) {
//         // base
//         if count > maxLen {maxLen=count}
        
//         // logic
//         for i := start; i < len(nums); i++ {
//             if prev == math.MaxInt64{
//                 dfs(i+1,count+1,nums[i])
//             } else if nums[i] > prev {
//                 dfs(i+1, count+1, nums[i])
//             }
//         }
//     }
    
//     dfs(0, 0, math.MaxInt64)
//     return maxLen
// }
    

// approach: bottom up dp and form every single subsequence for each subproblem and pick longest one for each subproblem
// time : o(n^2)
// space: o(n) == dp array
// func lengthOfLIS(nums []int) int {
//     if nums == nil || len(nums) == 0 {
//         return 0
//     }
//     dp := make([]int, len(nums))
//     for idx, _ := range dp {dp[idx] = 1}
//     maxLen := 1
//     for i := 1; i < len(nums); i++ {
//         currNum := nums[i]
//         for j := 0; j < i; j++ {
//             if currNum > nums[j] {
//                 dp[i] = max(dp[j]+1, dp[i])
//             }
//             maxLen = max(maxLen, dp[i])
//         }
//     }
//     return maxLen
// }

// func max(x, y int) int {
//     if x > y {return x}
//     return y
// }


// time: o(nlogn)
// space: o(n)
func lengthOfLIS(nums []int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }
    effective := []int{nums[0]}
    
    for i := 1; i < len(nums); i++ {
        lastEffectiveIdx := len(effective)-1

        if nums[i] > effective[lastEffectiveIdx] {
            effective = append(effective, nums[i])
        } else {
            idx := nextSmallestBinarySearch(nums[i],effective)
            effective[idx] = nums[i]
        }
    }
    return len(effective)
}

func nextSmallestBinarySearch(target int, nums []int) int {
    ans := -1
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            if nums[mid] == target {return mid}
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
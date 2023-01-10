/*
    
    approach: bottom up dp
    - DP matrix or DP array?
        - dp array works because we only have 1 decision making variable provided as input
    - What decision do we need to make at each ith subproblem ?
        1. add this number to current subarray sum
            - nums[i] + previous subarray sum ( which is 1 step back in dp array )
            - nums[i] + dp[i-1]
        2. choose this number on its own because adding to current subarray sum would be smaller than current number on its own
            - nums[i]        
        
        max(1, 2)
    - Draw the decision tree on the above decisions to see MANY repeated subproblems
            
    - Why dont we have a not choose case?
    - because if we have a not choose case, then we are simply saying, stop this subarray and skip this number.
    - we cannot skip any number because we are looking for contagious subarray
    - if we were looking for a subsequence sum, sure, then we have a not choose case as well since we can skip numbers in subsequences!
    - we can either add this number to current subarray sum OR start a new subarray sum and abandon the previous subarray sum.
    - we will also keep track of a max at each subproblem
        - max will not be at the last position of the dp array
        - because each subproblem will have its own max answer, we do not want the max answer for last idx, we want max whenever we had seen a max
        - there is a chance we will stop and make a new subarray, and the max was in the middle somewhere
        - example: 
            nums = [5,4,-1,7,-80]
            dp = [5 9 8 15 -65]
            
            we want to return 15 (5+4+-1+7)
    
    time: o(n)
    space: o(n)
*/
// func maxSubArray(nums []int) int {
//     if nums == nil || len(nums) == 0 {return -1}
//     if len(nums) == 1 {return nums[0]}
    
//     dp := make([]int, len(nums))
//     dp[0] = nums[0]
//     out := dp[0]

//     for i := 1; i < len(nums); i++ {
//         dp[i] = max(nums[i], nums[i]+dp[i-1])
//         out = max(out, dp[i])
//     }
//     return out
// }

/*
    approach: bottom up dp + space optimization
    - since we only need to look at 1 step back in 1d dp array
    - we can save that in a variable ( lets call it current or whatever )
    
    time: o(n)
    space: o(1)
*/

func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return -1}
    if len(nums) == 1 {return nums[0]}
    
    curr := nums[0]
    out := curr

    for i := 1; i < len(nums); i++ {
        curr = max(nums[i], nums[i]+curr)
        out = max(out, curr)
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}



/*
    approach: brute force
    - nested loop
    - keep track of running sum
    - save running into max whenever needed
    
    time: o(n^2)
    space: o(1)
    

func maxSubArray(nums []int) int {
    max := math.MinInt64
    for i := 0; i < len(nums); i++ {
        running := 0
        for j := i; j < len(nums); j++ {
            running += nums[j]
            if running > max {
                max = running
            }
        }
    }
    return max
}
*/
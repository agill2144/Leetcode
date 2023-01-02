
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


/*
    approach: linear running sum
    - we do what we used to with runningSum pattern
    - but instead of continuing adding nums as we loop
    - we will still add current num to runningSum
    - but since the inputs can be negative integers
    - the current number on its own may be bigger than currentRunningSum
        - this is when we will reset our window and our window starts from this number onwards ( since current number is > than the prev runningSum )
    - so, in that case , if current number > runningSum = then runningSum becomes currentNumber
    - Save the max between max,runningSum
    - So the idea becomes instead of blindly adding to runningSum , we will pick and choose whether
        - adding current number to runningSum is better or
        - currentNumber is better than the above addition result
    - Then we will save the max between previous max and new updated runningSum in each iteration
    
    - TLDR: 
        - runningSum pattern ( but less trickier )
        - runningSum = max(runningSum, runningSum+nums[i])
        - max = max(max, runningSum)
    
    time: o(n)
    space: o(1)
*/
func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return 0}
//     curr := nums[0]
//     out := nums[0]

//     for i := 1; i < len(nums); i++ {
//         curr = max(curr+nums[i], nums[i])
//         out = max(out, curr)
//     }
//     return out
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    out := dp[0]
    
    for i := 1; i < len(nums); i++ {
        // we do not have a zero case - this i dont understand why we dont have a zero case ( instinct says it has to do with "subarray" but havent proven to myself yet )
        // its between 1. choose this number on it own 2. this number + previous sub problem answer
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        out = max(dp[i], out)
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
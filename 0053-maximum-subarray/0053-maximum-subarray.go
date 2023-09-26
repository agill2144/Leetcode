func maxSubArray(nums []int) int {
    runningSum := 0
    maxSum := math.MinInt64
    for i := 0; i < len(nums); i++ {
        runningSum += nums[i]
        if runningSum > maxSum {
            maxSum = runningSum
        }
        // greedy-like intuition
        // any subarray whose sum is positive is worth keeping,
        // as soon as subarray becomes negative
        // we can drop this subarray ( i.e reset sum = 0 )
        // there is no point in carrying a negative subarray forward
        // this means from next idx, we will be exploring a new subarray
        
        // what if all the numbers were negative
        // - does not matter
        // - we add each number to runningSum, compare with maxSum, and save it if needed
        // therefore if all numbers were negative [-10,-9,-1]
        // we will first save, -10, then reset to 0, 
        // then we will have -9, save it and reset to 0 
        // and then we will have -1, save it and reset to 0
        if runningSum < 0 {
            runningSum = 0
        }
    }
    return maxSum
}

func max(x, y int) int {
    if x > y {return x}
    return y
}

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
        - this is when we will reset our window and our window starts from this number onwards
    - so, in that case , if current number > runningSum = then runningSum becomes currentNumber
    - Save the max between max,runningSum
    - So the idea becomes instead of blindly adding to runningSum , we will pick and choose whether
        - adding current number to runningSum is better or
        - currentNumber is better than the above addition result
    - Then we will save the max between previous max and new updated runningSum in each iteration
    
    - TLDR: 
        - runningSum pattern
        - runningSum = max(runningSum, runningSum+nums[i])
        - max = max(max, runningSum)
    
    time: o(n)
    space: o(1)
*/

func maxSubArray(nums []int) int {
    running := 0
    max := math.MinInt64
    for i := 0; i < len(nums); i++ {
        running += nums[i]
        running = getMax(running, nums[i])
        max = getMax(running, max)
    }
    return max
}

func getMax(n,m int) int {
    if n > m {return n}
    return m
}
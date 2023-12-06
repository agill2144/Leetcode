/*
    approach: kadanes algo
    - greey like
    - keep a running sum
    - save running sum if better than maxSum seen so far
    - now decide whether this subarray runningSum is worth keeping or not ?
    
    - its not worth keep if subarray runningSum is negative
        - reset runningSum and start forming a new runningSum (i.e new subArray)
        - what if all the next numbers are negative ?
            - [-5,-4,-1]; runningSum = -5
            - if all next numbers are negative then we care about smallest negative
            - which is -1
            - continuing with -5+(-4)+(-1) will be smaller -1
            - therefore if runningSum becomes negative, its not worth keeping
        - what if the next number is positive ?
            - [-5,6]; runningSum = -5
            - continuing with runningSum of -5 and adding next positive number 6; will make runningSum = 1
            - but infact our best subArray with largest sum is just [6] on its own
            - therefore even if the next number is positive, drop the runningSum if its negative
    - its worth keeping / continuing with a runningSum thats is positive
        - [10,-1,9];
        - runningSum at idx0 = 10
        - runningSum at idx1 = 9
        - runningSum at idx2 = 18
        - if we had dropped the runningSum when we saw a negative number and that negative number didnt make our runningSum negative
        - then our max subarray sum would have just be [10], when the ans is [10,-1,9]
        - but if instead of -1 it was -20, then that would have made our runningSum negative = -10, thats when we would have dropped it
    
    time = o(n)
    space = o(1)
    
    
    Follow up; print the subarray whose sum is max
    - we need to keep track of start and end of subarray that has maxSum
    - we need another ptr that keeps track of runningStart of a subarray
        - runningStart moves slowly
        - runningStart moves ahead when we start forming a new subarray
        - therefore runningStart is like a slow ptr
        - while i ptr is the end of a subarray ( i.e the fast ptr )
    - when we save a runningSum is better than maxSum
        - we have slow ptr = start of subarray
        - we have ith ptr = end of subarray
    - slow ptr moves forward when we have decided to form a new subarray
        - i.e if runningSum is negative, reset runningSum and move slow to i+1 ( forming new subarray )
        
*/
func maxSubArray(nums []int) int {
    sum := 0
    maxSum := math.MinInt64
    start, end := -1, -1
    slow := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum > maxSum {
            maxSum = sum
            start = slow
            end = i
        }
        
        // is it worth keeping this rolling sum / subarray ?
        // its not worth keeping if the sum is negative 
        // its worth keeping if the sum in positive
        if sum < 0 {
            slow = i+1
            sum = 0
        }
    }
    fmt.Println(nums[start:end+1])
    return maxSum
}
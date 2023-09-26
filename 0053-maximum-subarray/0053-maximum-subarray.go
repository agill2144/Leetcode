/*
    approach: kadanes algorithm
    - greedy like
    - keep track of a sum and a maxSum
    - for each num, add it to sum
    - save this sum if better than prev maxSum
    - if sum goes negative, its not worth keeping / continuing this subarray, reset to 0
        - what happens if next number is positive? 
            - would we rather keep (-sum + next-positive num) or just (next-positive num)?
        - what happens if all numbers are negative?
            - -sum + -nextNegativeNum is only going to make the sum farther away from 0.
            - i.e smaller in sum, we want the maxSum.
            - therefore if all numbers are negative, than we will end up saving single number thats max among all
    - if the sum is positive, its worth continuing with this subarray
        - if this this subarray was an ans , we have already saved it in maxSum
        - We are hoping next numbers dont make the sum negative, 
        - and hoping carrying the positive sum may lead to a larger sum
    
    time = o(n)
    space = o(1)
    
*/
func maxSubArray(nums []int) int {
    sum := 0
    maxSum := math.MinInt64
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        maxSum = max(sum, maxSum)
        if sum < 0 {
            sum = 0
        }
    }
    return maxSum
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
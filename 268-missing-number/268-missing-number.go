/*
    approach: sum!
    - take a curr sum of all nums
    - take the expected sum from 1 to n
        - n * (n+1)/2
    - return the diff
*/
func missingNumber(nums []int) int {
    currSum := 0
    for i := 0; i < len(nums); i++ {
        currSum += nums[i]
    }
    n := len(nums)
    expectedSum := n*(n+1)/2
    return expectedSum-currSum
}

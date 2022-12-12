/*
    This is sight variation of house robber - just not very obvious

    approach: Top down / Bottom up DP
    - First we need to figure out the range of numbers in nums array
        - We only really need a max 
        - Why do we need this?
        - The idea is that we will have a values array 
        - where each idx maps to a value we have seen in nums array
        - and each idx will be total sum of $value from nums array
        - for example
            - nums = [3,2,4,2]
            - max = 4
            - initial values array of size max+1 = [0,0,0,0,0]
            - then we will loop over nums array and add each ith value to its idx in values array
            - updated values array = [0,0,4,3,4]
                - 0 does not exist in nums - therefore idx 0 has a value of 0 in values
                - 1 does not exist in nums - therefore idx 1 has a value of 0 in values
                - 2 does exist in nums ( 2 times ) - therefore idx 2 has a value of 4(2+2) in values
                - 3 does exist in nums ( 1 time ) - therefore idx 3 has a value of 3 in values
                - 4 does exist in nums ( 1 time ) - therefore idx 4 has a value of 4 in values
            - now translate the values array to house robber problem
            - where if you rob a house, you cannot rob the next immediate house + value from current house
            - and if you do not rob this house, then you will pick a max between prev 2 values ( choose || notChoose )
            - values array becomes the houses array
    - Now we have a values array, for which we can apply house robber DP on it.
        - choose || notChoose case
        - we will return max between the two
        - We can take a dp array of size 2 ( choose and not choose case )
        - And start looping over values array
            - 0th case will be max between the prev row
            - 1 case will be current value + 0th case from prev row
    
    time: o(n) + o(n) + o(maxNumberInArray) = o(2n) + o(maxNumberInArray)
    space:  o(maxNumberInArray)
*/
func deleteAndEarn(nums []int) int {
    max := math.MinInt64
    // time: o(n)
    // space: o(1)
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    
    // time: o(n)
    // space: o(maxNumberInArray)
    values := make([]int, max+1)
    for _, num := range nums {
        values[num] += num
    }
    
    // top down
    dp := make([]int, 2)
    dp[0] = 0
    dp[1] = values[0]
    // time: o(maxNumberInArray)
    // space: o(1)
    for i := 1; i < len(values); i++ {
        choose := dp[1] // prev choose case
        notChoose := dp[0] // prev notChoose case
        dp[0] = int(math.Max(float64(choose), float64(notChoose)))
        dp[1] = notChoose + values[i]
    }
    return int(math.Max(float64(dp[0]), float64(dp[1])))
}
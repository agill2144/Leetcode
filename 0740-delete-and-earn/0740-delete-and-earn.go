/*
    approach: DP
    - DP hints: 
        - We are given *choices* to pick and decide from
            - So go explore all choices! DFS!
            - Draw the recursion tree and notice repeating subproblems
            - if we start by solving the smallest subproblem, and work our selves up - we will have the answer to the bigger problem!
        - We also have to *maximize* on the points earned by making decisions
    - The problem says if we picked nums[i], then we cannot use nums[i]-1 and nums[i]+1
        - Initially thought if the numbers were sorted, linear(no gaps in between) and were not dupes then we can just say, do not choose the next position.
        - But we cannot do that because 1. not sorted 2. we have dupes 3. numbers are not linear so we cannot make a decision based on idx positions.
    
    - Instead what we can do is fill in the gaps of missing number so we have a linear sequence of numbers represented by their idx positions
    - if a value does not exist, then its value will be 0
    - for example;
        - [1,5]
        - then we can fill in the missing numbers will 0's, such that 1 is sitting at idx 1 and 5 is sitting at idx 5
        - [0,1,0,0,0,5] ( a new array )
        - what about dupes? For example; [1,1,5,5]
        - Then we will simply treat our new array as a sum holder for each value at each respective idx position
        - at idx 1, we originally had 1 because there was only 1 
        - now at idx 1 we will have 2 because there are two one's
        - at idx 5 we originally had 5 because there was only one 5
        - now at idx 5 we will have sum of all 5's (i.e 10) becuase now we have two 5's
        - [0,2,0,0,0,10]
    - Now this problem becomes like the house robber problem
    - Do this bottom up with 0,1 cases and you will have an answer at the very top of your dp matrix
    
    - We can also optimize our space used by our dp matrix
    - since we only need to look 1 level down, we can just have 1d dp array
    - and overrwrite in place.
    
    Time: o(n) + o(maxVal) + o(maxVal)
    Space: o(maxVal) for the new array
*/
func deleteAndEarn(nums []int) int {
    maxVal := math.MinInt64
    for i := 0; i < len(nums); i++ {
        maxVal = max(maxVal, nums[i])
    }
    values := make([]int, maxVal+1)
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        values[num] += num
    }
    
    dp := [2]int{0, values[len(values)-1]}
    for i := len(values)-2; i >= 0; i-- {
        choose := values[i] + dp[0]
        notChoose := max(dp[0], dp[1])
        dp[0],dp[1] = notChoose, choose
    }
    
    return max(dp[0], dp[1])
}


func max(x, y int) int {
    if x > y {return x}
    return y
}
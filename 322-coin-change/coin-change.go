/*
    1. try greedy
    2. if it does not work, explore all paths using recursion
    3. see repeated subproblems in recursion, then optimize with dp
    4. dp matrix vs dp array ?
        - depends on number of constraints, decisions making variables
        - we have 2, i.e more than 1, therefore matrix
        - can use matrix pretty much everytime, i think, even


    approach: dfs
    - explore all paths
    numberOfBranchesPerNode ^ maxDepthOfRecursion * extraWorkPerRecursion
    numberOfBranchesPerNode = 2
    maxDepthOfRecursion = amount
    extraWorkPerRecursion = none
    tc = 2^n
    sc = o(n) for recursive stack

    approach: Bottom up dp
    - start with the smallest subproblem
    - and look left for pre-solved subproblem
    tc = o(m*n)
    sc = o(m*n)

    approach: bottom up dp, space optimized
    - in 2d dp, we only looked up 1 row for an old value and looked left in the same row for an updated value
    - we can use a 1D array
    - each cell before being updated has the prev row value ( well visually )
    - and to the left, in the same row, we get the updated value
    tc = o(m*n)
    sc = o(n)
*/
func coinChange(coins []int, amount int) int {
    m := len(coins)+1
    n := amount+1
    dp := make([]int, n)
    for i := 0; i < n; i++ {dp[i] = math.MaxInt64-1000}
    dp[0] = 0

    for i := 1; i < m; i++ { // coins
        for j := 1; j < n; j++ { // amount
            coin := coins[i-1]
            am := j
            if coin <= am {
                // not choose = j
                // choose = 1 + (dp[j-coinStepsBack])
                dp[j] = min(dp[j], 1+dp[j-coin])
            }
        }
    }
    if dp[n-1] == math.MaxInt64-1000 {return -1}
    return dp[n-1]
}
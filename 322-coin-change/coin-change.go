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

    approach: 2D DP
    - 
*/
func coinChange(coins []int, amount int) int {
    m := len(coins)+1
    n := amount+1
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        if i == 0 {
            for j := 1; j < n; j++ {
                dp[i][j] = math.MaxInt64-1000
            }
        }
    }
    for i := 1; i < m; i++ { // coins
        for j := 1; j < n; j++ { // amount
            coin := coins[i-1]
            am := j
            if coin > am {
                // we do not have a choose case
                // not choose case is 1 step up
                dp[i][j] = dp[i-1][j]
            } else {
                // not choose = j-1
                // choose = 1 + (dp[i][j-coinStepsBack])
                dp[i][j] = min(dp[i-1][j], 1+dp[i][j-coin])
            }
        }
    }
    if dp[m-1][n-1] == math.MaxInt64-1000 {return -1}
    return dp[m-1][n-1]
}
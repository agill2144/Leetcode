/*
    The DP framework to follow
    1. Is the question asking for something optimal?
    2. See if greedy works
        - Explore a single path and see if it works for all types of inputs
    3. If greedy ( exploring single path ) does not work, then explore ALL PATHS ( using DFS or similar )
    4. For DFS/recursion
        a. Create your decision tree
        b. Code out your decision tree ( using recursion )
        c. Our DFS logic for this problem is;
            1. we have 2 choices - to use a coin value or to not use a coin value
            2. if we do choose to use a coin value, then
            3. We have to keep track of number of coins used so far
            4. We also have to keep track of amount created so far
            5. So becuase we have to 2 choices, we will have 2 recursive calls
            6. One with the current coin choosen and other without ( with the new state for each )
            7. Our base condition becomes responsible for ensuring that
                1. if our amount == target amount, then update some global min with current number of coins used so far.
                2. if our amount goes beyond target amount, return
                3. or if we are out of coins to use, return
    5. Notice repeated subproblems in your decision tree?
        - Now cache your answers using DP matrix
    6. When to use Top-Down vs Bottom-Up DP ? 
        - I dont know fully yet
    7. For this we will use bottom up DP
        - Bottom up DP is solving a smaller subproblem
        - Its like starting from a leaf node of your decision tree
        - and then bubbling up to the main/bigger problem
        - In bottom up DP, we usually use dp matrix
    8. DP matrix
        - a matrix has rows and columns
        - We need to find our problem constraints / decision making variables and map them to matrix row and col
        - If we only have 1 constraint, then we can simply use dp array over a matrix
        - If we have 2 or more constraints, then matrix makes more sense.
    9. What are our constraints?
        - we are limited to a set to denominations ( coins array )
        - we have to pick and choose certain coins that add up to target ( amount )
        - so we have 2 constraints - coins array and amount
        - coins array ( each idx ) will map to rows
        - amount - will map to number of columns ( if amount is 11, the number of cols will be 11, each idx mapping to a small subproblem amount )
        - since we are solving bottom up, and we will be solving the smallest subproblem first, ( the bottom leaf of your decision tree )
    10. The data within our DP matrix will be min number of coins in each cell for each subproblem
    11. Draw out your matrix and see if can spot a pattern for answering repeated subproblems
        - for this question
        - the 0 case ( not choose ) - is right above the current cell
        - the 1 case ( choose case ) - is $coin steps back ( row idx will be our way to identify what coin we are currently using ) + 1
            - why +1 ? -- because we picked this coin (so count++) and count of the new repeated subproblem which is exactly $coin steps back in the same row
        - then we need to return the min number of coins between choose and not choose case - thats what we will save in each cell.
        - Note: for some coin values, the target amount maybe too small. This is when our coin in hand > amount. 
            - Which means we do not have a choose case, but only choice is to not choose this coin ( 0th case - which is 1 row above - same col )
    12. Finally when we have the whole matrix filled out, the answer is the last row and last col in our matrix.
        - Is this always the case?
        - SPEAK OUT THE PROBLEM... we are trying to find min coins to use to make amount $n
        - because our coins are represented as rows, the last row indicates min coins - for all coins from 0th row till last row
        - the last row represents the final coins list ( i.e use all coins we have from input )
        - therefore last row, but why last col?
        - Our subproblem amounts start from 1 to $amount, which means our cols will be representing from $1 to $n ( left to right )
        - if right is the total amount we are trying to make, then thats why the last cell.
*/

// brute force - dfs - explore all paths
// result TLE
// func coinChange(coins []int, amount int) int {
//     min := math.MaxInt64
//     var dfs func(start int, amount int, count int)
//     dfs = func(start, amount, count int) {
//         // base
//         if amount == 0 {
//             if count < min {min=count}
//             return
//         }
//         if amount < 0 || start == len(coins) {
//             return
//         }
//         // logic
//         dfs(start, amount-coins[start], count+1)
//         dfs(start+1, amount, count)
//     }
//     dfs(0, amount, 0)
//     return min
// }


// // bottom up dp - using dp matrix
// time: o(mn); space: o(mn)
// func coinChange(coins []int, amount int) int {
//     // dp matrix constraints
//     m := len(coins)
//     n := amount
    
//     dp := make([][]int, m+1)
//     for idx, _ := range dp {
//         dp[idx] = make([]int, n+1)
//     }
//     for j := 1; j < len(dp[0]); j++ {
//         dp[0][j] = amount+1
//     }
    
//     for i := 1; i < len(dp); i++ {
//         for j := 1; j < len(dp[0]); j++ {
//             coin := coins[i-1]
//             amount := j
//             if coin > amount {
//                 dp[i][j] = dp[i-1][j]
//             } else {
//                 dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-coin]+1)))
//             }
//         }
//     }
//     result := dp[len(dp)-1][len(dp[0])-1]
//     if result == amount+1 {return -1}
//     return result
// }


// bottom up dp - using dp array
// optimization of dp matrix
// in dp matrix we can see that we are only ever referencing a row above value
// so we do not really need a matrix, we can use a 1d array and keep on overwriting
// the 0 case ( not choose case / when coin > amount ) - we were getting from same col, but a row above
// in dp 1d array, our current ith position will be previous row value ( since we have changed it yet )
// so 0 case - continue ( do not change the value and pretend that this is row above value - therefore we do not change it - since in matrix we simply took row above value)
// for 1 case ( choose ) - we will follow the same math.min ( because we still want min coins used between choose and not choose cases )
// our choose case for next smaller subproblem was $coin steps back in the same row and not choose case row above but same col
// in 1d dp array - not choose case is same ith position and choose case is i-$coinStepsBack
// this optimizes our dp space
// however filling out the array will follow same nested for loop format
// time: o(mn)
// space: o(n) = n is amount
func coinChange(coins []int, amount int) int {
    // dp matrix constraints
    m := len(coins)
    n := amount
    
    dp := make([]int, n+1)
    for j := 1; j < len(dp); j++ {
        dp[j] = amount+1
    }
    
    for i := 1; i < m+1; i++ {
        for j := 1; j < len(dp); j++ {
            coin := coins[i-1]
            am := j
            if coin > am {continue}
            dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coin]+1)))
        }
    }
    result := dp[len(dp)-1]
    if result == amount+1 {return -1}
    return result
}
// exactly like coin change 
// the only thing in coin change was we were given the coins allowed to be used
// here we have to come up with the coins ( i.e all perfect squares from 1 <= n  ----- as long as i*i <= n )
// brute force: dfs - TLE
// func numSquares(n int) int {
//     sqNums := []int{}
//     for i := 1; i <= n; i++ {
//         sq := int(math.Pow(float64(i), float64(2)))
//         if sq == i*i {
//             sqNums = append(sqNums, sq)
//         }
//     }
//     var dfs func(target int, ptr int, count int) (int, bool)
//     dfs = func(target, ptr, count int) (int, bool) {
//         // base
//         if target == 0 {return count, true}
//         if ptr == len(sqNums) || target < 0 {return math.MaxInt64, false}
        
//         // logic
//         notChoose, lFound := dfs(target, ptr+1, count)
//         choose, rFound := dfs(target-sqNums[ptr], ptr, count+1)
        
//         if lFound && !rFound {return notChoose, true}
//         if !lFound && rFound {return choose, true}
//         if lFound && rFound {return min(choose, notChoose), true}
//         return math.MaxInt64, false
//     }
//     val, _ := dfs(n, 0, 0)
//     return val
// } 


// dp bottom up
// time: o()
// func numSquares(target int) int {
//     sqNums := []int{}
//     maxSqrtIdx := int(math.Sqrt(float64(target)))+1
//     for i := 1; i < maxSqrtIdx; i++ {
//         sqNums = append(sqNums, i*i)
        
//     }
//     // fmt.Println(sqNums)
//     m := len(sqNums)+1
//     n := target+1
//     dp := make([][]int, m)
//     for i := 0; i < m; i++ {
//         dp[i] = make([]int, n)
//     }
    
//     for i := 1; i < n; i++ {
//         dp[0][i] = math.MaxInt64
//     }
    
//     for i := 1; i < m; i++ {
//         for j := 1; j < n; j++ {
//             sqNum := sqNums[i-1]
//             tg := j
//             if sqNum > tg {
//                 dp[i][j] = dp[i-1][j]
//             } else {
//                 dp[i][j] = min(dp[i-1][j], dp[i][j-sqNum]+1)
//             }
//         }
//     }
    
//     return dp[m-1][n-1]
// }



func numSquares(target int) int {
    sqNums := []int{}
    for i := 1; i <= target; i++ {
        if i*i > target {break}
        sqNums = append(sqNums, i*i)
        
    }
    // fmt.Println(sqNums)
    m := len(sqNums)+1
    n := target+1
    dp := make([]int, n)
    
    for i := 1; i < n; i++ {
        dp[i] = math.MaxInt64
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            sqNum := sqNums[i-1]
            tg := j
            if sqNum > tg {
                continue
            }
            dp[j] = min(dp[j], dp[j-sqNum]+1)
            
        }
    }
    
    return dp[n-1]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}
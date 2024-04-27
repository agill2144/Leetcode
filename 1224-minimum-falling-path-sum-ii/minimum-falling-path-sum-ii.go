func minFallingPathSum(grid [][]int) int {
    smallest := math.MaxInt64
    smallestIdx := -1
    secondSmallest := math.MaxInt64

    nSmallest := math.MaxInt64
    nSmallestIdx := -1
    nSecondSmallest := math.MaxInt64

    ans := math.MaxInt64
    m := len(grid)
    n := len(grid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i != 0 {
                if j == smallestIdx {
                    grid[i][j] += secondSmallest
                } else {
                    grid[i][j] += smallest
                }
            }
            
            val := grid[i][j]
            if val < nSmallest {
                nSecondSmallest = nSmallest
                nSmallest = val
                nSmallestIdx = j
            }
            if val < nSecondSmallest &&  j != nSmallestIdx {
                nSecondSmallest = val
            }

        }
        smallest = nSmallest
        smallestIdx = nSmallestIdx
        secondSmallest = nSecondSmallest
        nSmallest = math.MaxInt64
        nSmallestIdx = -1
        nSecondSmallest = math.MaxInt64
    }
    for j := 0; j < n; j++ {
        ans = min(ans, grid[m-1][j])
    }    
    return ans
}

// like min path falling sum 1
// but instead of only 2 choices per cell
// we have n choices per cell
// time: o(n^3)
// space = o(1)
// func minFallingPathSum(grid [][]int) int {
//     m := len(grid)
//     n := len(grid[0])
//     ans := math.MaxInt64
    
//     for i := 1; i < m; i++ {
//         for j := 0; j < n; j++ {        
//             minSum := math.MaxInt64
//             for k := 0; k < n; k++ {
//                 if k == j {continue}
//                 minSum = min(minSum, grid[i-1][k] + grid[i][j])
//             }
//             grid[i][j] = minSum
//         }
//     }
//     for j := 0; j < n; j++ {
//         ans = min(ans, grid[m-1][j])
//     }
//     return ans
// }

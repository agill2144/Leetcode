func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            curr := mat[i][j]
            left := math.MinInt64
            if j-1 >= 0 {left = mat[i][j-1]}
            right := math.MinInt64
            if j+1 < n {right = mat[i][j+1]}
            top := math.MinInt64
            if i-1 >= 0 {top = mat[i-1][j]}
            bottom := math.MinInt64
            if i+1 < m {bottom = mat[i+1][j]}
            if curr > left && curr > right &&
                curr > top && curr > bottom {
                    return []int{i, j}
            }
        }
    }
    return nil
}
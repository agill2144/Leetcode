func generate(numRows int) [][]int {
    // num of cols in a row i is i+1
    // col[0] and col[n-1] is reserved by 1
    // in-between col vals are out[i-1][j] + out[i-1][j-1]
    out := make([][]int, numRows)
    for i := 0; i < len(out); i++ {
        n := i+1
        out[i] = make([]int, n)
        out[i][0] = 1
        out[i][n-1] = 1
        for j := 1; j < n-1; j++ {
            out[i][j] = out[i-1][j] + out[i-1][j-1]
        }
    }
    return out
}
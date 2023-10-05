func generate(numRows int) [][]int {
    n := numRows
    out := make([][]int, n)
    for i := 0; i < len(out); i++ {
        out[i] = make([]int, i+1)
        for j := 0; j < len(out[i]); j++ {
            if j == 0 || j == len(out[i])-1 {
                out[i][j] = 1
            } else {
                out[i][j] = out[i-1][j] + out[i-1][j-1]
            }
        }
    }
    return out
}
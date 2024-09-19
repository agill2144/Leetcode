func restoreMatrix(rowSum []int, colSum []int) [][]int {
    m := len(rowSum)
    n := len(colSum)
    currColSum := 0
    mat := make([][]int, m)
    for i := 0; i < m; i++ {
        mat[i] = make([]int, n)
        mat[i][0] = rowSum[i]
        currColSum += rowSum[i]
    }
    c := 0
    for c < n-1 {
        r := 0
        desiredSum := colSum[c]
        toRemove := abs(currColSum-desiredSum)
        nextColSum := toRemove
        for r < m {
            if mat[r][c] <= toRemove {
                mat[r][c+1] = mat[r][c]
                toRemove -= mat[r][c]
                mat[r][c] = 0
            } else {
                mat[r][c+1] = toRemove
                mat[r][c] -= toRemove
                toRemove = 0
            }
            r++
        }
        c++
        currColSum = nextColSum
    }
    return mat
}

func abs(x int) int {
    if x < 0 {return x *- 1}
    return x
}
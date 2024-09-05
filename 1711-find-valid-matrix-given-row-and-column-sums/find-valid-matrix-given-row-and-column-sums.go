func restoreMatrix(rowSum []int, colSum []int) [][]int {
    m := len(rowSum)
    n := len(colSum)
    matrix := make([][]int, m)
    currColSum := 0
    for i := 0; i < m; i++ {
        matrix[i] = make([]int, n)
        matrix[i][0] = rowSum[i]
        currColSum += rowSum[i]
    }
    r := 0
    c := 0
    for c < n-1 {
        desiredColSum := colSum[c]
        // if this is what we are removing from curr col
        // and moving these value to next col ( next col is all 0 )
        // this also means, that the next col sum will be what we will remove from this col        
        toRemove := abs(currColSum-desiredColSum)
        // therefore nextCol sum == the sum we will remove from current col
        nextColSum := abs(currColSum-colSum[c])

        // go thru the current col, and start moving vals to the right
        for r < m && toRemove > 0 {
            if matrix[r][c] <= toRemove {
                // we can remove all
                toRemove -= matrix[r][c]
                matrix[r][c+1] = matrix[r][c]
                matrix[r][c] = 0
            } else if toRemove > 0 {
                // we dont need all
                // eg: matrix val = 15, to remove = 2
                matrix[r][c] -= toRemove
                matrix[r][c+1] = toRemove
                toRemove = 0
            }
            r++
        }
        c++
        r = 0
        currColSum = nextColSum
    }
    return matrix
}

func abs(x int) int {
    if x < 0 {return x *- 1}
    return x
}
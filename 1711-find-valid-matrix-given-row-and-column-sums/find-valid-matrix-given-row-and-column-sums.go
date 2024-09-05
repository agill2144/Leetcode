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
    j := 0
    for j < n-1 {
        toRemove := abs(currColSum-colSum[j])
        nextColSum := abs(currColSum-colSum[j])
        // go thru the current col, and start moving vals to the right
        for r < m {
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
                break
            }
            r++
        }
        // fmt.Println("after col: ", j)
        // fmt.Println(matrix)
        // fmt.Println("-----------------")
        c++
        j++
        r = 0
        currColSum = nextColSum
    }
    return matrix
}

func abs(x int) int {
    if x < 0 {return x *- 1}
    return x
}
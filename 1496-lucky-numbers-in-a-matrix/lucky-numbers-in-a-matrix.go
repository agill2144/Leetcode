func luckyNumbers (matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    minSet := map[int]struct{}{}
    for i := 0; i < m; i++ {
        minRow := math.MaxInt64
        for j := 0; j < n; j++ {
            minRow = min(minRow, matrix[i][j])
        }
        minSet[minRow] = struct{}{}
    }
    out := []int{}
    for j := 0; j < n; j++ {
        maxCol := math.MinInt64
        for i := 0; i < m; i++ {
            maxCol = max(maxCol, matrix[i][j])
        }
        if _, ok := minSet[maxCol]; ok {out = append(out, maxCol)}
    }
    return out
}
func checkValid(matrix [][]int) bool {
    n := len(matrix)
    for i := 0; i < n; i++ {
        rowSet := map[int]struct{}{}
        colSet := map[int]struct{}{}
        for j := 0; j < n; j++ {
            rowVal := matrix[i][j]
            colVal := matrix[j][i]
            
            if rowVal < 1 || rowVal > n {return false}
            if colVal < 1 || colVal > n {return false}
            
            if _, ok := rowSet[rowVal]; ok {return false}
            rowSet[rowVal] = struct{}{}

            if _, ok := colSet[colVal]; ok {return false}
            colSet[colVal] = struct{}{}

        }
    }
    return true
}
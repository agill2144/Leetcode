func isValidSudoku(board [][]byte) bool {
    n := len(board)
    boxes := map[string]map[byte]bool{}
    rows := map[int]map[byte]bool{}
    cols := map[int]map[byte]bool{}
    for i := 0; i < n; i++ {
        if rows[i] == nil {rows[i] = map[byte]bool{}}
        for j := 0; j < n; j++ {
            val := board[i][j]
            if val != '.'{
                if rows[i][val] {return false}
                rows[i][val] = true
                r, c := i/3, j/3
                key := fmt.Sprintf("%v_%v", r, c)
                if boxes[key] == nil {boxes[key] = map[byte]bool{}}
                if boxes[key][val] {return false}
                boxes[key][val] = true

            }
            if cols[i] == nil {cols[i] = map[byte]bool{}}
            val = board[j][i]
            if val != '.' {
                if cols[i][val] {return false}
                cols[i][val] = true
            }
        }
    }
    return true
}
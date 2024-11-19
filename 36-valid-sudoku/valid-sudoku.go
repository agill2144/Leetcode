func isValidSudoku(board [][]byte) bool {
    n := len(board)
    boxes := map[string]map[byte]bool{}
    for i := 0; i < n; i++ {
        row := map[byte]bool{}
        col := map[byte]bool{}
        for j := 0; j < n; j++ {
            val := board[i][j]
            if val != '.'{
                if row[val] {return false}
                row[val] = true
                r, c := i/3, j/3
                key := fmt.Sprintf("%v_%v", r, c)
                if boxes[key] == nil {boxes[key] = map[byte]bool{}}
                if boxes[key][val] {return false}
                boxes[key][val] = true

            }

            val = board[j][i]
            if val != '.' {
                if col[val] {return false}
                col[val] = true
            }
        }
    }
    return true
}
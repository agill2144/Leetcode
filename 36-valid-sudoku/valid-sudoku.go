func isValidSudoku(board [][]byte) bool {
    n := len(board)
    box := map[string]map[byte]bool{}
    for i := 0; i < n; i++ {
        row := map[byte]bool{}
        col := map[byte]bool{}
        for j := 0; j < n; j++ {
            ij := board[i][j]
            key := fmt.Sprintf("%v-%v", i/3,j/3)
            if box[key] == nil {box[key]=map[byte]bool{}}
            if ij != '.' {
                if row[ij] {return false}
                row[ij] = true
                if box[key][ij] {return false}
                box[key][ij] = true
            }

            ji := board[j][i]
            if ji != '.' {
                if col[ji] {return false}
                col[ji] = true
            }

        }
    }
    return true
}
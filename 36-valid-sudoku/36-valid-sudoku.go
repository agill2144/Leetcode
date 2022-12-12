/*
    approach: Using multiple sets
    - set for each row
    - set for each col
    - list of sets for each 3x3 box
    
    - Each set will ensure we do not have a dupe for that set's purpose
*/
func isValidSudoku(board [][]byte) bool {
    m := len(board)
    n := len(board[0])
    boxSet := make([]map[byte]struct{}, m)
    
    for i := 0; i < m; i++ {
        rowSet := map[byte]struct{}{}
        colSet := map[byte]struct{}{}
        for j := 0; j < n; j++ {
            
            boxSetIdx :=  (i/3) * 3 + (j/3)
            
            if board[i][j] != '.' {
                _, existsInRowSet := rowSet[board[i][j]]
                if existsInRowSet {
                    return false
                } 
                rowSet[board[i][j]] = struct{}{}
                
                if _, ok := boxSet[boxSetIdx][board[i][j]]; ok {
                    return false
                }
                if boxSet[boxSetIdx] == nil {boxSet[boxSetIdx] = map[byte]struct{}{}}
                boxSet[boxSetIdx][board[i][j]] = struct{}{}            
            }

            if  board[j][i] != '.' {
                _, existsInColSet := colSet[board[j][i]]
                if existsInColSet {
                    return false
                }
                colSet[board[j][i]] = struct{}{}
            }            
        } 
    }
    
    return true
}
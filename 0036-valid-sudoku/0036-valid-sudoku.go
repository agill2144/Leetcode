func isValidSudoku(board [][]byte) bool {
    m := len(board)
    n := len(board[0])
    boxMap := map[string]*set{}
    
    for i := 0; i<m; i++ {
        rowSet := NewSet()
        colSet := NewSet()
        for j := 0; j < n; j++ {
            rowVal := board[i][j]
            colVal := board[j][i]
            if rowVal != '.' {
                if rowSet.contains(rowVal) {return false}
                rowSet.add(rowVal)
                
                boxKey := fmt.Sprint("%v-%v", i/3, j/3)
                boxSet := boxMap[boxKey]
                if boxSet == nil {
                    boxSet = NewSet()
                    boxMap[boxKey] = boxSet
                }
                if boxSet.contains(rowVal) {return false}
                boxSet.add(rowVal)
            
            }
            if colVal != '.' {
                if colSet.contains(colVal) {return false}
                colSet.add(colVal)
            }
        }
    }
    return true
    
}

type set struct{
    items map[byte]struct{}
}
func NewSet() *set {
    return &set{items: map[byte]struct{}{}}
}
func (s *set) add(x byte) {
    s.items[x] = struct{}{}
}
func (s *set) contains(x byte) bool {
    _, ok := s.items[x]
    return ok
}
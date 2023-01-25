func isValidSudoku(board [][]byte) bool {
    m := len(board)
    n := len(board[0])
    boxMap := map[string]*set{}
    
    for i := 0; i < m; i++ {
        rowSet := newSet()
        colSet := newSet()

        for j := 0; j < n; j++ {
            
            if board[i][j] != '.' {
                if rowSet.contains(board[i][j]) {
                    return false
                }
                rowSet.add(board[i][j])
                
                boxKey := fmt.Sprintf("%v:%v", i/3, j/3)
                boxSet , ok := boxMap[boxKey]
                if !ok {
                    boxSet = newSet()
                    boxMap[boxKey] = boxSet
                }
                if boxSet.contains(board[i][j]) {return false}
                boxSet.add(board[i][j])
            }
            
            if board[j][i] != '.' {
                if colSet.contains(board[j][i]) {
                    return false
                }
                colSet.add(board[j][i])
            }
        }
    }
    return true
    
}

type set struct {
    items map[byte]struct{}
}

func newSet() *set {
    return &set{items: map[byte]struct{}{}}
}
func (this *set) add(x byte) {
    this.items[x] = struct{}{}
}
func (this *set) contains(x byte) bool {
    _, ok := this.items[x]
    return ok
}
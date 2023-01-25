// time: o(1) - the matrix is given to be 9x9 in all test cases
// space: o(1) -  the matrix is given to be 9x9 in all test cases
// because worst case board is all filled out and our box map has 9^2 elements... 
// so its constant and never grows more than o(numOfElementsInRowSet) + o(numOfElementsInColSet) + o(9^2)
// space: o(9)+o(9)+o(81) -- so o(1) - these are constant numbers IMO

func isValidSudoku(board [][]byte) bool {
    m := len(board)
    n := len(board[0])
    boxMap := map[int]*set{}
    
    for i := 0; i < m; i++ {
        rowSet := newSet()
        colSet := newSet()
        for j := 0; j < n; j++ {
            if board[i][j] != '.' {
                if rowSet.contains(board[i][j]) {
                    return false
                }
                rowSet.add(board[i][j])
                
                boxKey := 0
                boxKey = boxKey * 10 + (i/3)
                boxKey = boxKey * 10 + (j/3)                
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
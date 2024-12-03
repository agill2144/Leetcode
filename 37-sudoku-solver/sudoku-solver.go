func solveSudoku(board [][]byte)  {
    n := len(board)
    rowSet := make([]*set, 9)
    colSet := make([]*set, 9)
    boxSet := map[string]*set{}
    sr, sc := -1, -1
    for i := 0; i < n; i++ {
        if rowSet[i] == nil {rowSet[i] = newSet()}
        if colSet[i] == nil {colSet[i] = newSet()}
        for j := 0; j < n; j++ {
            boxKey := fmt.Sprintf("%v-%v", i/3, j/3)
            if boxSet[boxKey] == nil { boxSet[boxKey] = newSet() }
            if board[i][j] != '.' {
                rowSet[i].add(board[i][j])
                boxSet[boxKey].add(board[i][j])            
            }
            if board[j][i] != '.' {
                colSet[i].add(board[j][i])
            }
            if board[i][j] == '.' && sr == -1 {sr = i; sc = j}
        }
    }
    choices := []byte{'1','2','3','4','5','6','7','8','9'}
    var dfs func(r, c int) bool
    dfs = func(r, c int)bool {
        // base
        if r == n {return true}
        if c == n {
            return dfs(r+1, 0)
        }
        if board[r][c] != '.' {
            return dfs(r, c+1)
        }


        // logic
        boxKey := fmt.Sprintf("%v-%v",r/3,c/3)
        rowData := rowSet[r]
        colData := colSet[c]
        boxData := boxSet[boxKey]
        for i := 0; i < len(choices); i++ {
            choice := choices[i]
            inRow := rowData.contains(choice)
            inCol := colData.contains(choice)
            inBox := boxData.contains(choice)
            if !inRow && !inCol && !inBox {
                board[r][c] = choice
                rowData.add(choice)
                colData.add(choice)
                boxData.add(choice)
                if dfs(r, c+1) {return true}
                boxData.remove(choice)
                colData.remove(choice)
                rowData.remove(choice)
                board[r][c] = '.'
            }
        }
        return false
    }
    dfs(sr,sc)
}


type set struct {
    items map[byte]struct{}
}

func newSet() *set {
    return &set{items: map[byte]struct{}{}}
}

func (s *set) add(key byte) {
    s.items[key] = struct{}{}
}

func (s *set) remove(key byte) {
    delete(s.items, key)
}

func (s *set) contains(key byte) bool {
    _, ok := s.items[key]
    return ok
}


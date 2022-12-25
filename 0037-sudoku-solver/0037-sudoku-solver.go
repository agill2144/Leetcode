/*
    time: something exponential (if someone asks me this question, I am walking out...)
*/

func solveSudoku(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    choices := []byte{'1','2','3','4','5','6','7','8','9'}
    rows := make([]*set, 9)
    cols := make([]*set, 9)
    boxMap := map[string]*set{}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rows[i] == nil {rows[i] = newSet()}
            if cols[i] == nil {cols[i] = newSet()}
            rows[i].add(board[i][j])
            cols[i].add(board[j][i])
            boxIdx := fmt.Sprintf("%v:%v", i/3,j/3)
            if val := boxMap[boxIdx]; val == nil {
                boxMap[boxIdx] = newSet()
            }
            boxMap[boxIdx].add(board[i][j])
        }
    }
    
    
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if r == m-1 && c == n {return true}
        if c == n {
            r++
            c = 0
        }
        
        if board[r][c] != '.' {
            return dfs(r, c+1)
        }
        
        boxIdx := fmt.Sprintf("%v:%v", r/3,c/3)
        boxSet := boxMap[boxIdx]
        rowSet := rows[r]
        colSet := cols[c]
        
        // logic
        for i := 0; i < len(choices); i++ {
            choice := choices[i]

            if !rowSet.contains(choice) && !colSet.contains(choice) &&
            (boxSet == nil || !boxSet.contains(choice)) {
                // action
                
                board[r][c] = choices[i]
                rowSet.add(choices[i])
                colSet.add(choices[i])
                if boxSet == nil {boxMap[boxIdx] = newSet()}
                boxMap[boxIdx].add(choices[i])

                // recurse
                if done := dfs(r, c+1); done {return true}

                // backtrack
                board[r][c] = '.'
                rowSet.remove(choices[i])
                colSet.remove(choices[i])
                boxMap[boxIdx].remove(choices[i])
            }
        
        }
        return false
    }
    dfs(0,0)
}



type set struct{
    items map[byte]struct{}
}
func newSet() *set {
    return &set{items: map[byte]struct{}{}}
}
func (s *set)contains(x byte) bool {
    _, ok := s.items[x]
    return ok == true
}
func (s *set)add(x byte) {
    s.items[x]  = struct{}{}
}
func (s *set)remove(x byte) {
    delete(s.items, x)
} 

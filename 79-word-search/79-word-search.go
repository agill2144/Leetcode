func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
    }
    
    
    var dfs func(r, c, wp int) bool
    dfs = func(r,c,wp int) bool {
        // base
        if wp == len(word) {
            return true
        }
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != word[wp] {
            return false
        }
        
        
        
        // logic
        tmp := board[r][c]
        for _, dir := range dirs {
            // action
            board[r][c] = '.'
            nr := r+dir[0]
            nc := c+dir[1]
            // recurse
            ok := dfs(nr,nc,wp+1)
            if ok {return true}
            // backtrack
            board[r][c] = tmp
        }
        return false
    }
    
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                found := dfs(i,j,0)
                if found {
                    return true
                }
            } 
        }
    }
    
    return false
}
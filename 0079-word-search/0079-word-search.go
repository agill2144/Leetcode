func exist(board [][]byte, word string) bool {
    dirs := [][]int{
        {1,0},{-1,0},
        {0,-1}, {0,1},
    }
    m := len(board)
    n := len(board[0])
    var dfs func(r, c, ptr int) bool
    dfs = func(r,c ,ptr int) bool {
        // base
        if ptr == len(word) {return true}
        
        // logic
        for _, dir := range dirs {
            nr := r + dir[0]
            nc := c + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == word[ptr] {
                // action
                oldChar := board[nr][nc]
                board[nr][nc] = '.'
                // recurse
                found := dfs(nr,nc,ptr+1)
                if found {return true}
                // backtrack
                board[nr][nc] = oldChar
            }
        }
        return false
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                board[i][j] = '.'
                ok := dfs(i, j, 1)
                if ok {return true}
                board[i][j] = word[0]
            }
        }
    }
    return false
}
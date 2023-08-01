func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    
    // used to mark all O's that are not flippable
    var dfs func(r , c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != 'O' {return}
        
        // logic
        board[r][c] = '.' // . reprensents a previous O that cannot be flipped
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    
    for i := 0; i < m; i++ {
        if board[i][0] == 'O' {dfs(i, 0)}
        if board[i][n-1] == 'O' {dfs(i, n-1)}
    }
    for j := 0; j < n; j++ {
        if board[0][j] == 'O' {dfs(0, j)}
        if board[m-1][j] == 'O' {dfs(m-1, j)}
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == '.' {
                board[i][j] = 'O'
            }
        }
    }
    
    
}
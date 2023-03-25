func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != 'O' {return}
        
        // logic
        board[r][c] = '0'
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    
    // row 0 and m-1
    for i := 0 ; i < n; i++ {
        if board[0][i] == 'O' {dfs(0,i)}
        if board[m-1][i] == 'O' {dfs(m-1,i)}
    }
    // col 0 and n-1
    for i := 0; i < m; i++ {
        if board[i][0] == 'O' {dfs(i,0)}
        if board[i][n-1] == 'O' {dfs(i,n-1)}
    }

    
    for i := 0; i < m ; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {board[i][j] = 'X'}
            if board[i][j] == '0' {board[i][j] = 'O'}
        }
    }
    
}
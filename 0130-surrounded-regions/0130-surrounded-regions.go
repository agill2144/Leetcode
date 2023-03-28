/*
    approach: bfs/dfs
    - we have to flip all O's to X if they are surrounded by X in all 4 dirs
    - The only O's that cannot be flipped are at the boundary cells
    - O's connected to boundary O's also cannot be flipped since the boundary ones cannot be flipped
    
    - So we need to find all O's that are connected to a boundary O's, because these are the ones that can never be flipped
    - therefore we can run a bfs / dfs to on boundary cells that are O's
        - mark all the connected O's with some identifier that denotes these cells cannot be replaced
    - because there are can many such connected components, we have to run dfs / bfs on all
    - then finally go over the board and flip the remaining O's to X's and revert back the irreplaceable identifiers back to O's
    
    time : o(mn)
    space : o(mn)
*/

func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    // use a dfs to mark all O's that can never be flipped
    // because they are connected to a boundary cell
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != 'O' {return}
        
        // logic
        // mark a O as E to identify as can-never-be-flipped
        board[r][c] = 'E'
        
        // find all other O's connected to this, because if this
        // cell cannot be flipped, no one else connected to it can be.
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    
    
    // now find all Os that cannot be flipped ( i.e the ones that are on the boundaries of the matrix )
    for j := 0; j < n; j++ {
        if board[0][j] == 'O' {dfs(0, j)}
        if board[m-1][j] == 'O' {dfs(m-1, j)}
    }
    for i := 0; i < m; i++ {
        if board[i][0] == 'O' {dfs(i, 0)}
        if board[i][n-1] == 'O' {dfs(i, n-1)}
    }
    
    // now reflip the E's ( never flippable ) back to O and flip the remaining O's to to X
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'E' {board[i][j] = 'O'} else if board[i][j] == 'O' {board[i][j] = 'X'} 
        }
    }
}

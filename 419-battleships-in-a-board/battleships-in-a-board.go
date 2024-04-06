func countBattleships(board [][]byte) int {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,-1},
        {0,1},
    }
    var dfs func(r, c int) 
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != 'X'{
            return
        } 
        
        // logic
        board[r][c] = '.'
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'X'{
                dfs(i,j)
                count++
            }
        }
    }
    return count
    
    
}
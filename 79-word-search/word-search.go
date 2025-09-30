func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{
        {-1,0},
        {1,0},
        {0,-1},
        {0,1},
    }
    var dfs func(r, c, ptr int) bool
    dfs = func(r, c, ptr int) bool {
        // base
        if ptr == len(word) {return true}
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != word[ptr] {return false}

        // visit current cell
        tmp := board[r][c]
        board[r][c] = '.'
        for _, dir := range dirs {
            if dfs(r+dir[0],c+dir[1],ptr+1) {return true}
        }
        board[r][c] = tmp
        return false
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                if dfs(i,j,0) {return true}
            }
        }
    }
    return false
}
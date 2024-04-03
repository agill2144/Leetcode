func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c, ptr int) bool
    dfs = func(r, c, ptr int) bool {
        // base
        if ptr == len(word) {return true}


        // logic
        for _, dir := range dirs {
            nr, nc := r+dir[0], c+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && word[ptr] == board[nr][nc] {
                // action
                tmp := board[nr][nc]
                board[nr][nc] = '.'
                // recurse
                if dfs(nr,nc,ptr+1) {
                    return true
                }
                // backtrack
                board[nr][nc] = tmp
            }
        }
        return false
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                board[i][j] = '.'
                if dfs(i, j, 1) {return true}
                board[i][j] = word[0]
            }
        }
    }
    return false
}

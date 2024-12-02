func exist(board [][]byte, word string) bool {
    var visited byte = '.'
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int, ptr int) bool
    dfs = func(r, c int, ptr int) bool {
        // base
        if ptr == len(word) {return true}
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != word[ptr] {return false}

        // logic
        // action
        tmp := board[r][c]
        board[r][c] = visited
        for _, dir := range dirs {
            nr := r+dir[0]
            nc := c+dir[1]
            if dfs(nr, nc, ptr+1) {board[r][c] = tmp; return true}
        }
        board[r][c] = tmp
        return false
    }

    // m = len(board)
    // n = len(board[0])
    // w = len(word)
    // o(m*n * 3^w)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                if dfs(i, j, 0) {return true}
            }
        }
    }
    return false
}
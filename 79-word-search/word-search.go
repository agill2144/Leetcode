func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var visited byte = '0'
    var dfs func(r, c, w int)bool
    dfs = func(r,c,w int) bool {
        // base
        if w == len(word) {return true}
        if r < 0 || r == m || c < 0 || c == n || board[r][c] != word[w] {return false}

        // logic
        tmp := board[r][c]
        board[r][c] = visited
        for _, dir := range dirs {
            if dfs(r+dir[0], c+dir[1], w+1) {board[r][c] = tmp; return true}
        }
        board[r][c] = tmp
        return false
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] && dfs(i,j,0) {return true}
        }
    }
    return false
}
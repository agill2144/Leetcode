/*
    approach: bfs/dfs
    - we have to flip all O's to X if they are surrounded by X in all 4 dirs
    - The only O's that cannot be flipped are at the boundary cells
    - So we need to find all O's that are connected to a boundary O's, because these are the ones that can never be flipped
    - therefore we can run a bfs / dfs to on boundary cells that are O's
        - mark all the connected O's with some identifier that denotes these cells cannot be replaced
    - because there are can many such connected components, we have to run dfs / bfs on all
    - then finally go over the board and flip the remaining O's to X's and revert back the irreplaceable identifiers back to O's
    
    time : o(mn)
    space : o(mn)
*/

func solve(board [][]byte)  {
    
    // I tried using queue independant nodes ( X ) and finding all 0s next to them but thats incorrect
    // this problem is not as straightforward as flip all 0s next to any X
    
    // We need to mark all 0's cannot be flipped.
    // The 0s that cannot be flipped are on the border + ALL 4 directions from that border 0 + its neighbours + its neighbours
    // More like, the 0s that cannot be flipped are the 0's that form an island from a border 0.
    
    // So we look for all 0 islands starting from a border, and mark all lands of that island "not-flippable" because its connected to a 0 that is on the border.
    // Once thats done, we loop over entire matrix and look for 0's that ARE NOT marked "not-flippable" and flip them, and as soon as we run into a 
    // "not-flippable" 0, we will remove that marker ( undo our action )
    
    
    m := len(board)
    n := len(board[0])
    fmt.Println(m,n)
    // go thru all borders and find all connected 0's and mark them not-flippable
    q := [][]int{}
    // row 0 and m-1
    for i := 0 ; i < n; i++ {
        if board[0][i] == 'O' {board[0][i] = 'E';       q = append(q, []int{0,i})}
        if board[m-1][i] == 'O' {board[m-1][i] = 'E';   q = append(q, []int{m-1,i})}
    }
    // col 0 and n-1
    for i := 0; i < m; i++ {
        if board[i][0] == 'O' {board[i][0] = 'E';q = append(q, []int{i,0})}
        if board[i][n-1] == 'O' {board[i][n-1] = 'E';q = append(q, []int{i,n-1})}
    }
    
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        for _, dir := range dirs {
            r := dq[0] + dir[0]
            c := dq[1] + dir[1]
            if r >= 0 && r < m && c >=0 && c < n && board[r][c] == 'O'{
                board[r][c] = 'E'
                q = append(q, []int{r,c})
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {board[i][j] = 'X'} else if board[i][j] == 'E' {board[i][j] = 'O'} 
        }
    }
    
}

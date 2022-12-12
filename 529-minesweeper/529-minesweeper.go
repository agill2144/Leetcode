/*
    approach: BFS
    - This is like applying flood fill / domino effect after starting from a start position based on some rules
    - If the start position is a mine, then we reveal the mine on the board and end the game right away, otherwise
    - We will initially enqueue this start position 
    - And process our queue in a BFS manner ( no need to take the queue size as we do not need to do anything after a level )
    - For a position being processed, we will count the number of mines around this cell ( all 8 dirs BUT ONLY IMMEDIATE CELLS )
    - if no mines around immediate cells, then expose this cell with a 'B'
        - also enqueue all childs
    - if mines are found, the place the count in that cell
        - do not enqueue any of this nodes child
    - All this mutation will be happening on the board itself, and we will return the board at the end.
    
    Time: o(mn) 
    - we only ever see a cell once
    space: o(mn)
    - We will enqueue a child and then process that child and then this child will add its other 8 childs depending on the matrix if possible
    - so somewhere at somepoint the space of the queue is proportional to the input matrix size
*/
// func updateBoard(board [][]byte, click []int) [][]byte {
    
//     sr := click[0]
//     sc := click[1]
//     if board[sr][sc] == 'M' {
//         board[sr][sc] = 'X'
//         return board
//     }
//     dirs := [][]int{
//         {1,0},
//         {-1,0},
//         {0,1},
//         {0,-1},
//         {-1,-1},
//         {-1,1},
//         {1,-1},
//         {1,1},
//     }
    
//     q := [][]int{{sr,sc}}
    
    
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         cr := dq[0]
//         cc := dq[1]
//         board[cr][cc] = 'B'
//         numMines, childs := countNumMinesAround(cr, cc, board, dirs)
//         if numMines > 0 {
//             board[cr][cc] = byte(numMines + '0')
//         } else {
//             for _, child := range childs {
//                 board[child[0]][child[1]] = 'B'
//                 q = append(q, child)
//             }
//         }
//     }
//     return board
    
// }



// func countNumMinesAround(r, c int, board [][]byte, dirs [][]int) (int, [][]int) {
//     m := len(board)
//     n := len(board[0])
//     count := 0
//     childs := [][]int{}
    
//     for _, dir := range dirs {
//         nr := r+dir[0]
//         nc := c+dir[1]
//         if nr >= 0 && nr < m && nc >=0 && nc < n {
            
//             if board[nr][nc] == 'M' {
//                 count++
//             }
            
//             if board[nr][nc] == 'E' {
//                 childs = append(childs, []int{nr,nc})
//             }
//         }
//     }
//     return count, childs
// }



/*
    approach: DFS
    - Same logic as BFS but using recursion
    time: o(mn)
    space: o(mn)
*/
func updateBoard(board [][]byte, click []int) [][]byte {
    
    sr := click[0]
    sc := click[1]
    if board[sr][sc] == 'M' {
        board[sr][sc] = 'X'
        return board
    }
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
        {-1,-1},
        {-1,1},
        {1,-1},
        {1,1},
    }
    m := len(board)
    n := len(board[0])
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] == 'B' {return}
        
        // logic
        board[r][c] = 'B'
        numMines := countNumMinesAround(r,c, board, dirs)
        if numMines > 0 {
            board[r][c] = byte(numMines+'0')
        } else {
            for _, dir := range dirs {
                dfs(r+dir[0], c+dir[1])
            }
        }
        
    }
    dfs(sr,sc)
    
    return board
    
}



func countNumMinesAround(r, c int, board [][]byte, dirs [][]int) int {
    m := len(board)
    n := len(board[0])
    count := 0
    
    for _, dir := range dirs {
        nr := r+dir[0]
        nc := c+dir[1]
        if nr >= 0 && nr < m && nc >=0 && nc < n && board[nr][nc] == 'M'  {
            count++ 
        }
    }
    return count
}
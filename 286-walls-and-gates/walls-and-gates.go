func wallsAndGates(rooms [][]int)  {
    empty := 2147483647
    gate := 0
    q := [][]int{}
    m := len(rooms)
    n := len(rooms[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == gate {
                q = append(q, []int{i,j})
            }
        }
    }
    dirs := [][]int{
        {-1,0},{1,0},
        {0,-1},{0,1},
    }
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr, cc := dq[0], dq[1]
        cd := rooms[cr][cc]
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            nd := cd+1
            if nr >= 0 && nr < m && nc >= 0 && nc < n && rooms[nr][nc] == empty {
                q = append(q, []int{nr,nc})
                rooms[nr][nc] = nd
            }
        }
    }

}
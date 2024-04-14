// multi-src bfs
// start from gates and find/fill nearest rooms
// nearest / shortest path without any custom weights = vanilla bfs
func wallsAndGates(rooms [][]int)  {
    m := len(rooms)
    n := len(rooms[0])
    emptyRoomID := 2147483647
    haveEmptyRooms := false
    q := [][]int{} // [r,c,dist]
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == 0 {
                q = append(q, []int{i,j,0})
            }
            if rooms[i][j] == emptyRoomID {
                haveEmptyRooms = true
            }
        }
    }
    if len(q) == 0 {return}
    if !haveEmptyRooms {return}
    
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    // from each gate, look for empty rooms
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            cd := dq[2]
            for _, dir := range dirs {
                nr := cr + dir[0]
                nc := cc + dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && rooms[nr][nc] == emptyRoomID {
                    rooms[nr][nc] = cd+1
                    q = append(q, []int{nr,nc,cd+1})
                }
            }
            qSize--
        }
    }
}
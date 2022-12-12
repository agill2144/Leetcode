func wallsAndGates(rooms [][]int)  {
    if rooms == nil {
        return
    }
    
    m := len(rooms)
    n := len(rooms[0])
    dirs := [][]int{ {1,0}, {-1,0}, {0,1}, {0,-1} }
    q := [][]int{}
    
    // fill up the queue with gates
    for i := 0 ; i < m; i++ {
        for j := 0 ; j < n; j++ {
            if rooms[i][j] == 0 {
                q = append(q, []int{i,j})
            }
        }
    }
    
    // proccess
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        
        // this could be a gate
        // or a empty room
        
        // go find all of its neighbours
        for _, dir := range dirs {
            r := dq[0] + dir[0]
            c := dq[1] + dir[1]
            
            if r >= 0 && r < m && c >= 0 && c < n && rooms[r][c] == 2147483647 {
                currDist := rooms[dq[0]][dq[1]]
                rooms[r][c] = currDist + 1
                q = append(q, []int{r,c})
            }
            
        }
    }
    
}
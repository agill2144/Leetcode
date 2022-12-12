func minKnightMoves(x int, y int) int {

    dirs := [][]int{
        {-2,-1},
        {-2,1},
        {-1,2},
        {1,2},
        {2,1},
        {2,-1},
        {1,-2},
        {-1,-2},
    }

    type position struct{ r,c int }
    steps := 0
    startPos := position{r:0,c:0}
    visited := map[position]struct{}{ startPos: struct{}{} }
    q := []position{startPos}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            
            dq := q[0]
            q = q[1:]
            
            cr := dq.r
            cc := dq.c
            if cr == x && cc == y {
                return steps
            }
            
            for _, dir := range dirs {
                r := cr + dir[0]
                c := cc + dir[1]
                rc := position{r: r, c:c}
                if _, ok := visited[rc]; !ok {
                    visited[rc] = struct{}{}
                    q = append(q, rc)
                }
            }
            qSize--
        }
        steps++
    }
    return -1
}

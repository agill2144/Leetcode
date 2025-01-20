func minKnightMoves(x int, y int) int {
    dirs := [][]int{
        {1,-2},{1,2},{2,-1},{2,1},{-1,-2},{-1,2},{-2,-1},{-2,1},
    }
    startRow := -300
    endRow := 300
    startCol := -300
    endCol := 300
    visited := map[[2]int]bool{[2]int{0,0}:true}
    q := [][]int{{0,0,0}} // <r,c,steps>
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr, cc, cs := dq[0], dq[1], dq[2]
        if cr == x && cc == y {return cs}
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            ns := cs + 1
            if nr >= startRow && nr <= endRow && nc >= startCol && nc <= endCol && !visited[[2]int{nr,nc}] {
                if nr == x && nc == y {return ns}
                visited[[2]int{nr,nc}] = true
                q = append(q, []int{nr,nc,ns})
            }
        }
    }
    return -1
}
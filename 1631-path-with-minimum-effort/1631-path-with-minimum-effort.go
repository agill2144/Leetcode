func minimumEffortPath(heights [][]int) int {
    m := len(heights)
    n := len(heights[0])
    
    dist := make([][]int,m)
    for i := 0; i < len(dist); i++ {
        dist[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dist[i][j] = math.MaxInt64
        }
    }
    dist[0][0] = 0
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    q := [][]int{{0,0}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        currDist := dist[cr][cc]
        
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n {
                maxDiffInPath := max(currDist, abs(heights[nr][nc] - heights[cr][cc]))
                if maxDiffInPath < dist[nr][nc] {
                    dist[nr][nc] = maxDiffInPath
                    q = append(q, []int{nr,nc})
                }
            }
        }

    }
    return dist[m-1][n-1]
}

func abs(x int) int {
    if x < 0 {return x *- 1}
    return x
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
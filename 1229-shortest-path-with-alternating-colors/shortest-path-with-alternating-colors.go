func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
    // start = 0
    // red = 1
    // blue = 2
    type adjListNode struct {
        node int
        color int
    }
    
    adjList := map[int][]adjListNode{}
    for i := 0; i < len(redEdges); i++ {
        u,v := redEdges[i][0], redEdges[i][1]
        adjList[u] = append(adjList[u], adjListNode{v,1})
    }
    for i := 0; i < len(blueEdges); i++ {
        u,v := blueEdges[i][0], blueEdges[i][1]
        adjList[u] = append(adjList[u], adjListNode{v,2})
    }
    
    // [ [S,R,B], [], [] ]
    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, 3)
    }

    out := make([]int, n)
    for i := 0; i < n; i++ {
        out[i] = math.MaxInt64
    }
    
    q := [][]int{{0,0,0}} // node, color dist
    visited[0][0] = true
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        node, color, dist := dq[0], dq[1], dq[2]
        out[node] = min(dist, out[node])
        
        for _, nei := range adjList[node] {
            neiNode := nei.node
            neiColor := nei.color
            if neiColor == color {continue}
            if !visited[neiNode][neiColor] {
                visited[neiNode][neiColor] = true
                q = append(q, []int{neiNode, neiColor, dist+1})
            }
        }
    }
    
    for i := 0; i < len(out); i++ {
        if out[i] == math.MaxInt64 {
            out[i] = -1
        }
    }
    return out

}
func closestMeetingNode(edges []int, node1 int, node2 int) int {
    n := len(edges)
    node1Dist := bfs(node1, n, edges)
    node2Dist := bfs(node2, n, edges)
    minDist := math.MaxInt64
    minDistNode := -1
    for i := 0; i < n; i++ {
        if node1Dist[i] != math.MaxInt64 && node2Dist[i] != math.MaxInt64 {
            if max(node1Dist[i], node2Dist[i]) < minDist {
                minDistNode = i
                minDist = max(node1Dist[i], node2Dist[i])
            }
        }
    }
    return minDistNode
}

func max(x, y int) int {
    if x > y {return x}
    return y
}

func bfs(start int, n int, edges []int) []int {
    visited := make([]bool, n)
    dist := make([]int, n)
    for i := 0; i < len(dist); i++ {
        dist[i] = math.MaxInt64
    }
    
    q := []int{start}
    visited[start] = true
    dist[start] = 0
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        currDist := dist[dq]
        nei := edges[dq] 
        neiDist := currDist+1
        if nei != -1 && (!visited[nei] ) {
            q = append(q, nei)
            dist[nei] = neiDist
            visited[nei] = true
        }
        
    }
    return dist
}
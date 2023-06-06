func closestMeetingNode(edges []int, node1 int, node2 int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := i
        dest := edges[i]
        if dest == -1 {continue}
        adjList[src] = append(adjList[src], dest) 
    }
    n := len(edges)
    node1Dist := bfs(node1, n, adjList)
    node2Dist := bfs(node2, n, adjList)
    minDist := math.MaxInt64
    minDistNode := -1
    for i := 0; i < n; i++ {
        if node1Dist[i] != math.MaxInt64 && node2Dist[i] != math.MaxInt64 {
            max := int(math.Max(float64(node1Dist[i]), float64(node2Dist[i])))
            if  max < minDist {
                minDistNode = i
                minDist = max
            }
        }
    }
    
    return minDistNode
}

func bfs(start int, n int, adjList map[int][]int) []int {
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
        
        node := dq
        currDist := dist[node]
        
        for _, nei := range adjList[node] {
            neiDist := currDist+1
            if !visited[nei] || neiDist < dist[nei] {
                q = append(q, nei)
                dist[nei] = neiDist
                visited[nei] = true
            }
        }
    }
    return dist
}
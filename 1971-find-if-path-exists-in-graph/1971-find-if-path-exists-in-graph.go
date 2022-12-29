func validPath(n int, edges [][]int, source int, destination int) bool {
    adjList := map[int][]int{}
    m := len(edges)
    
    for i := 0; i < m; i++ {
        start := edges[i][0]
        end := edges[i][1]
        adjList[start] = append(adjList[start], end)
        adjList[end] = append(adjList[end], start)
    }
    
    q := []int{source}
    visited := map[int]struct{}{source: {}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == destination {return true}
        childs := adjList[dq]
        for _, child := range childs {
            _, ok := visited[child]
            if !ok {
                q = append(q, child)
                visited[child] = struct{}{}
            }
        }
    }
    
    return false
}
func findCenter(edges [][]int) int {
    adjList := map[int][]int{}
    degrees := map[int]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
        degrees[u]++
        degrees[v]++
    }
    visited := map[int]struct{}{}
    q := []int{}
    for k,v := range degrees {
        if v == 1 {
            q = append(q, k)
            degrees[k] = 0
            visited[k] = struct{}{}
        }
    }
    if len(q) == 0 {panic("what")}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            for _, nei := range adjList[dq] {
                degrees[nei]--
                _, ok := visited[nei]
                if degrees[nei] == 0 && !ok {
                    q = append(q, nei)
                    visited[nei] = struct{}{}
                }
            }
            qSize--
        }
        if len(q) == 1 {
            break
        }
    }
    return q[0]
}
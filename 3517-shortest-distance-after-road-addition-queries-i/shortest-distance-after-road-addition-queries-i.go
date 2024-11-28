func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    adjList := map[int][]int{}    
    for i := 0; i < n; i++ {adjList[i] = append(adjList[i], i+1)}
    res := []int{}
    for i := 0; i < len(queries); i++ {
        u,v := queries[i][0], queries[i][1]
        adjList[u] = append(adjList[u], v)
        res = append(res, bfs(adjList, n)) 
    }
    return res
}

func bfs(adjList map[int][]int, n int) int {
    q := [][]int{{0,0}} // <node, dist-so-far>
    visited := make([]bool, n)
    visited[0] = true
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        node := dq[0]
        dist := dq[1]
        if node == n-1 {return dist}
        for _, nei := range adjList[node] {
            if visited[nei] {continue}
            q = append(q, []int{nei, dist+1})
            visited[nei] = true
            if nei == n-1 {return dist+1}
        }
    }
    return -1
}
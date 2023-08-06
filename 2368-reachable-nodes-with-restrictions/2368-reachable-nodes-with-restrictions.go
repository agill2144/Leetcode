func reachableNodes(n int, edges [][]int, restricted []int) int {
    restrictedSet := map[int]struct{}{}
    for i := 0; i < len(restricted); i++ {
        restrictedSet[restricted[i]] = struct{}{}
    }
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    visited := make([]bool, n+1)
    count := 0
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if visited[node] {return}
        
        // logic
        visited[node] = true
        count++
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            if _, ok := restrictedSet[nei]; ok {continue}
            dfs(nei, node)
        }
    }
    dfs(0, -1)
    return count
}
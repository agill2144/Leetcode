func validTree(n int, edges [][]int) bool {
    // graph is a valid tree if there is no cycle, and all nodes are connected
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)
    }

    count := 0
    visited := make([]bool, n)
    var dfs func(node, prev int, path []bool) bool
    dfs = func(node, prev int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return true}
        
        // logic
        visited[node] = true
        path[node] = true
        count++
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            if !dfs(nei, node, path) {return false}
        }
        path[node] = false
        return true
    }
    
    if !dfs(0, -1, make([]bool, n)) {return false}
    return count == n
    
}
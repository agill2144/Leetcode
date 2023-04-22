func validTree(n int, edges [][]int) bool {
    // a valid tree is a tree where there are no circles
    // and all components are connected
    
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)
    }
    
    visited := make([]bool, n)
    count := 0
    var dfs func(node, prev int, path []bool) bool
    dfs = func(node, prev int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return true}
        
        // logic
        visited[node] = true
        count++
        path[node] = true
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            if !dfs(nei, node, path) {return false}
        }
        path[node] = false
        return true
    }
    for i := 0; i < n; i++ {
        if !visited[i] {
            if count != 0 {return false}
            if !dfs(i, -1, make([]bool, n)) {return false}
        }
    }
    return count == n
}
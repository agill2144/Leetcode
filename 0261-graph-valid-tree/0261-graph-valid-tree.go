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
    var dfs func(node, prev int) bool
    dfs = func(node, prev int) bool {
        // base
        if visited[node] {return false}
        
        // logic
        visited[node] = true
        count++
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            if !dfs(nei, node) {return false}
        }
        return true
    }
    if !dfs(0, -1) {return false}
    return count == n
}
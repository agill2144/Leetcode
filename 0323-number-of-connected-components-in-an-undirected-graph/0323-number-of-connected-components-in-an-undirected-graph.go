func countComponents(n int, edges [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)
    }
    
    visited := make([]bool, n)
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if visited[node] {return}
        
        // logic
        visited[node] = true
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node)
        }
        
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            count++
            dfs(i, -1)
        }
    }
    return count
}
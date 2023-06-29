func countCompleteComponents(n int, edges [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)
    }
    nodesCount := 0
    edgesCount := 0
    visited := make([]bool, n)
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if visited[node] {return}
        
        // logic
        visited[node] = true
        nodesCount++
        edgesCount += len(adjList[node])
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node)
        }
    }
    completeCount := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i, -1)
            expected := nodesCount*(nodesCount-1)
            if edgesCount == expected {completeCount++}
            nodesCount = 0
            edgesCount = 0
        }
    }
    return completeCount
}
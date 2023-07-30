func findCircleNum(isConnected [][]int) int {
    adjList := map[int][]int{}
    n := len(isConnected)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 && i != j {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
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
        if !visited[i] {dfs(i, -1); count++}
    }
    return count
}

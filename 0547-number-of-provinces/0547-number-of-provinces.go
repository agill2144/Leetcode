func findCircleNum(isConnected [][]int) int {
    adjList := map[int][]int{}
    n := len(isConnected)
    visited := map[int]struct{}{}
    for i := 0; i < n ; i++ {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 && i != j {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)                
            }
        }
    }
    
    var dfs func(node, prev int) 
    dfs = func(node, prev int) {
        // base
        if _, ok := visited[node]; ok {return}
        visited[node] = struct{}{}
        
        // logic
        for _, neighbor := range adjList[node] {
            if neighbor == prev {continue}
            dfs(neighbor, node)
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            count++
            dfs(i, -1)
        }
    }
    return count
}
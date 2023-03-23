func makeConnected(n int, connections [][]int) int {
    if len(connections) < n-1 {return -1}
    adjList := map[int][]int{}
    for i := 0; i < len(connections); i++ {
        adjList[connections[i][0]] = append(adjList[connections[i][0]], connections[i][1])
        adjList[connections[i][1]] = append(adjList[connections[i][1]], connections[i][0])        
    }
    
    visited := map[int]struct{}{}
    var dfs func(curr, prev int)
    dfs = func(curr, prev int) {
        // base
        if _, ok := visited[curr]; ok {
            return
        }
        
        // logic
        visited[curr] = struct{}{}
        for _, node := range adjList[curr] {
            if node == prev {continue}
            dfs(node, curr)
        }
    }
    numConnected := 0
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            numConnected++
            dfs(i, -1)
        }
    }
    return numConnected-1
}
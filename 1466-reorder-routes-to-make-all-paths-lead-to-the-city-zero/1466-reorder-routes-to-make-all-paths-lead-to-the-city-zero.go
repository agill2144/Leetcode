func minReorder(n int, connections [][]int) int {
    set := map[[2]int]struct{}{}
    adjList := map[int][]int{}
    for i := 0; i < len(connections); i++ {
        adjList[connections[i][0]] = append(adjList[connections[i][0]], connections[i][1])
        adjList[connections[i][1]] = append(adjList[connections[i][1]], connections[i][0])
        set[[2]int{connections[i][0], connections[i][1]}] = struct{}{}
    }
    
    count := 0
    visited := map[int]struct{}{}
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        
        
        // logic
        visited[node] = struct{}{}
        if _, ok := set[[2]int{node, prev}]; !ok && prev != -1 {
            count++
        }
        for _, edge := range adjList[node] {
            if edge == prev {continue}
            dfs(edge, node)
        }
    }
    dfs(0,-1)
    return count
}
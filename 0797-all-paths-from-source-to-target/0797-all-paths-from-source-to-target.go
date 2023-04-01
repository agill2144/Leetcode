func allPathsSourceTarget(graph [][]int) [][]int {
    dest := len(graph)-1
    out := [][]int{}
    var dfs func(node int, path []int)
    dfs = func(node int, path []int) {
        // base
        if node == dest {
            newL := make([]int, len(path))
            copy(newL, path)
            newL = append(newL, dest)
            out = append(out, newL)
            return
        }
        
        // logic
        path = append(path, node)
        for _, nei := range graph[node] {
            dfs(nei, path)
        }
    }
    dfs(0, nil)
    return out
    
}
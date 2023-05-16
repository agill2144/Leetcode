func allPathsSourceTarget(graph [][]int) [][]int {
    out := [][]int{}
    dest := len(graph)-1
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
        path = path[:len(path)-1]
    }
    dfs(0, []int{})
    return out
}
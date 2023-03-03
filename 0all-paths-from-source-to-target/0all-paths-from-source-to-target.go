func allPathsSourceTarget(graph [][]int) [][]int {
    out := [][]int{}
    var dfs func(node int, path []int)
    dfs = func(node int, path []int) {
        // base
        if node == len(graph)-1 {
            newL := make([]int, len(path))
            copy(newL, path)
            newL = append(newL, len(graph)-1)
            out = append(out, newL)
            return
        }
        
        // logic
        path = append(path, node)
        for _, neighbor := range graph[node]{
            dfs(neighbor, path)
        }
        path = path[:len(path)-1]
    }
    dfs(0, nil)
    return out
}
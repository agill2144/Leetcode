func eventualSafeNodes(graph [][]int) []int {
    visited := make([]bool, len(graph))
    out := map[int]*bool{}
    toPtrBool := func(b bool)*bool {return &b }
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {
            val := out[node]
            if val != nil {return *val}
            return true
        }
        
        // logic
        visited[node] =true
        path[node] = true 
        for _, nei := range graph[node] {
            if !dfs(nei, path) {
                out[node] = toPtrBool(false)
                return false
            }
        }
        path[node] = false
        out[node] = toPtrBool(true)
        return true
    }
    for i := 0; i < len(graph); i++ {
        if !visited[i] {
            dfs(i, make([]bool, len(graph)))
        }
    }
    result := []int{}
    for i := 0; i < len(graph); i++ {
        if out[i] != nil && *out[i] {result = append(result, i)}
    }
    return result
}
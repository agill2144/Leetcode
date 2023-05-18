func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    visited := make([]bool, n)
    valid := make([]*bool, n)
    toPtrBool := func(b bool) *bool {return &b}    
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {
            if valid[node] != nil {return *valid[node]}
            return true
        }
        
        // logic
        visited[node] = true
        path[node] = true
        for _, nei := range graph[node] {
            if !dfs(nei, path){
                valid[node] = toPtrBool(false)
                path[node] = false
                return false
            }
        }
        path[node] = false
        valid[node] = toPtrBool(true)
        return true
    }
    
    for i := 0 ; i < n; i++ {
        if !visited[i] {
            dfs(i, make([]bool, n))
        }
    }
    out := []int{}
    for i := 0; i < n; i++ {
        if valid[i] != nil && *valid[i] {
            out = append(out, i)
        }
    }
    return out
}
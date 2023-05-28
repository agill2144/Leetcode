func eventualSafeNodes(graph [][]int) []int {
    safeNodes := make([]*bool, len(graph))
    toPtrBool := func(b bool) *bool {return &b} 
    for i := 0; i < len(graph); i++ {
        if len(graph[i]) == 0 {
            safeNodes[i] = toPtrBool(true)
        }
    }
    
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if safeNodes[node] != nil {
            return *safeNodes[node]
        }
        
        // logic
        path[node] = true
        for _, nei := range graph[node] {
            if !dfs(nei, path) {
                path[node] = false
                safeNodes[node] = toPtrBool(false)
                return false
            }
        }
        path[node] = false
        safeNodes[node] = toPtrBool(true)
        return true
    }

    path := make([]bool, len(graph))
    for i := 0; i < len(graph); i++ {
        if safeNodes[i] == nil {
            dfs(i, path)
        }
    }
    
    out := []int{}
    for i := 0; i < len(safeNodes); i++ {
        if safeNodes[i] != nil && *safeNodes[i] {
            out = append(out, i)
        }
    }
    return out
}

// o(v) + o(v+e) + o(v) + o(vlogv)
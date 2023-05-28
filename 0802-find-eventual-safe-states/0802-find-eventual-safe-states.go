func eventualSafeNodes(graph [][]int) []int {
    safeNodes := map[int]bool{}
    for i := 0; i < len(graph); i++ {
        if len(graph[i]) == 0 {
            safeNodes[i] = true
        }
    }
    
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if val, ok := safeNodes[node]; ok {
            return val
        }
        
        // logic
        path[node] = true
        for _, nei := range graph[node] {
            if !dfs(nei, path) {
                path[node] = false
                safeNodes[node] = false
                return false
            }
        }
        path[node] = false
        safeNodes[node] = true
        return true
    }

    path := make([]bool, len(graph))
    for i := 0; i < len(graph); i++ {
        if _, ok := safeNodes[i]; !ok {
            dfs(i, path)
        }
    }
    
    out := []int{}
    for k, v := range safeNodes {
        if v {
            out = append(out, k)
        }
    }
    sort.Ints(out)
    return out
}

// o(v) + o(v+e) + o(v) + o(vlogv)
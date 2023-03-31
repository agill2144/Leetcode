func eventualSafeNodes(graph [][]int) []int {
    toPtrBool := func(b bool) *bool {return &b}

    safeNodes := make([]*bool, len(graph))
    visited := make([]bool, len(graph))
    var dfs func(node int, path []bool) bool 
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return true}
        if visited[node] {
            // a visited node could be connected to a cycle, if thats the case, its also not safe...
            if safeNodes[node] != nil && *safeNodes[node] == false {return true}
            return false
        }
        
        // logic
        path[node] = true
        visited[node] = true
        for _, nei := range graph[node] {
            ok := dfs(nei, path)
            if ok {safeNodes[node] = toPtrBool(false); return true}
        }
        path[node] = false
        safeNodes[node] = toPtrBool(true)
        return false
    }
    for i := 0; i < len(graph); i++ {
        if !visited[i] {
            path := make([]bool, len(graph))
            dfs(i, path)
        }
    }
        
    out := []int{}
    for i := 0; i < len(graph); i++ {
        if *safeNodes[i] {out = append(out, i)}
    }
    return out
}
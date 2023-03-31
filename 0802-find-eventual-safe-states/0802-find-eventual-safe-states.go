func eventualSafeNodes(graph [][]int) []int {
    toPtrBool := func(b bool) *bool {return &b}

    safeNodes := make([]*bool, len(graph))
    var dfs func(node int, path []bool) bool 

    // hasCycleOrConnectedToCycle
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return true}
        if safeNodes[node] != nil {
            // a visited node could be connected to a cyclic node, if thats the case, its also not safe...
            return !*safeNodes[node]
        }
        
        // logic
        path[node] = true
        safeNodes[node] = toPtrBool(false)
        for _, nei := range graph[node] {
            if dfs(nei, path) {safeNodes[node] = toPtrBool(false); return true}
        }
        path[node] = false
        safeNodes[node] = toPtrBool(true)
        return false
    }
    for i := 0; i < len(graph); i++ {
        if safeNodes[i] == nil {
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
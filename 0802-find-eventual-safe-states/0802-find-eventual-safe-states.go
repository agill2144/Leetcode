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
        // action
        visited[node] = true
        path[node] = true
        for _, nei := range graph[node] {
            // if any of this node's neighbor, is invalid (part of a cycle)
            // then this makes this node invalid, and we need to return false
            // so that other nodes connected to this node can also see that they
            // are connected to an invalid node and therefore they are also invalid
            // basically propagate the invalid path all the way up in the current recursion stack
            // recursion
            if !dfs(nei, path) {
                valid[node] = toPtrBool(false)
                return false
            }
        }
        
        // if this node never exited with a false from exploring all of its neighbors,
        // then this node is a valid node
        valid[node] = toPtrBool(true)
        // backtrack
        path[node] = false
        
        // since this node is valid, we can return true
        return true
    }
    for i := 0; i < n; i++ {
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
func eventualSafeNodes(graph [][]int) []int {
    visited := make([]bool, len(graph))
    valid := make([]*bool, len(graph))
    
    toPtrBool := func(b bool) *bool {return &b}
    var isValid func(node int, path []bool) bool
    isValid = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {
            if valid[node] != nil {return *valid[node]}
            return true
        }
        
        // logic
        path[node] = true
        visited[node] = true
        for _, nei := range graph[node] {
            if !isValid(nei, path) {
                valid[node] = toPtrBool(false)
                return false
            }
        }
        
        path[node] = false
        valid[node] = toPtrBool(true)
        return true
    }
    
    for i := 0; i < len(graph); i++ {
        if !visited[i] {
            path := make([]bool, len(graph))
            isValid(i, path)
        }
    }
    out := []int{}
    for i := 0; i < len(graph); i++ {
        if *valid[i] {out = append(out, i)}
    }
    return out
}
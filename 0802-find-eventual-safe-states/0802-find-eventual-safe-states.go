func eventualSafeNodes(graph [][]int) []int {
    safeNodes := make([]bool, len(graph))
    for i := 0; i < len(graph); i++ {
        if len(graph[i]) == 0 {
            safeNodes[i] = true
        }
    }
    visited := make([]bool, len(graph))
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return safeNodes[node]}
        
        // logic
        path[node] = true
        visited[node] = true
        for _, nei := range graph[node] {
            if !dfs(nei, path) {
                safeNodes[nei] = false
                path[node] = false
                return false
            }
        }
        path[node] = false
        safeNodes[node]  = true
        return true
    }
    p := make([]bool, len(graph))
    for i := 0; i < len(graph); i++ {
        if !visited[i] {
            dfs(i,p)
        }
    }
    
    out := []int{}
    for i := 0; i < len(safeNodes); i++ {
        if safeNodes[i] {
            out = append(out, i)
        }
    }
    return out
}
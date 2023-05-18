// All paths of a node must lead to an end without running into a cycle
// only then a node is considered valid
// if a node path went into a cycle, then this node is not valid
// as well as all other parent nodes that are part of recursion, also have a path that went to a node that ran into a cycle
// therefore its parent and its parents are also invalid because they have a path somewhere that went to a cycle.
// time: o(v+e) + o(v)
// space: o(v) + o(v) + o(v)
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
    // o(v+e)
    p := make([]bool, n)
    for i := 0 ; i < n; i++ {
        if !visited[i] {
            dfs(i, p)
        }
    }
    out := []int{}
    // o(v)
    for i := 0; i < n; i++ {
        if valid[i] != nil && *valid[i] {
            out = append(out, i)
        }
    }
    return out
}
func eventualSafeNodes(graph [][]int) []int {
    adjList := map[int][]int{}
    indegrees := make([]int, len(graph))
    for i := 0; i < len(graph); i++ {
        child := i
        indegrees[child] += len(graph[i])
        for _, parent := range graph[i] {
            adjList[parent] = append(adjList[parent], child)
            
        }
    }
    
    q := []int{}
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }
    if len(q) == 0 {return nil}

    valid := make([]bool, len(graph))
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        valid[dq] = true
        for _, nei := range adjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 {
                q = append(q, nei)
            }
        }
    }
    out := []int{}
    for i := 0; i < len(valid); i++ {
        if valid[i] {
            out = append(out, i)
        }
    }
    return out
    
}
// func eventualSafeNodes(graph [][]int) []int {
//     safeNodes := make([]*bool, len(graph))
//     toPtrBool := func(b bool) *bool {return &b} 
//     for i := 0; i < len(graph); i++ {
//         if len(graph[i]) == 0 {
//             safeNodes[i] = toPtrBool(true)
//         }
//     }
    
//     var dfs func(node int, path []bool) bool
//     dfs = func(node int, path []bool) bool {
//         // base
//         if path[node] {return false}
//         if safeNodes[node] != nil {
//             return *safeNodes[node]
//         }
        
//         // logic
//         path[node] = true
//         for _, nei := range graph[node] {
//             if !dfs(nei, path) {
//                 path[node] = false
//                 safeNodes[node] = toPtrBool(false)
//                 return false
//             }
//         }
//         path[node] = false
//         safeNodes[node] = toPtrBool(true)
//         return true
//     }

//     path := make([]bool, len(graph))
//     for i := 0; i < len(graph); i++ {
//         if safeNodes[i] == nil {
//             dfs(i, path)
//         }
//     }
    
//     out := []int{}
//     for i := 0; i < len(safeNodes); i++ {
//         if safeNodes[i] != nil && *safeNodes[i] {
//             out = append(out, i)
//         }
//     }
//     return out
// }

// o(v) + o(v+e) + o(v) + o(vlogv) --- if sorting
// o(v) + o(v+e) + o(v) -- if not sorting
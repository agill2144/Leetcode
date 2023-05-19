func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    outdegrees := make([]int, n)
    revAdjList := map[int][]int{}
    
    for i := 0; i < n ; i++ {
        outdegrees[i] += len(graph[i])
        for _, ele := range graph[i] {
            revAdjList[ele] = append(revAdjList[ele], i)
        }
    }
    valid := make([]bool, n)
    q := []int{}
    visited := make([]bool, n)
    for i := 0; i < n; i++ {
        if outdegrees[i] == 0 {
            visited[i]=true
            valid[i]= true
            q = append(q, i)
        }
    }
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        for _, nei := range revAdjList[dq] {
            outdegrees[nei]--
            if outdegrees[nei] == 0 {
                visited[nei] = true
                valid[nei] = true
                q = append(q, nei)
            }
        }
    }
    
    out := []int{}
    for i := 0; i < len(valid); i++ {
        if valid[i] {out = append(out, i)}
    }
    return out
    
}
// All paths of a node must lead to an end without running into a cycle
// only then a node is considered valid
// if a node path went into a cycle, then this node is not valid
// as well as all other parent nodes that are part of recursion, also have a path that went to a node that ran into a cycle
// therefore its parent and its parents are also invalid because they have a path somewhere that went to a cycle.
// time: o(v+e) + o(v)
// space: o(v) + o(v) + o(v) + o(v+e)
// func eventualSafeNodes(graph [][]int) []int {
//     n := len(graph)
//     visited := make([]bool, n)
//     // why ptr to a bool ?
//     // a node could be visited ( ie in the current recursion stack )
//     // but we have not yet identified whether a visited node is valid or not
//     // using a regular bool arr, will be filled with false
//     // and if a node is visited, we cannot assume the value in valid arry is correct
//     // therefore we need a third state to determine whether a node is visited and if it is
//     // do we have an answer for it or not ? 
//     // ptr to bool helps us in that
//     // if we have answered a node, then its value must not be nil, use that value as a memoized value
//     // if this node is in our recursion stack but we have not yet identified whether its a safe node or not, then its value will be nil instead of false
//     valid := make([]*bool, n)
//     toPtrBool := func(b bool) *bool {return &b}
    
//     var dfs func(node int, path []bool) bool
//     dfs = func(node int, path []bool) bool {
//         // base
//         if path[node] {return false}
//         if visited[node] {
//             if valid[node] != nil {return *valid[node]}
//             return true
//         }
        
//         // logic
//         visited[node] = true
//         path[node] = true
//         for _, nei := range graph[node] {
//             if !dfs(nei, path){
//                 valid[node] = toPtrBool(false)
//                 path[node] = false
//                 return false
//             }
//         }
//         path[node] = false
//         valid[node] = toPtrBool(true)
//         return true
//     }
//     // o(v+e)
//     p := make([]bool, n)
//     for i := 0 ; i < n; i++ {
//         if !visited[i] {
//             dfs(i, p)
//         }
//     }
//     out := []int{}
//     // o(v)
//     for i := 0; i < n; i++ {
//         if valid[i] != nil && *valid[i] {
//             out = append(out, i)
//         }
//     }
//     return out
// }
func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    indegrees := make([]int, n)
    revAdjList := map[int][]int{}
    q := []int{}
    out := []int{}
    for i := 0; i < n; i++ {
        indegrees[i] = len(graph[i])
        if indegrees[i] == 0{q = append(q, i)}
        for _, edge := range graph[i] {
            revAdjList[edge] = append(revAdjList[edge], i)            
        }
    }
    if len(q) == 0 {return nil}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        out = append(out, dq)
        for _, nei := range revAdjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 {
                q = append(q, nei)
            }
        }
    }
    sort.Ints(out)
    return out
}

/*
    a node is eventually safe if all it's outgoing edges are to nodes that are eventually safe.
    
    approach: DFS cycle detection
    - A node is a safe node if all their paths lead to a terminal or another safe node
    - i.e all of a node's path did not end up in a cycle.
    - Therefore we can apply cycle detection using DFS on a directed graph
    - Then if the neigbhor exploration returns false, we do not mark this node safe
    - If all the neighbor explorations exited successfully, than we can mark this node safe
    
    - There can be disconnected components and this is a directed graph, therefore need to run dfs on all nodes
    - if a node is already visited, we cannot exit early and say that this node is valid ( i.e will not run into a cycle at all )
        - therefore we need to continue a path exploration even if a node is already visited but not in our current recursion stack
    - This is where we can "memoize" an answer for each node
    - if a node exited with false, then we can save this answer that this node is not safe
    - if a node exited with true, then we can save this answer that this nide is safe
    - now if a node is already visited, we can check with memoized answers - if none return true
    - thats where we need to maintain a valid array ( ptr to bool array )
        - so that we have 3 identifiers
            - nil = not explored
            - false = explored and invalid node
            - true = explored and valid node
        - if we dont use ptr to bool, the whole array will have false values
        - and then we wont be able to identify invalid vs not-yet-explored nodes.
        
    time: o(v+e)
    space: o(v+e)
*/
// func eventualSafeNodes(graph [][]int) []int {
//     n := len(graph)
//     visited := make([]bool, n)
//     valid := make([]*bool, n)
//     toPtrBool := func(b bool) *bool {return &b}
    
//     var dfs func(node int, path []bool ) bool
//     dfs = func(node int, path []bool ) bool {
//         // base
//         if path[node] {return false}
//         if visited[node] {
//             if valid[node] != nil {return *valid[node] }
//             return true
//         }
        
//         // logic
//         visited[node] = true
//         path[node] = true
//         for _, nei := range graph[node] {
//             if !dfs(nei, path) {
//                 valid[node] = toPtrBool(false)
//                 path[node] = false
//                 return false
//             }
//         }
//         valid[node] = toPtrBool(true)
//         path[node] = false
//         return true
//     }
//     for i := 0; i < n; i++ {
//         if !visited[i] {
//             dfs(i, make([]bool, n))
//         }
//     }
//     out := []int{}
//     for i := 0; i < n; i++ {
//         if *valid[i] {out = append(out, i)}
//     }
//     return out
// }
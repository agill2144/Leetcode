func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    indegrees := make([]int, n)
    revAdjList := map[int][]int{}
    q := []int{}
    out := []int{}
    for i := 0; i < n; i++ {
        for _, nei := range graph[i] {
            revAdjList[nei] = append(revAdjList[nei], i)
            indegrees[i]++
        }
        if indegrees[i] == 0 {
            q = append(q, i)
            out = append(out, i)
        }
    }
    if len(q) == 0 {return nil}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        for _, nei := range revAdjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 {out = append(out, nei); q = append(q, nei)}
        }
    }
    sort.Ints(out)
    return out
}
// /*
//     approach: DFS cycle detection
//     - as soon as we find a cycle, all the nodes connected to this cycle are invalid
//     - since using DFS means, connected nodes (i.e recursion path ) is in our recursive stack
//     - once a cycle is detected , we can return false
//     - once a false is detected from its neighbor, than that means this node is also invalid;
//         - mark it invalid, and propagate the false result up by returning false and exiting early
//     - if none of the neighbors of a node returned false, that means all of this nodes path exited without getting into a cycle.
//     - when this happens, this means this node is valid.
    
//     - Detecting cycle on a directed graph 
//         - use dfs
//         - use path to keep track of "who is in my path"
//         - a path becomes invalid, if we run into a node that's already in the path ( i.e cycle detected )
    
//     - Since this is a directed graph, we have to run dfs on all nodes
//         - we can use a visited set / array to not re-run dfs on already processed nodes
//     - But if dfs runs, and its neighbors is already visited, check if they are valid 
//     - thats where we need to maintain a valid array ( ptr to bool array )
//         - so that we have 3 identifiers
//             - nil = not explored
//             - false = invalid node
//             - true = valid node
//         - if we dont use ptr to bool, the whole array will have false values
//         - and then we wont be able to identify invalid vs not-yet-explored nodes.
        
//     time: o(v+e) + o(v)
//     space: o(v) + o(v) + o(v+e)
// */
// func eventualSafeNodes(graph [][]int) []int {
//     n := len(graph)
//     visited := make([]bool, n)
//     valid := make([]*bool, n)
//     toPtrBool := func(b bool) *bool {return &b}

//     // returns false if a cycle is detected
//     var dfs func(node int, path []bool) bool
//     dfs = func(node int, path []bool) bool {
//         // base
//         if path[node] {return false}
//         if visited[node] {
//             if valid[node] != nil {return *valid[node]}
//             return true
//         }
        
//         // logic
//         // action
//         visited[node] = true
//         path[node] = true
//         for _, nei := range graph[node] {
//             // if any of this node's neighbor, is invalid (part of a cycle)
//             // then this makes this node invalid, and we need to return false
//             // so that other nodes connected to this node can also see that they
//             // are connected to an invalid node and therefore they are also invalid
//             // basically propagate the invalid path all the way up in the current recursion stack
//             // recursion
//             if !dfs(nei, path) {
//                 valid[node] = toPtrBool(false)
//                 return false
//             }
//         }
        
//         // if this node never exited with a false from exploring all of its neighbors,
//         // then this node is a valid node
//         valid[node] = toPtrBool(true)
//         // backtrack
//         path[node] = false
        
//         // since this node is valid, we can return true
//         return true
//     }
    
//     for i := 0; i < n; i++ {
//         if !visited[i] {
//             dfs(i, make([]bool, n))
//         }
//     }
//     out := []int{}
//     for i := 0; i < n; i++ {
//         if valid[i] != nil && *valid[i] {
//             out = append(out, i)
//         }
//     }
//     return out
// }
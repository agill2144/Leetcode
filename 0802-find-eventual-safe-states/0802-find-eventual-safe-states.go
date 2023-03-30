func eventualSafeNodes(graph [][]int) []int {
    outdegrees := make([]int, len(graph))
    for i := 0; i < len(graph); i++ {
        outdegrees[i] += len(graph[i])
    }
    terminalNodes := map[int]struct{}{}
    for i := 0; i < len(outdegrees); i++ {
        if outdegrees[i] == 0 {terminalNodes[i] = struct{}{}}
    }
    var dfs func(node int, path map[int]struct{}) bool
    dfs = func(node int, path map[int]struct{}) bool {
        // base
        if _, ok := terminalNodes[node]; ok {return true}
        
        
        // logic
        path[node] = struct{}{}
        for _, nei := range graph[node] {
            if _, ok := path[nei]; ok {return false}            
            if isSafe := dfs(nei, path); !isSafe {return false}
        }
        delete(path, node)
        return true
    }
    out := []int{}
    for i := len(graph)-1; i >= 0; i-- {
        isSafe := dfs(i, map[int]struct{}{})
        if isSafe {terminalNodes[i] = struct{}{}; out = append(out, i)}
        // fmt.Println("end of: ", i, safeNodes)
    }
    sort.Ints(out)
    return out
}
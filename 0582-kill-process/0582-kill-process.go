func killProcess(pid []int, ppid []int, kill int) []int {
    adjList := map[int][]int{}
    for i := 0; i < len(pid); i++ {
        process := pid[i]
        parent := ppid[i]
        if parent == 0 {continue}
        adjList[parent] = append(adjList[parent], process)
    }
    out := []int{}
    visited := map[int]struct{}{}
    var dfs func(node int)
    dfs = func(node int) {
        // base
        if _, ok := visited[node]; ok  {return}
        
        // logic
        visited[node] = struct{}{}
        out = append(out, node) 
        for _, nei := range adjList[node] {
            dfs(nei)
        }
    }
    dfs(kill)
    return out
}
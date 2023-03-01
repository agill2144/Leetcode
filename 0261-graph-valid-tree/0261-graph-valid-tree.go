func validTree(n int, edges [][]int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1]) 
        adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0]) 
    }
    visited := map[int]struct{}{}
    var dfs func(node int, parent int) bool
    dfs = func(node, parent int) bool {
        // base
        if _, ok := visited[node]; ok {return false}
        
        // logic
        visited[node] = struct{}{}
        for _, child := range adjList[node] {
            if child == parent {continue}
            if ok := dfs(child, node); !ok {return false} 
        }
        return true
    }
    out := dfs(0, -1)
    if !out {return false}
    for i := 0; i < n; i++ {if _, ok :=  visited[i]; !ok {return false}}
    return true
}
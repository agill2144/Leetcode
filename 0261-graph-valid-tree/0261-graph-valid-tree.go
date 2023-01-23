func validTree(n int, edges [][]int) bool {
    if n  == 0 {return true}
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        n1 := edges[i][0]
        n2 := edges[i][1]
        adjList[n1] = append(adjList[n1], n2)
        adjList[n2] = append(adjList[n2], n1)
    }
    visited := map[int]struct{}{}
    var dfs func(node int, prev int) bool 
    dfs = func(node, prev int) bool {
        // base
        if _, ok := visited[node]; ok {return false}
        
        // logic
        visited[node] = struct{}{}
        for _, child := range adjList[node] {
            if child == prev {continue}
            if ok := dfs(child, node); !ok {return false}
        }
        return true
    }
    ok := dfs(0, -1)
    if !ok {return false}
    return len(visited) == n
}
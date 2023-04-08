func longestCycle(edges []int) int {
    out := -1    
    visited := map[int]struct{}{}
    
    var dfs func(node int, path map[int]int)
    dfs = func(node int, path map[int]int) {
        // base
        if node == -1 {return}
        if sizeAtThatNode, ok := path[node]; ok {
            out = max(out, len(path)-sizeAtThatNode)
            return
        }
        if _, ok := visited[node]; ok {return}
        
        // logic
        visited[node] = struct{}{}
        currSize := len(path)
        path[node] = currSize
        dfs(edges[node], path)
    }
    for i := 0; i < len(edges); i++ {
        if _, ok := visited[i]; !ok {
            dfs(i, map[int]int{})
        }
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
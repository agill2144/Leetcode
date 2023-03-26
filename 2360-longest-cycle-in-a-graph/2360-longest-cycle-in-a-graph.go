func longestCycle(edges []int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        if edges[i] == -1 { adjList[i] = []int{}; continue}
        adjList[i] = append(adjList[i], edges[i])
    }
    out := -1    
    visited := map[int]struct{}{}
    
    var dfs func(node int, path map[int]int)
    dfs = func(node int, path map[int]int) {
        // base
        if sizeAtThatNode, ok := path[node]; ok {
            out = max(out, len(path)-sizeAtThatNode)
            return
        }
        if _, ok := visited[node]; ok {return}
        
        // logic
        visited[node] = struct{}{}
        currSize := len(path)
        path[node] = currSize
        for _, nei := range adjList[node] {
            dfs(nei, path)
        }
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
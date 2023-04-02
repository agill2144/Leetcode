func possibleBipartition(n int, dislikes [][]int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(dislikes); i++ {
        a := dislikes[i][0]
        b := dislikes[i][1]
        adjList[a] = append(adjList[a], b)
        adjList[b] = append(adjList[b], a)        
    }
    visited := make([]bool, n+1)
    colors := make([]int, n+1)
    
    var dfs func(node, prev, color int) bool
    dfs = func(node, prev, color int) bool {
        // base
        if colors[node] != 0 && colors[node] != color {return false}
        if visited[node] {return true}
        
        // logic
        visited[node] = true
        colors[node] = color
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            neiColor := 1
            if color == 1 {neiColor = 2}
            if !dfs(nei, node, neiColor) {return false}
        }
        return true
    }
    for i := 1; i <= n; i++ {
        if !visited[i] {
            if !dfs(i, -1, 1) {return false}
        }
    }
    return true
}
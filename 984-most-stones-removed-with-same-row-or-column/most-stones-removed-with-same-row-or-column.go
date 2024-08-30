func removeStones(stones [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(stones); i++ {
        parentR, parentC := stones[i][0], stones[i][1]
        for j := 0; j < len(stones); j++ {
            if i == j {continue}
            childR, childC := stones[j][0], stones[j][1]
            if childR == parentR  || childC == parentC {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
    }
    visited := map[int]bool{}
    var dfs func(node int)
    dfs = func(node int) {
        // base
        if visited[node] {return}
        // logic
        visited[node] = true
        for _, nei := range adjList[node] {
            dfs(nei)
        }
    }
    count := 0
    for i := 0; i < len(stones); i++ {
        if !visited[i] {
            dfs(i)
            count++
        }
    }
    return len(stones)-count
}
func maximalNetworkRank(n int, roads [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(roads); i++ {
        u,v := roads[i][0], roads[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    
    
    maxCount := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if j == i {continue}
            u, v := i, j
            uEdges := adjList[u]
            vEdges := adjList[v]
            total := len(uEdges) + len(vEdges)
            if sliceContains(uEdges, v) || sliceContains(vEdges, u) {
                total--
            }
            if total > maxCount {maxCount = total}        
        }
        
        // fmt.Println("Pair: ", u,v, total)
        // fmt.Println("U count: ", u, uEdges)
        // fmt.Println("V count: ", v, vEdges)
        // fmt.Println("--------------------------")        
    }
    return maxCount
}

func sliceContains(slice []int, n int) bool {
    for _, ele := range slice {
        if ele == n {return true}
    }
    return false
}
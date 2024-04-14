func filterEdges(edges [][]int) [][]int {
    // Stores one mininum weight for all V
    mapV := make(map[int]int)
    for _, egde := range edges {
        v, weight := egde[0], egde[1]
        if _, ok := mapV[v]; !ok {
            mapV[v] = weight
        } else {
            mapV[v] = min(mapV[v], weight)
        }
    }
    
    ans := [][]int{}
    for v, weight := range mapV {
        ans = append(ans, []int{v, weight})
    }
    // Sorting edges by his lenght
    sort.Slice(ans, func(i, j int) bool {
        return ans[i][1] < ans[j][1]
    })
    return ans
}

func filterGraph(g map[int][][]int) map[int][][]int {
    ansG := make(map[int][][]int)
    for v, edges := range g {
        // For all key (v) in g (graph) get the processed, 
        // correct edges
        correctEdges := filterEdges(edges)
        // and update it
        ansG[v] = correctEdges
    }
    
    return ansG
}

func updateAns(currV int, g *map[int][][]int, dis *[]int, ans *[]int) {
    // For each edge
    for _, edge := range (*g)[currV] {
        nextV, weight := edge[0], edge[1]
        nextWeight := (*ans)[currV] + weight
        // Checking, is it correct to update ans and not updated before
        if nextWeight < (*ans)[nextV] && nextWeight < (*dis)[nextV] {
            // Is it, fill updated ans
            (*ans)[nextV] = nextWeight
            // And start next DFS func calling
            updateAns(nextV, g, dis, ans)
        }
    }
}

func minimumTime(n int, edges [][]int, disappear []int) []int {
    // Fill ans
    ans := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        ans[i] = math.MaxInt32
    }
    ans[0] = 0
    
    // Create graph
    graph := make(map[int][][]int)
    for _, egde := range edges {
        v, u := egde[0], egde[1]
        weight := egde[2]
        graph[v] = append(graph[v], []int{u, weight})
        graph[u] = append(graph[u], []int{v, weight})
    }
    
    graph = filterGraph(graph)
    
    // Start DFS
    updateAns(0, &graph, &disappear, &ans)
    
    // If ans[i] has a MAX val, i has't update, then, -1
    for i, val := range ans {
        if val == math.MaxInt32 {
            ans[i] = -1
        }
    }
    
    // return answer slice
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
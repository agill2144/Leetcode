func countPairs(n int, edges [][]int) int64 {
    // the question is easy, count number of nodes in a connected components
    // and the missing edges for each pair in all disconnected components
    // the math piece at the end stumbled me :/ 
    
    
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)
    }
    
    
    visited := map[int]struct{}{}
    var dfs func(node, prev int) int64
    dfs = func(node, prev int) int64 {
        // base
        if _, ok := visited[node]; ok {return 0}
        
        // logic
        var count int64 = 1
        visited[node] = struct{}{}
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            count += dfs(nei, node)
        }
        return count
    }

    var total int64
    nInt64 := int64(n)
    var totalComponentsSoFar int64
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            size := dfs(i, -1)
            total += (nInt64-size-totalComponentsSoFar)*size
            totalComponentsSoFar += size
        }
    }
    return total
    
}
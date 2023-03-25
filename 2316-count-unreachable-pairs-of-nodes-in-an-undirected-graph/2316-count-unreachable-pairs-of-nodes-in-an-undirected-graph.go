func countPairs(n int, edges [][]int) int64 {
    // the question is easy, count number of nodes in a connected component
    // and the missing edges for each pair in all disconnected components

    // for example if we have the 3 connected components of size [ 2,3,1 ]
    // then the number of missing edges from each node to other node in all other components are: (2x3)+(2x1)+(3x1)
    // we can brute force the math like the above ^    
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
            // the math logic is
            // we have n total nodes
            // we have located size nodes
            // remaining are n-size, that means n-size * size are total missing edges
            // however once we have considered a component already, we need to remove its size from the above math
            // therefore maintaining totalComponentsSoFar so we can remove it
            total += (nInt64-size-totalComponentsSoFar)*size
            totalComponentsSoFar += size
        }
    }
    return total
    
}
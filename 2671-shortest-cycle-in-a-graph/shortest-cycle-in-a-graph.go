func findShortestCycle(n int, edges [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)

    }
    ans := math.MaxInt64
    visited := make([]bool, n)
    dists := make([]int, n)
    var dfs func(node, prev, dist int)
    dfs = func(node, prev, dist int) {
        // base
        if visited[node] {
            ans = min(ans, dist-dists[node])
            return
        }
        // no point in going down a path whose dist is already > what we have saved
        if dist >= ans {return}

        // logic
        visited[node] = true
        dists[node] = dist
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node, dist+1)
        }
        visited[node] = false
        dists[node] = 0
    }
    for i := 0; i < n; i++ {
        dfs(i, -1, 0)        
    }
    if ans == math.MaxInt64 {return -1}
    return ans
}
func longestCycle(edges []int) int {
    visited := make([]bool, len(edges))
    largestCycle := -1
    var dfs func(node int, distToThisNode int, dist []int, path []bool)
    dfs = func(node, distToThisNode int, dist []int, path []bool) {
        // base
        if node == -1 {return}
        if path[node] {
            largestCycle = max(largestCycle, distToThisNode-dist[node])
            return
        }
        if visited[node] {return}

        // logic
        visited[node] = true
        path[node] = true
        dist[node] = distToThisNode
        dfs(edges[node], distToThisNode+1, dist, path)
        path[node] = false
        dist[node] = 0
    }
    d := make([]int, len(edges))
    p := make([]bool, len(edges))
    for i := 0; i < len(edges); i++ {
        if !visited[i] {
            dfs(i, 0, d, p)
        }
    }

    return largestCycle

}
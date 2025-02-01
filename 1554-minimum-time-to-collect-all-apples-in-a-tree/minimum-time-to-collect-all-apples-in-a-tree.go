func minTime(n int, edges [][]int, hasApple []bool) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    total := 0
    var dfs func(node, prev int) bool
    dfs = func(node, prev int)bool{
        // base


        // logic
        childHasApples := false
        for _, child := range adjList[node] {
            if child == prev {continue}
            if dfs(child, node) {
                childHasApples = true
                total += 2
            }
        }
        return childHasApples || hasApple[node]
    }
    dfs(0,-1)
    return total
}
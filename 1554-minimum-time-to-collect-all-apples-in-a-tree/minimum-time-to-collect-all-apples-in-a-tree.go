func minTime(n int, edges [][]int, hasApple []bool) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    var dfs func(node, prev int) (int)
    dfs = func(node, prev int) (int) {
        // base

        // logic
        total := 0
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            childTime := dfs(nei, node)
            // we are standing at a parent and looking down
            // curr node is a parent node
            // and we are back from child node recursion
            // now we ask, child has returned any time? ( meaning child had apples underneath itself )
            // or the child itself has apples
            // meaning a node does not check itself, it only looks down at its child and asks the question
            if childTime != 0 || hasApple[nei] {
                total += childTime + 2
            }
        }
        return total
    }
    return dfs(0,-1)
}
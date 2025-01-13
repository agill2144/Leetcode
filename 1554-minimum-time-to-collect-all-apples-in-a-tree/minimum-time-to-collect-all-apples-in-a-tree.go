func minTime(n int, edges [][]int, hasApple []bool) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        a, b := edges[i][0], edges[i][1]
        adjList[a] = append(adjList[a], b)
        adjList[b] = append(adjList[b], a)
    }
    total := 0
    var dfs func(node, prev int) bool
    dfs = func(node, prev int) bool{
        // base

        // logic
        childHasApples := false
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            // ask child whether we need to travel on this edge?
            if dfs(nei, node) {
                // if child says yes, then add a travel time +2
                total += 2
                // if even 1 child out of n said yes, then this node is also a child of some parent
                // meaning, this node now also need to inform its parent that its parent need 
                // to travel on this edge
                // therefore update a flag to tell the parent we need to travel to me(curr node)
                childHasApples = true
            }
        }
        // if either had apples, then this node's parent needs to travel to it
        // or if none of the child had apples, this node itself could have apples
        // in either cases, we need to tell this node's parent node to add travel time 
        // between parent to this node
        if childHasApples || hasApple[node] {return true}
        return false
    }
    dfs(0,-1)
    return total
}
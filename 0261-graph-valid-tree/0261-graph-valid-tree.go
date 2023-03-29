/*
    graph is a valid tree if
    1. fully connected
    2. contains no cycles

    we can simply do a dfs to check for cycles from node 0
    if graph is fully connected then the number of visited nodes
    in visited set should equal to n

    time: o(v+e)
    space: o(v+e)
*/
func validTree(n int, edges [][]int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
        adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])        
    }
    visited := map[int]struct{}{}
    var dfs func(node, prev int) bool
    dfs = func(node, prev int) bool {
        // base
        if _, ok := visited[node]; ok {return true} // cycle detected
        
        // logic
        visited[node] = struct{}{}
        for _, nei := range adjList[node]{
            if nei == prev {continue}
            if hasCycle := dfs(nei, node); hasCycle {return true}
        }
        return false
    }
    
    ok := dfs(0, -1)
    if ok {return false}
    return n == len(visited)
}
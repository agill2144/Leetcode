/*
    - there are disconnected components
    - meaning dfs on all
    - there can be cycles and this is undirected, so maintain a visited array and a prev node as part of dfs
    
    A connected graph with the fewest edges is a tree with n - 1 edges connecting n nodes. 
    So, in order for a graph with n nodes to be connected, 
    the graph must have at least n - 1 edges. 
    It is impossible to connect a graph that has fewer than n - 1 edges. 
    Otherwise, it is always possible to connect the graph.
    
    - count number of connected components
    - to connect all individual connected components,
    - we would need to numConnectedComponents-1 edges
    
    1---2
    |       
    |      4---5
    3
           6
           
    The above input has 3 connected components
    TO CONNECT ALL, how many edges do we need?
    3-1 = 2
    connect 3-4
    connect 3-6
    
    Therefore to make number of individual connected components connect to each other, we need numConnectedComponents-1 edges in total
    
    time: o(v+e)
    space: o(v+e)
*/
func makeConnected(n int, connections [][]int) int {
    if len(connections) < n-1 {return -1}
    adjList := map[int][]int{}
    for i := 0; i < len(connections); i++ {
        adjList[connections[i][0]] = append(adjList[connections[i][0]], connections[i][1])
        adjList[connections[i][1]] = append(adjList[connections[i][1]], connections[i][0])        
    }
    
    visited := map[int]struct{}{}
    var dfs func(curr, prev int)
    dfs = func(curr, prev int) {
        // base
        if _, ok := visited[curr]; ok {
            return
        }
        
        // logic
        visited[curr] = struct{}{}
        for _, node := range adjList[curr] {
            if node == prev {continue}
            dfs(node, curr)
        }
    }
    numConnected := 0
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            numConnected++
            dfs(i, -1)
        }
    }
    return numConnected-1
}
/*
    - there are disconnected components
    - meaning dfs on all
    - there can be cycles and this is undirected, so maintain a visited array and a prev node as part of dfs
    
    A connected graph with the fewest edges is a tree with n - 1 edges connecting n nodes. 
    So, in order for a graph with n nodes to be connected, 
    the graph must have at least n - 1 edges. 
    It is impossible to connect a graph that has fewer than n - 1 edges. 
    Otherwise, it is always possible to connect the graph.
    
    What I didnt understand
    - I tried to find redundant connections ( i.e cycles ) and disconnect ONLY those and connect those to other components
    - no, we are looking for number of connectedComponents and count number of edges to connect them
    - This line "You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected"
        suggests to disconnect existing ones to connect new ones, but logically we can only disconnect existing edge if it was redundant.
    
    - count number of connected components
    - to connect all individual connected components,
    - we would need to numConnectedComponents-1 edges
    
    1--2
    | /    4---5
    |/
    3 
        6
        |
        7
           
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
    cycles := 0
    visited := map[int]struct{}{}
    var dfs func(curr, prev int)
    dfs = func(curr, prev int) {
        // base
        if _, ok := visited[curr]; ok {
            cycles++            
            return
        }
        
        // logic
        visited[curr] = struct{}{}
        for _, node := range adjList[curr] {
            if node == prev {continue}
            dfs(node, curr)
        }
    }
    disconnectedComponentsCount := 0
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            disconnectedComponentsCount++
            dfs(i, -1)
        }
    }
    
    // if there are no disconnected components
    // we do not need to connect anything
    if disconnectedComponentsCount == 0 {return 0}
    
    // or if there are disconnected components
    // and we do not have enough cyclic eges(disconnectedComponentsCount-1);
    // we cannot connect it
    if cycles < disconnectedComponentsCount-1  {return -1}

    return disconnectedComponentsCount-1
}
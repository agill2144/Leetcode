func makeConnected(n int, connections [][]int) int {
    
    // if there are already less than n-1 edges, aint no way
    // we can connect all
    if len(connections) < n-1 {return -1}
    
    
    adjList := map[int][]int{}
    for i := 0; i < len(connections); i++ {
        u,v := connections[i][0], connections[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    
    visited := make([]bool, n)
    cyclicEdgesCount := 0
    var dfs func(node, prev int)
    dfs = func(node, prev int ) {
        // base
        if visited[node] {
            cyclicEdgesCount++
            return
        }
        
        // logic
        visited[node] = true
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node)
        }
    }
    
    connectedComponentsCount := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i, -1)
            connectedComponentsCount++
        }
    }
    
    // to connect 2 individual connected components we need 1 edge ( connectedComponentsCount-1 )
    // to connect 3 individual connected components we need 2 edges ( connectedComponentsCount-1 )
    // to connect 4 individual connected components we need 3 edges  ( connectedComponentsCount-1 )
    extraEdgesNeeded := connectedComponentsCount-1
    
    // we must use existing redundant edges ( we cant create new ones, we have to disconnect existing redundant ones )
    // thats what cyclicEdgesCount is keeping track of
    // if we dont have enough, return and exit early with -1
    if cyclicEdgesCount < extraEdgesNeeded {return -1}
    
    // otherwise we have enough redundant edges that can be used
    // we need to know the min edges needed which is what extraEdgesNeeded is telling us
    return extraEdgesNeeded
    
}
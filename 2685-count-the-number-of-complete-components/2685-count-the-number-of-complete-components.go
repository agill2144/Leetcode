/*
    For a component to be a connected component, there must be n-1 edges at minimum
    so if there are 4 nodes, there must be at min 3 (n-1) edges for this component to be connected
    
    example:
    0-1-2
    num nodes = 3
    num of edges to be connected = 3-1 = 2 (i.e n-1)
    and yes there are 2 edges
    
    example:
    0-1 2
    num nodes = 3
    num of edges to be considered connected = n-1 = 3-1 = 2
    total edges in actual graph = 1
    therefore not a connected component
    
    complete connected component is when EACH NODE MUST HAVE AN EDGE TO ALL OTHER NODES IN THE GRAPH
    ( not adjacent nodes , but ALL nodes)
    there must be an edge from a node to all other nodes in this graph
    
    so then what would be total number of edges to be expected in a complete connected graph?
    assume n ( num nodes ) = 3
    to be considered connected, it must have n-1 edges = 3-1 = 2
    to be considered complete connected, it must have n(n-1) edges = 3(3-1) = 3(2) = 6 edges
    
    for each node there must be n-1 edges
    and if there are n nodes in total
    then total edges must be n * (n-1) edges
    
    example:
    n = 3 ( 0,1,2 )

    for node 0, it must have 2 edges ( 3-1 = 2)
    0 - 2  
     \
      1
    for node 1, it must have 2 edges ( 3-1 = 2)
    1 - 0  
     \
      2
    for node 2, it must have 2 edges ( 3-1 = 2)
    2 - 0  
     \
      1
    
    0 should have 2 edges
    1 should have 2 edges
    2 should have 2 edges
    total = 6
    
    in other words, each node must have an explicit edge to ALL other nodes
    n nodes, each node should have n-1 edges
    therefore total edges = n(n-1)
    
    time:
    o(v+e) to create adjList  + o(v+e) for dfs
    
    space:
    o(v+e) for adjList + o(v) for recursive stack in dfs

*/
func countCompleteComponents(n int, edges [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u, v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    
    countNodes := 0
    countEdges := 0
    visited := make([]bool, n)
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if visited[node] {return}
        
        // logic
        visited[node] = true
        countNodes++
        countEdges += len(adjList[node])
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node)
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i, -1)
            expectedEdges := countNodes*(countNodes-1)
            if expectedEdges == countEdges {
                count++
            }
            countNodes = 0
            countEdges = 0
        }
    }
    return count    
}









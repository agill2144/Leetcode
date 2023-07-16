func distanceToCycle(n int, edges [][]int) []int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
    }
    
    // find nodes part of the cycle
    cyclicNodes := make([]bool, n)
    visited := make([]bool, n)
    cyclicNode := -1

    // N nodes will be connected to a cycle
    // but only handful of those N make the actual cycle
    // therefore we want to make sure, we only mark the nodes that FORM
    // a cycle
    // as soon as we run into a node that creates a cycle, we mark that node,
    // and return back telling previous nodes that they are part of a cycle
    // Once our recursion comes to the marked cyclic node, we mark identifying the cyclic nodes as "done"
    // thats what this flag is for
    
    done := false
    var dfs func(node, prev int) bool
    dfs = func(node, prev int) bool {
        // base
        if visited[node] {cyclicNode = node; return false}
        
        // logic
        visited[node] = true
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            if !dfs(nei, node) && !done {
                cyclicNodes[node] = true
                if cyclicNode == node {done = true}
                return false
            }
        }
        return true
    }
    dfs(0, -1)
    dist := make([]int, n)
    q := [][]int{}

    for i := 0; i < len(cyclicNodes); i++ {
        if cyclicNodes[i] {
            dist[i] = 0
            q = append(q, []int{i, 0})
        } else {
            dist[i] = -1 // not yet visited
        }
    }
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        node := dq[0]
        currDist := dq[1]
        for _, nei := range adjList[node] {
            if dist[nei] == -1 {
                dist[nei] = currDist+1
                q = append(q, []int{nei, currDist+1})
            }
        }
    }
    return dist
}
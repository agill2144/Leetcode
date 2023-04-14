/*
    approach:
    - convert directed to undirected graph
    - store input given edges in a set so they are searchable
    - and run a dfs from node 0 on a directed graph
        - where we have access to prev node as part of its signature
    - the idea becomes, we can search in our set whether an edge from node->prev exists
    - and since we are trying to reorient some roads such that each city can visit the city 0
    - then we just need to make sure adjacent node of 0 can visit 0
        if not, count++
        - then their adjacent nodes need to be able to visit current node and then continue recursing
    
    time: o(v+e)
    space: o(v+e)
*/
func minReorder(n int, connections [][]int) int {
    adjList := map[int][]int{}
    set := map[[2]int]struct{}{}
    for i := 0; i < len(connections); i++ {
        src := connections[i][0]
        dest := connections[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)      
        set[[2]int{src,dest}] = struct{}{}
    }
    visited := make([]bool, n)
    count := 0
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if visited[node] {return}
                
        // logic
        visited[node] = true
        if _, ok := set[[2]int{node,prev}]; !ok && prev != -1 {count++}
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node)
        }
    }
    dfs(0,-1)
    return count
}
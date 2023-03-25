/*
    approach: dfs
    - find number of connected components in a undirected graph
    - given graphs that are disconnected 
    - count number of graphs
    - the input is a adj matrix
        - in adj matrix, i<->j is an edge if at matrix[i][j] = 1
        - because its undirected, we have to add the other direction too, therefore matrix[j][i] will also be 1
        - there will be nodes in the adjMatrix, who do not have an edge to anyone, and thats a valid graph too
            - these are called self nodes where i == j
*/
func findCircleNum(isConnected [][]int) int {
    adjList := map[int][]int{}
    n := len(isConnected)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 && i != j { 
                // add both edges because this is an undirected graph
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
    }
    
    visited := map[int]struct{}{}
    var dfs func(node int)
    dfs = func(node int) {
        // base
        if _, ok := visited[node]; ok {return}
        
        // logic
        visited[node] = struct{}{}
        for _, x := range adjList[node] {
            // optimization: no need to further recurse if x is already visited
            if _, ok := visited[x]; ok {continue}
            dfs(x)
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            dfs(i)
            count++
        }
    }
    return count
}
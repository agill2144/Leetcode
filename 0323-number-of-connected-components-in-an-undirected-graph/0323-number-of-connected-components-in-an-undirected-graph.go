func countComponents(n int, edges [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
        adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])        
    }
    visited := map[int]struct{}{}
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if _, ok := visited[node]; ok {return}
        
        // logic
        visited[node]=struct{}{}
        for _, neighbor := range adjList[node] {
            if neighbor == prev {continue}
            dfs(neighbor, node)
        }
    }
    count := 0
    for i := 0 ; i < n; i++ {
        if _, ok := visited[i]; !ok {
            count++
            dfs(i, -1)
        }
    } 
    return count
}
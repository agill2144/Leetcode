func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
    }
    visited := map[int]struct{}{}
    var dfs func(node int, path map[int]struct{}) bool
    dfs = func(node int, path map[int]struct{}) bool {
        // base
        if _, inPath := path[node]; inPath {return false}
        if len(adjList[node]) == 0 {
            return destination == node
        }
        if _, isVisited := visited[node]; isVisited{return true}
        
        // logic
        visited[node] = struct{}{}
        path[node] = struct{}{}
        if len(adjList[node]) == 0 {return false}
        for _, neighbor := range adjList[node] {
            if ok := dfs(neighbor, path); !ok {return false}
        }
        delete(path, node)
        return true
    }
    return dfs(source, map[int]struct{}{})
}
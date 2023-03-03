func validPath(n int, edges [][]int, source int, destination int) bool {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)

    }
    visited := map[int]struct{}{}
    var dfs func(start int) bool
    dfs = func(start int) bool {
        // base
        if start == destination { return true }
        if _, ok := visited[start]; ok {return false}
        
        // logic
        visited[start] = struct{}{}
        for _, edge := range adjList[start] {
            if ok := dfs(edge); ok {return true}
        }
        return false
    }
    return dfs(source)
    
}

// func validPath(n int, edges [][]int, source int, destination int) bool {
//     adjList := map[int][]int{}
//     for i := 0; i < len(edges); i++ {
//         src := edges[i][0]
//         dest := edges[i][1]
//         adjList[src] = append(adjList[src], dest)
//         adjList[dest] = append(adjList[dest], src)

//     }
//     memo := map[int]bool{}
//     visited := map[int]struct{}{}
//     var dfs func(start int) bool
//     dfs = func(start int) bool {
//         // base
//         if start == destination { return true }
//         if _, ok := visited[start]; ok {return false}
        
//         // logic
//         if _, ok := memo[start]; !ok {
//             visited[start] = struct{}{}
//             for _, edge := range adjList[start] {
//                 if ok := dfs(edge); ok {memo[start]=true; return true}
//             }
//             delete(visited, start)
//             memo[start] = false
//         }
//         return false
//     }
//     return dfs(source)
    
// }
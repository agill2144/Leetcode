func minReorder(n int, connections [][]int) int {
    adjList := map[int][][]int{}
    for i := 0; i < len(connections); i++ {
        src := connections[i][0]
        dest := connections[i][1]
        adjList[src] = append(adjList[src], []int{dest, 1})
        adjList[dest] = append(adjList[dest], []int{src, 0})
    }
    
    count := 0
    visited := map[int]struct{}{}
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if _, ok := visited[node]; ok {return}
        
        // logic
        visited[node] = struct{}{}
        for _, edge := range adjList[node] {
            if edge[0] == prev {continue}
            count += edge[1]
            dfs(edge[0], node)
        }
    }
    dfs(0,-1)
    return count
}

// some what brute force space ( extra set of connections )
// func minReorder(n int, connections [][]int) int {
//     set := map[[2]int]struct{}{}
//     adjList := map[int][]int{}
//     for i := 0; i < len(connections); i++ {
//         adjList[connections[i][0]] = append(adjList[connections[i][0]], connections[i][1])
//         adjList[connections[i][1]] = append(adjList[connections[i][1]], connections[i][0])
//         set[[2]int{connections[i][0], connections[i][1]}] = struct{}{}
//     }
    
//     count := 0
//     visited := map[int]struct{}{}
//     var dfs func(node, prev int)
//     dfs = func(node, prev int) {
//         // base
//         if _, ok := visited[node]; ok {return}
        
//         // logic
//         visited[node] = struct{}{} 
//         // is there a connection/edge from node to prev?
//         // this is why connectons are saved in a  
//         if _, ok := set[[2]int{node, prev}]; !ok && prev != -1 {
//             count++
//         }
//         for _, edge := range adjList[node] {
//             if edge == prev {continue}
//             dfs(edge, node)
//         }
//     }
//     dfs(0,-1)
//     return count
// }
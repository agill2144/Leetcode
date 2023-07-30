// traverse the input without building adjList
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
   
    visited := make([]bool, n)
    var dfs func(row int)
    dfs = func(row int) {
        // base
        if visited[row] {return}
        
        // logic
        visited[row] = true
        for j := 0; j < n; j++ {
            if isConnected[row][j] == 1 && row != j {
                dfs(j)
            }
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i)
            count++
        }
    }
   
    return count
}

// build adjList and traverse
// func findCircleNum(isConnected [][]int) int {
//     adjList := map[int][]int{}
//     n := len(isConnected)
//     for i := 0; i < n; i++ {
//         for j := 0; j < n; j++ {
//             if isConnected[i][j] == 1 && i != j {
//                 adjList[i] = append(adjList[i], j)
//                 adjList[j] = append(adjList[j], i)
//             }
//         }
//     }
//     visited := make([]bool, n)
//     var dfs func(node, prev int)
//     dfs = func(node, prev int) {
//         // base
//         if visited[node] {return}
        
//         // logic
//         visited[node] = true
//         for _, nei := range adjList[node] {
//             if nei == prev {continue}
//             dfs(nei, node)
//         }
//     }
//     count := 0
//     for i := 0; i < n; i++ {
//         if !visited[i] {dfs(i, -1); count++}
//     }
//     return count
// }

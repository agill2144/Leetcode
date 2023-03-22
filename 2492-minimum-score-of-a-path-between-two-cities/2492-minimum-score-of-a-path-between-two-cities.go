

func minScore(n int, roads [][]int) int {
    adjList := map[int][][]int{}
    for i := 0; i < len(roads); i++ {
        src := roads[i][0]
        dest := roads[i][1]
        dist := roads[i][2]
        adjList[src] = append(adjList[src], []int{dest, dist})
        adjList[dest] = append(adjList[dest], []int{src, dist})
    }
    visited := map[int]struct{}{}
    ans := math.MaxInt64
    var dfs func(node []int)
    dfs = func(node []int) {
        // base
        if node == nil {return}
        n := node[0]
        w := node[1]
        if w < ans {
            ans = w
        }
        if _, ok := visited[n]; ok {return}
        visited[n] = struct{}{}
        for _, edge := range adjList[n] {
            dfs(edge)
        } 
    }
    dfs([]int{1, math.MaxInt64})
    return ans
}



// time: o(v+e)
// space: o(v+e) for adjList
// approach: BFS
// func minScore(n int, roads [][]int) int {
//     adjList := map[int][][]int{}
//     for i := 0; i < len(roads); i++ {
//         src := roads[i][0]
//         dest := roads[i][1]
//         dist := roads[i][2]
//         adjList[src] = append(adjList[src], []int{dest, dist})
//         adjList[dest] = append(adjList[dest], []int{src, dist})
//     }
//     // we just need min edge weight from a path that start from 1
//     // there is a gurantee that there is a path from 1 to dest(len of roads)
//     ans := math.MaxInt64
//     visited := make([]bool, n+1)
//     visited[1] = true
//     //              <edge, weight>
//     q := [][]int{{1,math.MaxInt64}}
    
    
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         node := dq[0]
//         for _, edge := range adjList[node] {
//             // this edge may be visited, but if its weight is smaller, save it
//             // if the edge is already visited and we do not add to queue, then we miss a potential smaller ans
//             // therefore the edge weight check is being done here instead of outside this loop ( above )
//             edgeNode := edge[0]
//             edgeWeight := edge[1]
//             if edgeWeight < ans {ans = edgeWeight}

//             if !visited[edgeNode] {
//                 q = append(q, edge)
//                 visited[edgeNode] = true
//             }
//         }
//     }
//     return ans
// }
func findMinHeightTrees(n int, edges [][]int) []int {
    if n < 2 {
        out := []int{}
        for i := 0; i < n; i++ {out = append(out, i)}
        return out
    }

    adjList := map[int][]int{}
    degrees := make([]int, n)
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        adjList[u] = append(adjList[u], v)
        adjList[v] = append(adjList[v], u)
        degrees[u]++
        degrees[v]++
    }

    // find the leaves 
    // leaves are nodes with only 1 edge
    // leaves can never be our MHT roots because they are at the far bottom
    // going from root to leaf will have a max height and we want min height
    // therefore leaves will never be our answer
    // now the idea becomes, cut the leaves and process next layer of leaves
    // once a layer of leaves have been processed, find the parent/adj neighbors
    // once those layer of leaves have been cut, their adj neis edge count drops too
    // therefore the adj nei degrees drop by 1
    // while we drop them, we can check whether this adj nei is now a leaf ? (degrees[nei] == 1)
    // enqueue it if it is.
    // when we do know we have answer ?
    // when the number of nodes left to be processed are <= 2 
    // there will never be more than 2 nodes as our answer
    // draw it out ; 
    // if there is only 1 node in the input, its the answer 
    // if there are 2 nodes in the input, both nodes max depth will be 1, therefore both will be our answers
    // if there are 3 nodes in the input, only 1 node's max depth will be the smallest, that node will be our answer
    // therefore its either 1 node or 2 nodes at max
    leaves := []int{}
    for i := 0; i < len(degrees); i++ {
        if degrees[i] == 1 {
            leaves = append(leaves, i)
        }
    }
    toBeProcessed := n
    for len(leaves) != 0 {
        qSize := len(leaves)
        if toBeProcessed <= 2 {return leaves}
        for qSize != 0 {
            dq := leaves[0]
            leaves = leaves[1:]
            toBeProcessed--
            for _, nei := range adjList[dq] {
                degrees[nei]--
                if degrees[nei] == 1 {
                    leaves = append(leaves, nei)
                }
            }
            qSize--
        }
    }
    return nil
}


// brute force: make each node a root and check its max height
// then find the min height
// then find nodes with that height, those nodes are our answers
// func findMinHeightTrees(n int, edges [][]int) []int {
//     adjList := map[int][]int{}
//     heights := map[int]int{}
//     for i := 0; i < len(edges); i++ {
//         u,v := edges[i][0], edges[i][1]
//         adjList[u] = append(adjList[u], v)
//         adjList[v] = append(adjList[v], u)
//         heights[u] = 0
//         heights[v] = 0
//     }

//     visited := make([]bool, n)
//     var dfs func(r, node, prev, height int)
//     dfs = func(r, node, prev, height int) {
//         // base
//         if visited[node] {return}        

//         // logic
//         visited[node] = true
//         heights[r] = max(heights[r], height)
//         for _, nei := range adjList[node] {
//             if node == prev {continue}
//             dfs(r, nei, node, height+1)
//         }
//     }

//     for i := 0; i < n; i++ {
//         dfs(i, i, -1, 0)
//         visited = make([]bool, n)
//     }

//     minH := math.MaxInt64
//     for _, h := range heights {
//         minH = min(minH, h)
//     }

//     out := []int{}
//     for node, h := range heights {
//         if h == minH {
//             out = append(out, node)
//         }
//     }
//     return out
// }

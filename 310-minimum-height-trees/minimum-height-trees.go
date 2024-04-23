func findMinHeightTrees(n int, edges [][]int) []int {
    if n <= 2 {
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
    // leaves are ones who have exactly 1 edge
    leaves := []int{}
    leftNodes := n
    for i := 0; i < len(degrees); i++ {
        if degrees[i] == 1 {
            leaves = append(leaves, i)
            degrees[i] = 0
        }
    }
    // process the leaves using multi source bfs
    for len(leaves) != 0 {
        if leftNodes <= 2 {return leaves}
        qSize := len(leaves)
        for qSize != 0 {
            dq := leaves[0]
            leaves = leaves[1:]            
            leftNodes--
            for _, nei := range adjList[dq] {
                degrees[nei]--
                // find new adj leaves
                if degrees[nei] == 1 {
                    leaves = append(leaves, nei)
                }
            }
            qSize--
        }
    }
    return nil
}
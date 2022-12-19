func validPath(n int, edges [][]int, source int, destination int) bool {
    adjMatrix := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjMatrix[edges[i][0]] = append(adjMatrix[edges[i][0]], edges[i][1])
        adjMatrix[edges[i][1]] = append(adjMatrix[edges[i][1]], edges[i][0])
    }

    q := []int{source}
    visited := map[int]bool{source:true}
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == destination {return true}
        
        childs := adjMatrix[dq]
        for _, child := range childs {
            _, ok := visited[child]
            if !ok {
                visited[child] = true
                q = append(q, child)
            }
        }

    }
    return false
}
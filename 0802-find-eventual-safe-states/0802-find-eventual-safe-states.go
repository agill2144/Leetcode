func eventualSafeNodes(graph [][]int) []int {
    indegrees := make([]int, len(graph))
    revAdjList := map[int][]int{}
    q := []int{}
    for i := 0; i < len(graph); i++ {
        edges := graph[i]
        for _, edge := range edges {
            revAdjList[edge] = append(revAdjList[edge], i)
        }
        indegrees[i] = len(edges)
        if indegrees[i] == 0 {q = append(q, i)}
    }
    if len(q) == 0 {return nil}
    
    out := []int{}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        out = append(out, dq)
        for _, nei := range revAdjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 {q = append(q, nei)}
        }
    }
    sort.Ints(out)
    return out
}
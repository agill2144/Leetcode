func eventualSafeNodes(graph [][]int) []int {
    adjList := map[int][]int{}
    outdegrees := make([]int, len(graph))
    for i := 0; i < len(graph); i++ {
        outdegrees[i] = len(graph[i])
        for j := 0; j < len(graph[i]); j++ {
            adjList[graph[i][j]] = append(adjList[graph[i][j]], i)
        }
    }
    out := []int{}
    q := []int{}
    for i := 0; i < len(outdegrees); i++ {
        if outdegrees[i] == 0 {
            q = append(q, i)
            out = append(out, i)
        }
    }
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        for _, nei := range adjList[dq] {
            outdegrees[nei]--
            if outdegrees[nei] == 0 {
                q = append(q, nei)
                out = append(out, nei)
            }
        }
    }
    sort.Ints(out)
    return out
}
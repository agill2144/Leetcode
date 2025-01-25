func eventualSafeNodes(graph [][]int) []int {
    adjList := map[int][]int{}
    outdegrees := make([]int, len(graph))
    for i := 0; i < len(graph); i++ {
        outdegrees[i] = len(graph[i])
        for j := 0; j < len(graph[i]); j++ {
            adjList[graph[i][j]] = append(adjList[graph[i][j]], i)
        }
    }
    safeNodes := map[int]bool{}
    q := []int{}
    for i := 0; i < len(outdegrees); i++ {
        if outdegrees[i] == 0 {
            q = append(q, i)
            safeNodes[i] = true
        }
    }
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        for _, nei := range adjList[dq] {
            outdegrees[nei]--
            if outdegrees[nei] == 0 {
                q = append(q, nei)
                safeNodes[nei] = true
            }
        }
    }
    out := []int{}
    for i := 0; i < len(graph); i++ {
        if safeNodes[i] {out = append(out, i)}
    }
    return out
}
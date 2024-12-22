func canFinish(numCourses int, pre [][]int) bool {
    indegrees := make([]int, numCourses)
    adjList := map[int][]int{}
    for i := 0; i < len(pre); i++ {
        a, b := pre[i][0], pre[i][1]
        // a depends on b
        // b is independent / parent
        // { independent -> [dependents...] }
        indegrees[a]++
        adjList[b] = append(adjList[b], a)
    }
    q := []int{}
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {q = append(q, i)}
    }
    completed := 0
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        completed++
        for _, nei := range adjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 {q = append(q, nei)}
        }
    }
    return completed == numCourses
}
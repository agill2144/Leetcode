func canFinish(numCourses int, prerequisites [][]int) bool {
    adjList := map[int][]int{}
    indegrees := make([]int, numCourses)
    for i := 0; i < len(prerequisites); i++ {
        a, b := prerequisites[i][0], prerequisites[i][1]
        adjList[b] = append(adjList[b], a)
        indegrees[a]++
    }
    q := []int{}
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {q = append(q, i)}
    }
    if len(q) == 0 {return false}
    taken := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            taken++
            for _, nei := range adjList[dq] {
                indegrees[nei]--
                if indegrees[nei] == 0 {q = append(q, nei)}
            }
            qSize--
        }
    }
    return numCourses == taken
}
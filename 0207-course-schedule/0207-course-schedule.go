func canFinish(numCourses int, prerequisites [][]int) bool {
    adjList := map[int][]int{}
    indegrees := make([]int, numCourses)
    for i := 0 ; i < len(prerequisites); i++ {
        prereq := prerequisites[i][1]
        course := prerequisites[i][0]
        indegrees[course]++
        adjList[prereq] = append(adjList[prereq], course)
    }
    q := []int{}
    for i := 0 ; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }
    taken := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            taken++
            for _, ele := range adjList[dq] {
                indegrees[ele]--
                if indegrees[ele] == 0 {
                    q = append(q, ele)
                }
            }
            qSize--
        }
    }
    return taken == numCourses
}
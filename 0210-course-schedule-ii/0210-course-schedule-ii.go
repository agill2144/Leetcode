func findOrder(numCourses int, prerequisites [][]int) []int {
    indegrees := make([]int, numCourses)
    adjList := map[int][]int{}
    for i := 0; i < len(prerequisites); i++ {
        course := prerequisites[i][0]
        prereq := prerequisites[i][1]
        adjList[prereq] = append(adjList[prereq], course)
        indegrees[course]++
    }
    
    q := []int{}
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }

    // all courses have 1 or more dependencies, therefore cannot start from anywhere
    if len(q) == 0 {
        return nil
    }
    out := []int{}
    coursesTaken := 0
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        coursesTaken++
        out = append(out, dq)
        childs := adjList[dq]
        for _, child := range childs {
            indegrees[child]--
            if indegrees[child] == 0 {
                q = append(q, child)
            }
        }
    }
    if coursesTaken != numCourses {return nil}
    return out
    
}
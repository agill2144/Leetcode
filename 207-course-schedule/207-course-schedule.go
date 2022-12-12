func canFinish(numCourses int, prerequisites [][]int) bool {
    adjList := map[int][]int{}
    indegrees := make([]int, numCourses)
    
    for i := 0; i < len(prerequisites); i++ {
        independent := prerequisites[i][0] // parent
        dependent := prerequisites[i][1] // child
        adjList[independent] = append(adjList[independent], dependent)
        indegrees[dependent]++
    }
    total := 0
    q := []int{}
    
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }
    if len(q) == 0 {return false}
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        childs := adjList[dq]
        total++
        for i := 0; i < len(childs); i++ {
            child := childs[i]
            indegrees[child]--
            if indegrees[child] == 0 {
                q = append(q, child)
            }
        }
    }
    
    if total != numCourses {return false}
    return true
}
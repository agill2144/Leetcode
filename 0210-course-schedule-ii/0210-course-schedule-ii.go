func findOrder(numCourses int, prerequisites [][]int) []int {
    adjList := map[int][]int{}
    indegrees := make([]int, numCourses)
    for i := 0; i < len(prerequisites); i++ {
        a := prerequisites[i][1]
        b := prerequisites[i][0]
        indegrees[b]++
        adjList[a] = append(adjList[a], b)
    }
    
    q := []int{}
    totalTaken := 0
    // fmt.Println(indegrees, adjList)
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
            totalTaken++
        }
    }
    if len(q) == 0 {return nil}
    
    out := []int{}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        out = append(out, dq)
        for _, nei := range adjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 {
                totalTaken++
                q = append(q, nei)
            }
        }
    }
    
    if totalTaken != numCourses {return nil}
    return out
}
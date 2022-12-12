func findOrder(numCourses int, prerequisites [][]int) []int {
    
    if numCourses == 0 {return nil}
    result := []int{}
    
    graph := map[int][]int{}
    indegrees := make([]int, numCourses)
    
    for i := 0; i < len(prerequisites); i++{
        parent := prerequisites[i][1]
        child := prerequisites[i][0]
        graph[parent] = append(graph[parent], child)
        indegrees[child]++
    }
    
    q := []int{}
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }
    
    if len(q) == 0 {return nil}
    total := 0
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        total++
        result = append(result, dq)
        childs := graph[dq]
        for i := 0; i < len(childs); i++ {
            child := childs[i]
            indegrees[child]--
            if indegrees[child] == 0 {q = append(q, child)}
        }
    }
    if total != numCourses {return nil}
    
//     if len(result) == 0 {
//         return []int{0}
//     }
    return result
}
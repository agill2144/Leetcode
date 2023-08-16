func findOrder(numCourses int, prerequisites [][]int) []int {
    // [a,b] ; a depends on b ; b is independent b->a
    
    
    adjList := map[int][]int{}
    for i := 0; i < len(prerequisites); i++ {
        a,b := prerequisites[i][0], prerequisites[i][1]
        adjList[b] = append(adjList[b], a)
    }
    
    visited := make([]bool, numCourses)
    st := []int{}
    
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return true}
        
        // logic
        visited[node] = true
        path[node] = true
        for _, nei := range adjList[node] {
            if !dfs(nei, path) {return false}
        }
        st = append(st, node)
        path[node] = false
        return true
    }
    
    p := make([]bool, numCourses)
    for i := 0; i < numCourses; i++ {
        if !visited[i] {
            if !dfs(i, p) {return nil}
        }
    }
    // could not take all courses
    if len(st) != numCourses {return nil}
    for i := 0; i < len(st)/2; i++ {
        st[i], st[len(st)-1-i] = st[len(st)-1-i], st[i]
    }
    return st
    
}
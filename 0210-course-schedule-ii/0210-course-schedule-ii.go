func findOrder(numCourses int, prerequisites [][]int) []int {
    adjList := map[int][]int{}
    for i := 0; i < len(prerequisites); i++ {
        a, b := prerequisites[i][0], prerequisites[i][1]
        // a depends on b
        // b is independent
        adjList[b] = append(adjList[b], a)
    }
    
    st := []int{}
    visited := make([]bool, numCourses)
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
        path[node] = false
        st = append(st, node)
        return true
    }

    p := make([]bool, numCourses)
    for i := 0; i < numCourses; i++ {
        if !visited[i] && !dfs(i, p) {return nil}
    }
    if len(st) != numCourses {return nil}
    for i := 0; i < len(st)/2 ; i++ {
        st[i], st[len(st)-1-i] = st[len(st)-1-i], st[i] 
    }
    return st
}
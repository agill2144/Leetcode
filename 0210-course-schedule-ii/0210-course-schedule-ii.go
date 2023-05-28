func findOrder(numCourses int, prerequisites [][]int) []int {
    adjList := map[int][]int{}
    indegrees := make([]int, numCourses)
    for i := 0; i < len(prerequisites); i++ {
        a := prerequisites[i][1]
        b := prerequisites[i][0]
        indegrees[b]++
        adjList[a] = append(adjList[a], b)
    }
    visited := make([]bool, numCourses)
    var dfs func(node int, path []bool) bool
    st := []int{}
    dfs = func(node int, path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return true}
        
        // logic
        path[node] = true
        visited[node] = true
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
    out := []int{}
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        out = append(out,top)
    }
    return out
}
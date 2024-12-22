func findOrder(numCourses int, pre [][]int) []int {
    adjList := map[int][]int{}
    for i := 0; i < len(pre); i++ {
        a, b := pre[i][0], pre[i][1]
        adjList[b] = append(adjList[b], a)
    }
    visited := make([]bool, numCourses)
    path := make([]bool, numCourses)
    completed := 0
    st := []int{}
    var dfs func(curr, prev int, path []bool) bool
    dfs = func(curr, prev int, path []bool) bool {
        // base
        if path[curr] {return false}
        if visited[curr] {return true}

        // logic
        path[curr] = true
        visited[curr] = true
        completed++
        for _, nei := range adjList[curr] {
            if !dfs(nei, curr, path) {return false}
        }
        st = append(st, curr)
        path[curr] = false
        return true
    }
    for i := 0; i < numCourses; i++ {
        if !visited[i] {
            if !dfs(i, -1, path) {return nil}
        }
    }
    if completed != numCourses {return nil}
    for i := 0; i < len(st)/2; i++ {
        st[i], st[len(st)-1-i] = st[len(st)-1-i], st[i]
    }
    return st
}
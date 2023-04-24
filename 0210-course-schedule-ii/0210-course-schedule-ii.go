func findOrder(numCourses int, prerequisites [][]int) []int {
    adjList := map[int][]int{}
    for i := 0; i < len(prerequisites); i++ {
        src := prerequisites[i][0]
        dest := prerequisites[i][1]
        adjList[src] = append(adjList[src], dest)
    }
    visited := make([]bool, numCourses)
    st := []int{}
    var dfs func(node int, path []bool) bool
    dfs = func(node int , path []bool) bool {
        // base
        if path[node] {return false}
        if visited[node] {return true}
        
        
        // logic
        visited[node] = true
        
        // action
        path[node] = true
        for _, nei := range adjList[node] {
            // recurse
            if !dfs(nei, path) {return false}
        }
        
        // backtrack
        path[node] = false
        st = append(st, node)
        return true
    }
    for i := 0; i < numCourses; i++ {
        if !visited[i] {
            if !dfs(i, make([]bool, numCourses)) {return nil}
        }
    }
    return st
}
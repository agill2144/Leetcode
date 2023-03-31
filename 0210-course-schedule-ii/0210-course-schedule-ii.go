// top sort + cycle detection using dfs

func findOrder(numCourses int, prerequisites [][]int) []int {
    adjList := map[int][]int{}
    for i := 0; i < len(prerequisites); i++ {
        // a depends on b
        // b is indepndent
        a := prerequisites[i][0]
        b := prerequisites[i][1]
        adjList[b] = append(adjList[b], a)
    }
    // count := 0
    st := []int{}
    visited := make([]bool, numCourses)
    
    var dfs func(node int, path []bool) bool
    dfs = func(node int, path []bool) bool {
        // base
        // found a cycle
        if path[node] {return true}
        
        // skip already visited node
        if visited[node] {return false}
        
        // logic
        path[node] = true
        visited[node] = true
        for _, nei := range adjList[node] {
            if dfs(nei, path) {return true}
        }
        
        path[node] = false
        st = append(st, node)
        return false
    }
    for i := 0; i < numCourses; i++ {
        if !visited[i] {
            path := make([]bool, numCourses)
            if dfs(i,path) {return nil}
        }
    }
    if len(st) != numCourses {return nil}
    for i := 0; i < len(st)/2; i++ {
        st[i], st[len(st)-1-i] = st[len(st)-1-i], st[i]
    }
    return st
}
func findOrder(n int, pre [][]int) []int {
    adjList := map[int][]int{}    
    for i := 0; i < len(pre); i++ {
        a,b := pre[i][0], pre[i][1]
        adjList[b] = append(adjList[b], a) 
    }
    order := []int{}
    visited := make([]bool, n)
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
        order = append(order, node)
        path[node] = false
        return true
    }
    p := make([]bool, n)
    for i := 0; i < n; i++ {
        if !visited[i] {
            if !dfs(i, p) {return nil}
        }
    }
    if len(order) != n {return nil}
    for i := 0; i < len(order)/2; i++ {
        order[i], order[len(order)-1-i] = order[len(order)-1-i], order[i]
    }
    return order
}
func findOrder(n int, pre [][]int) []int {
    indegrees := make([]int, n)
    adjList := map[int][]int{}    
    for i := 0; i < len(pre); i++ {
        a,b := pre[i][0], pre[i][1]
        adjList[b] = append(adjList[b], a) 
        indegrees[a]++
    }
    visited := make([]bool, n)
    q := []int{}
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            visited[i] = true
            q = append(q, i)
        }
    }
    if len(q) == 0 {return nil}
    order := []int{}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        order = append(order, dq)
        for _, nei := range adjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 && !visited[nei] {
                visited[nei] = true
                q = append(q, nei)
            }
        }
    }
    if len(order) != n {return nil}
    return order
}
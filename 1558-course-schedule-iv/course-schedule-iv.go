// Pre-process a visited matrix
// Start a bfs from each course i and assign for each course j you visit true.
// v = num of vertices
// e = total num of edges
// tc = o( (v+e) * (v+e) ) = o(v+e)^2
// sc = o(2 (v+e)) for adjList and visited matrix and for queue
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    m := numCourses
    n := numCourses
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {visited[i] = make([]bool, n)}
    adjList := map[int][]int{}
    for i := 0; i < len(prerequisites); i++ {
        a,b := prerequisites[i][0], prerequisites[i][1]
        // you must take course a first if you want to take course b.
        // b depends on a
        // a is independent
        adjList[a] = append(adjList[a], b)
    }
    for i := 0; i < n; i++ {
        bfs(i, adjList, visited)
    }

    out := make([]bool, len(queries))
    for i := 0; i < len(queries); i++ {
        u,v := queries[i][0], queries[i][1]
        out[i] = visited[u][v]
    }
    return out
}

func bfs(start int, adjList map[int][]int, visited [][]bool) {
    q := []int{start}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        visited[start][dq] = true
        for _, nei := range adjList[dq] {
            if !visited[start][nei] {
                q = append(q, nei)
                visited[start][nei] = true
            }
        }
    }

}

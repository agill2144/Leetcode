func minimumSemesters(n int, relations [][]int) int {
    adjList := map[int][]int{}
    indegrees := make([]int, n+1)
    for i := 0; i < len(relations); i++ {
        prev := relations[i][0]
        next := relations[i][1]
        indegrees[next]++
        adjList[prev] = append(adjList[prev], next)
    }
    
    q := []int{}
    totalTaken := 0
    
    for i := 1; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
            totalTaken++
        }
    }
    if len(q) == 0 {return -1}
    
    semesters := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            for _, nei := range adjList[dq] {
                indegrees[nei]--
                if indegrees[nei] == 0 {
                    totalTaken++
                    q = append(q, nei)
                }
            }
            qSize--
        }
        semesters++
    }
    
    if totalTaken != n {return -1}
    return semesters
}
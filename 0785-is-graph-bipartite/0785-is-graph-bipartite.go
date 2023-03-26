func isBipartite(graph [][]int) bool {
    colors := make([]int, len(graph))
    q := []int{}
    for i := 0; i < len(graph); i++ {
        
        if colors[i] == 0 {
            // color it
            colors[i] = 1
            q = append(q, i)
            for len(q) != 0 {
                dq := q[0]
                q = q[1:]
                currColor := colors[dq]
                for _, nei := range graph[dq] {
                    neiColor := colors[nei]
                    if neiColor == 0 {
                        if currColor == 1 {colors[nei] = 2} else {colors[nei] = 1}
                        q = append(q, nei)
                    } else if neiColor == currColor {return false}
                }
            }
        }
    }
    return true
}
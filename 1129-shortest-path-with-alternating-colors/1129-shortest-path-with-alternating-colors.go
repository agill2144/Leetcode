func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
    /*
        colors
        0 -> Start
        1 -> blue
        2 -> red
    */
    
    
    adjList := map[int][][]int{}
    for i := 0; i < len(redEdges); i++ {
        adjList[redEdges[i][0]] = append(adjList[redEdges[i][0]], []int{redEdges[i][1], 2})
    }

    for i := 0; i < len(blueEdges); i++ {
        adjList[blueEdges[i][0]] = append(adjList[blueEdges[i][0]], []int{blueEdges[i][1], 1})
    }
    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, 3)
    }
    
    res := make([]int, n)
    for i := 0; i < len(res); i++ {res[i] = math.MaxInt64}
    q := [][]int{{0,0,0}}
    visited[0][0] = true

    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        node := dq[0]
        dist := dq[1]
        color := dq[2]
        if dist < res[node] {res[node] = dist}
        for _, nei := range adjList[node] {
            neiNode := nei[0]
            neiColor := nei[1]
            if !visited[neiNode][neiColor] && neiColor != color {
                q = append(q, []int{neiNode, dist+1, neiColor})            
                visited[neiNode][neiColor] = true
            }
        }
    }
    for i := 0; i < len(res); i++ {if res[i] == math.MaxInt64 { res[i] = -1} }
    return res    
}
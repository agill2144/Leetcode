func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
    /*
        colors;
        0 -> start
        1 -> red
        2 -> blue
    */
    type node struct {
        val int
        color int
    }
    adjList := map[int][]*node{}
    for i := 0; i < len(redEdges); i++ {
        u,v := redEdges[i][0], redEdges[i][1]
        adjList[u] = append(adjList[u], &node{v,1})
    }
    for i := 0; i < len(blueEdges); i++ {
        u,v := blueEdges[i][0], blueEdges[i][1]
        adjList[u] = append(adjList[u], &node{v,2})
    }

    visited := make([][]bool, n)
    for i := 0; i < len(visited); i++ {
        visited[i] = make([]bool, 3)
    }
    dist := make([]int, n)
    for i := 0; i < len(dist); i++ {dist[i] = math.MaxInt64}
    q := [][]int{{0,0,0}} // < node, color, dist >
    visited[0][0] = true
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        currNode, currColor, currDist := dq[0], dq[1], dq[2]
        dist[currNode] = min(currDist, dist[currNode])
        for _, nei := range adjList[currNode] {
            neiNode := nei.val
            neiColor := nei.color
            neiDist := currDist+1
            if neiColor != currColor && !visited[neiNode][neiColor]{
                visited[neiNode][neiColor] = true
                q = append(q,[]int{neiNode,neiColor, neiDist})
            }
        }
    }
    for i := 0; i < len(dist); i++ {if dist[i] == math.MaxInt64 {dist[i] = -1}}
    return dist
}
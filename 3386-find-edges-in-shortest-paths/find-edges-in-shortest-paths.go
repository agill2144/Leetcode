func findAnswer(n int, edges [][]int) []bool {
    adjList := map[int][][]int{}
    for i := 0; i < len(edges); i++ {
        u,v,w := edges[i][0], edges[i][1], edges[i][2]
        adjList[u] = append(adjList[u], []int{v,w})
        adjList[v] = append(adjList[v], []int{u,w})
    }
    fromSrc := dijkstra(0, n, adjList)
    fromDest := dijkstra(n-1, n, adjList)
    shortestDist := fromSrc[n-1]
    out := make([]bool, len(edges))
    for i := 0; i < len(edges); i++ {
        u,v,w := edges[i][0], edges[i][1], edges[i][2]
        if fromSrc[u]+w+fromDest[v] == shortestDist || fromSrc[v]+w+fromDest[u] == shortestDist {
            out[i] = true
        }
    }
    return out
}

// shortest dist from start to all other nodes
func dijkstra(start int, n int, adjList map[int][][]int) []int {
    dist := make([]int, n)
    for i := 0; i < n; i++ {
        dist[i] = math.MaxInt64-100
    }
    dist[start] = 0
    pq := &minHeap{items: []node{}}
    heap.Push(pq, node{start, 0})
    for pq.Len() != 0 {
        dq := heap.Pop(pq).(node)
        currNode := dq.val
        currDist := dq.dist
        if currDist > dist[currNode] {
            continue
        }
        for _, nei := range adjList[currNode] {
            neiNode := nei[0]
            neiDist := nei[1] + currDist
            if neiDist < dist[neiNode] {
                dist[neiNode] = neiDist
                heap.Push(pq, node{val: neiNode, dist: neiDist})
            }
        }
    }
    return dist
}


type node struct {
    val int
    dist int
}
type minHeap struct {
	items []node
}
func (m *minHeap) Less(i, j int) bool {
	return m.items[i].dist < m.items[j].dist
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(node))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}

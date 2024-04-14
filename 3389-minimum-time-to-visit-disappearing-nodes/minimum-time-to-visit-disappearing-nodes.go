func minimumTime(n int, edges [][]int, disappear []int) []int {
    adjList := map[int][][]int{} // u: {v, w}
    for i := 0; i < len(edges); i++ {
        u,v,w := edges[i][0], edges[i][1], edges[i][2]
        adjList[u] = append(adjList[u], []int{v,w})
        adjList[v] = append(adjList[v], []int{u,w})
    }
    dist := make([]int, n)
    for i := 0; i < n; i++ {
        dist[i] = math.MaxInt64
    }
    dist[0] = 0
    pq := &minHeap{items: [][]int{{0,0}}}
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        node := dq[0]
        time := dq[1]
        if time >= disappear[node] || time > dist[node] {continue}

        dist[node] = min(dist[node], time)
        for _, nei := range adjList[node] {
            neiNode := nei[0]
            neiTime := nei[1] + time
            if neiTime > dist[neiNode] {continue}
            heap.Push(pq, []int{neiNode, neiTime})
        }
    }
    for i := 0; i < n; i++ {if dist[i] == math.MaxInt64{dist[i]= -1}}
    return dist
}


type minHeap struct {
    items [][]int // [ [node, weight] ] 
}

func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minHeap) Less(i, j int) bool {
    return m.items[i][1] < m.items[j][1]
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
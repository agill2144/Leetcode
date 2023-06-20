package main


func dijkstra(v int, adj [][]int, start int) []int {
    pq := &minHeap{items: [][]int{}}
    dist := make([]int, v)
    for i := 0; i < len(dist); i++ {
        dist[i] = math.MaxInt64
    }
    dist[start] = 0
    heap.Push(pq, []int{start, 0})
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        node := dq[0]
        distance := dq[1]
        for _, nei := range adj[node] {
            neiNode := nei[0]
            neiDist := nei[1]+distance
            if neiDist < dist[neiNode] {
                dist[neiNode] = neiDist
                heap.Push(pq, []int{neiNode, neiDist})
            }
        }
    }
    return dist
}

type minHeap struct {
    items [][]int // [ [node, dist] ]
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



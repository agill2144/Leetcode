func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
    adjList := map[float64][][]float64{}
    for i := 0; i < len(edges); i++ {
        u,v,w := float64(edges[i][0]), float64(edges[i][1]), succProb[i]
        adjList[u] = append(adjList[u], []float64{v,w})
        adjList[v] = append(adjList[v], []float64{u,w})
    }
    probs := make([]float64, n)
    pq := &maxHeap{items: [][]float64{}}
    heap.Push(pq, []float64{float64(start_node),1.0})
    destNode := float64(end_node)
    for pq.Len() != 0 {
        top := heap.Pop(pq).([]float64)
        currNode := top[0]
        currW := top[1]
        if currNode == destNode {return currW}
        for _, nei := range adjList[currNode] {
            neiNode := nei[0]
            neiWeight := currW * nei[1]
            if neiWeight > probs[int(neiNode)] {
                probs[int(neiNode)] = neiWeight
                heap.Push(pq, []float64{neiNode,neiWeight})
            }
        } 
    }
    return 0.0

}


type maxHeap struct {
    items [][]float64 // [ [node, probSoFar] ]
}

func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.([]float64))}
func (m *maxHeap) Less(i, j int) bool {
    return m.items[i][1] > m.items[j][1]
}
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}


func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    adjList := map[float64][][]float64{}
    for i := 0; i < len(edges); i++ {
        src := float64(edges[i][0])
        dest := float64(edges[i][1])
        w := succProb[i]
        adjList[src] = append(adjList[src], []float64{dest,w})
        adjList[dest] = append(adjList[dest], []float64{src,w})
    }
    probs := make([]float64, n)
    pq := &maxHeap{items: [][]float64{ {float64(start), 1.0} }}
    probs[start] = 1.0
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]float64)
        node := dq[0]
        currProb := dq[1]
        if node == float64(end) {return currProb}
        
        for _, nei := range adjList[node] {
            neiNode := nei[0]
            neiProb := currProb * nei[1]
            if neiProb > probs[int(neiNode)] {
                probs[int(neiNode)] = neiProb
                heap.Push(pq, []float64{neiNode, neiProb})
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
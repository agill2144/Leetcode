func networkDelayTime(times [][]int, n int, k int) int {
    adjList := map[int][][]int{}
    for i := 0; i < len(times); i++ {
        src := times[i][0]
        dest := times[i][1]
        time := times[i][2]
        adjList[src] = append(adjList[src], []int{dest, time})
    }
    allTimes := make([]int, n+1)
    for i := 0; i < n+1; i++ {allTimes[i] = math.MaxInt64}
    pq := &minHeap{items: [][]int{}}
    heap.Push(pq, []int{k, 0})
    allTimes[k] = 0
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        node := dq[0]
        time := dq[1]
        // fmt.Println("dq: ",dq)
        for _, nei := range adjList[node] {
            neiNode := nei[0]
            neiTime := nei[1]
            // fmt.Println(nei, allTimes[neiNode], neiTime+time)
            if time+neiTime < allTimes[neiNode] {
                allTimes[neiNode] = time+neiTime
                heap.Push(pq, []int{neiNode, time+neiTime})
            }
        }
    }
    // fmt.Println(allTimes)
    out := 0
    for i := 1; i < len(allTimes); i++ {
        if allTimes[i] > out {out = allTimes[i]}
    }
    if out == math.MaxInt64 {return -1}
    return out
}

type minHeap struct {
    items [][]int // node, time
}

/*
    type Interface interface {
        Sort
        Pop() T
        Push(x T)
    }
*/

func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minHeap) Less(i, j int) bool {
    if m.items[i][1] == m.items[j][1] {
        return m.items[i][0] < m.items[j][0]
    }
    return m.items[i][1] < m.items[j][1]
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
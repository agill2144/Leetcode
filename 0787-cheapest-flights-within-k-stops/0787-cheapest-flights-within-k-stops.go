func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    adjList := map[int][][]int{}
    for i := 0; i < len(flights); i++ {
        from := flights[i][0]
        to := flights[i][1]
        cost := flights[i][2]
        adjList[from] = append(adjList[from], []int{to, cost})
    }
    steps := make([]int, n)
    for i := 0; i < n; i++ {steps[i] = math.MaxInt64}
    steps[src] = 0
    pq := &minHeap{items: [][]int{{src,0,0}}} // <node, cost, k>
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        currNode := dq[0]
        currCost := dq[1]
        currSteps := dq[2]
        if currNode == dst {return currCost}
        steps[currNode] = currSteps
        for _, adjNodeList := range adjList[currNode] {
            adjNode := adjNodeList[0]
            adjNodeCost := adjNodeList[1]+currCost
            adjNodeSteps := currSteps+1
            // if this node is already visited, only process if we are reaching with shorter amount of steps
            // and make sure we are within k+1 steps limit
            if adjNodeSteps <= k+1 && adjNodeSteps < steps[adjNode] {
                heap.Push(pq, []int{adjNode, adjNodeCost, adjNodeSteps})
            }
        }
    }
    return -1
}


type minHeap struct {
    items [][]int // <node, cost, k>
}

func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minHeap) Less(i, j int) bool {
    //  prioritize min cost
    return m.items[i][1] < m.items[j][1]
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
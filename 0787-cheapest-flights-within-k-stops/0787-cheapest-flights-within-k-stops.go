// 2d dijkstra ( k adds another dimension ), consider all possibilities
// this is brute force, there is a better solution
// time = o(v+e) + o(vk) + o(vk Log vk)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    
    adjList := map[int][][]int{}
    // v = n ( number of verticies)
    // e = total number of edges
    
    // time = o(v+e)
    // space = o(v+e)
    for i := 0; i < len(flights); i++ {
        from := flights[i][0]
        to := flights[i][1]
        cost := flights[i][2]
        adjList[from] = append(adjList[from], []int{to, cost})
    }
    
    // time = o(vk) 
    // space = o(vk)
    nodeStops := make([][]int, n)
    for i := 0; i < len(nodeStops); i++ {
        nodeStops[i] = make([]int, k+1)
        for j := 0; j < len(nodeStops[i]); j++ {
            nodeStops[i][j] = math.MaxInt64
        }
    }
    
    nodeStops[src][0] = 0
    // time = o(vk log vk)
    pq := &minHeap{items: [][]int{{src, 0, k}}}
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        currNode := dq[0]
        currCost := dq[1]
        remaining := dq[2]
        if currNode == dst {
            return currCost
        }
        
        for _, nei := range adjList[currNode] {
            node := nei[0]
            cost := currCost + nei[1]
            if remaining >= 0 && cost < nodeStops[node][remaining] {
                nodeStops[node][remaining] = cost
                heap.Push(pq, []int{node, cost, remaining-1})
            }
        }
    }
    
    // total time = o()
    return -1
}


type minHeap struct {
    items [][]int // <node, cost, k>
}

func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minHeap) Less(i, j int) bool {
    //  minimize cost
    return m.items[i][1] < m.items[j][1]
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
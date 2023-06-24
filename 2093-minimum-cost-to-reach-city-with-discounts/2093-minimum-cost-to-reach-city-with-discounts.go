func minimumCost(n int, highways [][]int, discounts int) int {
    adjList := map[int][][]int{}
    for i := 0; i < len(highways); i++ {
        u,v,w := highways[i][0], highways[i][1], highways[i][2]
        adjList[u] = append(adjList[u], []int{v,w})
        adjList[v] = append(adjList[v], []int{u,w})
    }
    src := 0
    dest := n-1
    cityDiscounts := make([][]int, n)
    for i := 0; i < len(cityDiscounts); i++ {
        cityDiscounts[i] = make([]int, discounts+1)
        for j := 0; j < len(cityDiscounts[i]); j++ {
            cityDiscounts[i][j] = math.MaxInt64
        }
    }
    cityDiscounts[src][0] = 0
    pq := &minHeap{items: [][]int{{src,0,discounts}}}
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        city := dq[0]
        cost := dq[1]
        remaining := dq[2]
        if city == dest {return cost}
        
        for _, nei := range adjList[city] {
            adjCity := nei[0]
            adjCityCost := cost+nei[1]
            
            if  adjCityCost < cityDiscounts[adjCity][remaining] {
                cityDiscounts[adjCity][remaining] = adjCityCost
                heap.Push(pq, []int{adjCity, adjCityCost, remaining})
            }
            
            adjCityCostDiscounted := cost + (nei[1]/2)
            if  remaining > 0 && adjCityCostDiscounted < cityDiscounts[adjCity][remaining-1] {
                cityDiscounts[adjCity][remaining-1] = adjCityCostDiscounted
                heap.Push(pq, []int{adjCity, adjCityCostDiscounted, remaining-1})                                
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
    return m.items[i][1] < m.items[j][1]
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
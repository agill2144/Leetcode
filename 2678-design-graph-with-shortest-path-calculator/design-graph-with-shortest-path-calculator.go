type Graph struct {
    n int
    adjList map[int][][]int
}


func Constructor(n int, edges [][]int) Graph {
    adjList := map[int][][]int{}
    for i := 0; i < len(edges); i++ {
        u,v,w := edges[i][0], edges[i][1], edges[i][2]
        adjList[u] = append(adjList[u], []int{v,w})
    }
    return Graph{n:n, adjList: adjList}
}


func (this *Graph) AddEdge(edge []int)  {
    u,v,w := edge[0],edge[1],edge[2]
    this.adjList[u] = append(this.adjList[u], []int{v,w})
}


func (this *Graph) ShortestPath(node1 int, node2 int) int {
    pq := &minHeap{items: [][]int{{node1,0}}}
    cost := make([]int, this.n)
    for i := 0; i < len(cost); i++ {cost[i] = math.MaxInt64}
    cost[node1] = 0
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        node := dq[0]
        currCost := dq[1]
        if currCost > cost[node] {continue}
        if node == node2 {return cost[node2] }

        for _, nei := range this.adjList[node] {
            neiNode := nei[0]
            neiCost := nei[1]
            totalNeiCost := neiCost + currCost
            if totalNeiCost < cost[neiNode] {
                cost[neiNode] = totalNeiCost
                heap.Push(pq, []int{neiNode, totalNeiCost})
            }
        }
    }
    return -1
}

type minHeap struct {
    items [][]int // [ [node, weight] ]
}
func (m *minHeap) Less(i,j int) bool {
    return m.items[i][1] < m.items[j][1]
}
func (m *minHeap) Swap(i,j int) {
    m.items[i],m.items[j] = m.items[j],m.items[i]
}
func (m *minHeap) Len() int {
    return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
    m.items = append(m.items, x.([]int))
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}




/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
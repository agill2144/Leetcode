func findAnswer(n int, edges [][]int) []bool {
    edgeIdx := map[string]int{}
    adjList := map[int][][]int{}
    for i := 0; i < len(edges); i++ {
        u,v,w := edges[i][0], edges[i][1], edges[i][2]
        adjList[u] = append(adjList[u], []int{v,w})
        adjList[v] = append(adjList[v], []int{u,w})
        edgeIdx[fmt.Sprintf("%v-%v",u,v)] = i
        edgeIdx[fmt.Sprintf("%v-%v",v,u)] = i
    }
    dist := make([]node, n)
    dist[0] = node{0,0,"0"}
    for i := 1; i < n; i++ {
        dist[i] = node{i,math.MaxInt64, ""}
    }
    dest := n-1
    shortestDistToDest := math.MaxInt64
    out := make([]bool, len(edges))

    pq := &minHeap{items: []node{dist[0]}}
    for pq.Len() != 0 {
        dq := heap.Pop(pq).(node)
        currNode := dq.val
        currDist := dq.dist
        currPath := dq.path
        if currNode == dest {
            if currDist < shortestDistToDest {
                shortestDistToDest = currDist
                out = make([]bool, len(edges))
                splitPath := strings.Split(currPath, "-")
                for i := 0; i < len(splitPath); i++ {
                    if i+1 > len(splitPath)-1 {continue}
                    // is this an edge in our input ?
                    pair := fmt.Sprintf("%v-%v", splitPath[i], splitPath[i+1])
                    if idx, ok := edgeIdx[pair]; ok {
                        out[idx] = true
                    }
                }
            } else if currDist == shortestDistToDest {
                splitPath := strings.Split(currPath, "-")
                for i := 0; i < len(splitPath); i++ {
                    // is this an edge in our input ?
                    if i+1 == len(splitPath) {continue}
                    pair := fmt.Sprintf("%v-%v", splitPath[i], splitPath[i+1])
                    if idx, ok := edgeIdx[pair]; ok {
                        out[idx] = true
                    }
                    
                }            }
        }
        
        if currDist > dist[currNode].dist {continue}
        for _, nei := range adjList[currNode] {
            neiNode := nei[0]
            neiDist := currDist + nei[1]
            neiPath := currPath + "-" + fmt.Sprintf("%v", neiNode)
            neiTypedNode := node{neiNode, neiDist, neiPath}
            if neiDist < dist[neiNode].dist {
                heap.Push(pq, neiTypedNode)
                dist[neiNode] = neiTypedNode
            } else if neiDist == dist[neiNode].dist && neiPath != dist[neiNode].path {
                heap.Push(pq, neiTypedNode)
                dist[neiNode] = neiTypedNode
            }
        }
    }
    return out
}


type node struct {
    val int
    dist int
    path string
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

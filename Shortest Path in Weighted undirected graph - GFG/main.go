func main() {
    fmt.Println(shortestPath(5,6,[][]int{ {1,2,2}, {2,5,5}, {2,3,4}, {1,4,1},{4,3,3},{3,5,1} }))
}


/*
You are given a weighted undirected graph having n+1 vertices numbered from 0 to n and m edges describing there are edges between a to b with some weight, find the shortest path between the vertex 1 and the vertex n and if path does not exist then return a list consisting of only -1.

Example:
Input:
n = 5, m= 6
edges = [[1,2,2], [2,5,5], [2,3,4], [1,4,1],[4,3,3],[3,5,1]]
Output:
1 4 3 5
Explaination:
Shortest path from 1 to n is by the path 1 4 3 5

Your Task:
You don't need to read input or print anything. Your task is to complete the function shortestPath() which takes n vertex and m edges and vector of edges having weight as inputs and returns the shortest path between vertex 1 to n.

Expected Time Complexity: O(m* log(n))
Expected Space Complexity: O(n)

Constraint:
2 <= n <= 105
0 <= m <= 105
0<= a, b <= n
1 <= w <= 105
*/

func shortestPath(n int, m int, edges [][]int) []int {
    adjList := map[int][][]int{} // {parent: [ [child, weight], [child, weight] ]}
    for i := 0; i < len(edges); i++ {
        u,v,w := edges[i][0], edges[i][1], edges[i][2]
        adjList[u] = append(adjList[u], []int{v, w})
        adjList[v] = append(adjList[v], []int{u, w})
    }
    start := 1
    end := n
    // dist array will also act as our visited array
    dist := make([]int, n+1)
    // why ? 
    // we need to find and return the path from start->end
    // we can put entire path in pq ( like word ladder 2) and use the last node as current node each time,
    // but whenever we have to add another entry in pq, we would have to create a deep copy of the entire path
    // because arrays are pass by refernce...
    // if we had a new nei to dq'd path and add to pq, then the original dq'd will need to backtrack the last added node to add another valid nei 
    // as soon as it backtracks the last node, and adds a new one, the first modified dq'd entry added to pq will also get modified
    // therefore we have to deep copy
    // other option is to use a parent array
    // parent array acts like a memoization array which helps us find which node a specific node came from
    // so that when pq is done processing, our parent array[end] will say which node it came from
    // and then we can backtrack all the way back to start node using the parent array
    parent := make([]int, n+1)
    
    // fill dist array with inf
    // fill parent array with idx itself
    for i := 0; i < len(dist); i++ {
        dist[i] = math.MaxInt64
        parent[i] = i // this identifies that ith node does not have a parent yet
    }
    
    // distance from start -> start will always be 0
    dist[start] = 0
    // parent of start will be start itself ( i.e this is root node )
    parent[start] = start
    
    // enqueue start node
    pq := &minHeap{items: [][]int{}}
    heap.Push(pq, []int{start, 0})
    
    // process pq
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        node := dq[0]
        weight := dq[1]
        
        for _, nei := range adjList[node] {
            adjNode := nei[0]
            adjNodeWeight := nei[1] + weight
            // if this node weight is better than what we have seen so far, enqueue it and update its parent, dist
            if adjNodeWeight < dist[adjNode] {
                dist[adjNode] = adjNodeWeight
                parent[adjNode] = node
                heap.Push(pq, []int{adjNode, adjNodeWeight})
            }
        }
    }
    
    // how do we ensure we reached end ?
    // well if dist[end] is still math.MaxInt64, that means we never reached the end
    if dist[end] == math.MaxInt64 {return []int{-1}}
    
    // backtrack parent array from end node until the value and idx is not the same or until idx == start node
    ptr := end
    path := []int{}
    for ptr != parent[ptr] {
        path = append(path, ptr)
        ptr = parent[ptr]
    }
    path = append(path, start)
    // reverse the path since we started from the back
    for i := 0; i < len(path)/2; i++ {
        path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
    }
    return path
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



// Word ladder 2 inspired solution
// TLE on GFG ; likely because of the copy we have make each time
// func shortestPath(n int, m int, edges [][]int) []int {
//     adjList := map[int][][]int{}
//     for i := 0; i < len(edges); i++ {
//         src := edges[i][0]
//         dest := edges[i][1]
//         w := edges[i][2]
//         adjList[src] = append(adjList[src], []int{dest,w})
//         adjList[dest] = append(adjList[dest], []int{src,w})
//     }
//     start := 1
//     target := n
//     pq := &minHeap{items: []*minHeapNode{} }
//     heap.Push(pq, &minHeapNode{path: []int{start}, weight: 0})
//     dist := make([]int, n+1)
//     for i := 0; i < len(dist); i++ {
//         dist[i] = math.MaxInt64
//     }
//     dist[start] = 0
//     visited := make([]bool, n+1)
//     visited[start] = true
//     for pq.Len() != 0 {
//         dq := heap.Pop(pq).(*minHeapNode)
//         currPath := dq.path
//         newPath := make([]int, len(currPath))
//         copy(newPath, currPath)
//         currW := dq.weight
//         currNode := currPath[len(currPath)-1]
//         if currNode == target {return currPath}
        
//         for _, nei := range adjList[currNode] {
//             neiNode := nei[0]
//             neiWeight := nei[1]+currW
//             if !visited[neiNode] || neiWeight < dist[neiNode] {
//                 newPath = append(newPath, neiNode)
//                 heap.Push(pq, &minHeapNode{path: newPath, weight: neiWeight})
//                 visited[neiNode] = true
//                 newPath = newPath[:len(newPath)-1]
//             }
//         }
//     }
//     return []int{-1}
// }


// implements heap.Interface

// type minHeapNode struct {
//     path []int
//     weight int
// }

// type minHeap struct {
//     items []*minHeapNode 
// }

// func (m *minHeap) Len() int {return len(m.items)}
// func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
// func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(*minHeapNode))}
// func (m *minHeap) Less(i, j int) bool {
//     iPath := m.items[i].path
//     iNode := iPath[len(iPath)-1]
//     iWeight := m.items[i].weight

//     jPath := m.items[j].path
//     jNode := jPath[len(jPath)-1]
//     jWeight := m.items[j].weight
    
//     if iWeight == jWeight {
//         return iNode < jNode
//     }
//     return iWeight < jWeight
// }
// func (m *minHeap) Pop() interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }

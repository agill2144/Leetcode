func minimumEffortPath(heights [][]int) int {
    m := len(heights)
    n := len(heights[0])
    
    // in a path, we have to take abs diff between adjNode and keep max rolling
    // find a path from src to dest and find the max effort 
    // max effort is the max diff between any 2 adjnode within this path
    // 1 2 2 2 5 = max diff = 3 (5-3)
    // 1 2 8 3 5 = max diff = 6 (8-2)
    // 1 3 5 3 5 = max diff = 2 (5-3)
    // then we have to the min out of all the path 
    // we can brute this using dfs and there is likely repeating subproblems ( DP hint )
    // however this is also like a BFS / graph question
    // kind of like ; find shortest path in binary matrix
    // the only diff is we have weights involved in this question vs binary matrix was vanilla bfs ( level by level )
    // shortest path with weights = dijkstra
    
    // dijsktra we need
    // 1. pq
    // 2. distance array that helps checked if a element is visited or not and if yes, whether we have a reached with smaller distance
    pq := &minHeap{items: [][]int{}}
    distance := make([][]int, m)
    for i := 0; i < m; i++ {
        distance[i] = make([]int, n)
        for j := 0; j < n; j++ {
            distance[i][j] = math.MaxInt64
        }
    }
    // distance of start node will always be 0
    distance[0][0] = 0
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    heap.Push(pq, []int{0,0,0})
    
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        cr := dq[0]
        cc := dq[1]
        currentEffort := dq[2]
        if cr == m-1 && cc == n-1 {return currentEffort}
        
        for _, dir := range dirs {
            nr := cr+dir[0]
            nc := cc+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n {
                maxEffortInThisPath := max(abs(heights[cr][cc]-heights[nr][nc]),currentEffort)
                if maxEffortInThisPath < distance[nr][nc] {
                    distance[nr][nc] = maxEffortInThisPath
                    heap.Push(pq, []int{nr,nc, maxEffortInThisPath})
                }
            }
        }
    }
    return -1
}

// follow up, give me the path, use parent array to backtrack from end node :) 
// or brute force ( put path in pq but deal with time and space complex of deep-copy)

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}
func max(x, y int) int {
    if x > y {return x}
    return y
}
 
type minHeap struct {
    items [][]int // [ [r,c,distance] ] 
}

func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minHeap) Less(i, j int) bool {
    return m.items[i][2] < m.items[j][2]
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
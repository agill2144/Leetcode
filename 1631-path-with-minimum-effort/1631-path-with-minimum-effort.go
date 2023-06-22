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
        currentDist := dq[2]
        // as soon as we run into our destination
        // we CAN exit because this is PQ
        // PQ will prioritize smaller distances therefore
        // we are guranteed that if there are multiple ans in PQ, the smaller one will be processed first
        // and we need the smallest ans
        if cr == m-1 && cc == n-1 {return currentDist}
        
        for _, dir := range dirs {
            nr := cr+dir[0]
            nc := cc+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n {
                neiDist := abs(heights[nr][nc] - heights[cr][cc])
                maxDistInPath := max(neiDist, currentDist)
                if maxDistInPath < distance[nr][nc] {
                    distance[nr][nc] = maxDistInPath
                    heap.Push(pq, []int{nr,nc, maxDistInPath})
                }
            }
        }
    }
    return -1

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

// Brute force BFS, we process cells multiple times if we run into a path that results into smaller effort
// this is like a graph with edges weights, we need to minimize final efforts
// in a graph with edge weights, we may already have an answer for a cell via a path
// another path could come in with a better ans for the same cell, and that means, re-process an already processed cell
// which also means, re-process all other cells that are connected this cell again, because a better answer for a cell
// also means better ans for other cells connected to it
// Instead we need to prioritize smaller efforts first instead to avoid redundant processing of the same cell over and over again
// Graph with edge weights + shortest path + minimize something + vanilla bfs has to reprocess same cells again = Dijkstra is the ans
// func minimumEffortPath(heights [][]int) int {
//     m := len(heights)
//     n := len(heights[0])
    
//     dist := make([][]int,m)
//     for i := 0; i < len(dist); i++ {
//         dist[i] = make([]int, n)
//         for j := 0; j < n; j++ {
//             dist[i][j] = math.MaxInt64
//         }
//     }
//     dist[0][0] = 0
//     dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
//     q := [][]int{{0,0}}
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         cr := dq[0]
//         cc := dq[1]
//         currDist := dist[cr][cc]
        
//         for _, dir := range dirs {
//             nr := cr + dir[0]
//             nc := cc + dir[1]
//             if nr >= 0 && nr < m && nc >= 0 && nc < n {
//                 maxDiffInPath := max(currDist, abs(heights[nr][nc] - heights[cr][cc]))
//                 if maxDiffInPath < dist[nr][nc] {
//                     dist[nr][nc] = maxDiffInPath
//                     q = append(q, []int{nr,nc})
//                 }
//             }
//         }

//     }
//     return dist[m-1][n-1]
// }

func abs(x int) int {
    if x < 0 {return x *- 1}
    return x
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
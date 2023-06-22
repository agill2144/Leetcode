func minimumEffortPath(heights [][]int) int {
    m := len(heights)
    n := len(heights[0])
    
    distances := make([][]int,m)
    for i := 0; i < len(distances); i++ {
        distances[i] = make([]int, n)
        for j := 0; j < n; j++ {
            distances[i][j] = math.MaxInt64
        }
    }
    distances[0][0] = 0
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    pq := &minHeap{items: [][]int{{0,0,0}}}
    
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        cr := dq[0]
        cc := dq[1]
        // or we also use dq[2] becuase pq has distances as part of each path ( so it knows how to sort )
        currentDist := distances[cr][cc] 
    
        if cr == m-1 && cc == n-1 {return currentDist}
        
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n {
                neiDist := abs(heights[cr][cc] - heights[nr][nc])
                maxDistInPath := max(currentDist, neiDist)
                if maxDistInPath < distances[nr][nc] {
                    distances[nr][nc] = maxDistInPath
                    heap.Push(pq, []int{nr,nc,maxDistInPath})
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
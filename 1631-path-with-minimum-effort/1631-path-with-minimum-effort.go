
// Brute force BFS, we process cells multiple times if we run into a path that results into smaller effort
// this is like a graph with edges weights, we need to minimize final efforts
// in a graph with edge weights, we may already have an answer for a cell via a path
// another path could come in with a better ans for the same cell, and that means, re-process an already processed cell
// which also means, re-process all other cells that are connected this cell again, because a better answer for a cell
// also means better ans for other cells connected to it
// Instead we need to prioritize smaller efforts first instead to avoid redundant processing of the same cell over and over again
// Graph with edge weights + shortest path + minimize something + vanilla bfs has to reprocess same cells again = Dijkstra is the ans
func minimumEffortPath(heights [][]int) int {
    m := len(heights)
    n := len(heights[0])
    
    // this is like classic dist array but rebranded as "Efforts"
    efforts := make([][]int, m)
    for i := 0; i < m; i++ {
        efforts[i] = make([]int, n)
        for j := 0; j < n; j++ {
            efforts[i][j] = math.MaxInt64
        }
    }
    efforts[0][0] = 0
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    pq := &minHeap{items: [][]int{{0,0,0}}}
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        cr,cc,currEffort := dq[0],dq[1],dq[2]
        if cr == m-1 && cc == n-1 {return currEffort}
        
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n {
                effort := abs(heights[cr][cc]-heights[nr][nc])
                effort = max(effort, currEffort)
                if effort < efforts[nr][nc] {
                    efforts[nr][nc] = effort
                    heap.Push(pq, []int{nr,nc,effort})
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

// vanilla bfs
// instead of going depth first, and increasing redundant processing of nodes,
// lets go level by level and see all adjacent nodes at the same time
// func minimumEffortPath(heights [][]int) int {
//     m := len(heights)
//     n := len(heights[0])
    
//     dist := make([][]int, m)
//     for i := 0; i < m; i++ {
//         dist[i] = make([]int, n)
//         for j := 0; j < n; j++ {
//             dist[i][j] = math.MaxInt64
//         }
//     }
//     dist[0][0] = 0
//     dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
//     minEffort := math.MaxInt64
    
//     q := [][]int{{0,0,0}} // r,c,dist
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         cr, cc, currEffort := dq[0], dq[1], dq[2]

//         if dist[cr][cc] < currEffort {
//             // we have a better one in the queue, skip this
//             continue
//         }
        
//         if cr == m-1 && cc == n-1 {
//             minEffort = min(currEffort, minEffort)
//             continue
//         }
        
//         for _, dir := range dirs {
//             nr := cr+dir[0]
//             nc := cc+dir[1]
//             if nr >= 0 && nr < m && nc >= 0 && nc < n {
//                 effort := abs(heights[cr][cc]-heights[nr][nc])
//                 effort = max(currEffort, effort)
//                 if effort < dist[nr][nc] {
//                     dist[nr][nc] = effort
//                     q = append(q, []int{nr,nc,effort})
//                 }
//             }
//         }
//     }
//     return minEffort
// }


// brute force, try all paths, backtrack
// for each cell, we have 4 choices/dirs
// there are m*n cells
// time = 4^mn
// space = mn (recursive stack)
// func minimumEffortPath(heights [][]int) int {
//     m := len(heights)
//     n := len(heights[0])
    
//     dist := make([][]int, m)
//     for i := 0; i < m; i++ {
//         dist[i] = make([]int, n)
//         for j := 0; j < n; j++ {
//             dist[i][j] = math.MaxInt64
//         }
//     }
//     dist[0][0] = 0
//     dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
//     minEffort := math.MaxInt64
//     var dfs func(r, c int)
//     dfs = func(r, c int) {
//         // base
//         if r == m-1 && c == n-1 {
//             minEffort = min(dist[r][c], minEffort)
//             return
//         }
        
        
//         // logic
//         for _, dir := range dirs {
//             nr := r+dir[0]
//             nc := c+dir[1]
//             if nr >= 0 && nr < m && nc >= 0 && nc < n {
//                 absDiff := abs(heights[r][c]-heights[nr][nc])
//                 absDiff  = max(dist[r][c], absDiff)
//                 if dist[nr][nc] == math.MaxInt64 {
//                     // action
//                     dist[nr][nc] = absDiff
//                     // recurse
//                     dfs(nr, nc)
//                     // backtrack
//                     dist[nr][nc] = math.MaxInt64
//                 }
//             }
//         }
//     }
//     dfs(0,0)
//     return minEffort
// }

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}

func max(x, y int) int {
    if x > y {return x}
    return y
}

func min(x, y int) int {
    if x < y {return x}
    return y
}
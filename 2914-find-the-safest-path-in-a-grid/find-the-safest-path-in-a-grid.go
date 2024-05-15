func maximumSafenessFactor(grid [][]int) int {
    n := len(grid)
    if n == 0 {return 0}
    if grid[0][0] == 1 {return 0}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    q := [][]int{}
    thief := -1
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                q = append(q, []int{i,j})
                grid[i][j] = thief
            }
        }
    }

    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0] // a
        cc := dq[1] // b
        val := grid[cr][cc]
        if val == thief {val=0}
        for _, dir := range dirs {
            nr := cr+dir[0] // x
            nc := cc+dir[1] // y
            // looking for empty cells ( val = 0)
            if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 0 {
                // |a - x| + |b - y|
                dist := abs(cr-nr) + abs(cc-nc)
                grid[nr][nc] = dist+val
                q = append(q, []int{nr,nc})
            }
        }
    }

    visited := -2
    pq := &maxHeap{items: [][]int{ {0,0,grid[0][0]} }}
    minDist := grid[0][0]
    fmt.Println(grid)
    grid[0][0] = visited
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        cr := dq[0]
        cc := dq[1]
        dist := dq[2]
        minDist = min(minDist, dist)
        if cr == n-1 && cc == n-1 {break}
        for _, dir := range dirs {
            nr := cr+dir[0] 
            nc := cc+dir[1] 
            if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] != visited {
                nd := grid[nr][nc]
                if nd == thief {nd = 0}            
                heap.Push(pq, []int{nr,nc,nd})
                grid[nr][nc] = visited
            }
        }
    }
    
    return minDist
}



type maxHeap struct {
    items [][]int // [ [r,c,distance] ] 
}

func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *maxHeap) Less(i, j int) bool {
    return m.items[i][2] > m.items[j][2]
}
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}


func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}
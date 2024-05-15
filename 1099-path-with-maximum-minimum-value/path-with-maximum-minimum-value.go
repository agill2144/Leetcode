func maximumMinimumPath(grid [][]int) int {
    minVal := math.MaxInt64
    m := len(grid)
    n := len(grid[0])
    sr, sc := 0,0
    dr, dc := m-1, n-1
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    pq := &maxHeap{items: [][]int{{sr,sc,grid[sr][sc]}}}
    visited := -1
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        cr, cc, cv := dq[0], dq[1], dq[2]
        minVal = min(minVal, cv)
        if cr == dr && cc == dc {return minVal}
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != visited {
                heap.Push(pq, []int{nr,nc, grid[nr][nc]})
                grid[nr][nc] = visited
            }
        }
    }
    return minVal
}



type maxHeap struct {
	items [][]int // [ [r,c,val]
}

func (m *maxHeap) Len() int {
	return len(m.items)
}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i][2] > m.items[j][2]
}

func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func (m *maxHeap) Push(x any) {
	m.items = append(m.items, x.([]int))
}

func (m *maxHeap) Pop() any {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}

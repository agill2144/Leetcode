func minimumObstacles(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    dr, dc := m-1, n-1
    visited := -1
    obsCount := 0
    if grid[0][0] == 1 {obsCount = 1}
    pq := &minHeap{items: [][]int{{0,0,obsCount}}}
    grid[0][0] = visited
    for pq.Len() != 0 {
        dq := heap.Pop(pq).([]int)
        cr := dq[0]
        cc := dq[1]
        currObsCount := dq[2]
        if cr == dr && cc == dc {return currObsCount}
        for _, dir := range dirs {
            nr := cr + dir[0]
            nc := cc + dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != visited {
                neiObsCount := currObsCount
                if grid[nr][nc] == 1 {neiObsCount++}
                grid[nr][nc] = visited
                heap.Push(pq, []int{nr,nc,neiObsCount})
            }
        }
    }
    return -1
}
type minHeap struct {
	items [][]int // [ [r,c,obsCount] ]
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i][2] < m.items[j][2]
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
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

func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })
    // tracks meeting end times
    rooms := &minHeap{items: []int{}}
    
    for i := 0; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]
        if rooms.Len() == 0 {
            heap.Push(rooms, end)
        } else {
            earliestEnd := rooms.items[0]
            if start >= earliestEnd {
                heap.Pop(rooms)
            }
            heap.Push(rooms, end)
        }
    }
    return rooms.Len()
}


type minHeap struct {
	items []int
}
func (m *minHeap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Less(i, j int) bool {return m.items[i] < m.items[j]}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
func (m *minHeap) Pop()interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
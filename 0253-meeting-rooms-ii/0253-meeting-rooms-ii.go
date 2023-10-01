func minMeetingRooms(intervals [][]int) int {
    sort.SliceStable(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0] 
    })
    // keeps track of earliest end times
    // we use this to identify whether a meeting room can be re-used or not
    // we can re-use a meeting room if current start time is after the earliest end time
    mn := &minHeap{items: []int{}}
    for i := 0; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]
        if mn.Len() != 0 {
            earliestEndTime := mn.items[0]
            if start >= earliestEndTime {
                heap.Pop(mn)
            }
        }
        heap.Push(mn, end)
    }
    return mn.Len()
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
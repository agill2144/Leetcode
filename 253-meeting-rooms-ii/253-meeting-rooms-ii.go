/*
    time: o(nlogn) to sort + o(nlogk) for minHeap (avg case)
    time: o(nlogn) to sort + o(nlogn) for minHeap (worst case all meetings started at the same time)
    
    space: o(n) for heap (worst case) and o(k) for avg case
*/
func minMeetingRooms(intervals [][]int) int {
    if intervals == nil || len(intervals) == 0 {
       return 0
   }
   sort.Slice(intervals, func(i,j int)bool {
       return intervals[i][0] < intervals[j][0]
   })

   mn := &minHeap{items: []int{}}
   for i := 0; i < len(intervals); i++ {
       startTime := intervals[i][0]
       endTime := intervals[i][1]
       if len(mn.items) == 0 {
           heap.Push(mn, endTime)
       } else {
           earliestEndTime := mn.items[0]
           if startTime < earliestEndTime {
               heap.Push(mn, endTime)
           } else {
               heap.Pop(mn)
               heap.Push(mn, endTime)
           }
       }
   }

    return len(mn.items)
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
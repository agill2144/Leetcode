func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    rooms := &minHeap{items: []int{}}
    for i := 0; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]
        // can this meeting take an existing room?
        if rooms.Len() != 0 {
            if start >= rooms.items[0] {
                // yes, make this room avail ( i.e remove from heap )
                heap.Pop(rooms)
            }
        }
        heap.Push(rooms, end)
    }
    return len(rooms.items)
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
// brute force
// TC = o(nlogn) + o(n^2)
// space = o(n)
// func minMeetingRooms(intervals [][]int) int {
//     sort.Slice(intervals, func(i, j int)bool{
//         return intervals[i][0] < intervals[j][0]
//     })

//     // rooms keeping track of which meeting is ending at what time
//     // if we have a new meeting, we can check with what rooms are
//     // free, and the rooms that will be free are the ones whose 
//     // meetings have already ended compared currentTime
//     rooms := []int{}
//     for i := 0; i < len(intervals); i++ {
//         start := intervals[i][0]
//         end := intervals[i][1]
//         if len(rooms) == 0 {
//             rooms = append(rooms, end)
//         } else {
//             foundExistingRoom := false
//             for j := 0; j < len(rooms); j++ {
//                 if start >= rooms[j] {
//                     rooms[j] = end
//                     foundExistingRoom = true
//                     break
//                 }
//             }
//             // if none of the rooms are avail, we need a new one
//             if !foundExistingRoom {
//                 rooms = append(rooms, end)
//             }
//         }
//     }
//     return len(rooms)
// }
/*
    we have some time slots in hand
    - lets make sure they are sorted by start time
    - so we can correctly assign them rooms
    
    - we need some data structure that helps us identify total rooms used so far
    - when we have a time slot in hand, and none of the rooms have been used at, assign this meeting a room
    - when we have a time slot in hand, and there are some meetings happening, we need to know, whether
        - we can re-use any old room ?
            - this can only happen if some meeting has ended and that room is empty now
        - if we cant re-use any old room, then we must allocate another room...
    
    - we can toss each interval pair in another array
        - each node in the array acts as a meeting room
    - and then when we have to identify whether we can re-use any old room
        - we have to check if current meeting start time is happening before any old meeting end time
        - if there was any such meeting that ended before this one is starting, we can take that room
        - i.e change the value of that element to current meeting's end time
    
    - finally once all time slots are processed, the len of array is the min meeting rooms needed
    
    time = o(n^2) because for each slot we had to go search for a room whose meeting had ended ( i.e endTime < currentStartTime )
    
    - the bottle neck is searching the array for each time-slot
    - if we kept the array sorted while inserting endTimes for each time-slot, then we can make this time better
    - OR
    - we can use a data structure like "minHeap" which sorts by endTime, therefore the oldest endTime will be at the top
    - instead of each element in array being as meeting room, each node in minHeap is a meeting room
    - now when we have a time-slot in hand, and top of the heap shows earliest end time
    - we can easily check whether we can re-use the same room or not.
    - if we can, pop the top
    - and insert the new one!
    - if we cannot, do not pop anything, it means all meetings are still going on and we need another room
        - just insert in this case
        
    time = o(nlogn) + o(n*logHeapSize) = o(nlogn) ;
        - worst case all time-slots are overlapping, and we allocate each slot its own room
        - therefore heapSize becomes n , therefore logHeapSize = logn
    space = o(n)
    
*/
func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
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
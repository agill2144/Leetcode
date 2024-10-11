func mostBooked(n int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int)bool{
        return meetings[i][0] < meetings[j][0]
    })
    usedRooms := &minHeap{rooms:[]*room{}}
    unusedRooms := &unusedRoomMinHeap{rooms: []*room{}}
    count := make([]int, n)
    for i := 0; i < n; i++ { heap.Push(unusedRooms, &room{id:i}) }
    
    for i := 0; i < len(meetings); i++ {
        start := meetings[i][0]
        end := meetings[i][1]

        for usedRooms.Len() > 0 && start >= usedRooms.rooms[0].endTime {
            heap.Push(unusedRooms, heap.Pop(usedRooms).(*room))
        }
        if unusedRooms.Len() > 0 {
            newRoom := heap.Pop(unusedRooms).(*room)
            count[newRoom.id]++
            newRoom.endTime = end
            heap.Push(usedRooms, newRoom)
        } else {
            earliestEndingRoom := heap.Pop(usedRooms).(*room)
            count[earliestEndingRoom.id]++
            earliestEndingRoom.endTime += (end-start)
            heap.Push(usedRooms, earliestEndingRoom)
        }

        

    }
    max := -1
    id := -1
    for i := 0; i < len(count); i++ {
        if count[i] > max {
            max = count[i]
            id = i
        }
    }
    return id
}

type room struct {
    endTime int
    count int
    id int
}
type unusedRoomMinHeap struct {
    rooms []*room
}
type minHeap struct {
    rooms []*room
}


func (m *minHeap) Less(i,j int) bool {
    if m.rooms[i].endTime == m.rooms[j].endTime {
        return m.rooms[i].id < m.rooms[j].id
    }
    return m.rooms[i].endTime < m.rooms[j].endTime
}
func (m *minHeap) Swap(i,j int) {
    m.rooms[i],m.rooms[j] = m.rooms[j],m.rooms[i]
}
func (m *minHeap) Len() int {
    return len(m.rooms)
}
func (m *minHeap) Push(x interface{}) {
    m.rooms = append(m.rooms, x.(*room))
}
func (m *minHeap) Pop() interface{} {
    out := m.rooms[len(m.rooms)-1]
    m.rooms = m.rooms[:len(m.rooms)-1]
    return out
}


func (m *unusedRoomMinHeap) Less(i,j int) bool {
    return m.rooms[i].id < m.rooms[j].id
}
func (m *unusedRoomMinHeap) Swap(i,j int) {
    m.rooms[i],m.rooms[j] = m.rooms[j],m.rooms[i]
}
func (m *unusedRoomMinHeap) Len() int {
    return len(m.rooms)
}
func (m *unusedRoomMinHeap) Push(x interface{}) {
    m.rooms = append(m.rooms, x.(*room))
}
func (m *unusedRoomMinHeap) Pop() interface{} {
    out := m.rooms[len(m.rooms)-1]
    m.rooms = m.rooms[:len(m.rooms)-1]
    return out
}
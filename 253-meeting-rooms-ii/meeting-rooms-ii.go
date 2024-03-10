func minMeetingRooms(intervals [][]int) int {
    mn := &minHeap{items: []int{}}
    sort.Slice(intervals, func(i,j int)bool{
        return intervals[i][0] < intervals[j][0]
    })

    for i := 0; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if len(mn.items) == 0 {
            heap.Push(mn, end)
            continue
        }
        earliestEndTime := mn.items[0]
        if start >= earliestEndTime {
            heap.Pop(mn)
            heap.Push(mn, end)
        } else {
            heap.Push(mn, end)
        }
    }
    return len(mn.items)
}

/*
    type Interface interface {
        sort.Interface
        Pop() int
        Push(x int)
    }
*/  

type minHeap struct {
    items []int
}
func (m *minHeap) Len() int {
    return len(m.items)
}

func (m *minHeap) Less(i, j int) bool {
    return m.items[i] < m.items[j]
}

func (m *minHeap) Swap(i, j int) {
    m.items[i],m.items[j] = m.items[j], m.items[i]
}

func (m *minHeap) Push(x interface{}) {
    m.items = append(m.items, x.(int))
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
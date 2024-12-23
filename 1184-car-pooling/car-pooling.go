func carPooling(trips [][]int, capacity int) bool {
    // trips[i] = [numPassengersi, fromi, toi]
    sort.Slice(trips, func(i, j int)bool{
        return trips[i][1] < trips[j][1]
    })
    count := 0
    mn := &minHeap{items: [][]int{} } // [ [dropOffLocation, numOfPassengers ]]
    for i := 0; i < len(trips); i++ {
        pickUp, dropOff, c := trips[i][1], trips[i][2], trips[i][0]
        // drop off all passengers whose dropOff point was earlier than current point
        for mn.Len() > 0 && pickUp >= mn.items[0][0] {
            top := heap.Pop(mn).([]int)
            count -= top[1]
        }
        count += c
        if count > capacity {return false}
        heap.Push(mn, []int{dropOff, c})
    }
    return true
}

type minHeap struct {
	items [][]int
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i][0] < m.items[j][0]
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

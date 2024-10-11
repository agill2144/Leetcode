func smallestChair(times [][]int, targetFriend int) int {
    targetFriendTime := times[targetFriend]
    sort.Slice(times, func(i, j int)bool {
        return times[i][0] < times[j][0]
    })
    // create chairs
    chairs := &chairMinHeap{items: []int{}}
    for i := 0; i < len(times); i++ {heap.Push(chairs, i)}
    mn := &minHeap{items: [][]int{}}
    for i := 0; i < len(times); i++ {
        arrive, depart := times[i][0], times[i][1]
        // free up used chairs
        for mn.Len() > 0 && arrive >= mn.items[0][1] {
            heap.Push(chairs, heap.Pop(mn).([]int)[0])
        }
        // give chair to friend
        cid := heap.Pop(chairs).(int)
        if arrive == targetFriendTime[0] {return cid}
        heap.Push(mn, []int{cid, depart})
        // if this friend is our target friend
        // we have its chairIdx (cid), return it
    }
    return -1

}

type chairMinHeap struct {
	items []int
}

func (m *chairMinHeap) Less(i, j int) bool {
	return m.items[i] < m.items[j]
}
func (m *chairMinHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *chairMinHeap) Len() int {
	return len(m.items)
}
func (m *chairMinHeap) Push(x interface{}) {
	m.items = append(m.items, x.(int))
}
func (m *chairMinHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}



type minHeap struct {
	items [][]int // [[chairIdx, endTime]]
}

func (m *minHeap) Less(i, j int) bool {
	if m.items[i][1] == m.items[j][1] {
        return m.items[i][0] < m.items[j][0]
    }
    return m.items[i][1] < m.items[j][1]
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

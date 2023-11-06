type SeatManager struct {
    mn *minHeap
    maxSize int
}

/*
requirement : reserve() Fetches the smallest-numbered unreserved seat, reserves it, and returns its number.
we need the earliest unserved seat, therefore minHeap
*/

func Constructor(n int) SeatManager {
    mn := &minHeap{}
    for i := 1; i <= n; i++ {
        heap.Push(mn, i)
    }
    return SeatManager{mn:mn, maxSize:n}
}


func (this *SeatManager) Reserve() int {
    top := heap.Pop(this.mn).(int)
    return top
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(this.mn, seatNumber)
}


type minHeap struct {
    items []int // seats that are not-yet-reserved
}

func (m *minHeap) Less(i,j int) bool {
    return m.items[i] < m.items[j]
}
func (m *minHeap) Swap(i,j int) {
    m.items[i],m.items[j] = m.items[j],m.items[i]
}
func (m *minHeap) Len() int {
    return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
    m.items = append(m.items, x.(int))
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
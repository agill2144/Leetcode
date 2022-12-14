type MedianFinder struct {
    left *maxHeap
    right *minHeap
    items []int
}


func Constructor() MedianFinder {
    return MedianFinder{
        left: &maxHeap{items: []int{}},
        right: &minHeap{items: []int{}},
        items: []int{},
    }
}


func (this *MedianFinder) AddNum(num int)  {
    this.items = append(this.items, num)
    
    for i := len(this.items)-2; i >= 0; i-- {
        if this.items[i+1] < this.items[i] {
            this.items[i+1], this.items[i] = this.items[i], this.items[i+1]
        } else {
            break
        }
    }
    
//     heap.Push(this.left, num)
    
//     // ensure left top <= right top
//     if this.left.Len() != 0 && this.right.Len() != 0 && 
//     this.left.items[0] > this.right.items[0] {
//         heap.Push(this.right, heap.Pop(this.left))
//     }
    
//     // ensure its balanced, so that the median is at the top of surface of both heaps
//     // otherwise median could get burried somewhere deep inside a heap
//     if this.left.Len() > this.right.Len() + 1 {
//         heap.Push(this.right, heap.Pop(this.left))
//     } else if this.right.Len() > this.left.Len() {
//         heap.Push(this.left, heap.Pop(this.right))
//     }
        
}


func (this *MedianFinder) FindMedian() float64 {
    mid := (len(this.items)-1)/2
    if len(this.items) % 2 == 0 {
        return float64(this.items[mid]+this.items[mid+1])/2
    }
    return float64(this.items[mid])
    // if this.left.Len() > this.right.Len() {
    //     return float64(this.left.items[0])
    // } else if this.right.Len() > this.left.Len() {
    //     return float64(this.right.items[0])
    // }
    // return float64(this.left.items[0]+this.right.items[0]) / 2.0
}


type minHeap struct {items []int}
func (m *minHeap) Less(i, j int) bool { return m.items[i] < m.items[j] }
func (m *minHeap) Swap(i, j int)  { m.items[i], m.items[j] = m.items[j], m.items[i] }
func (m *minHeap) Len() int  { return len(m.items) }
func (m *minHeap) Push(x interface{}) { m.items = append(m.items, x.(int)) }
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

type maxHeap struct {items []int}
func (m *maxHeap) Less(i, j int) bool { return m.items[i] > m.items[j] }
func (m *maxHeap) Swap(i, j int)  { m.items[i], m.items[j] = m.items[j], m.items[i] }
func (m *maxHeap) Len() int  { return len(m.items) }
func (m *maxHeap) Push(x interface{}) { m.items = append(m.items, x.(int)) }
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}



/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
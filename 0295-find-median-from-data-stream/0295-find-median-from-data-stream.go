type MedianFinder struct {
    left *maxHeap
    right *minHeap
}


func Constructor() MedianFinder {
    return MedianFinder{
        left: &maxHeap{items: []int{}},
        right : &minHeap{items: []int{}},
    }
}

func abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}

func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.left, num)
    
    if this.left.Len() != 0 && this.right.Len() != 0 {
        if this.left.items[0] > this.right.items[0] {
            heap.Push(this.right, heap.Pop(this.left))        
        }
    }
    
    if abs(len(this.left.items)-len(this.right.items)) >= 2 {
        if len(this.left.items) > len(this.right.items) { 
            heap.Push(this.right, heap.Pop(this.left))
        } else {
            heap.Push(this.left, heap.Pop(this.right))
        }
    }
    
}


func (this *MedianFinder) FindMedian() float64 {
    if len(this.left.items) > len(this.right.items) {
        return float64(this.left.items[0])
    } else if len(this.right.items) > len(this.left.items) {
        return float64(this.right.items[0])
    }
    return float64(this.left.items[0]+this.right.items[0]) / 2.0
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
type SmallestInfiniteSet struct {
    set map[int]struct{}
    mn *minHeap
    curr int
}


func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet{
        set: map[int]struct{}{},
        mn: &minHeap{items: []int{}},
        curr: 1,
    }
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    curr := this.curr
    if len(this.mn.items) == 0 {
        this.curr++
        return curr
    }
    
    top := heap.Pop(this.mn).(int)
    if top < curr {
        delete(this.set, top)
        return top
    }
    
    this.curr++
    return curr
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if num < this.curr {
        if _, ok := this.set[num]; !ok {
            heap.Push(this.mn, num)
            this.set[num] = struct{}{}
        }
    }
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */



type minHeap struct {
    items []int
}

func (m *minHeap) Less(i, j int) bool {return m.items[i] < m.items[j]}
func (m *minHeap) Swap(i, j int) {    
    m.items[i],m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Push(x interface{}) { m.items = append(m.items, x.(int))}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

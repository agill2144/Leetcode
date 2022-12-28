func minStoneSum(piles []int, k int) int {
    mh := &maxHeap{items: [][]int{}}
    total := 0
    for i := 0; i < len(piles); i++ {
        total += piles[i]
        heap.Push(mh, []int{i, piles[i]})
    }
    for k != 0 {
        top := heap.Pop(mh).([]int)
        idx, value := top[0],top[1]
        if value % 2 != 0 {
            piles[idx] = (value / 2)+1
        } else {
            piles[idx] = value / 2
        }
        total -= value-piles[idx]
        k--
        if k != 0 {heap.Push(mh, []int{idx, piles[idx]})}
    }
    return total
}

type maxHeap struct {
    items [][]int // [idx, value]
}

func (m *maxHeap) Less(i, j int)bool{return m.items[i][1] > m.items[j][1]}
func (m *maxHeap) Swap(i, j int){m.items[i], m.items[j] = m.items[j],m.items[i]}
func (m *maxHeap) Len()int{return len(m.items)}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *maxHeap) Pop()interface{}{
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
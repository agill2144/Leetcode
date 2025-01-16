/*
    approach: naive brute force
    - for each k
    - run a n traversal and find max idx
    
    n = len(gifts)
    tc = o(k*n)
    sc = o(1)


    approach: maxHeap
    - we have to take the largest val
    - and update it with sqrt of that val
    - then get the next highest val ( could be the same val we just updated )
    - therefore we need to constantly sort
    - and since we want the max each time
    - heap / max heap

    n = len(gifts)
    tc: o(n) + o(klogn)
    sc: o(n)
*/

func pickGifts(gifts []int, k int) int64 {
    var total int64
    for k != 0 {
        maxIdx := 0
        for i := 1; i < len(gifts); i++ {
            if gifts[i] > gifts[maxIdx] {maxIdx = i}
        }
        gifts[maxIdx] = int(math.Sqrt(float64(gifts[maxIdx])))
        k--
    }
    for i := 0; i < len(gifts); i++ {
        total += int64(gifts[i])
    }
    return total
}

// heap impl
// func pickGifts(gifts []int, k int) int64 {
//     mx := &maxHeap{items: []int{}}
//     var total int64
//     for i := 0; i < len(gifts); i++ {
//         total += int64(gifts[i])
//         heap.Push(mx, gifts[i])
//     }
//     for k != 0 {
//         top := heap.Pop(mx).(int)
//         sqrt := int(math.Sqrt(float64(top)))
//         diff := int64(top)-int64(sqrt)
//         total -= diff
//         heap.Push(mx, sqrt)
//         k--
//     }
//     return total
// }

// type maxHeap struct {
// 	items []int
// }

// func (m *maxHeap) Less(i, j int) bool {
// 	return m.items[i] > m.items[j]
// }
// func (m *maxHeap) Swap(i, j int) {
// 	m.items[i], m.items[j] = m.items[j], m.items[i]
// }
// func (m *maxHeap) Len() int {
// 	return len(m.items)
// }
// func (m *maxHeap) Push(x interface{}) {
// 	m.items = append(m.items, x.(int))
// }
// func (m *maxHeap) Pop() interface{} {
// 	out := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return out
// }

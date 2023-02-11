func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := []int{}
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] < nums2[n2] {
            merged = append(merged, nums1[n1])
            n1++
        } else {
            merged = append(merged, nums2[n2])
            n2++            
        }
    }
    for n1 < len(nums1) { merged = append(merged, nums1[n1]) ; n1++ }
    for n2 < len(nums2) { merged = append(merged, nums2[n2]) ; n2++ }
    
    left := &maxHeap{items: []int{}}
    right := &minHeap{items: []int{}}
    
    
    for i := 0; i < len(merged); i++ {
        heap.Push(left, merged[i])
        // ensure left top <= right top
        if left.Len() != 0 && right.Len() != 0 && 
        left.items[0] > right.items[0] {
            heap.Push(right, heap.Pop(left))
        }
        // ensure its balanced, so that the median is at the top of surface of both heaps
        // otherwise median could get burried somewhere deep inside a heap
        if left.Len() > right.Len() + 1 {
            heap.Push(right, heap.Pop(left))
        } else if right.Len() > left.Len() {
            heap.Push(left, heap.Pop(right))
        }

    }
    if len(left.items) > len(right.items) {
        return float64(left.items[0])
    } else if len(right.items) > len(left.items) {
        return float64(right.items[0])
    }
    return (float64(left.items[0]) + float64(right.items[0]))/2.0
}


func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}




type maxHeap struct {items []int}
func(m *maxHeap) Less(i, j int)bool {return m.items[i] > m.items[j]}
func(m *maxHeap) Swap(i, j int)     {m.items[i], m.items[j] = m.items[j], m.items[i]}
func(m *maxHeap) Len()int           {return len(m.items)}
func(m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
func(m *maxHeap) Pop()interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

type minHeap struct {items []int}
func(m *minHeap) Less(i, j int)bool {return m.items[i] < m.items[j]}
func(m *minHeap) Swap(i, j int)     {m.items[i], m.items[j] = m.items[j], m.items[i]}
func(m *minHeap) Len()int           {return len(m.items)}
func(m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
func(m *minHeap) Pop()interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
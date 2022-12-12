
// manual min heap implementation
// type MinHeap struct {
// 	items []int
// }

// func (mn *MinHeap) insert(item int) {
// 	mn.items = append(mn.items, item)
// 	mn.heapifyUp(len(mn.items)-1)
// }

// func (mn *MinHeap) extract()int {
// 	mn.swap(0, len(mn.items)-1)
// 	out := mn.items[len(mn.items)-1]
// 	mn.items = mn.items[:len(mn.items)-1]
// 	mn.heapifyDown(0)

// 	return out
// }

// func (mn *MinHeap) swap(a,b int) {
// 	mn.items[a], mn.items[b] = mn.items[b], mn.items[a]
// }

// func (mn *MinHeap) heapifyDown(idx int) {
// 	l := len(mn.items)-1
// 	if idx < 0 || idx > l {
// 		return
// 	}
// 	left := (2*idx)+1
// 	right := (2*idx)+2
// 	for (left <= l && mn.items[left] < mn.items[idx]) || (right <= l && mn.items[right] < mn.items[idx]) {
// 		swappingWith := idx
// 		if left <= l && mn.items[left] < mn.items[swappingWith] {
// 			swappingWith = left
// 		}
// 		if right <= l && mn.items[right] < mn.items[swappingWith] {
// 			swappingWith = right
// 		}
// 		if swappingWith == idx {return}
// 		mn.swap(swappingWith, idx)
// 		idx = swappingWith
// 		left = (2*idx)+1
// 		right = (2*idx)+2
// 	}
// }

// func (mn *MinHeap) heapifyUp(idx int) {
// 	l := len(mn.items)-1
// 	if idx < 0 || idx > l {
// 		return
// 	}
// 	parent := int(math.Floor(float64((idx-1)/2)))
// 	for parent <= l && mn.items[idx] < mn.items[parent] {
// 		mn.swap(idx, parent)
// 		idx = parent
// 		parent = int(math.Floor(float64((idx-1)/2)))
// 	}
// }

// func findKthLargest(nums []int, k int) int {
// 	/*
// 		create a min heap ( biggest num is on the top )
// 		put each num in max heap if min heap size != k -- if minHeap size == k and top item > current item then extract and insert
// 		return the top valie
// 	*/
// 	mn := &MinHeap{items: []int{}}
// 	for _, e := range nums {
// 		if len(mn.items) != k {
// 			mn.insert(e)
// 		} else {
// 			peek := mn.items[0]
// 			if e > peek {
// 				mn.extract()
// 				mn.insert(e)
// 			}
// 		}
// 	}
// 	return mn.items[0]
// }

//-----------------------------------------------------------------------------------------------------------------------
// container lib max heap implementation
// type MaxHeap struct {
// 	items []int
// }

// func (m *MaxHeap) Len() int {return len(m.items)}
// func (m *MaxHeap) Less(i, j int) bool {return m.items[i] > m.items[j]}
// func (m *MaxHeap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
// func (m *MaxHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
// func (m *MaxHeap) Pop() interface{} {
// 	i := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return i
// }

// func findKthLargest(nums []int, k int) int {
// 	var out int
// 	if nums == nil {
// 		return -1
// 	}
// 	mx := &MaxHeap{items: []int{}}
// 	for _, e := range nums {
// 		heap.Push(mx, e)
// 	}
// 	fmt.Println(mx.items)
// 	for i:=0; i<k; i++ {
// 		out = heap.Pop(mx).(int)
// 	}
// 	return out
// }
//-----------------------------------------------------------------------------------------------------------------------
// container lib min heap implementation
// sorting is always there -- sort nums and then return last kth element - but time would be o(nlogn)
// instead we can use minHeap
// time: o(nlogk)
// we loop over all numbers in nums o(n)
// we store them in heap but our insert/extract only ever has to deal with k items in heap therefore logk
// therefore total time: o(n) * o(logk) = o(nlogk)
// space: o(k) - at worse we store k elements in our min heap
func findKthLargest(nums []int, k int) int {
	var out int
	if nums == nil {
		return -1
	}
	mn := &MinHeap{items: []int{}}
	for _, e := range nums {
        heap.Push(mn, e)
        if len(mn.items) > k {heap.Pop(mn)}
	}
	out = mn.items[0]
	return out
}

// implements container heap interface : https://pkg.go.dev/container/heap#Interface
type MinHeap struct {
	items []int
}

func (m *MinHeap) Len() int {return len(m.items)}
func (m *MinHeap) Less(i, j int) bool {return m.items[i] < m.items[j]}
func (m *MinHeap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *MinHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
func (m *MinHeap) Pop() interface{} {
	i := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return i
}

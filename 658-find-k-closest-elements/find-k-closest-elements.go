func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    if x < arr[0] {
        return arr[:k]
    } else if x > arr[len(arr)-1] {
        return arr[n-k:n]
    }
    // find the left most on right side of x
    left := 0
    right := n-1
    idx := -1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] >= x {
            idx = mid
            right = mid-1
            if arr[mid] == x {break}
        } else {
            left = mid+1
        }
    }
    fmt.Println(idx, arr[idx])

    p1 := idx // go right
    p2 := idx-1 // go left
    out := []int{}
    for len(out) != k && p1 < n && p2 >= 0 {
        p1Dist := abs(arr[p1]-x)
        p2Dist := abs(arr[p2]-x)

        if p1Dist < p2Dist {
            out = append(out, arr[p1])
            p1++
        } else {
            out = append(out, arr[p2])
            p2--
        }
    }
    for len(out) != k && p1 < n{out = append(out, arr[p1]); p1++}
    for len(out) != k && p2 >= 0 {out = append(out, arr[p2]); p2--}
    sort.Ints(out)

    return out
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}

/*
    approach: heap
    - find k closest
    - heap vibes
    - closest = smallest = max heap
        - we will keep only k elements
        - sort by dist
        - if two distances are same, sort by value

    time = o(nlogk) + o(k) + o(klogk)
    space = o(k)

*/
// func findClosestElements(arr []int, k int, x int) []int {
//     mx := &maxHeap{items: []*heapNode{}}
//     // o(n * logk)
//     for i := 0; i < len(arr); i++ {
//         val := arr[i]
//         dist := abs(val-x)
//         heap.Push(mx, &heapNode{dist:dist, val:val})
//         if mx.Len() > k {
//             heap.Pop(mx)
//         }
//     }
//     out := []int{}
//     // o(k)
//     for mx.Len() != 0 {
//         out = append(out, heap.Pop(mx).(*heapNode).val)
//     }
//     // o(klogk)
//     sort.Ints(out)
//     return out
// }



// type heapNode struct {
//     dist int
//     val int
// }
// type maxHeap struct {
// 	items []*heapNode
// }

// func (m *maxHeap) Less(i, j int) bool {
// 	if m.items[i].dist == m.items[j].dist {
//         return m.items[i].val > m.items[j].val
//     }
//     return m.items[i].dist > m.items[j].dist
// }
// func (m *maxHeap) Swap(i, j int) {
// 	m.items[i], m.items[j] = m.items[j], m.items[i]
// }
// func (m *maxHeap) Len() int {
// 	return len(m.items)
// }
// func (m *maxHeap) Push(x interface{}) {
// 	m.items = append(m.items, x.(*heapNode))
// }
// func (m *maxHeap) Pop() interface{} {
// 	out := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return out
// }

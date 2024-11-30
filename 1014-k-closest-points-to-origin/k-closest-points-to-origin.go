type distNode struct {
    dist float64
    cord []int
}
// appraoch: quick select
// we want k closest to origin
// origin is 0,0 -> euclidean dist of origin is 0
// i.e we want k smallest distances
// using quick select, we can make k size partition on left
func kClosest(points [][]int, k int) [][]int {
    distances := []*distNode{}
    for i := 0; i < len(points); i++ {
        node := &distNode{
            dist: calcDist(points[i][0],points[i][1]),
            cord: points[i],
        }
        distances = append(distances, node)
    }
    left := 0
    right := len(distances)-1
    for left <= right {
        pivot := right
        nextSmaller := left
        for i := left; i < pivot; i++ {
            if distances[i].dist <= distances[pivot].dist {
                distances[i], distances[nextSmaller] = distances[nextSmaller], distances[i]
                nextSmaller++
            }
        }
        distances[pivot], distances[nextSmaller] = distances[nextSmaller], distances[pivot]
        if nextSmaller == k-1 {
            break
        } else if k-1 > nextSmaller {
            left = nextSmaller+1
        } else {
            right = nextSmaller-1
        }
    }

    out := [][]int{}
    for i := 0; i < k; i++ {
        out = append(out, distances[i].cord)
    }
    return out
}

// approach: max heap
// tc = o(n * logk) + o(k)
// sc = o(k)
// func kClosest(points [][]int, k int) [][]int {
//     mx := &maxHeap{items: []*distNode{}}
//     for i := 0; i < len(points); i++ {
//         node := &distNode{
//             dist: calcDist(points[i][0],points[i][1]),
//             cord: points[i],
//         }
//         heap.Push(mx, node)
//         if mx.Len() > k {
//             heap.Pop(mx)
//         }
//     }
//     out := [][]int{}
//     for mx.Len() != 0 {
//         top := heap.Pop(mx).(*distNode)
//         out = append(out, top.cord)
//     }
//     return out
// }
// type maxHeap struct {
// 	items []*distNode
// }

// func (m *maxHeap) Less(i, j int) bool {
// 	return m.items[i].dist > m.items[j].dist
// }
// func (m *maxHeap) Swap(i, j int) {
// 	m.items[i], m.items[j] = m.items[j], m.items[i]
// }
// func (m *maxHeap) Len() int {
// 	return len(m.items)
// }
// func (m *maxHeap) Push(x interface{}) {
// 	m.items = append(m.items, x.(*distNode))
// }
// func (m *maxHeap) Pop() interface{} {
// 	out := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return out
// }

// approach: brute force / sort
// tc = o(n) + o(nlogn) + o(k) = o(nlogn)
// sc = o(n) + o(n) = o(n)
// func kClosest(points [][]int, k int) [][]int {
//     type distNode struct {
//         dist float64
//         cord []int
//     }

//     // sc = o(n)
//     distances := []*distNode{}

//     // n = len(points)
//     // tc = o(n)
//     for i := 0; i < len(points); i++ {
//         distances = append(distances, &distNode{
//             dist: calcDist(points[i][0], points[i][1]),
//             cord: points[i],
//         })
//     }

//     // tc = o(nlogn)
//     sort.Slice(distances, func(i, j int)bool{
//         return distances[i].dist < distances[j].dist
//     })

//     out := [][]int{}
//     //tc = o(k)
//     for i := 0; i < k; i++ {
//         out = append(out, distances[i].cord)
//     }
//     return out
// }



func calcDist(x, y int) float64 {
    return math.Sqrt(float64(x*x) + float64(y*y))
}
func kClosest(points [][]int, k int) [][]int {
    if points == nil || len(points) == 0 {
        return nil
    }
    
    mx := &maxHeap{items: []item{}}
    for i := 0; i < len(points); i++ {
        x := points[i][0]
        y := points[i][1]
        dist := int(math.Pow(float64(x-0), 2)) + int(math.Pow(float64(y-0), 2))
        heap.Push(mx, item{dist: dist, points: points[i]})
        if mx.Len() > k {
            heap.Pop(mx)
        }
    }
    
    out := [][]int{}
    for k != 0 || mx.Len() != 0 {
        out = append(out, heap.Pop(mx).(item).points)
        k--
    }
    
    return out
    
}


type item struct {
    dist int
    points []int
}
type maxHeap struct{
    items []item
}

func(mx *maxHeap) Len() int {return len(mx.items)}
func(mx *maxHeap) Less(i,j int) bool {
    return mx.items[i].dist > mx.items[j].dist
}
func(mx *maxHeap) Swap(i, j int) {
    mx.items[i], mx.items[j] = mx.items[j], mx.items[i]
}
func(mx *maxHeap) Pop() interface{} {
    out := mx.items[len(mx.items)-1]
    mx.items = mx.items[:len(mx.items)-1]
    return out
}
func(mx *maxHeap) Push(x interface{}) {
    mx.items = append(mx.items, x.(item))
}

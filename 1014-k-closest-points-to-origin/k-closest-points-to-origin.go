
func kClosest(points [][]int, k int) [][]int {
    dists := []*distNode{}
    for i := 0; i < len(points); i++ {
        node := &distNode{
            dist: calcDist(points[i][0], points[i][1]),
            point: points[i],
        }
        dists = append(dists, node)
    }
    left := 0
    right := len(dists)-1
    for left <= right {
        pivot := right
        ns := left
        for i := left; i < pivot; i++ {
            if dists[i].dist <= dists[pivot].dist {
                dists[i], dists[ns] = dists[ns], dists[i]
                ns++
            }
        }
        dists[ns], dists[pivot] = dists[pivot], dists[ns]
        if ns == k-1 {break}
        if k-1 < ns {
            right = ns-1
        } else {
            left = ns+1
        }
    }

    out := [][]int{}
    for i := 0; i < k; i++ {
        out = append(out, dists[i].point)
    }
    return out
}

func calcDist(x, y int) float64 {
    return math.Sqrt(float64(x*x)+float64(y*y))
}

type distNode struct {
    dist float64
    point []int
}

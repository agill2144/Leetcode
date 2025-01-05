func kClosest(points [][]int, k int) [][]int {
    type distNode struct {
        dist float64
        point []int
    }
    dists := []*distNode{}
    for i := 0; i < len(points); i++ {
        dist := calcDist(points[i][0], points[i][1])
        dists = append(dists, &distNode{dist,points[i]})
    }
    targetIdx := k-1
    left := 0
    right := len(dists)-1
    for left <= right {
        ns := left
        pivot := right
        for i := left; i < pivot; i++ {
            if dists[i].dist <= dists[pivot].dist {
                dists[i], dists[ns] = dists[ns], dists[i]
                ns++
            }
        }
        dists[ns], dists[pivot] = dists[pivot], dists[ns]
        if ns == targetIdx {
            break
        }
        if targetIdx > ns {
            left = ns+1
        } else {
            right = ns-1
        }
    }
    out := [][]int{}
    for i := 0; i < len(dists) && len(out) != k; i++ {
        out = append(out, dists[i].point)
    }
    return out
}

func calcDist(x, y int) float64 {
    return math.Sqrt(math.Pow(float64(x), 2.0)+math.Pow(float64(y),2))
}

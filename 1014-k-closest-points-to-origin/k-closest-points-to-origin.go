func kClosest(points [][]int, k int) [][]int {
    type distNode struct {
        dist int
        cord []int
    }
    dists := []distNode{}
    n := len(points)
    for i := 0; i < n; i++ {
        dists = append(dists, distNode{
            dist: calcDist(points[i][0], points[i][1]),
            cord: points[i],
        })
    }
    left := 0
    right := len(dists)-1
    targetIdx := k-1
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
        if ns == targetIdx {break}
        if targetIdx > ns {
            left = ns+1
        } else {
            right = ns-1
        }
    }
    out := [][]int{}
    for i := 0; i < k; i++ {
        out = append(out, dists[i].cord)
    }
    return out
}

func calcDist(x, y int) int {
    return (x*x) + (y*y)
}
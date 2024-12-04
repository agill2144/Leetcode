func kClosest(points [][]int, k int) [][]int {
    type distNode struct {
        dist float64
        point []int
    }
    dists := []*distNode{}
    for i := 0; i < len(points); i++ {
        dists = append(dists, &distNode{
            dist: calcDist(points[i][0], points[i][1]),
            point: points[i],
        })
    }
    sort.Slice(dists, func(i, j int)bool{
        return dists[i].dist < dists[j].dist
    })
    out := [][]int{}
    for i := 0; i < k; i++ {
        out = append(out, dists[i].point)
    }
    return out
}

func calcDist(x, y int) float64 {
    return math.Sqrt(float64(x*x)+float64(y*y))
}
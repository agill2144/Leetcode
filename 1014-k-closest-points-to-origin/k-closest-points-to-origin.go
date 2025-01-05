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
    sort.Slice(dists, func(i, j int)bool{
        return dists[i].dist < dists[j].dist
    })
    out := [][]int{}
    for i := 0; i < len(dists) && len(out) != k; i++ {
        out = append(out, dists[i].point)
    }
    return out
}

func calcDist(x, y int) float64 {
    return math.Sqrt(math.Pow(float64(x), 2.0)+math.Pow(float64(y),2))
}

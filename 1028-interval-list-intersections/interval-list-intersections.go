func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    out := [][]int{}
    a, b := 0,0
    for a < len(firstList) && b < len(secondList) {
        aStart, aEnd := firstList[a][0], firstList[a][1]
        bStart, bEnd := secondList[b][0], secondList[b][1]
        start := max(aStart, bStart)
        end := min(aEnd, bEnd)
        if start <= end {
            out = append(out, []int{start, end})
        }
        if aEnd < bEnd {a++} else {b++}
    }
    return out
}
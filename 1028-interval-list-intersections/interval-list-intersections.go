func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    out := [][]int{}
    f, s := 0, 0
    for f < len(firstList) && s < len(secondList) {
        aStart, aEnd := firstList[f][0], firstList[f][1]
        bStart, bEnd := secondList[s][0], secondList[s][1]
        start, end := max(aStart, bStart), min(aEnd, bEnd)
        if start <= end {
            out = append(out, []int{start, end})
        }
        if aEnd < bEnd {f++} else {s++}
    }
    return out
}
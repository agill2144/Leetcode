func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    a, b := 0, 0
    out := [][]int{}
    for a < len(firstList) && b < len(secondList) {
        aStart, aEnd := firstList[a][0], firstList[a][1]
        bStart, bEnd := secondList[b][0], secondList[b][1]
        if aStart > bEnd || bStart > aEnd {
            if aEnd < bEnd {a++} else {b++}
        } else {
            out = append(out, []int{max(aStart, bStart), min(aEnd, bEnd)})
            if aEnd < bEnd {a++} else {b++}
        }
    }
    return out
}
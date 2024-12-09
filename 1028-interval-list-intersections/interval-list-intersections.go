func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    out := [][]int{}
    a, b := 0,0
    for a < len(firstList) && b < len(secondList) {
        aStart, aEnd := firstList[a][0], firstList[a][1]
        bStart, bEnd := secondList[b][0], secondList[b][1]
        // 1. if there is no intersection, move on
        if aStart > bEnd {
            // if a starts after b ends
            b++
        } else if bStart > aEnd {
            // b starts after a ends
            a++
        } else {
            // there is an intersection!
            // we want the latest start time
            // and we want the earlist end time
            out = append(out, []int{max(aStart,bStart), min(aEnd, bEnd)})
            if aEnd < bEnd {a++} else {b++}
        }

    }
    return out
}
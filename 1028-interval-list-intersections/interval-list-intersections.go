func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    a, b := 0, 0
    out := [][]int{}
    for a < len(firstList) && b < len(secondList) {
        start := max(firstList[a][0], secondList[b][0])
        end := min(firstList[a][1], secondList[b][1])
        if start <= end {
            out = append(out, []int{start, end})
        }
        if firstList[a][1] < secondList[b][1] {
            a++
        } else {
            b++
        }
    }
    return out
}
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    if len(firstList) == 0 || len(secondList) == 0 {return nil}
    out := [][]int{}
    f, s := 0, 0
    for f < len(firstList) && s < len(secondList) {
        start := max(firstList[f][0], secondList[s][0])
        end := min(firstList[f][1], secondList[s][1])
        if start <= end {
            out = append(out, []int{start, end})
        }
        if firstList[f][1] < secondList[s][1] {
            f++
        } else {
            s++
        }
    }
    return out
}
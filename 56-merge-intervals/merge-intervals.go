func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {return nil}
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        // when is it not overlapping ?
        // when i starts after prev ends ( prev = last element in output array )
        if intervals[i][0] > out[len(out)-1][1] {
            out = append(out, intervals[i])
        } else {
            out[len(out)-1][1] = max(intervals[i][1], out[len(out)-1][1])
        }
    }
    return out
}
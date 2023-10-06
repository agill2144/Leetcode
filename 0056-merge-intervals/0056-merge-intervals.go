func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    st := [][]int{}
    for i := 0; i < len(intervals); i++ {
        if len(st) != 0 && intervals[i][0] <= st[len(st)-1][1] {
            st[len(st)-1][0] = min(intervals[i][0], st[len(st)-1][0])
            st[len(st)-1][1] = max(intervals[i][1], st[len(st)-1][1])
        } else {
            st = append(st, intervals[i])
        }
    }
    return st
}


func min(x, y int) int {
    if x < y {return x}
    return y
}
func max(x, y int) int {
    if x > y {return x}
    return y
}
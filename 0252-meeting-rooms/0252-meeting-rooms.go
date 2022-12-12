/*
    time : o(nlogn) + o(n)
    space: o(n) if sorting algo used merge sort
*/
func canAttendMeetings(intervals [][]int) bool {
    if intervals == nil || len(intervals) <= 1 {return true}
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i-1][1]
        if start < end {
            return false
        }
    }
    return true
}
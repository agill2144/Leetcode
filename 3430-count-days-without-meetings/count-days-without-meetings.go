func countDays(days int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int)bool{return meetings[i][0] < meetings[j][0]})
    merged := [][]int{}
    for i := 0; i < len(meetings); i++ {
        start, end := meetings[i][0], meetings[i][1]
        if len(merged) == 0 {
            merged = append(merged, meetings[i])
        } else {
            prev := merged[len(merged)-1]
            prevStart, prevEnd := prev[0],prev[1]
            if start <= prevEnd {
                merged = merged[:len(merged)-1]
                newStart := prevStart
                newEnd := max(end, prevEnd)
                merged = append(merged, []int{newStart, newEnd})
            } else {
                merged = append(merged, meetings[i])
            }
        }
    }
    totalMeetingDays := 0
    for i := 0; i < len(merged); i++ {
        start, end := merged[i][0], merged[i][1]
        totalMeetingDays += (end-start+1)
    }
    return days-totalMeetingDays
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}
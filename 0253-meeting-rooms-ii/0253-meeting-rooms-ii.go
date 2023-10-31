func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })
    // tracks meeting end times
    rooms := []int{}
    
    for i := 0; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]
        if len(rooms) == 0 {
            rooms = append(rooms, end)
        } else {
            // find which room's meeting has ended
            // -- take existing room if any
            found := false
            for j := 0; j < len(rooms); j++ {
                if start >= rooms[j] {
                    // found an existing room that we can take
                    rooms[j] = end
                    found = true
                    break
                }
            }
            // -- otherwise create a new room
            if !found {
                rooms = append(rooms, end)                
            }
        }
    }
    return len(rooms)
}
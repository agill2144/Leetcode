func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })

    // rooms keeping track of which meeting is ending at what time
    // if we have a new meeting, we can check with what rooms are
    // free, and the rooms that will be free are the ones whose 
    // meetings have already ended compared currentTime
    rooms := []int{}
    for i := 0; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]
        if len(rooms) == 0 {
            rooms = append(rooms, end)
        } else {
            foundExistingRoom := false
            for j := 0; j < len(rooms); j++ {
                if start >= rooms[j] {
                    rooms[j] = end
                    foundExistingRoom = true
                    break
                }
            }
            // if none of the rooms are avail, we need a new one
            if !foundExistingRoom {
                rooms = append(rooms, end)
            }
        }
    }
    return len(rooms)
}
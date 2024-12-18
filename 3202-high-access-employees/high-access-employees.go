func findHighAccessEmployees(access_times [][]string) []string {
    parsedTimings := map[string][]int{}
    for i := 0; i < len(access_times); i++ {
        name := access_times[i][0]
        t := access_times[i][1]
        hh, _ := strconv.Atoi(t[:2])
        mm, _ := strconv.Atoi(t[2:])
        mins := (hh*60) + mm
        parsedTimings[name] = append(parsedTimings[name], mins)
    }
    out := []string{}
    for name, times := range parsedTimings {
        if len(times) < 3 {continue}
        sort.Ints(times)
        left := 0
        for i := 0; i < len(times); i++ {
            // get a window of size 3 first
            if i-left+1 < 3 {continue}
            // when we have a window of size 3
            // then make sure its the 2 entreme ends
            // are within the limit (highest time, lowest time)
            // if they are not < 60 mins, shrink window from left
            for left <= i && times[i] >= times[left]+60 {
                left++
            }
            // we dont need to check adj times, 
            // because, the largest possible diff in a sorted array
            // will always be the lastElement - firstElement
            // if our extreme end is satisfied ( right Time - left time < 60 mins )
            // the entire subarray is satisfied
            // if we have a subarray of size 3, save this ans
            if i-left+1 == 3 {out =append(out, name); break}
        }
    }
    return out
}
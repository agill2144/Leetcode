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
        count := 1
        slow := 0
        for i := 1; i < len(times); i++ {
            curr := times[i]
            prev := times[i-1]
            for slow <= i && curr >= times[slow]+60 {
                slow++
                count--
                if count < 0 {count=0}
            }
            if curr < prev+60 && curr < times[slow]+60 {
                count++
                if count == 3 {
                    out = append(out, name)
                    break
                }
            } else {
                slow = i
                count = 1
            }
        }
    }
    return out
}
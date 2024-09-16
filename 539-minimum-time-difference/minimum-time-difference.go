func findMinDifference(timePoints []string) int {
    minutes := []int{}
    for i := 0; i < len(timePoints); i++ {
        t := strings.Split(timePoints[i], ":")
        hh, _ := strconv.Atoi(t[0])
        mm, _ := strconv.Atoi(t[1])

        // hour to mins 
        // 60 mins in 1 hour
        // 120 mins in 2 hour ( 2 * 60 )
        // 180 mins in 3 hour ( 3 * 60 )
        tInMins := hh*60
        // now add the remaining mins
        tInMins += mm
        minutes = append(minutes, tInMins)
    }
    sort.Ints(minutes)
    fmt.Println(minutes)
    // smallest diff ( adj numbers or 0 and n-1 element)

    diff := (minutes[0]+ (24*60)) - minutes[len(minutes)-1]
    for i := 1; i < len(minutes); i++ {
        diff = min(diff, minutes[i]-minutes[i-1])
    }

    return diff
}
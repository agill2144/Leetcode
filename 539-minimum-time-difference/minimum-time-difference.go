func findMinDifference(timePoints []string) int {
    // mins in 1 hour = 60
    // eg; 2hr = 120 mins
    minutes := []int{}
    for i := 0; i < len(timePoints); i++ {
        t := strings.Split(timePoints[i], ":")
        hh, mm := t[0], t[1]
        hhInt, _ := strconv.Atoi(hh)
        mmInt, _ := strconv.Atoi(mm)
        m := (hhInt * 60) + mmInt
        if m == 0 {m = 1440}
        minutes = append(minutes, m)
    }
    sort.Ints(minutes)
    smallestDiff := minutes[0]+(24*60) - minutes[len(minutes)-1]
    for i := 1; i < len(minutes); i++ {
        smallestDiff = min(smallestDiff, minutes[i]-minutes[i-1])
    }
    return smallestDiff
}
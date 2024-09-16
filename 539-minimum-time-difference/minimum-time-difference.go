// convert to minutes
// sort
// find smallest minutes differnce
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
    // smallest diff ( adj numbers or 0 and n-1 element)
    // since time is circule (i.e after 23:59 -> 00:00 )
    // the 0th minutes needs to be relative to the n-1th minutes element
    // 00:00 -> 0
    // 23:50 -> 1439
    // if we did 1439-0 = 1439
    // but really ... 00:00 to minutes should be 1400
    // therefore we need to make sure that 0th time is after midnight
    // therefore minutes[0] + 24hr*60min 
    diff := (minutes[0]+ (24*60)) - minutes[len(minutes)-1]
    for i := 1; i < len(minutes); i++ {
        // rest is just comparing mins difference between adj elements
        diff = min(diff, minutes[i]-minutes[i-1])
    }

    return diff
}
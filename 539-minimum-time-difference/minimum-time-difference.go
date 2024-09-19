// convert to minutes
// sort
// find smallest minutes differnce

func findMinDifference(timePoints []string) int {
    mins := []int{}
    for i := 0; i < len(timePoints); i++ {
        t := strings.Split(timePoints[i], ":")
        hh, _ := strconv.Atoi(t[0])
        mm, _ := strconv.Atoi(t[1])
        mins = append(mins, (hh*60)+mm)
    }
    sort.Ints(mins)
    // smallest diff ( adj numbers or 0 and n-1 element)
    // since time is circule (i.e after 23:59 -> 00:00 )
    // the 0th minutes needs to be relative to the n-1th minutes element
    // 00:00 -> 0
    // 23:50 -> 1439
    // if we did 1439-0 = 1439
    // but really ... 00:00 to minutes should be 1400
    // therefore we need to make sure that 0th time is after midnight
    // therefore minutes[0] + 24hr*60min 
    minDiff := (mins[0]+(24*60)) - mins[len(mins)-1]
    for i := 1; i < len(mins); i++ {
        minDiff = min(minDiff, mins[i]-mins[i-1])
    }
    return minDiff
}

/*
    1hr = 60 mins
    2hr = 120 mins ( 2 * 60 )
    5hr = 5*60 = 300 mins

    05:34 = (5hr*60) + 34 = 300 + 34 = 334 
*/
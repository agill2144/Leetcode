func findMinDifference(timePoints []string) int {
    mins := []int{}
    for i := 0; i < len(timePoints); i++ {
        hhmm := strings.Split(timePoints[i], ":")
        hh, _ := strconv.Atoi(hhmm[0])
        mm, _ := strconv.Atoi(hhmm[1])
        mins = append(mins, (hh*60)+mm)
    }
    sort.Ints(mins)
    // edge case ( walk forward from last time till mins[0] )
    // mins[-1] = 23:50 (i.e 1430 mins)
    // mins[0] = 02:00 (i.e 120 mins)
    // the difference between these 2 is not 1430-120 = 1320 mins ..
    // but really its (10 mins from 1430 -> 1440) + 120 mins
    // we can get remaining mins left until 1440 is hit by; 1440-mins[-1]
    // then we can simply add the mins at mins[0]
    minDiff := (1440-mins[len(mins)-1])+mins[0]
    for i := 1; i < len(mins); i++ {
        minDiff = min(minDiff, mins[i]-mins[i-1])
    }
    return minDiff
}  
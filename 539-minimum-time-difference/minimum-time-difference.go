func findMinDifference(timePoints []string) int {
    mins := []int{}
    for i := 0; i < len(timePoints); i++ {
        hhmm := strings.Split(timePoints[i], ":")
        hh, _ := strconv.Atoi(hhmm[0])
        mm, _ := strconv.Atoi(hhmm[1])
        m := (hh*60)+mm
        mins = append(mins, m)
    }
    sort.Ints(mins)
    // its possible that ele[0] is 0 or some early time in morning
    // and the last element is closer to midnight
    // but their diff would be large as is
    // 2:00am, 11:50pm ; 2am = 120 mins ; 11:50pm = 1430 min
    // 1430 - 120 = 1220 mins
    // but if walked back from 2am ( anti-clock wise )
    // 2am->1am (60 mins) + 1am -> midnight (60 mins) + midnight -> 11:50 (10 mins)
    // 60+60+10 = 130 mins
    // therefore we need to handle the walking back calculation and the only walk back 
    // we need to handle is smallest possible time ( 0th idx )
    // i.e walk back from time at 0th idx to time at n-1th idx
    // there are 2 possible times at idx 0
    // time[0] = 0 ( i.e assume 1440 mins )
    // time[0] = not 0 ( ex: 120 mins )
    // walkBack := 0
    // if mins[0] != 0 {walkBack += mins[0]}
    // // assume we walked back to midnight and mightnight converted to mins = 1440
    // // now get the diff from midnight to last mins element
    // walkBack += (1440-mins[len(mins)-1])
    // minDiff := walkBack
    // OR walk forward from last element to mins[0]
    minDiff := (1440-mins[len(mins)-1]) + mins[0]
    for i := 1; i < len(mins); i++ {
        minDiff = min(minDiff, mins[i]-mins[i-1])
    }
    return minDiff
}
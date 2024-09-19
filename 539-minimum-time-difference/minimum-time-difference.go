func findMinDifference(timePoints []string) int {
    mins := []int{}
    for i := 0; i < len(timePoints); i++ {
        t := strings.Split(timePoints[i], ":")
        hh, _ := strconv.Atoi(t[0])
        mm, _ := strconv.Atoi(t[1])
        mins = append(mins, (hh*60)+mm)
    }
    sort.Ints(mins)
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
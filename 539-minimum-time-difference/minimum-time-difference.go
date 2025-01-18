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
    // its possible that
    minDiff := mins[0]+(24*60) - mins[len(mins)-1]
    for i := 1; i < len(mins); i++ {
        minDiff = min(minDiff, mins[i]-mins[i-1])
    }
    return minDiff
}
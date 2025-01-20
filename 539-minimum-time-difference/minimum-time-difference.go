func findMinDifference(timePoints []string) int {
    mins := []int{}
    for i := 0; i < len(timePoints); i++ {
        hhmm := strings.Split(timePoints[i], ":")
        hh, _ := strconv.Atoi(hhmm[0])
        mm, _ := strconv.Atoi(hhmm[1])
        mins = append(mins, (hh*60)+mm)
    }
    sort.Ints(mins)
    minDiff := (1440-mins[len(mins)-1])+mins[0]
    for i := 1; i < len(mins); i++ {
        minDiff = min(minDiff, mins[i]-mins[i-1])
    }
    return minDiff
}  
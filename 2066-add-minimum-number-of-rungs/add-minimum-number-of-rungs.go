func addRungs(rungs []int, dist int) int {
    count := 0
    idx := -1
    for i := 0; i < len(rungs); i++ {
        rung := rungs[i]
        pos := 0
        if idx != -1 {pos = rungs[idx]}
        if rung - pos > dist {
            rungsNeeded := (rungs[i]-1-pos)/dist
            count += rungsNeeded
        }
        idx = i
    }
    return count
}
func addRungs(rungs []int, dist int) int {
    count := 0
    curr := 0
    for i := 0; i < len(rungs); i++ {
        if curr + dist < rungs[i] {
            diff := rungs[i]-curr-1
            count += int(math.Floor(float64(diff)/float64(dist)))
        }
        curr = rungs[i]
    }
    return count
}

// dist = 2
// 25 100

// 
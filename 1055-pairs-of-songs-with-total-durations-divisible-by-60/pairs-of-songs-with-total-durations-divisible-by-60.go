func numPairsDivisibleBy60(time []int) int {
    rem := map[int]int{}
    count := 0
    for i := 0; i < len(time); i++ {
        r := time[i] % 60
        target := 60-r
        if r == 0 {target = 0}
        count += rem[target]
        rem[r]++
    }
    return count
}
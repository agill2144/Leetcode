func numPairsDivisibleBy60(time []int) int {
    rem := map[int]int{}
    count := 0
    for i := 0; i < len(time); i++ {
        r := time[i] % 60
        if r == 0 {
            count += rem[0]
        } else {
            count += rem[60-r]
        }
        rem[r]++
    }
    return count
}
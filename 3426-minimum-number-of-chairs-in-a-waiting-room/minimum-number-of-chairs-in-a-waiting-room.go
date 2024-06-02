func minimumChairs(s string) int {
    maxCount := 0
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'E' {
            count++
            maxCount = max(maxCount, count)
        } else {
            count--
        }
    }
    return maxCount
}
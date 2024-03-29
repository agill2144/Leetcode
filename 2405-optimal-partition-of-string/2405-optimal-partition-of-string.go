func partitionString(s string) int {
    left := 0
    charIdx := map[byte]int{}
    count := 1
    for i := 0; i < len(s); i++ {
        char := s[i]
        if lastSeenAt, ok := charIdx[char]; ok && left <= lastSeenAt {
            count++
            left = i
        }
        charIdx[char] = i
    }
    return count
}
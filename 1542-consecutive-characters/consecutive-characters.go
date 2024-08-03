func maxPower(s string) int {
    if len(s) == 1 {return 1}
    maxLen := 0
    i := 0
    for i < len(s) {
        start := i
        char := s[i]
        for i+1 < len(s) && s[i+1] == char {
            i++
        }
        maxLen = max(maxLen, i-start+1)
        i++
    }
    return maxLen
}
func secondsToRemoveOccurrences(s string) int {
    c := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '0' && i+1 < len(s) && s[i+1] == '1' {
            c++
        }
    }
    if c == 0 {return 0}
    ss := strings.ReplaceAll(s, "01","10")
    return 1 + secondsToRemoveOccurrences(ss)
}
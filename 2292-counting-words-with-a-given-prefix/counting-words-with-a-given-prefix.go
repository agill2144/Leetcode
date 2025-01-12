func prefixCount(words []string, pref string) int {
    count := 0
    for i := 0; i < len(words); i++ {
        if strings.HasPrefix(words[i], pref) {count++}
    }
    return count
}
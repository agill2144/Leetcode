func countPrefixSuffixPairs(words []string) int {
    count := 0
    for i := 0; i < len(words); i++ {
        for j := i+1; j < len(words); j++ {
            if isPrefixAndSuffix(words[i], words[j]) {
                count++
            }
        }
    }
    return count
}

func isPrefixAndSuffix(a, b string) bool {
    return strings.HasPrefix(b, a) &&
        strings.HasSuffix(b,a)
}
func strStr(haystack string, needle string) int {
    for i := 0; i <= len(haystack)-len(needle); i++ {
        subStr := haystack[i:i+len(needle)]
        if subStr == needle {return i}
    }
    return -1
}
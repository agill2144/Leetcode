func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1 // no way a bigger word can exist in a smaller word
    }
    if len(needle) == 0 {return 0}
    for i := 0; i < len(haystack)+1-len(needle); i++ {
        if string(haystack[i:i+len(needle)]) == needle {
            return i
        }
    }
    return -1
}
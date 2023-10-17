func strStr(haystack string, needle string) int {
    lps := getLPS(needle)
    i := 0
    j := 0
    for i < len(haystack) {
        if haystack[i] == needle[j] {
            i++
            j++
            if j == len(needle) {
                return i-len(needle)
            }
        } else if haystack[i] != needle[j] && j > 0 {
            j = lps[j-1]
        } else {
            i++
        }
    }
    return -1
}


func getLPS(pattern string) []int {
    lps := make([]int, len(pattern))
    i := 1
    j := 0
    for i < len(pattern) {
        if pattern[j] == pattern[i] {
            j++
            lps[i] = j
            i++
        } else if pattern[j] != pattern[i] && j > 0{
            j = lps[j-1]
        } else {
            i++
        }
    }
    return lps
}